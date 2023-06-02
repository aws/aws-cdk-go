package awsdatasync


// The network settings that DataSync Discovery uses to connect with your on-premises storage system's management interface.
//
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
type CfnStorageSystem_ServerConfigurationProperty struct {
	// The domain name or IP address of your storage system's management interface.
	ServerHostname *string `field:"required" json:"serverHostname" yaml:"serverHostname"`
	// The network port for accessing the storage system's management interface.
	ServerPort *float64 `field:"optional" json:"serverPort" yaml:"serverPort"`
}

