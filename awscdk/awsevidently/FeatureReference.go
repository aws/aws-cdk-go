package awsevidently


// A reference to a Feature resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   featureReference := &FeatureReference{
//   	FeatureArn: jsii.String("featureArn"),
//   }
//
type FeatureReference struct {
	// The Arn of the Feature resource.
	FeatureArn *string `field:"required" json:"featureArn" yaml:"featureArn"`
}

