## We can perform consistency checks per aggregate level with DynamoDB

The assumption here is that we can use the aggregate id as the partition key, and aggregate id and version for the event as the composite primary key. This should give us the optimistic concurrency behaviour we want when we try to insert data into DynamoDB and it will give us an error if the expected version is already there.

## It's OK to increase the version per each event

If we were to get multiple events per command execution to be stored, we need to increase the version per each event. As the version is squential, this should really be fine (TM).

## We can gurantee storage of multiple events in a single transaction by guranteeing atomicity and durability

We should be able to do this with [`TransactWriteItems`](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/transaction-apis.html#transaction-apis-txwriteitems) API.

see https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/transactions.html

## It's OK to cap that we can store up to 10 events in a single transaction

As you can see in [the doc of `TransactWriteItems` API](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/transaction-apis.html#transaction-apis-txwriteitems):

> TransactWriteItems is a synchronous and idempotent write operation that groups up to 10 write actions in a single all-or-nothing operation.

The assumption is that this is going to be OK.

## We can guarantee causal ordering for the events at the aggregate level

Just scan based on the the range key which we can use UNIX timestamp for the event save request + the version for.

## Event dispatch tracking is someone else's concern

This is about persisting the events and being able to use it as the source of truth for the events. As dispatching is another transaction and cannot be guaranteed to be performed, we won't be dealing witht that.

It's usually the concern of some other part of the consumer system to check this and make sure that events are distpacthed as expected.

### How can this be tracked though?

Keeping an external state per each event and marking the events as dispatched in that case. For example, we can store the aggregate id and event version in Redis before we commit them. Then, we can make sure that we dispatch these and remove them in Redis after the dispatch is completed.

In order to ensure that we dispatch all the events (i.e. cover the cases where we fail to dispatch after storage), we can have a side enqueer which can ensure that all the events are dispatched after a certain while.

We can also store the versions as sorted list per aggregate id on redis and ensure that on each dispatch we dispatch the events we haven't dispatched so far in order.

Some fundemental assumptions here:

 - We can fail after we store the event on Redis and before storing them on DynamoDB. so, the side enqueer needs to ensure that events actually exist before dispatching.
 - We may fail to delete the disptached event from Redis even if the event is dispatched. Therefore, the system should be able to deal with events which are dispatched multiple times unexpectedly.