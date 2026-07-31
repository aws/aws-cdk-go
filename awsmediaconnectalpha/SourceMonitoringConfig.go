package awsmediaconnectalpha


// Source monitoring settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awsmediaconnectalpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sourceMonitoringConfig := &SourceMonitoringConfig{
//   	BlackFrames: &MonitoringMetric{
//   		State: mediaconnect_alpha.State_ENABLED,
//   		Threshold: cdk.Duration_Minutes(jsii.Number(30)),
//   	},
//   	ContentQualityAnalysisState: mediaconnect_alpha.State_ENABLED,
//   	FrozenFrames: &MonitoringMetric{
//   		State: mediaconnect_alpha.State_ENABLED,
//   		Threshold: cdk.Duration_*Minutes(jsii.Number(30)),
//   	},
//   	SilentAudio: &MonitoringMetric{
//   		State: mediaconnect_alpha.State_ENABLED,
//   		Threshold: cdk.Duration_*Minutes(jsii.Number(30)),
//   	},
//   	ThumbnailState: mediaconnect_alpha.State_ENABLED,
//   }
//
// Experimental.
type SourceMonitoringConfig struct {
	// Black-frames detection on the source.
	// Default: - black frames monitoring is not configured.
	//
	// Experimental.
	BlackFrames *MonitoringMetric `field:"optional" json:"blackFrames" yaml:"blackFrames"`
	// Indicates whether content quality analysis is enabled or disabled.
	// Default: - content quality analysis is disabled.
	//
	// Experimental.
	ContentQualityAnalysisState State `field:"optional" json:"contentQualityAnalysisState" yaml:"contentQualityAnalysisState"`
	// Frozen-frames detection on the source.
	// Default: - frozen frames monitoring is not configured.
	//
	// Experimental.
	FrozenFrames *MonitoringMetric `field:"optional" json:"frozenFrames" yaml:"frozenFrames"`
	// Silent-audio detection on the source.
	// Default: - silent audio monitoring is not configured.
	//
	// Experimental.
	SilentAudio *MonitoringMetric `field:"optional" json:"silentAudio" yaml:"silentAudio"`
	// The current state of the thumbnail monitoring.
	// Default: - thumbnail monitoring is disabled.
	//
	// Experimental.
	ThumbnailState State `field:"optional" json:"thumbnailState" yaml:"thumbnailState"`
}

