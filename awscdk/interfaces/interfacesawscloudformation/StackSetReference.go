package interfacesawscloudformation


// A reference to a StackSet resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   stackSetReference := &StackSetReference{
//   	StackSetId: jsii.String("stackSetId"),
//   }
//
type StackSetReference struct {
	// The StackSetId of the StackSet resource.
	StackSetId *string `field:"required" json:"stackSetId" yaml:"stackSetId"`
}

