package pipelines

import (
	"github.com/aws/aws-cdk-go/awscdk/awssns"
)

// Base options for a pipelines stage.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var topic topic
//
//   baseStageOptions := &baseStageOptions{
//   	confirmBroadeningPermissions: jsii.Boolean(false),
//   	securityNotificationTopic: topic,
//   }
//
// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
type BaseStageOptions struct {
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
}

