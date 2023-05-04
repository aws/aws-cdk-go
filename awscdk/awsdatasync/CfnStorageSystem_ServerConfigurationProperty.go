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
type CfnStorageSystem_ServerConfigurationProperty struct {
	// `CfnStorageSystem.ServerConfigurationProperty.ServerHostname`.
	ServerHostname *string `field:"required" json:"serverHostname" yaml:"serverHostname"`
	// `CfnStorageSystem.ServerConfigurationProperty.ServerPort`.
	ServerPort *float64 `field:"optional" json:"serverPort" yaml:"serverPort"`
}

