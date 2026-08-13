package awskinesisfirehose


// Describes the metadata sent to the Http endpoint destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   httpAttribute := &HttpAttribute{
//   	Name: jsii.String("name"),
//   	Value: jsii.String("value"),
//   }
//
type HttpAttribute struct {
	// The name of the Http endpoint common attribute.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The value of the Http endpoint common attribute.
	Value *string `field:"required" json:"value" yaml:"value"`
}

