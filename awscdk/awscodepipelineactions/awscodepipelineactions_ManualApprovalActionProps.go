package awscodepipelineactions

import (
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awssns"
)

// Construction properties of the {@link ManualApprovalAction}.
//
// Example:
//   pipeline := codepipeline.NewPipeline(this, jsii.String("MyPipeline"))
//   approveStage := pipeline.addStage(&stageOptions{
//   	stageName: jsii.String("Approve"),
//   })
//   manualApprovalAction := codepipeline_actions.NewManualApprovalAction(&manualApprovalActionProps{
//   	actionName: jsii.String("Approve"),
//   })
//   approveStage.addAction(manualApprovalAction)
//
//   role := iam.role.fromRoleArn(this, jsii.String("Admin"), awscdk.Arn.format(&arnComponents{
//   	service: jsii.String("iam"),
//   	resource: jsii.String("role"),
//   	resourceName: jsii.String("Admin"),
//   }, this))
//   manualApprovalAction.grantManualApproval(role)
//
// Experimental.
type ManualApprovalActionProps struct {
	// The physical, human-readable name of the Action.
	//
	// Note that Action names must be unique within a single Stage.
	// Experimental.
	ActionName *string `field:"required" json:"actionName" yaml:"actionName"`
	// The runOrder property for this Action.
	//
	// RunOrder determines the relative order in which multiple Actions in the same Stage execute.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html
	//
	// Experimental.
	RunOrder *float64 `field:"optional" json:"runOrder" yaml:"runOrder"`
	// The name of the namespace to use for variables emitted by this action.
	// Experimental.
	VariablesNamespace *string `field:"optional" json:"variablesNamespace" yaml:"variablesNamespace"`
	// The Role in which context's this Action will be executing in.
	//
	// The Pipeline's Role will assume this Role
	// (the required permissions for that will be granted automatically)
	// right before executing this Action.
	// This Action will be passed into your {@link IAction.bind}
	// method in the {@link ActionBindOptions.role} property.
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// Any additional information that you want to include in the notification email message.
	// Experimental.
	AdditionalInformation *string `field:"optional" json:"additionalInformation" yaml:"additionalInformation"`
	// URL you want to provide to the reviewer as part of the approval request.
	// Experimental.
	ExternalEntityLink *string `field:"optional" json:"externalEntityLink" yaml:"externalEntityLink"`
	// Optional SNS topic to send notifications to when an approval is pending.
	// Experimental.
	NotificationTopic awssns.ITopic `field:"optional" json:"notificationTopic" yaml:"notificationTopic"`
	// A list of email addresses to subscribe to notifications when this Action is pending approval.
	//
	// If this has been provided, but not `notificationTopic`,
	// a new Topic will be created.
	// Experimental.
	NotifyEmails *[]*string `field:"optional" json:"notifyEmails" yaml:"notifyEmails"`
}

