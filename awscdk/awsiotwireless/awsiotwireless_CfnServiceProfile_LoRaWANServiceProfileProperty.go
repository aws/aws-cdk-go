package awsiotwireless


// LoRaWANServiceProfile object.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loRaWANServiceProfileProperty := &loRaWANServiceProfileProperty{
//   	addGwMetadata: jsii.Boolean(false),
//   	channelMask: jsii.String("channelMask"),
//   	devStatusReqFreq: jsii.Number(123),
//   	dlBucketSize: jsii.Number(123),
//   	dlRate: jsii.Number(123),
//   	dlRatePolicy: jsii.String("dlRatePolicy"),
//   	drMax: jsii.Number(123),
//   	drMin: jsii.Number(123),
//   	hrAllowed: jsii.Boolean(false),
//   	minGwDiversity: jsii.Number(123),
//   	nwkGeoLoc: jsii.Boolean(false),
//   	prAllowed: jsii.Boolean(false),
//   	raAllowed: jsii.Boolean(false),
//   	reportDevStatusBattery: jsii.Boolean(false),
//   	reportDevStatusMargin: jsii.Boolean(false),
//   	targetPer: jsii.Number(123),
//   	ulBucketSize: jsii.Number(123),
//   	ulRate: jsii.Number(123),
//   	ulRatePolicy: jsii.String("ulRatePolicy"),
//   }
//
type CfnServiceProfile_LoRaWANServiceProfileProperty struct {
	// The AddGWMetaData value.
	AddGwMetadata interface{} `field:"optional" json:"addGwMetadata" yaml:"addGwMetadata"`
	// The ChannelMask value.
	//
	// This property is `ReadOnly` and can't be inputted for create. It's returned with `Fn::GetAtt`
	ChannelMask *string `field:"optional" json:"channelMask" yaml:"channelMask"`
	// The DevStatusReqFreq value.
	//
	// This property is `ReadOnly` and can't be inputted for create. It's returned with `Fn::GetAtt`
	DevStatusReqFreq *float64 `field:"optional" json:"devStatusReqFreq" yaml:"devStatusReqFreq"`
	// The DLBucketSize value.
	//
	// This property is `ReadOnly` and can't be inputted for create. It's returned with `Fn::GetAtt`
	DlBucketSize *float64 `field:"optional" json:"dlBucketSize" yaml:"dlBucketSize"`
	// The DLRate value.
	//
	// This property is `ReadOnly` and can't be inputted for create. It's returned with `Fn::GetAtt`
	DlRate *float64 `field:"optional" json:"dlRate" yaml:"dlRate"`
	// The DLRatePolicy value.
	//
	// This property is `ReadOnly` and can't be inputted for create. It's returned with `Fn::GetAtt`
	DlRatePolicy *string `field:"optional" json:"dlRatePolicy" yaml:"dlRatePolicy"`
	// The DRMax value.
	//
	// This property is `ReadOnly` and can't be inputted for create. It's returned with `Fn::GetAtt`
	DrMax *float64 `field:"optional" json:"drMax" yaml:"drMax"`
	// The DRMin value.
	//
	// This property is `ReadOnly` and can't be inputted for create. It's returned with `Fn::GetAtt`
	DrMin *float64 `field:"optional" json:"drMin" yaml:"drMin"`
	// The HRAllowed value that describes whether handover roaming is allowed.
	//
	// This property is `ReadOnly` and can't be inputted for create. It's returned with `Fn::GetAtt`
	HrAllowed interface{} `field:"optional" json:"hrAllowed" yaml:"hrAllowed"`
	// The MinGwDiversity value.
	//
	// This property is `ReadOnly` and can't be inputted for create. It's returned with `Fn::GetAtt`
	MinGwDiversity *float64 `field:"optional" json:"minGwDiversity" yaml:"minGwDiversity"`
	// The NwkGeoLoc value.
	//
	// This property is `ReadOnly` and can't be inputted for create. It's returned with `Fn::GetAtt`
	NwkGeoLoc interface{} `field:"optional" json:"nwkGeoLoc" yaml:"nwkGeoLoc"`
	// The PRAllowed value that describes whether passive roaming is allowed.
	//
	// This property is `ReadOnly` and can't be inputted for create. It's returned with `Fn::GetAtt`
	PrAllowed interface{} `field:"optional" json:"prAllowed" yaml:"prAllowed"`
	// The RAAllowed value that describes whether roaming activation is allowed.
	RaAllowed interface{} `field:"optional" json:"raAllowed" yaml:"raAllowed"`
	// The ReportDevStatusBattery value.
	//
	// This property is `ReadOnly` and can't be inputted for create. It's returned with `Fn::GetAtt`
	ReportDevStatusBattery interface{} `field:"optional" json:"reportDevStatusBattery" yaml:"reportDevStatusBattery"`
	// The ReportDevStatusMargin value.
	//
	// This property is `ReadOnly` and can't be inputted for create. It's returned with `Fn::GetAtt`
	ReportDevStatusMargin interface{} `field:"optional" json:"reportDevStatusMargin" yaml:"reportDevStatusMargin"`
	// The TargetPer value.
	//
	// This property is `ReadOnly` and can't be inputted for create. It's returned with `Fn::GetAtt`
	TargetPer *float64 `field:"optional" json:"targetPer" yaml:"targetPer"`
	// The UlBucketSize value.
	//
	// This property is `ReadOnly` and can't be inputted for create. It's returned with `Fn::GetAtt`
	UlBucketSize *float64 `field:"optional" json:"ulBucketSize" yaml:"ulBucketSize"`
	// The ULRate value.
	//
	// This property is `ReadOnly` and can't be inputted for create. It's returned with `Fn::GetAtt`
	UlRate *float64 `field:"optional" json:"ulRate" yaml:"ulRate"`
	// The ULRatePolicy value.
	//
	// This property is `ReadOnly` and can't be inputted for create. It's returned with `Fn::GetAtt`
	UlRatePolicy *string `field:"optional" json:"ulRatePolicy" yaml:"ulRatePolicy"`
}

