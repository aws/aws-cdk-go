package mixinsawsappsync


// The `DynamoDBConfig` property type specifies the `AwsRegion` and `TableName` for an Amazon DynamoDB table in your account for an AWS AppSync data source.
//
// `DynamoDBConfig` is a property of the [AWS::AppSync::DataSource](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-datasource.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dynamoDBConfigProperty := &DynamoDBConfigProperty{
//   	AwsRegion: jsii.String("awsRegion"),
//   	DeltaSyncConfig: &DeltaSyncConfigProperty{
//   		BaseTableTtl: jsii.String("baseTableTtl"),
//   		DeltaSyncTableName: jsii.String("deltaSyncTableName"),
//   		DeltaSyncTableTtl: jsii.String("deltaSyncTableTtl"),
//   	},
//   	TableName: jsii.String("tableName"),
//   	UseCallerCredentials: jsii.Boolean(false),
//   	Versioned: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-datasource-dynamodbconfig.html
//
type CfnDataSourcePropsMixin_DynamoDBConfigProperty struct {
	// The AWS Region.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-datasource-dynamodbconfig.html#cfn-appsync-datasource-dynamodbconfig-awsregion
	//
	AwsRegion *string `field:"optional" json:"awsRegion" yaml:"awsRegion"`
	// The `DeltaSyncConfig` for a versioned datasource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-datasource-dynamodbconfig.html#cfn-appsync-datasource-dynamodbconfig-deltasyncconfig
	//
	DeltaSyncConfig interface{} `field:"optional" json:"deltaSyncConfig" yaml:"deltaSyncConfig"`
	// The table name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-datasource-dynamodbconfig.html#cfn-appsync-datasource-dynamodbconfig-tablename
	//
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
	// Set to `TRUE` to use AWS Identity and Access Management with this data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-datasource-dynamodbconfig.html#cfn-appsync-datasource-dynamodbconfig-usecallercredentials
	//
	UseCallerCredentials interface{} `field:"optional" json:"useCallerCredentials" yaml:"useCallerCredentials"`
	// Set to TRUE to use Conflict Detection and Resolution with this data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-datasource-dynamodbconfig.html#cfn-appsync-datasource-dynamodbconfig-versioned
	//
	Versioned interface{} `field:"optional" json:"versioned" yaml:"versioned"`
}

