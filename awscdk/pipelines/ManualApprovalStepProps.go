package pipelines

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
)

// Construction properties for a `ManualApprovalStep`.
//
// Example:
//   var pipeline CodePipeline
//
//   preprod := NewMyApplicationStage(this, jsii.String("PreProd"))
//   prod := NewMyApplicationStage(this, jsii.String("Prod"))
//   topic := sns.NewTopic(this, jsii.String("ChangeApprovalTopic"))
//
//   pipeline.AddStage(preprod, &AddStageOpts{
//   	Post: []Step{
//   		pipelines.NewShellStep(jsii.String("Validate Endpoint"), &ShellStepProps{
//   			Commands: []*string{
//   				jsii.String("curl -Ssf https://my.webservice.com/"),
//   			},
//   		}),
//   	},
//   })
//   pipeline.AddStage(prod, &AddStageOpts{
//   	Pre: []Step{
//   		pipelines.NewManualApprovalStep(jsii.String("PromoteToProd"), &ManualApprovalStepProps{
//   			//All options below are optional
//   			Comment: jsii.String("Please validate changes"),
//   			ReviewUrl: jsii.String("https://my.webservice.com/"),
//   			NotificationTopic: topic,
//   		}),
//   	},
//   })
//
type ManualApprovalStepProps struct {
	// The comment to display with this manual approval.
	// Default: - No comment.
	//
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Optional SNS topic to send notifications to when an approval is pending.
	// Default: - No notifications.
	//
	NotificationTopic awssns.ITopic `field:"optional" json:"notificationTopic" yaml:"notificationTopic"`
	// The URL for review associated with this manual approval.
	// Default: - No URL.
	//
	ReviewUrl *string `field:"optional" json:"reviewUrl" yaml:"reviewUrl"`
}

