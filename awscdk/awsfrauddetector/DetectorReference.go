package awsfrauddetector


// A reference to a Detector resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   detectorReference := &DetectorReference{
//   	DetectorArn: jsii.String("detectorArn"),
//   }
//
type DetectorReference struct {
	// The Arn of the Detector resource.
	DetectorArn *string `field:"required" json:"detectorArn" yaml:"detectorArn"`
}

