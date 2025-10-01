package awsmedialive


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var colorSpacePassthroughSettings interface{}
//   var rec601Settings interface{}
//   var rec709Settings interface{}
//
//   av1SettingsProperty := &Av1SettingsProperty{
//   	AfdSignaling: jsii.String("afdSignaling"),
//   	Bitrate: jsii.Number(123),
//   	BufSize: jsii.Number(123),
//   	ColorSpaceSettings: &Av1ColorSpaceSettingsProperty{
//   		ColorSpacePassthroughSettings: colorSpacePassthroughSettings,
//   		Hdr10Settings: &Hdr10SettingsProperty{
//   			MaxCll: jsii.Number(123),
//   			MaxFall: jsii.Number(123),
//   		},
//   		Rec601Settings: rec601Settings,
//   		Rec709Settings: rec709Settings,
//   	},
//   	FixedAfd: jsii.String("fixedAfd"),
//   	FramerateDenominator: jsii.Number(123),
//   	FramerateNumerator: jsii.Number(123),
//   	GopSize: jsii.Number(123),
//   	GopSizeUnits: jsii.String("gopSizeUnits"),
//   	Level: jsii.String("level"),
//   	LookAheadRateControl: jsii.String("lookAheadRateControl"),
//   	MaxBitrate: jsii.Number(123),
//   	MinIInterval: jsii.Number(123),
//   	ParDenominator: jsii.Number(123),
//   	ParNumerator: jsii.Number(123),
//   	QvbrQualityLevel: jsii.Number(123),
//   	RateControlMode: jsii.String("rateControlMode"),
//   	SceneChangeDetect: jsii.String("sceneChangeDetect"),
//   	TimecodeBurninSettings: &TimecodeBurninSettingsProperty{
//   		FontSize: jsii.String("fontSize"),
//   		Position: jsii.String("position"),
//   		Prefix: jsii.String("prefix"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-av1settings.html
//
type CfnChannel_Av1SettingsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-av1settings.html#cfn-medialive-channel-av1settings-afdsignaling
	//
	AfdSignaling *string `field:"optional" json:"afdSignaling" yaml:"afdSignaling"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-av1settings.html#cfn-medialive-channel-av1settings-bitrate
	//
	Bitrate *float64 `field:"optional" json:"bitrate" yaml:"bitrate"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-av1settings.html#cfn-medialive-channel-av1settings-bufsize
	//
	BufSize *float64 `field:"optional" json:"bufSize" yaml:"bufSize"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-av1settings.html#cfn-medialive-channel-av1settings-colorspacesettings
	//
	ColorSpaceSettings interface{} `field:"optional" json:"colorSpaceSettings" yaml:"colorSpaceSettings"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-av1settings.html#cfn-medialive-channel-av1settings-fixedafd
	//
	FixedAfd *string `field:"optional" json:"fixedAfd" yaml:"fixedAfd"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-av1settings.html#cfn-medialive-channel-av1settings-frameratedenominator
	//
	FramerateDenominator *float64 `field:"optional" json:"framerateDenominator" yaml:"framerateDenominator"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-av1settings.html#cfn-medialive-channel-av1settings-frameratenumerator
	//
	FramerateNumerator *float64 `field:"optional" json:"framerateNumerator" yaml:"framerateNumerator"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-av1settings.html#cfn-medialive-channel-av1settings-gopsize
	//
	GopSize *float64 `field:"optional" json:"gopSize" yaml:"gopSize"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-av1settings.html#cfn-medialive-channel-av1settings-gopsizeunits
	//
	GopSizeUnits *string `field:"optional" json:"gopSizeUnits" yaml:"gopSizeUnits"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-av1settings.html#cfn-medialive-channel-av1settings-level
	//
	Level *string `field:"optional" json:"level" yaml:"level"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-av1settings.html#cfn-medialive-channel-av1settings-lookaheadratecontrol
	//
	LookAheadRateControl *string `field:"optional" json:"lookAheadRateControl" yaml:"lookAheadRateControl"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-av1settings.html#cfn-medialive-channel-av1settings-maxbitrate
	//
	MaxBitrate *float64 `field:"optional" json:"maxBitrate" yaml:"maxBitrate"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-av1settings.html#cfn-medialive-channel-av1settings-miniinterval
	//
	MinIInterval *float64 `field:"optional" json:"minIInterval" yaml:"minIInterval"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-av1settings.html#cfn-medialive-channel-av1settings-pardenominator
	//
	ParDenominator *float64 `field:"optional" json:"parDenominator" yaml:"parDenominator"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-av1settings.html#cfn-medialive-channel-av1settings-parnumerator
	//
	ParNumerator *float64 `field:"optional" json:"parNumerator" yaml:"parNumerator"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-av1settings.html#cfn-medialive-channel-av1settings-qvbrqualitylevel
	//
	QvbrQualityLevel *float64 `field:"optional" json:"qvbrQualityLevel" yaml:"qvbrQualityLevel"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-av1settings.html#cfn-medialive-channel-av1settings-ratecontrolmode
	//
	RateControlMode *string `field:"optional" json:"rateControlMode" yaml:"rateControlMode"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-av1settings.html#cfn-medialive-channel-av1settings-scenechangedetect
	//
	SceneChangeDetect *string `field:"optional" json:"sceneChangeDetect" yaml:"sceneChangeDetect"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-av1settings.html#cfn-medialive-channel-av1settings-timecodeburninsettings
	//
	TimecodeBurninSettings interface{} `field:"optional" json:"timecodeBurninSettings" yaml:"timecodeBurninSettings"`
}

