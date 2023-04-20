// An experiment to bundle the entire CDK into a single module
package awscdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// CloudFormation template options for a stack.
// Experimental.
type ITemplateOptions interface {
	// Gets or sets the description of this stack.
	//
	// If provided, it will be included in the CloudFormation template's "Description" attribute.
	// Experimental.
	Description() *string
	// Experimental.
	SetDescription(d *string)
	// Metadata associated with the CloudFormation template.
	// Experimental.
	Metadata() *map[string]interface{}
	// Experimental.
	SetMetadata(m *map[string]interface{})
	// Gets or sets the AWSTemplateFormatVersion field of the CloudFormation template.
	// Experimental.
	TemplateFormatVersion() *string
	// Experimental.
	SetTemplateFormatVersion(t *string)
	// Gets or sets the top-level template transform for this stack (e.g. "AWS::Serverless-2016-10-31").
	// Deprecated: use `transforms` instead.
	Transform() *string
	// Deprecated: use `transforms` instead.
	SetTransform(t *string)
	// Gets or sets the top-level template transform(s) for this stack (e.g. `["AWS::Serverless-2016-10-31"]`).
	// Experimental.
	Transforms() *[]*string
	// Experimental.
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

func (j *jsiiProxy_ITemplateOptions) Transform() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transform",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITemplateOptions)SetTransform(val *string) {
	_jsii_.Set(
		j,
		"transform",
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

