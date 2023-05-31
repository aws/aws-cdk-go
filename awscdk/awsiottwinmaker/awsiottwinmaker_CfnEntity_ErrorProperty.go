package awsiottwinmaker


// The entity error.
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
type CfnEntity_ErrorProperty struct {
	// The entity error code.
	Code *string `field:"optional" json:"code" yaml:"code"`
	// The entity error message.
	Message *string `field:"optional" json:"message" yaml:"message"`
}

