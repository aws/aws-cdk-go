package pipelines

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
)

// Properties for a `PermissionsBroadeningCheck`.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   var pipeline codePipeline
//
//   stage := NewMyApplicationStage(this, jsii.String("MyApplication"))
//   pipeline.AddStage(stage, &AddStageOpts{
//   	Pre: []step{
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
	NotificationTopic awssns.ITopic `field:"optional" json:"notificationTopic" yaml:"notificationTopic"`
}

