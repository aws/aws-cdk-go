package awsiotwireless


// LoRaWAN device profile object.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loRaWANDeviceProfileProperty := &loRaWANDeviceProfileProperty{
//   	classBTimeout: jsii.Number(123),
//   	classCTimeout: jsii.Number(123),
//   	factoryPresetFreqsList: []interface{}{
//   		jsii.Number(123),
//   	},
//   	macVersion: jsii.String("macVersion"),
//   	maxDutyCycle: jsii.Number(123),
//   	maxEirp: jsii.Number(123),
//   	pingSlotDr: jsii.Number(123),
//   	pingSlotFreq: jsii.Number(123),
//   	pingSlotPeriod: jsii.Number(123),
//   	regParamsRevision: jsii.String("regParamsRevision"),
//   	rfRegion: jsii.String("rfRegion"),
//   	rxDataRate2: jsii.Number(123),
//   	rxDelay1: jsii.Number(123),
//   	rxDrOffset1: jsii.Number(123),
//   	rxFreq2: jsii.Number(123),
//   	supports32BitFCnt: jsii.Boolean(false),
//   	supportsClassB: jsii.Boolean(false),
//   	supportsClassC: jsii.Boolean(false),
//   	supportsJoin: jsii.Boolean(false),
//   }
//
type CfnDeviceProfile_LoRaWANDeviceProfileProperty struct {
	// The ClassBTimeout value.
	ClassBTimeout *float64 `field:"optional" json:"classBTimeout" yaml:"classBTimeout"`
	// The ClassCTimeout value.
	ClassCTimeout *float64 `field:"optional" json:"classCTimeout" yaml:"classCTimeout"`
	// `CfnDeviceProfile.LoRaWANDeviceProfileProperty.FactoryPresetFreqsList`.
	FactoryPresetFreqsList interface{} `field:"optional" json:"factoryPresetFreqsList" yaml:"factoryPresetFreqsList"`
	// The MAC version (such as OTAA 1.1 or OTAA 1.0.3) to use with this device profile.
	MacVersion *string `field:"optional" json:"macVersion" yaml:"macVersion"`
	// The MaxDutyCycle value.
	MaxDutyCycle *float64 `field:"optional" json:"maxDutyCycle" yaml:"maxDutyCycle"`
	// The MaxEIRP value.
	MaxEirp *float64 `field:"optional" json:"maxEirp" yaml:"maxEirp"`
	// The PingSlotDR value.
	PingSlotDr *float64 `field:"optional" json:"pingSlotDr" yaml:"pingSlotDr"`
	// The PingSlotFreq value.
	PingSlotFreq *float64 `field:"optional" json:"pingSlotFreq" yaml:"pingSlotFreq"`
	// The PingSlotPeriod value.
	PingSlotPeriod *float64 `field:"optional" json:"pingSlotPeriod" yaml:"pingSlotPeriod"`
	// The version of regional parameters.
	RegParamsRevision *string `field:"optional" json:"regParamsRevision" yaml:"regParamsRevision"`
	// The frequency band (RFRegion) value.
	RfRegion *string `field:"optional" json:"rfRegion" yaml:"rfRegion"`
	// `CfnDeviceProfile.LoRaWANDeviceProfileProperty.RxDataRate2`.
	RxDataRate2 *float64 `field:"optional" json:"rxDataRate2" yaml:"rxDataRate2"`
	// `CfnDeviceProfile.LoRaWANDeviceProfileProperty.RxDelay1`.
	RxDelay1 *float64 `field:"optional" json:"rxDelay1" yaml:"rxDelay1"`
	// `CfnDeviceProfile.LoRaWANDeviceProfileProperty.RxDrOffset1`.
	RxDrOffset1 *float64 `field:"optional" json:"rxDrOffset1" yaml:"rxDrOffset1"`
	// `CfnDeviceProfile.LoRaWANDeviceProfileProperty.RxFreq2`.
	RxFreq2 *float64 `field:"optional" json:"rxFreq2" yaml:"rxFreq2"`
	// The Supports32BitFCnt value.
	Supports32BitFCnt interface{} `field:"optional" json:"supports32BitFCnt" yaml:"supports32BitFCnt"`
	// The SupportsClassB value.
	SupportsClassB interface{} `field:"optional" json:"supportsClassB" yaml:"supportsClassB"`
	// The SupportsClassC value.
	SupportsClassC interface{} `field:"optional" json:"supportsClassC" yaml:"supportsClassC"`
	// The SupportsJoin value.
	SupportsJoin interface{} `field:"optional" json:"supportsJoin" yaml:"supportsJoin"`
}

