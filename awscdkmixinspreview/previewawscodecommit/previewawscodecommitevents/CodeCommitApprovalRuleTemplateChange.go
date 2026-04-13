package previewawscodecommitevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.codecommit@CodeCommitApprovalRuleTemplateChange.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   codeCommitApprovalRuleTemplateChange := awscdkmixinspreview.Events.NewCodeCommitApprovalRuleTemplateChange()
//
// Experimental.
type CodeCommitApprovalRuleTemplateChange interface {
}

// The jsii proxy struct for CodeCommitApprovalRuleTemplateChange
type jsiiProxy_CodeCommitApprovalRuleTemplateChange struct {
	_ byte // padding
}

// Experimental.
func NewCodeCommitApprovalRuleTemplateChange() CodeCommitApprovalRuleTemplateChange {
	_init_.Initialize()

	j := jsiiProxy_CodeCommitApprovalRuleTemplateChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_codecommit.events.CodeCommitApprovalRuleTemplateChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCodeCommitApprovalRuleTemplateChange_Override(c CodeCommitApprovalRuleTemplateChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_codecommit.events.CodeCommitApprovalRuleTemplateChange",
		nil, // no parameters
		c,
	)
}

// EventBridge event pattern for CodeCommit Approval Rule Template Change.
// Experimental.
func CodeCommitApprovalRuleTemplateChange_EventPattern(options *CodeCommitApprovalRuleTemplateChange_CodeCommitApprovalRuleTemplateChangeProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateCodeCommitApprovalRuleTemplateChange_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_codecommit.events.CodeCommitApprovalRuleTemplateChange",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

