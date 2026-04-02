package awsmediapackagev2alpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Base properties common to all segment configurations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import mediapackagev2_alpha "github.com/aws/aws-cdk-go/awsmediapackagev2alpha"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   segmentPropsBase := &SegmentPropsBase{
//   	Duration: cdk.Duration_Minutes(jsii.Number(30)),
//   	IncludeIframeOnlyStreams: jsii.Boolean(false),
//   	Name: jsii.String("name"),
//   }
//
// Experimental.
type SegmentPropsBase struct {
	// Duration of each segment.
	// Default: Duration.seconds(6)
	//
	// Experimental.
	Duration awscdk.Duration `field:"optional" json:"duration" yaml:"duration"`
	// Whether to include I-frame-only streams.
	// Default: false.
	//
	// Experimental.
	IncludeIframeOnlyStreams *bool `field:"optional" json:"includeIframeOnlyStreams" yaml:"includeIframeOnlyStreams"`
	// Name of the segment.
	// Default: 'segment'.
	//
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

