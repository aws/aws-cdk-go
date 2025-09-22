package awsservicecatalog


// A reference to a StackSetConstraint resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   stackSetConstraintReference := &StackSetConstraintReference{
//   	StackSetConstraintId: jsii.String("stackSetConstraintId"),
//   }
//
type StackSetConstraintReference struct {
	// The Id of the StackSetConstraint resource.
	StackSetConstraintId *string `field:"required" json:"stackSetConstraintId" yaml:"stackSetConstraintId"`
}

