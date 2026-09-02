package awsmedialivealpha


// Options for a smart subtitle caption source (AI-generated subtitles via Elemental Inference).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   var captionSynchronizationMode CaptionSynchronizationMode
//
//   smartSubtitleSourceOptions := &SmartSubtitleSourceOptions{
//   	CaptionSynchronizationMode: captionSynchronizationMode,
//   	InferenceFeedOutput: jsii.String("inferenceFeedOutput"),
//   }
//
// Experimental.
type SmartSubtitleSourceOptions struct {
	// Controls whether MediaLive delays video to synchronize captions with audio and video output.
	// Default: - service default.
	//
	// Experimental.
	CaptionSynchronizationMode CaptionSynchronizationMode `field:"optional" json:"captionSynchronizationMode" yaml:"captionSynchronizationMode"`
	// The name of the Elemental Inference feed output that provides the subtitles.
	// Default: - service default.
	//
	// Experimental.
	InferenceFeedOutput *string `field:"optional" json:"inferenceFeedOutput" yaml:"inferenceFeedOutput"`
}

