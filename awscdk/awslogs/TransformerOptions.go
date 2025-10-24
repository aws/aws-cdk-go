package awslogs


// Properties for Transformer created from LogGroup.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var processor IProcessor
//
//   transformerOptions := &TransformerOptions{
//   	TransformerConfig: []IProcessor{
//   		processor,
//   	},
//   	TransformerName: jsii.String("transformerName"),
//   }
//
type TransformerOptions struct {
	// List of processors in a transformer.
	TransformerConfig *[]IProcessor `field:"required" json:"transformerConfig" yaml:"transformerConfig"`
	// Name of the transformer.
	TransformerName *string `field:"required" json:"transformerName" yaml:"transformerName"`
}

