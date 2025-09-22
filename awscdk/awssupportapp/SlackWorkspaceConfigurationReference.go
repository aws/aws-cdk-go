package awssupportapp


// A reference to a SlackWorkspaceConfiguration resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   slackWorkspaceConfigurationReference := &SlackWorkspaceConfigurationReference{
//   	TeamId: jsii.String("teamId"),
//   }
//
type SlackWorkspaceConfigurationReference struct {
	// The TeamId of the SlackWorkspaceConfiguration resource.
	TeamId *string `field:"required" json:"teamId" yaml:"teamId"`
}

