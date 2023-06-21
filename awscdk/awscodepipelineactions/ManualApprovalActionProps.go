package awscodepipelineactions

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
)

// Construction properties of the `ManualApprovalAction`.
//
// Example:
//   pipeline := codepipeline.NewPipeline(this, jsii.String("MyPipeline"))
//   approveStage := pipeline.AddStage(&StageOptions{
//   	StageName: jsii.String("Approve"),
//   })
//   manualApprovalAction := codepipeline_actions.NewManualApprovalAction(&ManualApprovalActionProps{
//   	ActionName: jsii.String("Approve"),
//   })
//   approveStage.AddAction(manualApprovalAction)
//
//   role := iam.Role_FromRoleArn(this, jsii.String("Admin"), awscdk.Arn_Format(&ArnComponents{
//   	Service: jsii.String("iam"),
//   	Resource: jsii.String("role"),
//   	ResourceName: jsii.String("Admin"),
//   }, this))
//   manualApprovalAction.GrantManualApproval(role)
//
type ManualApprovalActionProps struct {
	// The physical, human-readable name of the Action.
	//
	// Note that Action names must be unique within a single Stage.
	ActionName *string `field:"required" json:"actionName" yaml:"actionName"`
	// The runOrder property for this Action.
	//
	// RunOrder determines the relative order in which multiple Actions in the same Stage execute.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html
	//
	RunOrder *float64 `field:"optional" json:"runOrder" yaml:"runOrder"`
	// The name of the namespace to use for variables emitted by this action.
	VariablesNamespace *string `field:"optional" json:"variablesNamespace" yaml:"variablesNamespace"`
	// The Role in which context's this Action will be executing in.
	//
	// The Pipeline's Role will assume this Role
	// (the required permissions for that will be granted automatically)
	// right before executing this Action.
	// This Action will be passed into your `IAction.bind`
	// method in the `ActionBindOptions.role` property.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// Any additional information that you want to include in the notification email message.
	AdditionalInformation *string `field:"optional" json:"additionalInformation" yaml:"additionalInformation"`
	// URL you want to provide to the reviewer as part of the approval request.
	ExternalEntityLink *string `field:"optional" json:"externalEntityLink" yaml:"externalEntityLink"`
	// Optional SNS topic to send notifications to when an approval is pending.
	NotificationTopic awssns.ITopic `field:"optional" json:"notificationTopic" yaml:"notificationTopic"`
	// A list of email addresses to subscribe to notifications when this Action is pending approval.
	//
	// If this has been provided, but not `notificationTopic`,
	// a new Topic will be created.
	NotifyEmails *[]*string `field:"optional" json:"notifyEmails" yaml:"notifyEmails"`
}

