package esdynamodb

import (
	"github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
)

const tableName = "esdynamodb-events"

// EnsureInitialized ensures that the DynamoDB table is created
// for the given AWS session.
func EnsureInitialized(sess *session.Session, cfg *aws.Config) error {
	svc := dynamodb.New(sess, cfg)
	input := &dynamodb.CreateTableInput{
		// When you use non-key attribute in at "AttributeDefinitions", you must use it as index, otherwise it's against the way of dynamodb to work. see http://docs.aws.amazon.com/amazondynamodb/latest/gettingstartedguide/GettingStarted.JsShell.06.html
		// So no need to put non-key attribute in "AttributeDefinitions" if you're not gonna use it as index or primary key.
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("StreamID"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("Version"),
				AttributeType: aws.String("N"),
			},
			{
				AttributeName: aws.String("UNIXNanoTimestamp"),
				AttributeType: aws.String("N"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("StreamID"),
				KeyType:       aws.String("HASH"),
			},
			{
				AttributeName: aws.String("Version"),
				KeyType:       aws.String("RANGE"),
			},
		},
		GlobalSecondaryIndexes: []*dynamodb.GlobalSecondaryIndex{
			{
				IndexName: aws.String("GlobalStream"),
				KeySchema: []*dynamodb.KeySchemaElement{
					{
						AttributeName: aws.String("UNIXNanoTimestamp"),
						KeyType:       aws.String("HASH"),
					},
					{
						AttributeName: aws.String("Version"),
						KeyType:       aws.String("RANGE"),
					},
				},
				Projection: &dynamodb.Projection{
					ProjectionType: aws.String("ALL"),
				},
				ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
					ReadCapacityUnits:  aws.Int64(1),
					WriteCapacityUnits: aws.Int64(1),
				},
			},			
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(1),
			WriteCapacityUnits: aws.Int64(1),
		},
		TableName: aws.String(tableName),
	}

	_, err := svc.CreateTable(input)
	return err
}