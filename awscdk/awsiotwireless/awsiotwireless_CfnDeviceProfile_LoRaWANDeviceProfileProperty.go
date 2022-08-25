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
//   	macVersion: jsii.String("macVersion"),
//   	maxDutyCycle: jsii.Number(123),
//   	maxEirp: jsii.Number(123),
//   	pingSlotDr: jsii.Number(123),
//   	pingSlotFreq: jsii.Number(123),
//   	pingSlotPeriod: jsii.Number(123),
//   	regParamsRevision: jsii.String("regParamsRevision"),
//   	rfRegion: jsii.String("rfRegion"),
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
	// The Supports32BitFCnt value.
	Supports32BitFCnt interface{} `field:"optional" json:"supports32BitFCnt" yaml:"supports32BitFCnt"`
	// The SupportsClassB value.
	SupportsClassB interface{} `field:"optional" json:"supportsClassB" yaml:"supportsClassB"`
	// The SupportsClassC value.
	SupportsClassC interface{} `field:"optional" json:"supportsClassC" yaml:"supportsClassC"`
	// The SupportsJoin value.
	SupportsJoin interface{} `field:"optional" json:"supportsJoin" yaml:"supportsJoin"`
}

