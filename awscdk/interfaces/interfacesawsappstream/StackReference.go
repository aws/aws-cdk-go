package interfacesawsappstream


// A reference to a Stack resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   stackReference := &StackReference{
//   	StackName: jsii.String("stackName"),
//   }
//
type StackReference struct {
	// The Name of the Stack resource.
	StackName *string `field:"required" json:"stackName" yaml:"stackName"`
}

