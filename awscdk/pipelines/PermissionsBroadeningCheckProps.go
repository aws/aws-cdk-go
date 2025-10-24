package pipelines

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
)

// Properties for a `PermissionsBroadeningCheck`.
//
// Example:
//   var pipeline CodePipeline
//
//   stage := NewMyApplicationStage(this, jsii.String("MyApplication"))
//   pipeline.AddStage(stage, &AddStageOpts{
//   	Pre: []Step{
//   		pipelines.NewConfirmPermissionsBroadening(jsii.String("Check"), &PermissionsBroadeningCheckProps{
//   			Stage: *Stage,
//   		}),
//   	},
//   })
//
type PermissionsBroadeningCheckProps struct {
	// The CDK Stage object to check the stacks of.
	//
	// This should be the same Stage object you are passing to `addStage()`.
	Stage awscdk.Stage `field:"required" json:"stage" yaml:"stage"`
	// Topic to send notifications when a human needs to give manual confirmation.
	// Default: - no notification.
	//
	NotificationTopic awssns.ITopic `field:"optional" json:"notificationTopic" yaml:"notificationTopic"`
}

