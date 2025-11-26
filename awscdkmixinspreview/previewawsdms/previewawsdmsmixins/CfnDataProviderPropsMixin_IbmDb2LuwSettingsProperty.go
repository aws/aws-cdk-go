package previewawsdmsmixins


// IbmDb2LuwSettings property identifier.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   ibmDb2LuwSettingsProperty := &IbmDb2LuwSettingsProperty{
//   	CertificateArn: jsii.String("certificateArn"),
//   	DatabaseName: jsii.String("databaseName"),
//   	Port: jsii.Number(123),
//   	ServerName: jsii.String("serverName"),
//   	SslMode: jsii.String("sslMode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-dataprovider-ibmdb2luwsettings.html
//
type CfnDataProviderPropsMixin_IbmDb2LuwSettingsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-dataprovider-ibmdb2luwsettings.html#cfn-dms-dataprovider-ibmdb2luwsettings-certificatearn
	//
	CertificateArn *string `field:"optional" json:"certificateArn" yaml:"certificateArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-dataprovider-ibmdb2luwsettings.html#cfn-dms-dataprovider-ibmdb2luwsettings-databasename
	//
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-dataprovider-ibmdb2luwsettings.html#cfn-dms-dataprovider-ibmdb2luwsettings-port
	//
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-dataprovider-ibmdb2luwsettings.html#cfn-dms-dataprovider-ibmdb2luwsettings-servername
	//
	ServerName *string `field:"optional" json:"serverName" yaml:"serverName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-dataprovider-ibmdb2luwsettings.html#cfn-dms-dataprovider-ibmdb2luwsettings-sslmode
	//
	SslMode *string `field:"optional" json:"sslMode" yaml:"sslMode"`
}

