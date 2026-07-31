package awsmediaconnectalpha


// Base configuration for Media Stream.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import mediaconnect_alpha "github.com/aws/aws-cdk-go/awsmediaconnectalpha"
//
//   mediaStreamBase := &MediaStreamBase{
//   	MediaStreamId: jsii.Number(123),
//   	MediaStreamName: jsii.String("mediaStreamName"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   }
//
// Experimental.
type MediaStreamBase struct {
	// A unique identifier for the media stream.
	// Experimental.
	MediaStreamId *float64 `field:"required" json:"mediaStreamId" yaml:"mediaStreamId"`
	// A name that helps you distinguish one media stream from another.
	// Experimental.
	MediaStreamName *string `field:"required" json:"mediaStreamName" yaml:"mediaStreamName"`
	// A description that can help you quickly identify what your media stream is used for.
	// Default: - no description.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

