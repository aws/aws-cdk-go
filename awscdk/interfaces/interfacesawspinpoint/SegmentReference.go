package interfacesawspinpoint


// A reference to a Segment resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   segmentReference := &SegmentReference{
//   	SegmentArn: jsii.String("segmentArn"),
//   	SegmentId: jsii.String("segmentId"),
//   }
//
type SegmentReference struct {
	// The ARN of the Segment resource.
	SegmentArn *string `field:"required" json:"segmentArn" yaml:"segmentArn"`
	// The SegmentId of the Segment resource.
	SegmentId *string `field:"required" json:"segmentId" yaml:"segmentId"`
}

