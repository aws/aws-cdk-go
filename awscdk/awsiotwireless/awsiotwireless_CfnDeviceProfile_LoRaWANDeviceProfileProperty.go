package awsiotwireless


// LoRaWAN device profile object.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loRaWANDeviceProfileProperty := &LoRaWANDeviceProfileProperty{
//   	ClassBTimeout: jsii.Number(123),
//   	ClassCTimeout: jsii.Number(123),
//   	FactoryPresetFreqsList: []interface{}{
//   		jsii.Number(123),
//   	},
//   	MacVersion: jsii.String("macVersion"),
//   	MaxDutyCycle: jsii.Number(123),
//   	MaxEirp: jsii.Number(123),
//   	PingSlotDr: jsii.Number(123),
//   	PingSlotFreq: jsii.Number(123),
//   	PingSlotPeriod: jsii.Number(123),
//   	RegParamsRevision: jsii.String("regParamsRevision"),
//   	RfRegion: jsii.String("rfRegion"),
//   	RxDataRate2: jsii.Number(123),
//   	RxDelay1: jsii.Number(123),
//   	RxDrOffset1: jsii.Number(123),
//   	RxFreq2: jsii.Number(123),
//   	Supports32BitFCnt: jsii.Boolean(false),
//   	SupportsClassB: jsii.Boolean(false),
//   	SupportsClassC: jsii.Boolean(false),
//   	SupportsJoin: jsii.Boolean(false),
//   }
//
type CfnDeviceProfile_LoRaWANDeviceProfileProperty struct {
	// The ClassBTimeout value.
	ClassBTimeout *float64 `field:"optional" json:"classBTimeout" yaml:"classBTimeout"`
	// The ClassCTimeout value.
	ClassCTimeout *float64 `field:"optional" json:"classCTimeout" yaml:"classCTimeout"`
	// The list of values that make up the FactoryPresetFreqs value.
	//
	// Valid range of values include a minimum value of 1000000 and a maximum value of 16700000.
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
	// The RXDataRate2 value.
	RxDataRate2 *float64 `field:"optional" json:"rxDataRate2" yaml:"rxDataRate2"`
	// The RXDelay1 value.
	RxDelay1 *float64 `field:"optional" json:"rxDelay1" yaml:"rxDelay1"`
	// The RXDROffset1 value.
	RxDrOffset1 *float64 `field:"optional" json:"rxDrOffset1" yaml:"rxDrOffset1"`
	// The RXFreq2 value.
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

