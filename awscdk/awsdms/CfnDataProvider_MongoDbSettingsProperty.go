package awsdms


// Provides information that defines a MongoDB endpoint.
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
	// The port value for the MongoDB source endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-dataprovider-mongodbsettings.html#cfn-dms-dataprovider-mongodbsettings-port
	//
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// The name of the server on the MongoDB source endpoint.
	//
	// For MongoDB Atlas, provide the server name for any of the servers in the replication set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-dataprovider-mongodbsettings.html#cfn-dms-dataprovider-mongodbsettings-servername
	//
	ServerName *string `field:"required" json:"serverName" yaml:"serverName"`
	// The authentication mechanism you use to access the MongoDB source endpoint.
	//
	// For the default value, in MongoDB version 2.x, `"default"` is `"mongodb_cr"` . For MongoDB version 3.x or later, `"default"` is `"scram_sha_1"` . This setting isn't used when `AuthType` is set to `"no"` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-dataprovider-mongodbsettings.html#cfn-dms-dataprovider-mongodbsettings-authmechanism
	//
	AuthMechanism *string `field:"optional" json:"authMechanism" yaml:"authMechanism"`
	// The MongoDB database name. This setting isn't used when `AuthType` is set to `"no"` .
	//
	// The default is `"admin"` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-dataprovider-mongodbsettings.html#cfn-dms-dataprovider-mongodbsettings-authsource
	//
	AuthSource *string `field:"optional" json:"authSource" yaml:"authSource"`
	// The authentication type you use to access the MongoDB source endpoint.
	//
	// When when set to `"no"` , user name and password parameters are not used and can be empty.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-dataprovider-mongodbsettings.html#cfn-dms-dataprovider-mongodbsettings-authtype
	//
	AuthType *string `field:"optional" json:"authType" yaml:"authType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-dataprovider-mongodbsettings.html#cfn-dms-dataprovider-mongodbsettings-certificatearn
	//
	CertificateArn *string `field:"optional" json:"certificateArn" yaml:"certificateArn"`
	// The database name on the MongoDB source endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-dataprovider-mongodbsettings.html#cfn-dms-dataprovider-mongodbsettings-databasename
	//
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-dataprovider-mongodbsettings.html#cfn-dms-dataprovider-mongodbsettings-sslmode
	//
	SslMode *string `field:"optional" json:"sslMode" yaml:"sslMode"`
}

