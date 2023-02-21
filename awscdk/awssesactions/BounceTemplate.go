package awssesactions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A bounce template.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   bounceTemplate := awscdk.Aws_ses_actions.BounceTemplate_MAILBOX_DOES_NOT_EXIST()
//
type BounceTemplate interface {
	Props() *BounceTemplateProps
}

// The jsii proxy struct for BounceTemplate
type jsiiProxy_BounceTemplate struct {
	_ byte // padding
}

func (j *jsiiProxy_BounceTemplate) Props() *BounceTemplateProps {
	var returns *BounceTemplateProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}


func NewBounceTemplate(props *BounceTemplateProps) BounceTemplate {
	_init_.Initialize()

	if err := validateNewBounceTemplateParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_BounceTemplate{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ses_actions.BounceTemplate",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewBounceTemplate_Override(b BounceTemplate, props *BounceTemplateProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ses_actions.BounceTemplate",
		[]interface{}{props},
		b,
	)
}

func BounceTemplate_MAILBOX_DOES_NOT_EXIST() BounceTemplate {
	_init_.Initialize()
	var returns BounceTemplate
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ses_actions.BounceTemplate",
		"MAILBOX_DOES_NOT_EXIST",
		&returns,
	)
	return returns
}

func BounceTemplate_MAILBOX_FULL() BounceTemplate {
	_init_.Initialize()
	var returns BounceTemplate
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ses_actions.BounceTemplate",
		"MAILBOX_FULL",
		&returns,
	)
	return returns
}

func BounceTemplate_MESSAGE_CONTENT_REJECTED() BounceTemplate {
	_init_.Initialize()
	var returns BounceTemplate
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ses_actions.BounceTemplate",
		"MESSAGE_CONTENT_REJECTED",
		&returns,
	)
	return returns
}

func BounceTemplate_MESSAGE_TOO_LARGE() BounceTemplate {
	_init_.Initialize()
	var returns BounceTemplate
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ses_actions.BounceTemplate",
		"MESSAGE_TOO_LARGE",
		&returns,
	)
	return returns
}

func BounceTemplate_TEMPORARY_FAILURE() BounceTemplate {
	_init_.Initialize()
	var returns BounceTemplate
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ses_actions.BounceTemplate",
		"TEMPORARY_FAILURE",
		&returns,
	)
	return returns
}

