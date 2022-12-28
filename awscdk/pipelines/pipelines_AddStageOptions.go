package pipelines

import (
	"github.com/aws/aws-cdk-go/awscdk/awssns"
)

// Options for adding an application stage to a pipeline.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var topic topic
//
//   addStageOptions := &addStageOptions{
//   	confirmBroadeningPermissions: jsii.Boolean(false),
//   	extraRunOrderSpace: jsii.Number(123),
//   	manualApprovals: jsii.Boolean(false),
//   	securityNotificationTopic: topic,
//   }
//
// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
type AddStageOptions struct {
	// Runs a `cdk diff --security-only --fail` to pause the pipeline if there are any security changes.
	//
	// If the stage is configured with `confirmBroadeningPermissions` enabled, you can use this
	// property to override the stage configuration. For example, Pipeline Stage
	// "Prod" has confirmBroadeningPermissions enabled, with applications "A", "B", "C". All three
	// applications will run a security check, but if we want to disable the one for "C",
	// we run `stage.addApplication(C, { confirmBroadeningPermissions: false })` to override the pipeline
	// stage behavior.
	//
	// Adds 1 to the run order space.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	ConfirmBroadeningPermissions *bool `field:"optional" json:"confirmBroadeningPermissions" yaml:"confirmBroadeningPermissions"`
	// Optional SNS topic to send notifications to when the security check registers changes within the application.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	SecurityNotificationTopic awssns.ITopic `field:"optional" json:"securityNotificationTopic" yaml:"securityNotificationTopic"`
	// Add room for extra actions.
	//
	// You can use this to make extra room in the runOrder sequence between the
	// changeset 'prepare' and 'execute' actions and insert your own actions there.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	ExtraRunOrderSpace *float64 `field:"optional" json:"extraRunOrderSpace" yaml:"extraRunOrderSpace"`
	// Add manual approvals before executing change sets.
	//
	// This gives humans the opportunity to confirm the change set looks alright
	// before deploying it.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	ManualApprovals *bool `field:"optional" json:"manualApprovals" yaml:"manualApprovals"`
}

