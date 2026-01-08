package awsdms


// SybaseAseSettings property identifier.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sybaseAseSettingsProperty := &SybaseAseSettingsProperty{
//   	Port: jsii.Number(123),
//   	ServerName: jsii.String("serverName"),
//   	SslMode: jsii.String("sslMode"),
//
//   	// the properties below are optional
//   	CertificateArn: jsii.String("certificateArn"),
//   	DatabaseName: jsii.String("databaseName"),
//   	EncryptPassword: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-dataprovider-sybaseasesettings.html
//
type CfnDataProvider_SybaseAseSettingsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-dataprovider-sybaseasesettings.html#cfn-dms-dataprovider-sybaseasesettings-port
	//
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-dataprovider-sybaseasesettings.html#cfn-dms-dataprovider-sybaseasesettings-servername
	//
	ServerName *string `field:"required" json:"serverName" yaml:"serverName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-dataprovider-sybaseasesettings.html#cfn-dms-dataprovider-sybaseasesettings-sslmode
	//
	SslMode *string `field:"required" json:"sslMode" yaml:"sslMode"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-dataprovider-sybaseasesettings.html#cfn-dms-dataprovider-sybaseasesettings-certificatearn
	//
	CertificateArn *string `field:"optional" json:"certificateArn" yaml:"certificateArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-dataprovider-sybaseasesettings.html#cfn-dms-dataprovider-sybaseasesettings-databasename
	//
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-dataprovider-sybaseasesettings.html#cfn-dms-dataprovider-sybaseasesettings-encryptpassword
	//
	EncryptPassword interface{} `field:"optional" json:"encryptPassword" yaml:"encryptPassword"`
}

