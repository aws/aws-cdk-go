package awsglue


// Specifies data stores to crawl.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   targetsProperty := &targetsProperty{
//   	catalogTargets: []interface{}{
//   		&catalogTargetProperty{
//   			databaseName: jsii.String("databaseName"),
//   			tables: []*string{
//   				jsii.String("tables"),
//   			},
//   		},
//   	},
//   	dynamoDbTargets: []interface{}{
//   		&dynamoDBTargetProperty{
//   			path: jsii.String("path"),
//   		},
//   	},
//   	jdbcTargets: []interface{}{
//   		&jdbcTargetProperty{
//   			connectionName: jsii.String("connectionName"),
//   			exclusions: []*string{
//   				jsii.String("exclusions"),
//   			},
//   			path: jsii.String("path"),
//   		},
//   	},
//   	mongoDbTargets: []interface{}{
//   		&mongoDBTargetProperty{
//   			connectionName: jsii.String("connectionName"),
//   			path: jsii.String("path"),
//   		},
//   	},
//   	s3Targets: []interface{}{
//   		&s3TargetProperty{
//   			connectionName: jsii.String("connectionName"),
//   			dlqEventQueueArn: jsii.String("dlqEventQueueArn"),
//   			eventQueueArn: jsii.String("eventQueueArn"),
//   			exclusions: []*string{
//   				jsii.String("exclusions"),
//   			},
//   			path: jsii.String("path"),
//   			sampleSize: jsii.Number(123),
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

