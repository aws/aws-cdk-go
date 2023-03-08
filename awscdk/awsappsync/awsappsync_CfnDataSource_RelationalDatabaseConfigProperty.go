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
//   relationalDatabaseConfigProperty := &relationalDatabaseConfigProperty{
//   	relationalDatabaseSourceType: jsii.String("relationalDatabaseSourceType"),
//
//   	// the properties below are optional
//   	rdsHttpEndpointConfig: &rdsHttpEndpointConfigProperty{
//   		awsRegion: jsii.String("awsRegion"),
//   		awsSecretStoreArn: jsii.String("awsSecretStoreArn"),
//   		dbClusterIdentifier: jsii.String("dbClusterIdentifier"),
//
//   		// the properties below are optional
//   		databaseName: jsii.String("databaseName"),
//   		schema: jsii.String("schema"),
//   	},
//   }
//
type CfnDataSource_RelationalDatabaseConfigProperty struct {
	// The type of relational data source.
	RelationalDatabaseSourceType *string `field:"required" json:"relationalDatabaseSourceType" yaml:"relationalDatabaseSourceType"`
	// Information about the Amazon RDS resource.
	RdsHttpEndpointConfig interface{} `field:"optional" json:"rdsHttpEndpointConfig" yaml:"rdsHttpEndpointConfig"`
}

