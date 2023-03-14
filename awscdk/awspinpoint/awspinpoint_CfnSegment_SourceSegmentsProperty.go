package awspinpoint


// Specifies the base segment to build the segment on.
//
// A base segment, also called a *source segment* , defines the initial population of endpoints for a segment. When you add dimensions to the segment, Amazon Pinpoint filters the base segment by using the dimensions that you specify.
//
// You can specify more than one dimensional segment or only one imported segment. If you specify an imported segment, the segment size estimate that displays on the Amazon Pinpoint console indicates the size of the imported segment without any filters applied to it.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sourceSegmentsProperty := &sourceSegmentsProperty{
//   	id: jsii.String("id"),
//
//   	// the properties below are optional
//   	version: jsii.Number(123),
//   }
//
type CfnSegment_SourceSegmentsProperty struct {
	// The unique identifier for the source segment.
	Id *string `field:"required" json:"id" yaml:"id"`
	// The version number of the source segment.
	Version *float64 `field:"optional" json:"version" yaml:"version"`
}

