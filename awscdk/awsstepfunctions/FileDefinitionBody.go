package awsstepfunctions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/constructs-go/constructs/v10"
)

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var chainable iChainable
//
//   fileDefinitionBody := awscdk.Aws_stepfunctions.FileDefinitionBody_FromChainable(chainable)
//
type FileDefinitionBody interface {
	DefinitionBody
	Path() *string
	Bind(scope constructs.Construct, _sfnPrincipal awsiam.IPrincipal, _sfnProps *StateMachineProps, _graph StateGraph) *DefinitionConfig
}

// The jsii proxy struct for FileDefinitionBody
type jsiiProxy_FileDefinitionBody struct {
	jsiiProxy_DefinitionBody
}

func (j *jsiiProxy_FileDefinitionBody) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}


func NewFileDefinitionBody(path *string, options *awss3assets.AssetOptions) FileDefinitionBody {
	_init_.Initialize()

	if err := validateNewFileDefinitionBodyParameters(path, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_FileDefinitionBody{}

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions.FileDefinitionBody",
		[]interface{}{path, options},
		&j,
	)

	return &j
}

func NewFileDefinitionBody_Override(f FileDefinitionBody, path *string, options *awss3assets.AssetOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions.FileDefinitionBody",
		[]interface{}{path, options},
		f,
	)
}

func FileDefinitionBody_FromChainable(chainable IChainable) DefinitionBody {
	_init_.Initialize()

	if err := validateFileDefinitionBody_FromChainableParameters(chainable); err != nil {
		panic(err)
	}
	var returns DefinitionBody

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.FileDefinitionBody",
		"fromChainable",
		[]interface{}{chainable},
		&returns,
	)

	return returns
}

func FileDefinitionBody_FromFile(path *string, options *awss3assets.AssetOptions) DefinitionBody {
	_init_.Initialize()

	if err := validateFileDefinitionBody_FromFileParameters(path, options); err != nil {
		panic(err)
	}
	var returns DefinitionBody

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.FileDefinitionBody",
		"fromFile",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

func FileDefinitionBody_FromString(definition *string) DefinitionBody {
	_init_.Initialize()

	if err := validateFileDefinitionBody_FromStringParameters(definition); err != nil {
		panic(err)
	}
	var returns DefinitionBody

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.FileDefinitionBody",
		"fromString",
		[]interface{}{definition},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FileDefinitionBody) Bind(scope constructs.Construct, _sfnPrincipal awsiam.IPrincipal, _sfnProps *StateMachineProps, _graph StateGraph) *DefinitionConfig {
	if err := f.validateBindParameters(scope, _sfnPrincipal, _sfnProps); err != nil {
		panic(err)
	}
	var returns *DefinitionConfig

	_jsii_.Invoke(
		f,
		"bind",
		[]interface{}{scope, _sfnPrincipal, _sfnProps, _graph},
		&returns,
	)

	return returns
}

