package awsb2bi


// A reference to a Transformer resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   transformerReference := &TransformerReference{
//   	TransformerArn: jsii.String("transformerArn"),
//   	TransformerId: jsii.String("transformerId"),
//   }
//
type TransformerReference struct {
	// The ARN of the Transformer resource.
	TransformerArn *string `field:"required" json:"transformerArn" yaml:"transformerArn"`
	// The TransformerId of the Transformer resource.
	TransformerId *string `field:"required" json:"transformerId" yaml:"transformerId"`
}

