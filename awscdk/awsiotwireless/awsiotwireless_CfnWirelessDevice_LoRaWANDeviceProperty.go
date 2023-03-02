package awsiotwireless


// LoRaWAN object for create functions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loRaWANDeviceProperty := &loRaWANDeviceProperty{
//   	abpV10X: &abpV10xProperty{
//   		devAddr: jsii.String("devAddr"),
//   		sessionKeys: &sessionKeysAbpV10xProperty{
//   			appSKey: jsii.String("appSKey"),
//   			nwkSKey: jsii.String("nwkSKey"),
//   		},
//   	},
//   	abpV11: &abpV11Property{
//   		devAddr: jsii.String("devAddr"),
//   		sessionKeys: &sessionKeysAbpV11Property{
//   			appSKey: jsii.String("appSKey"),
//   			fNwkSIntKey: jsii.String("fNwkSIntKey"),
//   			nwkSEncKey: jsii.String("nwkSEncKey"),
//   			sNwkSIntKey: jsii.String("sNwkSIntKey"),
//   		},
//   	},
//   	devEui: jsii.String("devEui"),
//   	deviceProfileId: jsii.String("deviceProfileId"),
//   	otaaV10X: &otaaV10xProperty{
//   		appEui: jsii.String("appEui"),
//   		appKey: jsii.String("appKey"),
//   	},
//   	otaaV11: &otaaV11Property{
//   		appKey: jsii.String("appKey"),
//   		joinEui: jsii.String("joinEui"),
//   		nwkKey: jsii.String("nwkKey"),
//   	},
//   	serviceProfileId: jsii.String("serviceProfileId"),
//   }
//
type CfnWirelessDevice_LoRaWANDeviceProperty struct {
	// LoRaWAN object for create APIs.
	AbpV10X interface{} `field:"optional" json:"abpV10X" yaml:"abpV10X"`
	// ABP device object for create APIs for v1.1.
	AbpV11 interface{} `field:"optional" json:"abpV11" yaml:"abpV11"`
	// The DevEUI value.
	DevEui *string `field:"optional" json:"devEui" yaml:"devEui"`
	// The ID of the device profile for the new wireless device.
	DeviceProfileId *string `field:"optional" json:"deviceProfileId" yaml:"deviceProfileId"`
	// OTAA device object for create APIs for v1.0.x.
	OtaaV10X interface{} `field:"optional" json:"otaaV10X" yaml:"otaaV10X"`
	// OTAA device object for v1.1 for create APIs.
	OtaaV11 interface{} `field:"optional" json:"otaaV11" yaml:"otaaV11"`
	// The ID of the service profile.
	ServiceProfileId *string `field:"optional" json:"serviceProfileId" yaml:"serviceProfileId"`
}

