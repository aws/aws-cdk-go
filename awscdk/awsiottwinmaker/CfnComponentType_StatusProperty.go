package awsiottwinmaker


// The component type status.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   statusProperty := &StatusProperty{
//   	Error: &ErrorProperty{
//   		Code: jsii.String("code"),
//   		Message: jsii.String("message"),
//   	},
//   	State: jsii.String("state"),
//   }
//
type CfnComponentType_StatusProperty struct {
	// The component type error.
	Error interface{} `field:"optional" json:"error" yaml:"error"`
	// The component type status state.
	State *string `field:"optional" json:"state" yaml:"state"`
}

