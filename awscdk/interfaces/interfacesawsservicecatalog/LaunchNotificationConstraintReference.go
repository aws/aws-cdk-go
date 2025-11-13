package interfacesawsservicecatalog


// A reference to a LaunchNotificationConstraint resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   launchNotificationConstraintReference := &LaunchNotificationConstraintReference{
//   	LaunchNotificationConstraintId: jsii.String("launchNotificationConstraintId"),
//   }
//
type LaunchNotificationConstraintReference struct {
	// The Id of the LaunchNotificationConstraint resource.
	LaunchNotificationConstraintId *string `field:"required" json:"launchNotificationConstraintId" yaml:"launchNotificationConstraintId"`
}

