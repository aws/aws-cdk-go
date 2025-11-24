package mixinsawsappsync


// Use the `RelationalDatabaseConfig` property type to specify `RelationalDatabaseConfig` for an AWS AppSync data source.
//
// `RelationalDatabaseConfig` is a property of the [AWS::AppSync::DataSource](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-datasource.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   relationalDatabaseConfigProperty := &RelationalDatabaseConfigProperty{
//   	RdsHttpEndpointConfig: &RdsHttpEndpointConfigProperty{
//   		AwsRegion: jsii.String("awsRegion"),
//   		AwsSecretStoreArn: jsii.String("awsSecretStoreArn"),
//   		DatabaseName: jsii.String("databaseName"),
//   		DbClusterIdentifier: jsii.String("dbClusterIdentifier"),
//   		Schema: jsii.String("schema"),
//   	},
//   	RelationalDatabaseSourceType: jsii.String("relationalDatabaseSourceType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-datasource-relationaldatabaseconfig.html
//
type CfnDataSourcePropsMixin_RelationalDatabaseConfigProperty struct {
	// Information about the Amazon RDS resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-datasource-relationaldatabaseconfig.html#cfn-appsync-datasource-relationaldatabaseconfig-rdshttpendpointconfig
	//
	RdsHttpEndpointConfig interface{} `field:"optional" json:"rdsHttpEndpointConfig" yaml:"rdsHttpEndpointConfig"`
	// The type of relational data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-datasource-relationaldatabaseconfig.html#cfn-appsync-datasource-relationaldatabaseconfig-relationaldatabasesourcetype
	//
	RelationalDatabaseSourceType *string `field:"optional" json:"relationalDatabaseSourceType" yaml:"relationalDatabaseSourceType"`
}

