package awsappsync


// The `DynamoDBConfig` property type specifies the `AwsRegion` and `TableName` for an Amazon DynamoDB table in your account for an AWS AppSync data source.
//
// `DynamoDBConfig` is a property of the [AWS::AppSync::DataSource](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-datasource.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dynamoDBConfigProperty := &DynamoDBConfigProperty{
//   	AwsRegion: jsii.String("awsRegion"),
//   	TableName: jsii.String("tableName"),
//
//   	// the properties below are optional
//   	DeltaSyncConfig: &DeltaSyncConfigProperty{
//   		BaseTableTtl: jsii.String("baseTableTtl"),
//   		DeltaSyncTableName: jsii.String("deltaSyncTableName"),
//   		DeltaSyncTableTtl: jsii.String("deltaSyncTableTtl"),
//   	},
//   	UseCallerCredentials: jsii.Boolean(false),
//   	Versioned: jsii.Boolean(false),
//   }
//
type CfnDataSource_DynamoDBConfigProperty struct {
	// The AWS Region.
	AwsRegion *string `field:"required" json:"awsRegion" yaml:"awsRegion"`
	// The table name.
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
	// The `DeltaSyncConfig` for a versioned datasource.
	DeltaSyncConfig interface{} `field:"optional" json:"deltaSyncConfig" yaml:"deltaSyncConfig"`
	// Set to `TRUE` to use AWS Identity and Access Management with this data source.
	UseCallerCredentials interface{} `field:"optional" json:"useCallerCredentials" yaml:"useCallerCredentials"`
	// Set to TRUE to use Conflict Detection and Resolution with this data source.
	Versioned interface{} `field:"optional" json:"versioned" yaml:"versioned"`
}

