package awscdkpipesalpha


// The inputTemplate that is used to transform the input event payload with unquoted variables.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import pipes_alpha "github.com/aws/aws-cdk-go/awscdkpipesalpha"
//
//   inputTransformationConfig := &InputTransformationConfig{
//   	InputTemplate: jsii.String("inputTemplate"),
//   }
//
// Experimental.
type InputTransformationConfig struct {
	// The inputTemplate that is used to transform the input event payload.
	// Experimental.
	InputTemplate *string `field:"required" json:"inputTemplate" yaml:"inputTemplate"`
}

