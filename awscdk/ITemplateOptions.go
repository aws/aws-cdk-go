// Version 2 of the AWS Cloud Development Kit library
package awscdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// CloudFormation template options for a stack.
type ITemplateOptions interface {
	// Gets or sets the description of this stack.
	//
	// If provided, it will be included in the CloudFormation template's "Description" attribute.
	Description() *string
	SetDescription(d *string)
	// Metadata associated with the CloudFormation template.
	Metadata() *map[string]interface{}
	SetMetadata(m *map[string]interface{})
	// Gets or sets the AWSTemplateFormatVersion field of the CloudFormation template.
	TemplateFormatVersion() *string
	SetTemplateFormatVersion(t *string)
	// Gets or sets the top-level template transform(s) for this stack (e.g. `["AWS::Serverless-2016-10-31"]`).
	Transforms() *[]*string
	SetTransforms(t *[]*string)
}

// The jsii proxy for ITemplateOptions
type jsiiProxy_ITemplateOptions struct {
	_ byte // padding
}

func (j *jsiiProxy_ITemplateOptions) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITemplateOptions)SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ITemplateOptions) Metadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITemplateOptions)SetMetadata(val *map[string]interface{}) {
	_jsii_.Set(
		j,
		"metadata",
		val,
	)
}

func (j *jsiiProxy_ITemplateOptions) TemplateFormatVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateFormatVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITemplateOptions)SetTemplateFormatVersion(val *string) {
	_jsii_.Set(
		j,
		"templateFormatVersion",
		val,
	)
}

func (j *jsiiProxy_ITemplateOptions) Transforms() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"transforms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITemplateOptions)SetTransforms(val *[]*string) {
	_jsii_.Set(
		j,
		"transforms",
		val,
	)
}

