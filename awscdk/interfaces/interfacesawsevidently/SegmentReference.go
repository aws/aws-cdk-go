package interfacesawsevidently


// A reference to a Segment resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   segmentReference := &SegmentReference{
//   	SegmentArn: jsii.String("segmentArn"),
//   }
//
type SegmentReference struct {
	// The Arn of the Segment resource.
	SegmentArn *string `field:"required" json:"segmentArn" yaml:"segmentArn"`
}

