package awsservicecatalog


// A reference to a LaunchTemplateConstraint resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   launchTemplateConstraintReference := &LaunchTemplateConstraintReference{
//   	LaunchTemplateConstraintId: jsii.String("launchTemplateConstraintId"),
//   }
//
type LaunchTemplateConstraintReference struct {
	// The Id of the LaunchTemplateConstraint resource.
	LaunchTemplateConstraintId *string `field:"required" json:"launchTemplateConstraintId" yaml:"launchTemplateConstraintId"`
}

