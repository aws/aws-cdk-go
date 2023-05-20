package awsiottwinmaker


// The component type error.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   errorProperty := &ErrorProperty{
//   	Code: jsii.String("code"),
//   	Message: jsii.String("message"),
//   }
//
type CfnComponentType_ErrorProperty struct {
	// The component type error code.
	Code *string `field:"optional" json:"code" yaml:"code"`
	// The component type error message.
	Message *string `field:"optional" json:"message" yaml:"message"`
}

