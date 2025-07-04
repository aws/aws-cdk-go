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
//   			ConnectionName: jsii.String("connectionName"),
//   			DatabaseName: jsii.String("databaseName"),
//   			DlqEventQueueArn: jsii.String("dlqEventQueueArn"),
//   			EventQueueArn: jsii.String("eventQueueArn"),
//   			Tables: []*string{
//   				jsii.String("tables"),
//   			},
//   		},
//   	},
//   	DeltaTargets: []interface{}{
//   		&DeltaTargetProperty{
//   			ConnectionName: jsii.String("connectionName"),
//   			CreateNativeDeltaTable: jsii.Boolean(false),
//   			DeltaTables: []*string{
//   				jsii.String("deltaTables"),
//   			},
//   			WriteManifest: jsii.Boolean(false),
//   		},
//   	},
//   	DynamoDbTargets: []interface{}{
//   		&DynamoDBTargetProperty{
//   			Path: jsii.String("path"),
//   		},
//   	},
//   	HudiTargets: []interface{}{
//   		&HudiTargetProperty{
//   			ConnectionName: jsii.String("connectionName"),
//   			Exclusions: []*string{
//   				jsii.String("exclusions"),
//   			},
//   			MaximumTraversalDepth: jsii.Number(123),
//   			Paths: []*string{
//   				jsii.String("paths"),
//   			},
//   		},
//   	},
//   	IcebergTargets: []interface{}{
//   		&IcebergTargetProperty{
//   			ConnectionName: jsii.String("connectionName"),
//   			Exclusions: []*string{
//   				jsii.String("exclusions"),
//   			},
//   			MaximumTraversalDepth: jsii.Number(123),
//   			Paths: []*string{
//   				jsii.String("paths"),
//   			},
//   		},
//   	},
//   	JdbcTargets: []interface{}{
//   		&JdbcTargetProperty{
//   			ConnectionName: jsii.String("connectionName"),
//   			EnableAdditionalMetadata: []*string{
//   				jsii.String("enableAdditionalMetadata"),
//   			},
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-crawler-targets.html
//
type CfnCrawler_TargetsProperty struct {
	// Specifies AWS Glue Data Catalog targets.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-crawler-targets.html#cfn-glue-crawler-targets-catalogtargets
	//
	CatalogTargets interface{} `field:"optional" json:"catalogTargets" yaml:"catalogTargets"`
	// Specifies an array of Delta data store targets.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-crawler-targets.html#cfn-glue-crawler-targets-deltatargets
	//
	DeltaTargets interface{} `field:"optional" json:"deltaTargets" yaml:"deltaTargets"`
	// Specifies Amazon DynamoDB targets.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-crawler-targets.html#cfn-glue-crawler-targets-dynamodbtargets
	//
	DynamoDbTargets interface{} `field:"optional" json:"dynamoDbTargets" yaml:"dynamoDbTargets"`
	// Specifies Apache Hudi data store targets.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-crawler-targets.html#cfn-glue-crawler-targets-huditargets
	//
	HudiTargets interface{} `field:"optional" json:"hudiTargets" yaml:"hudiTargets"`
	// Specifies Apache Iceberg data store targets.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-crawler-targets.html#cfn-glue-crawler-targets-icebergtargets
	//
	IcebergTargets interface{} `field:"optional" json:"icebergTargets" yaml:"icebergTargets"`
	// Specifies JDBC targets.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-crawler-targets.html#cfn-glue-crawler-targets-jdbctargets
	//
	JdbcTargets interface{} `field:"optional" json:"jdbcTargets" yaml:"jdbcTargets"`
	// A list of Mongo DB targets.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-crawler-targets.html#cfn-glue-crawler-targets-mongodbtargets
	//
	MongoDbTargets interface{} `field:"optional" json:"mongoDbTargets" yaml:"mongoDbTargets"`
	// Specifies Amazon Simple Storage Service (Amazon S3) targets.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-crawler-targets.html#cfn-glue-crawler-targets-s3targets
	//
	S3Targets interface{} `field:"optional" json:"s3Targets" yaml:"s3Targets"`
}

