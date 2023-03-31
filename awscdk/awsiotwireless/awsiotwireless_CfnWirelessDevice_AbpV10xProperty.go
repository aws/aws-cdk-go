package awsiotwireless


// ABP device object for LoRaWAN specification v1.0.x.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   abpV10xProperty := &abpV10xProperty{
//   	devAddr: jsii.String("devAddr"),
//   	sessionKeys: &sessionKeysAbpV10xProperty{
//   		appSKey: jsii.String("appSKey"),
//   		nwkSKey: jsii.String("nwkSKey"),
//   	},
//   }
//
type CfnWirelessDevice_AbpV10xProperty struct {
	// The DevAddr value.
	DevAddr *string `field:"required" json:"devAddr" yaml:"devAddr"`
	// Session keys for ABP v1.0.x.
	SessionKeys interface{} `field:"required" json:"sessionKeys" yaml:"sessionKeys"`
}

