package awslogs


// A reference to a Transformer resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   transformerReference := &TransformerReference{
//   	LogGroupIdentifier: jsii.String("logGroupIdentifier"),
//   }
//
type TransformerReference struct {
	// The LogGroupIdentifier of the Transformer resource.
	LogGroupIdentifier *string `field:"required" json:"logGroupIdentifier" yaml:"logGroupIdentifier"`
}

