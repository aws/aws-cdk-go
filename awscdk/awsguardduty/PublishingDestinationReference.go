package awsguardduty


// A reference to a PublishingDestination resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   publishingDestinationReference := &PublishingDestinationReference{
//   	DetectorId: jsii.String("detectorId"),
//   	PublishingDestinationId: jsii.String("publishingDestinationId"),
//   }
//
type PublishingDestinationReference struct {
	// The DetectorId of the PublishingDestination resource.
	DetectorId *string `field:"required" json:"detectorId" yaml:"detectorId"`
	// The Id of the PublishingDestination resource.
	PublishingDestinationId *string `field:"required" json:"publishingDestinationId" yaml:"publishingDestinationId"`
}

