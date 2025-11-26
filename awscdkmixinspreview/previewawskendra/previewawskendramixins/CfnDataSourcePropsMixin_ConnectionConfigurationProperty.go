package previewawskendramixins


// Provides the configuration information that's required to connect to a database.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   connectionConfigurationProperty := &ConnectionConfigurationProperty{
//   	DatabaseHost: jsii.String("databaseHost"),
//   	DatabaseName: jsii.String("databaseName"),
//   	DatabasePort: jsii.Number(123),
//   	SecretArn: jsii.String("secretArn"),
//   	TableName: jsii.String("tableName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-connectionconfiguration.html
//
type CfnDataSourcePropsMixin_ConnectionConfigurationProperty struct {
	// The name of the host for the database.
	//
	// Can be either a string (host.subdomain.domain.tld) or an IPv4 or IPv6 address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-connectionconfiguration.html#cfn-kendra-datasource-connectionconfiguration-databasehost
	//
	DatabaseHost *string `field:"optional" json:"databaseHost" yaml:"databaseHost"`
	// The name of the database containing the document data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-connectionconfiguration.html#cfn-kendra-datasource-connectionconfiguration-databasename
	//
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// The port that the database uses for connections.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-connectionconfiguration.html#cfn-kendra-datasource-connectionconfiguration-databaseport
	//
	DatabasePort *float64 `field:"optional" json:"databasePort" yaml:"databasePort"`
	// The Amazon Resource Name (ARN) of an AWS Secrets Manager secret that stores the credentials.
	//
	// The credentials should be a user-password pair. For more information, see [Using a Database Data Source](https://docs.aws.amazon.com/kendra/latest/dg/data-source-database.html) . For more information about AWS Secrets Manager , see [What Is AWS Secrets Manager](https://docs.aws.amazon.com/secretsmanager/latest/userguide/intro.html) in the *AWS Secrets Manager* user guide.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-connectionconfiguration.html#cfn-kendra-datasource-connectionconfiguration-secretarn
	//
	SecretArn *string `field:"optional" json:"secretArn" yaml:"secretArn"`
	// The name of the table that contains the document data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-connectionconfiguration.html#cfn-kendra-datasource-connectionconfiguration-tablename
	//
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
}

