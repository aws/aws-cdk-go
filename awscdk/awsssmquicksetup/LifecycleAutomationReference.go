package awsssmquicksetup


// A reference to a LifecycleAutomation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lifecycleAutomationReference := &LifecycleAutomationReference{
//   	AssociationId: jsii.String("associationId"),
//   }
//
type LifecycleAutomationReference struct {
	// The AssociationId of the LifecycleAutomation resource.
	AssociationId *string `field:"required" json:"associationId" yaml:"associationId"`
}

