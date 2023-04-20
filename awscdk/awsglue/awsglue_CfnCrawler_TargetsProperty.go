package awsglue


// Specifies data stores to crawl.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   targetsProperty := &TargetsProperty{
//   	CatalogTargets: []interface{}{
//   		&CatalogTargetProperty{
//   			DatabaseName: jsii.String("databaseName"),
//   			Tables: []*string{
//   				jsii.String("tables"),
//   			},
//   		},
//   	},
//   	DynamoDbTargets: []interface{}{
//   		&DynamoDBTargetProperty{
//   			Path: jsii.String("path"),
//   		},
//   	},
//   	JdbcTargets: []interface{}{
//   		&JdbcTargetProperty{
//   			ConnectionName: jsii.String("connectionName"),
//   			Exclusions: []*string{
//   				jsii.String("exclusions"),
//   			},
//   			Path: jsii.String("path"),
//   		},
//   	},
//   	MongoDbTargets: []interface{}{
//   		&MongoDBTargetProperty{
//   			ConnectionName: jsii.String("connectionName"),
//   			Path: jsii.String("path"),
//   		},
//   	},
//   	S3Targets: []interface{}{
//   		&S3TargetProperty{
//   			ConnectionName: jsii.String("connectionName"),
//   			DlqEventQueueArn: jsii.String("dlqEventQueueArn"),
//   			EventQueueArn: jsii.String("eventQueueArn"),
//   			Exclusions: []*string{
//   				jsii.String("exclusions"),
//   			},
//   			Path: jsii.String("path"),
//   			SampleSize: jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnCrawler_TargetsProperty struct {
	// Specifies AWS Glue Data Catalog targets.
	CatalogTargets interface{} `field:"optional" json:"catalogTargets" yaml:"catalogTargets"`
	// Specifies Amazon DynamoDB targets.
	DynamoDbTargets interface{} `field:"optional" json:"dynamoDbTargets" yaml:"dynamoDbTargets"`
	// Specifies JDBC targets.
	JdbcTargets interface{} `field:"optional" json:"jdbcTargets" yaml:"jdbcTargets"`
	// A list of Mongo DB targets.
	MongoDbTargets interface{} `field:"optional" json:"mongoDbTargets" yaml:"mongoDbTargets"`
	// Specifies Amazon Simple Storage Service (Amazon S3) targets.
	S3Targets interface{} `field:"optional" json:"s3Targets" yaml:"s3Targets"`
}

