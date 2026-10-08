package awsdms


// Provides information that defines a Microsoft SQL Server endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   microsoftSqlServerSettingsProperty := &MicrosoftSqlServerSettingsProperty{
//   	CertificateArn: jsii.String("certificateArn"),
//   	DatabaseName: jsii.String("databaseName"),
//   	Port: jsii.Number(123),
//   	S3AccessRoleArn: jsii.String("s3AccessRoleArn"),
//   	S3Path: jsii.String("s3Path"),
//   	ServerName: jsii.String("serverName"),
//   	SslMode: jsii.String("sslMode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-dataprovider-microsoftsqlserversettings.html
//
type CfnDataProviderPropsMixin_MicrosoftSqlServerSettingsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-dataprovider-microsoftsqlserversettings.html#cfn-dms-dataprovider-microsoftsqlserversettings-certificatearn
	//
	CertificateArn *string `field:"optional" json:"certificateArn" yaml:"certificateArn"`
	// Database name for the endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-dataprovider-microsoftsqlserversettings.html#cfn-dms-dataprovider-microsoftsqlserversettings-databasename
	//
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// Endpoint TCP port.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-dataprovider-microsoftsqlserversettings.html#cfn-dms-dataprovider-microsoftsqlserversettings-port
	//
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The ARN for the role the application uses to access its Amazon S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-dataprovider-microsoftsqlserversettings.html#cfn-dms-dataprovider-microsoftsqlserversettings-s3accessrolearn
	//
	S3AccessRoleArn *string `field:"optional" json:"s3AccessRoleArn" yaml:"s3AccessRoleArn"`
	// The path for the Amazon S3 bucket that the application uses for accessing the user-defined schema.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-dataprovider-microsoftsqlserversettings.html#cfn-dms-dataprovider-microsoftsqlserversettings-s3path
	//
	S3Path *string `field:"optional" json:"s3Path" yaml:"s3Path"`
	// Fully qualified domain name of the endpoint.
	//
	// For an Amazon RDS SQL Server instance, this is the output of [DescribeDBInstances](https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_DescribeDBInstances.html) , in the `[Endpoint](https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_Endpoint.html) .Address` field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-dataprovider-microsoftsqlserversettings.html#cfn-dms-dataprovider-microsoftsqlserversettings-servername
	//
	ServerName *string `field:"optional" json:"serverName" yaml:"serverName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-dataprovider-microsoftsqlserversettings.html#cfn-dms-dataprovider-microsoftsqlserversettings-sslmode
	//
	SslMode *string `field:"optional" json:"sslMode" yaml:"sslMode"`
}

