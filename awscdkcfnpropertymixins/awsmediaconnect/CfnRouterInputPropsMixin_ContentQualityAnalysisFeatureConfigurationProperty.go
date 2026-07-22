package awsmediaconnect


// Configures the content quality analysis features for the router input.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   contentQualityAnalysisFeatureConfigurationProperty := &ContentQualityAnalysisFeatureConfigurationProperty{
//   	BlackFrames: &BlackFramesConfigurationProperty{
//   		State: jsii.String("state"),
//   		ThresholdSeconds: jsii.Number(123),
//   	},
//   	FrozenFrames: &FrozenFramesConfigurationProperty{
//   		State: jsii.String("state"),
//   		ThresholdSeconds: jsii.Number(123),
//   	},
//   	SilentAudio: &SilentAudioConfigurationProperty{
//   		State: jsii.String("state"),
//   		ThresholdSeconds: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-contentqualityanalysisfeatureconfiguration.html
//
type CfnRouterInputPropsMixin_ContentQualityAnalysisFeatureConfigurationProperty struct {
	// Detects black frames in the router input's source content and reports them through a CloudWatch metric, an EventBridge event, and a router input message.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-contentqualityanalysisfeatureconfiguration.html#cfn-mediaconnect-routerinput-contentqualityanalysisfeatureconfiguration-blackframes
	//
	BlackFrames interface{} `field:"optional" json:"blackFrames" yaml:"blackFrames"`
	// Detects frozen video frames in the router input's source content and reports them through a CloudWatch metric, an EventBridge event, and a router input message.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-contentqualityanalysisfeatureconfiguration.html#cfn-mediaconnect-routerinput-contentqualityanalysisfeatureconfiguration-frozenframes
	//
	FrozenFrames interface{} `field:"optional" json:"frozenFrames" yaml:"frozenFrames"`
	// Detects silent audio in the router input's source content and reports it through a CloudWatch metric, an EventBridge event, and a router input message.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-contentqualityanalysisfeatureconfiguration.html#cfn-mediaconnect-routerinput-contentqualityanalysisfeatureconfiguration-silentaudio
	//
	SilentAudio interface{} `field:"optional" json:"silentAudio" yaml:"silentAudio"`
}

