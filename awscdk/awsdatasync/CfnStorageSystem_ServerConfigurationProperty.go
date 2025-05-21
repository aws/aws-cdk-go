package awsdatasync


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serverConfigurationProperty := &ServerConfigurationProperty{
//   	ServerHostname: jsii.String("serverHostname"),
//
//   	// the properties below are optional
//   	ServerPort: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-storagesystem-serverconfiguration.html
//
type CfnStorageSystem_ServerConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-storagesystem-serverconfiguration.html#cfn-datasync-storagesystem-serverconfiguration-serverhostname
	//
	ServerHostname *string `field:"required" json:"serverHostname" yaml:"serverHostname"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-storagesystem-serverconfiguration.html#cfn-datasync-storagesystem-serverconfiguration-serverport
	//
	ServerPort *float64 `field:"optional" json:"serverPort" yaml:"serverPort"`
}

