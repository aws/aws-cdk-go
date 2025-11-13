package interfacesawsservicecatalog


// A reference to a LaunchRoleConstraint resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   launchRoleConstraintReference := &LaunchRoleConstraintReference{
//   	LaunchRoleConstraintId: jsii.String("launchRoleConstraintId"),
//   }
//
type LaunchRoleConstraintReference struct {
	// The Id of the LaunchRoleConstraint resource.
	LaunchRoleConstraintId *string `field:"required" json:"launchRoleConstraintId" yaml:"launchRoleConstraintId"`
}

