package awscodebuild

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// The abstract interface of a CodeBuild source.
//
// Implemented by {@link Source}.
type ISource interface {
	Bind(scope constructs.Construct, project IProject) *SourceConfig
	BadgeSupported() *bool
	Identifier() *string
	Type() *string
}

// The jsii proxy for ISource
type jsiiProxy_ISource struct {
	_ byte // padding
}

func (i *jsiiProxy_ISource) Bind(scope constructs.Construct, project IProject) *SourceConfig {
	if err := i.validateBindParameters(scope, project); err != nil {
		panic(err)
	}
	var returns *SourceConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{scope, project},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_ISource) BadgeSupported() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"badgeSupported",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISource) Identifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISource) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

