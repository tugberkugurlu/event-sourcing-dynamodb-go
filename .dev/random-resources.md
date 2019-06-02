 - https://github.com/golang/go/wiki/Modules
 - https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/transaction-apis.html
 - https://aws.amazon.com/blogs/aws/new-amazon-dynamodb-transactions/
 - [Amazon DynamoDB Examples Using the AWS SDK for Go](https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/using-dynamodb-with-go-sdk.html)
 - [Getting Started with the AWS SDK for Go](https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/setting-up.html)
 - https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_CreateTable.html
 - https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.CoreComponents.html
 
https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_AttributeDefinition.html

```
AttributeType
The data type for the attribute, where:

S - the attribute is of type String

N - the attribute is of type Number

B - the attribute is of type Binary

Type: String

Valid Values: S | N | B

Required: Yes
```

https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_GlobalSecondaryIndex.html
https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_Projection.html
```
ProjectionType
The set of attributes that are projected into the index:

KEYS_ONLY - Only the index and primary keys are projected into the index.

INCLUDE - Only the specified table attributes are projected into the index. The list of projected attributes are in NonKeyAttributes.

ALL - All of the table attributes are projected into the index.

Type: String

Valid Values: ALL | KEYS_ONLY | INCLUDE

Required: No
```

 - https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/DynamoDBLocal.UsageNotes.html
 - https://dwmkerr.com/run-amazon-dynamodb-locally-with-docker/
 - https://www.stevejgordon.co.uk/running-aws-dynamodb-locally-for-net-core-developers
 - https://medium.com/@mcleanjnathan/testing-with-dynamo-local-and-go-7b7000ef9602


 "2019/06/02 21:30:02 ValidationException: The number of attributes in key schema must match the number of attributesdefined in attribute definitions.": https://stackoverflow.com/questions/30866030/number-of-attributes-in-key-schema-must-match-the-number-of-attributes-defined-i