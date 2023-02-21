package awsiotwireless


// ABP device object for create APIs for v1.1.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   abpV11Property := &abpV11Property{
//   	devAddr: jsii.String("devAddr"),
//   	sessionKeys: &sessionKeysAbpV11Property{
//   		appSKey: jsii.String("appSKey"),
//   		fNwkSIntKey: jsii.String("fNwkSIntKey"),
//   		nwkSEncKey: jsii.String("nwkSEncKey"),
//   		sNwkSIntKey: jsii.String("sNwkSIntKey"),
//   	},
//   }
//
type CfnWirelessDevice_AbpV11Property struct {
	// The DevAddr value.
	DevAddr *string `field:"required" json:"devAddr" yaml:"devAddr"`
	// Session keys for ABP v1.1.
	SessionKeys interface{} `field:"required" json:"sessionKeys" yaml:"sessionKeys"`
}

