package previewawscodecommitevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for aws.codecommit@CodeCommitApprovalRuleTemplateChange event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   codeCommitApprovalRuleTemplateChangeProps := &CodeCommitApprovalRuleTemplateChangeProps{
//   	ApprovalRuleTemplateContentSha256: []*string{
//   		jsii.String("approvalRuleTemplateContentSha256"),
//   	},
//   	ApprovalRuleTemplateId: []*string{
//   		jsii.String("approvalRuleTemplateId"),
//   	},
//   	ApprovalRuleTemplateName: []*string{
//   		jsii.String("approvalRuleTemplateName"),
//   	},
//   	CallerUserArn: []*string{
//   		jsii.String("callerUserArn"),
//   	},
//   	CreationDate: []*string{
//   		jsii.String("creationDate"),
//   	},
//   	Event: []*string{
//   		jsii.String("event"),
//   	},
//   	EventMetadata: &AWSEventMetadataProps{
//   		Region: []*string{
//   			jsii.String("region"),
//   		},
//   		Resources: []*string{
//   			jsii.String("resources"),
//   		},
//   		Version: []*string{
//   			jsii.String("version"),
//   		},
//   	},
//   	LastModifiedDate: []*string{
//   		jsii.String("lastModifiedDate"),
//   	},
//   	NotificationBody: []*string{
//   		jsii.String("notificationBody"),
//   	},
//   	Repositories: map[string]*string{
//   		"repositoriesKey": jsii.String("repositories"),
//   	},
//   }
//
// Experimental.
type CodeCommitApprovalRuleTemplateChange_CodeCommitApprovalRuleTemplateChangeProps struct {
	// approvalRuleTemplateContentSha256 property.
	//
	// Specify an array of string values to match this event if the actual value of approvalRuleTemplateContentSha256 is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ApprovalRuleTemplateContentSha256 *[]*string `field:"optional" json:"approvalRuleTemplateContentSha256" yaml:"approvalRuleTemplateContentSha256"`
	// approvalRuleTemplateId property.
	//
	// Specify an array of string values to match this event if the actual value of approvalRuleTemplateId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ApprovalRuleTemplateId *[]*string `field:"optional" json:"approvalRuleTemplateId" yaml:"approvalRuleTemplateId"`
	// approvalRuleTemplateName property.
	//
	// Specify an array of string values to match this event if the actual value of approvalRuleTemplateName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ApprovalRuleTemplateName *[]*string `field:"optional" json:"approvalRuleTemplateName" yaml:"approvalRuleTemplateName"`
	// callerUserArn property.
	//
	// Specify an array of string values to match this event if the actual value of callerUserArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CallerUserArn *[]*string `field:"optional" json:"callerUserArn" yaml:"callerUserArn"`
	// creationDate property.
	//
	// Specify an array of string values to match this event if the actual value of creationDate is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CreationDate *[]*string `field:"optional" json:"creationDate" yaml:"creationDate"`
	// event property.
	//
	// Specify an array of string values to match this event if the actual value of event is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Event *[]*string `field:"optional" json:"event" yaml:"event"`
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// lastModifiedDate property.
	//
	// Specify an array of string values to match this event if the actual value of lastModifiedDate is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	LastModifiedDate *[]*string `field:"optional" json:"lastModifiedDate" yaml:"lastModifiedDate"`
	// notificationBody property.
	//
	// Specify an array of string values to match this event if the actual value of notificationBody is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	NotificationBody *[]*string `field:"optional" json:"notificationBody" yaml:"notificationBody"`
	// repositories property.
	//
	// Specify an array of string values to match this event if the actual value of repositories is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Repositories *map[string]*string `field:"optional" json:"repositories" yaml:"repositories"`
}

