package awsiotwireless


// LoRaWANServiceProfile object.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loRaWANServiceProfileProperty := &LoRaWANServiceProfileProperty{
//   	AddGwMetadata: jsii.Boolean(false),
//   	ChannelMask: jsii.String("channelMask"),
//   	DevStatusReqFreq: jsii.Number(123),
//   	DlBucketSize: jsii.Number(123),
//   	DlRate: jsii.Number(123),
//   	DlRatePolicy: jsii.String("dlRatePolicy"),
//   	DrMax: jsii.Number(123),
//   	DrMin: jsii.Number(123),
//   	HrAllowed: jsii.Boolean(false),
//   	MinGwDiversity: jsii.Number(123),
//   	NwkGeoLoc: jsii.Boolean(false),
//   	PrAllowed: jsii.Boolean(false),
//   	RaAllowed: jsii.Boolean(false),
//   	ReportDevStatusBattery: jsii.Boolean(false),
//   	ReportDevStatusMargin: jsii.Boolean(false),
//   	TargetPer: jsii.Number(123),
//   	UlBucketSize: jsii.Number(123),
//   	UlRate: jsii.Number(123),
//   	UlRatePolicy: jsii.String("ulRatePolicy"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-serviceprofile-lorawanserviceprofile.html
//
type CfnServiceProfile_LoRaWANServiceProfileProperty struct {
	// The AddGWMetaData value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-serviceprofile-lorawanserviceprofile.html#cfn-iotwireless-serviceprofile-lorawanserviceprofile-addgwmetadata
	//
	AddGwMetadata interface{} `field:"optional" json:"addGwMetadata" yaml:"addGwMetadata"`
	// The ChannelMask value.
	//
	// This property is `ReadOnly` and can't be inputted for create. It's returned with `Fn::GetAtt`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-serviceprofile-lorawanserviceprofile.html#cfn-iotwireless-serviceprofile-lorawanserviceprofile-channelmask
	//
	ChannelMask *string `field:"optional" json:"channelMask" yaml:"channelMask"`
	// The DevStatusReqFreq value.
	//
	// This property is `ReadOnly` and can't be inputted for create. It's returned with `Fn::GetAtt`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-serviceprofile-lorawanserviceprofile.html#cfn-iotwireless-serviceprofile-lorawanserviceprofile-devstatusreqfreq
	//
	DevStatusReqFreq *float64 `field:"optional" json:"devStatusReqFreq" yaml:"devStatusReqFreq"`
	// The DLBucketSize value.
	//
	// This property is `ReadOnly` and can't be inputted for create. It's returned with `Fn::GetAtt`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-serviceprofile-lorawanserviceprofile.html#cfn-iotwireless-serviceprofile-lorawanserviceprofile-dlbucketsize
	//
	DlBucketSize *float64 `field:"optional" json:"dlBucketSize" yaml:"dlBucketSize"`
	// The DLRate value.
	//
	// This property is `ReadOnly` and can't be inputted for create. It's returned with `Fn::GetAtt`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-serviceprofile-lorawanserviceprofile.html#cfn-iotwireless-serviceprofile-lorawanserviceprofile-dlrate
	//
	DlRate *float64 `field:"optional" json:"dlRate" yaml:"dlRate"`
	// The DLRatePolicy value.
	//
	// This property is `ReadOnly` and can't be inputted for create. It's returned with `Fn::GetAtt`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-serviceprofile-lorawanserviceprofile.html#cfn-iotwireless-serviceprofile-lorawanserviceprofile-dlratepolicy
	//
	DlRatePolicy *string `field:"optional" json:"dlRatePolicy" yaml:"dlRatePolicy"`
	// The DRMax value.
	//
	// This property is `ReadOnly` and can't be inputted for create. It's returned with `Fn::GetAtt`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-serviceprofile-lorawanserviceprofile.html#cfn-iotwireless-serviceprofile-lorawanserviceprofile-drmax
	//
	DrMax *float64 `field:"optional" json:"drMax" yaml:"drMax"`
	// The DRMin value.
	//
	// This property is `ReadOnly` and can't be inputted for create. It's returned with `Fn::GetAtt`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-serviceprofile-lorawanserviceprofile.html#cfn-iotwireless-serviceprofile-lorawanserviceprofile-drmin
	//
	DrMin *float64 `field:"optional" json:"drMin" yaml:"drMin"`
	// The HRAllowed value that describes whether handover roaming is allowed.
	//
	// This property is `ReadOnly` and can't be inputted for create. It's returned with `Fn::GetAtt`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-serviceprofile-lorawanserviceprofile.html#cfn-iotwireless-serviceprofile-lorawanserviceprofile-hrallowed
	//
	HrAllowed interface{} `field:"optional" json:"hrAllowed" yaml:"hrAllowed"`
	// The MinGwDiversity value.
	//
	// This property is `ReadOnly` and can't be inputted for create. It's returned with `Fn::GetAtt`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-serviceprofile-lorawanserviceprofile.html#cfn-iotwireless-serviceprofile-lorawanserviceprofile-mingwdiversity
	//
	MinGwDiversity *float64 `field:"optional" json:"minGwDiversity" yaml:"minGwDiversity"`
	// The NwkGeoLoc value.
	//
	// This property is `ReadOnly` and can't be inputted for create. It's returned with `Fn::GetAtt`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-serviceprofile-lorawanserviceprofile.html#cfn-iotwireless-serviceprofile-lorawanserviceprofile-nwkgeoloc
	//
	NwkGeoLoc interface{} `field:"optional" json:"nwkGeoLoc" yaml:"nwkGeoLoc"`
	// The PRAllowed value that describes whether passive roaming is allowed.
	//
	// This property is `ReadOnly` and can't be inputted for create. It's returned with `Fn::GetAtt`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-serviceprofile-lorawanserviceprofile.html#cfn-iotwireless-serviceprofile-lorawanserviceprofile-prallowed
	//
	PrAllowed interface{} `field:"optional" json:"prAllowed" yaml:"prAllowed"`
	// The RAAllowed value that describes whether roaming activation is allowed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-serviceprofile-lorawanserviceprofile.html#cfn-iotwireless-serviceprofile-lorawanserviceprofile-raallowed
	//
	RaAllowed interface{} `field:"optional" json:"raAllowed" yaml:"raAllowed"`
	// The ReportDevStatusBattery value.
	//
	// This property is `ReadOnly` and can't be inputted for create. It's returned with `Fn::GetAtt`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-serviceprofile-lorawanserviceprofile.html#cfn-iotwireless-serviceprofile-lorawanserviceprofile-reportdevstatusbattery
	//
	ReportDevStatusBattery interface{} `field:"optional" json:"reportDevStatusBattery" yaml:"reportDevStatusBattery"`
	// The ReportDevStatusMargin value.
	//
	// This property is `ReadOnly` and can't be inputted for create. It's returned with `Fn::GetAtt`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-serviceprofile-lorawanserviceprofile.html#cfn-iotwireless-serviceprofile-lorawanserviceprofile-reportdevstatusmargin
	//
	ReportDevStatusMargin interface{} `field:"optional" json:"reportDevStatusMargin" yaml:"reportDevStatusMargin"`
	// The TargetPer value.
	//
	// This property is `ReadOnly` and can't be inputted for create. It's returned with `Fn::GetAtt`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-serviceprofile-lorawanserviceprofile.html#cfn-iotwireless-serviceprofile-lorawanserviceprofile-targetper
	//
	TargetPer *float64 `field:"optional" json:"targetPer" yaml:"targetPer"`
	// The UlBucketSize value.
	//
	// This property is `ReadOnly` and can't be inputted for create. It's returned with `Fn::GetAtt`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-serviceprofile-lorawanserviceprofile.html#cfn-iotwireless-serviceprofile-lorawanserviceprofile-ulbucketsize
	//
	UlBucketSize *float64 `field:"optional" json:"ulBucketSize" yaml:"ulBucketSize"`
	// The ULRate value.
	//
	// This property is `ReadOnly` and can't be inputted for create. It's returned with `Fn::GetAtt`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-serviceprofile-lorawanserviceprofile.html#cfn-iotwireless-serviceprofile-lorawanserviceprofile-ulrate
	//
	UlRate *float64 `field:"optional" json:"ulRate" yaml:"ulRate"`
	// The ULRatePolicy value.
	//
	// This property is `ReadOnly` and can't be inputted for create. It's returned with `Fn::GetAtt`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-serviceprofile-lorawanserviceprofile.html#cfn-iotwireless-serviceprofile-lorawanserviceprofile-ulratepolicy
	//
	UlRatePolicy *string `field:"optional" json:"ulRatePolicy" yaml:"ulRatePolicy"`
}

