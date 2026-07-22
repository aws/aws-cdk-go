package awsmediaconnect


// The content quality analysis configuration for the router input.
//
// The content quality analysis feature only monitors the first video stream and the first audio stream it encounters within the router input source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   routerContentQualityAnalysisConfigurationProperty := &RouterContentQualityAnalysisConfigurationProperty{
//   	ContentLevel: &ContentQualityAnalysisFeatureConfigurationProperty{
//   		BlackFrames: &BlackFramesConfigurationProperty{
//   			State: jsii.String("state"),
//   			ThresholdSeconds: jsii.Number(123),
//   		},
//   		FrozenFrames: &FrozenFramesConfigurationProperty{
//   			State: jsii.String("state"),
//   			ThresholdSeconds: jsii.Number(123),
//   		},
//   		SilentAudio: &SilentAudioConfigurationProperty{
//   			State: jsii.String("state"),
//   			ThresholdSeconds: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-routercontentqualityanalysisconfiguration.html
//
type CfnRouterInputPropsMixin_RouterContentQualityAnalysisConfigurationProperty struct {
	// Configures the content quality analysis features for the router input.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-routercontentqualityanalysisconfiguration.html#cfn-mediaconnect-routerinput-routercontentqualityanalysisconfiguration-contentlevel
	//
	ContentLevel interface{} `field:"optional" json:"contentLevel" yaml:"contentLevel"`
}

