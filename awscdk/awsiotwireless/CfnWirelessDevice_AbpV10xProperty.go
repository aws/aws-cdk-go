package awsiotwireless


// ABP device object for LoRaWAN specification v1.0.x.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   abpV10xProperty := &AbpV10xProperty{
//   	DevAddr: jsii.String("devAddr"),
//   	SessionKeys: &SessionKeysAbpV10xProperty{
//   		AppSKey: jsii.String("appSKey"),
//   		NwkSKey: jsii.String("nwkSKey"),
//   	},
//   }
//
type CfnWirelessDevice_AbpV10xProperty struct {
	// The DevAddr value.
	DevAddr *string `field:"required" json:"devAddr" yaml:"devAddr"`
	// Session keys for ABP v1.0.x.
	SessionKeys interface{} `field:"required" json:"sessionKeys" yaml:"sessionKeys"`
}

