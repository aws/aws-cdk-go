package awsdevopsagent


// Preferences that configure behavior of this AgentSpace.
//
// This container fully owns the AgentSpace preferences: the properties supplied here replace the stored preferences in their entirety, and omitting the Preferences container reverts all preferences to their service defaults.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   preferencesProperty := &PreferencesProperty{
//   	ElevatedActionsEnabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-agentspace-preferences.html
//
type CfnAgentSpace_PreferencesProperty struct {
	// Indicates whether elevated directed actions are permitted in this AgentSpace.
	//
	// Defaults to false when not set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-agentspace-preferences.html#cfn-devopsagent-agentspace-preferences-elevatedactionsenabled
	//
	ElevatedActionsEnabled interface{} `field:"optional" json:"elevatedActionsEnabled" yaml:"elevatedActionsEnabled"`
}

