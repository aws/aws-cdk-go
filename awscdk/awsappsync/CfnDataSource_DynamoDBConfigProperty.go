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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-datasource-dynamodbconfig.html
//
type CfnDataSource_DynamoDBConfigProperty struct {
	// The AWS Region.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-datasource-dynamodbconfig.html#cfn-appsync-datasource-dynamodbconfig-awsregion
	//
	AwsRegion *string `field:"required" json:"awsRegion" yaml:"awsRegion"`
	// The table name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-datasource-dynamodbconfig.html#cfn-appsync-datasource-dynamodbconfig-tablename
	//
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
	// The `DeltaSyncConfig` for a versioned datasource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-datasource-dynamodbconfig.html#cfn-appsync-datasource-dynamodbconfig-deltasyncconfig
	//
	DeltaSyncConfig interface{} `field:"optional" json:"deltaSyncConfig" yaml:"deltaSyncConfig"`
	// Set to `TRUE` to use AWS Identity and Access Management with this data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-datasource-dynamodbconfig.html#cfn-appsync-datasource-dynamodbconfig-usecallercredentials
	//
	UseCallerCredentials interface{} `field:"optional" json:"useCallerCredentials" yaml:"useCallerCredentials"`
	// Set to TRUE to use Conflict Detection and Resolution with this data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-datasource-dynamodbconfig.html#cfn-appsync-datasource-dynamodbconfig-versioned
	//
	Versioned interface{} `field:"optional" json:"versioned" yaml:"versioned"`
}

