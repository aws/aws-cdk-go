package awsdms


// IbmDb2zOsSettings property identifier.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ibmDb2zOsSettingsProperty := &IbmDb2zOsSettingsProperty{
//   	DatabaseName: jsii.String("databaseName"),
//   	Port: jsii.Number(123),
//   	ServerName: jsii.String("serverName"),
//   	SslMode: jsii.String("sslMode"),
//
//   	// the properties below are optional
//   	CertificateArn: jsii.String("certificateArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-dataprovider-ibmdb2zossettings.html
//
type CfnDataProvider_IbmDb2zOsSettingsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-dataprovider-ibmdb2zossettings.html#cfn-dms-dataprovider-ibmdb2zossettings-databasename
	//
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-dataprovider-ibmdb2zossettings.html#cfn-dms-dataprovider-ibmdb2zossettings-port
	//
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-dataprovider-ibmdb2zossettings.html#cfn-dms-dataprovider-ibmdb2zossettings-servername
	//
	ServerName *string `field:"required" json:"serverName" yaml:"serverName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-dataprovider-ibmdb2zossettings.html#cfn-dms-dataprovider-ibmdb2zossettings-sslmode
	//
	SslMode *string `field:"required" json:"sslMode" yaml:"sslMode"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-dataprovider-ibmdb2zossettings.html#cfn-dms-dataprovider-ibmdb2zossettings-certificatearn
	//
	CertificateArn *string `field:"optional" json:"certificateArn" yaml:"certificateArn"`
}

