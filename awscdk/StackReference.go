package awscdk


// A reference to a Stack resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   stackReference := &StackReference{
//   	StackId: jsii.String("stackId"),
//   }
//
type StackReference struct {
	// The StackId of the Stack resource.
	StackId *string `field:"required" json:"stackId" yaml:"stackId"`
}

