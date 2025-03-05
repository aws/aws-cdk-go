package awsappsync


// Use the `RelationalDatabaseConfig` property type to specify `RelationalDatabaseConfig` for an AWS AppSync data source.
//
// `RelationalDatabaseConfig` is a property of the [AWS::AppSync::DataSource](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-datasource.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   relationalDatabaseConfigProperty := &RelationalDatabaseConfigProperty{
//   	RelationalDatabaseSourceType: jsii.String("relationalDatabaseSourceType"),
//
//   	// the properties below are optional
//   	RdsHttpEndpointConfig: &RdsHttpEndpointConfigProperty{
//   		AwsRegion: jsii.String("awsRegion"),
//   		AwsSecretStoreArn: jsii.String("awsSecretStoreArn"),
//   		DbClusterIdentifier: jsii.String("dbClusterIdentifier"),
//
//   		// the properties below are optional
//   		DatabaseName: jsii.String("databaseName"),
//   		Schema: jsii.String("schema"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-datasource-relationaldatabaseconfig.html
//
type CfnDataSource_RelationalDatabaseConfigProperty struct {
	// The type of relational data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-datasource-relationaldatabaseconfig.html#cfn-appsync-datasource-relationaldatabaseconfig-relationaldatabasesourcetype
	//
	RelationalDatabaseSourceType *string `field:"required" json:"relationalDatabaseSourceType" yaml:"relationalDatabaseSourceType"`
	// Information about the Amazon RDS resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-datasource-relationaldatabaseconfig.html#cfn-appsync-datasource-relationaldatabaseconfig-rdshttpendpointconfig
	//
	RdsHttpEndpointConfig interface{} `field:"optional" json:"rdsHttpEndpointConfig" yaml:"rdsHttpEndpointConfig"`
}

