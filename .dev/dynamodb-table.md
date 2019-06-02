Partition Key: BoundedContext name + aggregate id
Range Key: event store unix timestamp nano seconds + version
Version: version

Primary Key: Partition Key + version as a composite primary key
Event: name of the event
EventJsonMetadata: such as the name of the type etc. See https://stackoverflow.com/a/7855298/463785 and also see https://github.com/jetbasrawi/go.cqrs/blob/e4d812d57f090ecede016aa36d70c73626a8eb17/repository.go#L60-L67
EventJsonContent: serialized content
OccuredAt: event occurance time in unix timestamp nano seconds 