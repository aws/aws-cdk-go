package awsdms


// MongoDbSettings property identifier.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mongoDbSettingsProperty := &MongoDbSettingsProperty{
//   	Port: jsii.Number(123),
//   	ServerName: jsii.String("serverName"),
//
//   	// the properties below are optional
//   	AuthMechanism: jsii.String("authMechanism"),
//   	AuthSource: jsii.String("authSource"),
//   	AuthType: jsii.String("authType"),
//   	CertificateArn: jsii.String("certificateArn"),
//   	DatabaseName: jsii.String("databaseName"),
//   	SslMode: jsii.String("sslMode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-dataprovider-mongodbsettings.html
//
type CfnDataProvider_MongoDbSettingsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-dataprovider-mongodbsettings.html#cfn-dms-dataprovider-mongodbsettings-port
	//
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-dataprovider-mongodbsettings.html#cfn-dms-dataprovider-mongodbsettings-servername
	//
	ServerName *string `field:"required" json:"serverName" yaml:"serverName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-dataprovider-mongodbsettings.html#cfn-dms-dataprovider-mongodbsettings-authmechanism
	//
	AuthMechanism *string `field:"optional" json:"authMechanism" yaml:"authMechanism"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-dataprovider-mongodbsettings.html#cfn-dms-dataprovider-mongodbsettings-authsource
	//
	AuthSource *string `field:"optional" json:"authSource" yaml:"authSource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-dataprovider-mongodbsettings.html#cfn-dms-dataprovider-mongodbsettings-authtype
	//
	AuthType *string `field:"optional" json:"authType" yaml:"authType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-dataprovider-mongodbsettings.html#cfn-dms-dataprovider-mongodbsettings-certificatearn
	//
	CertificateArn *string `field:"optional" json:"certificateArn" yaml:"certificateArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-dataprovider-mongodbsettings.html#cfn-dms-dataprovider-mongodbsettings-databasename
	//
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-dataprovider-mongodbsettings.html#cfn-dms-dataprovider-mongodbsettings-sslmode
	//
	SslMode *string `field:"optional" json:"sslMode" yaml:"sslMode"`
}

