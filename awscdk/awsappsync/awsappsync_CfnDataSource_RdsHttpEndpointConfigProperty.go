package awsappsync


// Use the `RdsHttpEndpointConfig` property type to specify the `RdsHttpEndpoint` for an AWS AppSync relational database.
//
// `RdsHttpEndpointConfig` is a property of the [AWS AppSync DataSource RelationalDatabaseConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-datasource-relationaldatabaseconfig.html) resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   rdsHttpEndpointConfigProperty := &rdsHttpEndpointConfigProperty{
//   	awsRegion: jsii.String("awsRegion"),
//   	awsSecretStoreArn: jsii.String("awsSecretStoreArn"),
//   	dbClusterIdentifier: jsii.String("dbClusterIdentifier"),
//
//   	// the properties below are optional
//   	databaseName: jsii.String("databaseName"),
//   	schema: jsii.String("schema"),
//   }
//
type CfnDataSource_RdsHttpEndpointConfigProperty struct {
	// AWS Region for RDS HTTP endpoint.
	AwsRegion *string `field:"required" json:"awsRegion" yaml:"awsRegion"`
	// The ARN for database credentials stored in AWS Secrets Manager .
	AwsSecretStoreArn *string `field:"required" json:"awsSecretStoreArn" yaml:"awsSecretStoreArn"`
	// Amazon RDS cluster Amazon Resource Name (ARN).
	DbClusterIdentifier *string `field:"required" json:"dbClusterIdentifier" yaml:"dbClusterIdentifier"`
	// Logical database name.
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// Logical schema name.
	Schema *string `field:"optional" json:"schema" yaml:"schema"`
}

