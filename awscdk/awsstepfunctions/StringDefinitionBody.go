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
//   stringDefinitionBody := awscdk.Aws_stepfunctions.StringDefinitionBody_FromChainable(chainable)
//
type StringDefinitionBody interface {
	DefinitionBody
	Body() *string
	Bind(_scope constructs.Construct, _sfnPrincipal awsiam.IPrincipal, _sfnProps *StateMachineProps, _graph StateGraph) *DefinitionConfig
}

// The jsii proxy struct for StringDefinitionBody
type jsiiProxy_StringDefinitionBody struct {
	jsiiProxy_DefinitionBody
}

func (j *jsiiProxy_StringDefinitionBody) Body() *string {
	var returns *string
	_jsii_.Get(
		j,
		"body",
		&returns,
	)
	return returns
}


func NewStringDefinitionBody(body *string) StringDefinitionBody {
	_init_.Initialize()

	if err := validateNewStringDefinitionBodyParameters(body); err != nil {
		panic(err)
	}
	j := jsiiProxy_StringDefinitionBody{}

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions.StringDefinitionBody",
		[]interface{}{body},
		&j,
	)

	return &j
}

func NewStringDefinitionBody_Override(s StringDefinitionBody, body *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions.StringDefinitionBody",
		[]interface{}{body},
		s,
	)
}

func StringDefinitionBody_FromChainable(chainable IChainable) DefinitionBody {
	_init_.Initialize()

	if err := validateStringDefinitionBody_FromChainableParameters(chainable); err != nil {
		panic(err)
	}
	var returns DefinitionBody

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.StringDefinitionBody",
		"fromChainable",
		[]interface{}{chainable},
		&returns,
	)

	return returns
}

func StringDefinitionBody_FromFile(path *string, options *awss3assets.AssetOptions) DefinitionBody {
	_init_.Initialize()

	if err := validateStringDefinitionBody_FromFileParameters(path, options); err != nil {
		panic(err)
	}
	var returns DefinitionBody

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.StringDefinitionBody",
		"fromFile",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

func StringDefinitionBody_FromString(definition *string) DefinitionBody {
	_init_.Initialize()

	if err := validateStringDefinitionBody_FromStringParameters(definition); err != nil {
		panic(err)
	}
	var returns DefinitionBody

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.StringDefinitionBody",
		"fromString",
		[]interface{}{definition},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StringDefinitionBody) Bind(_scope constructs.Construct, _sfnPrincipal awsiam.IPrincipal, _sfnProps *StateMachineProps, _graph StateGraph) *DefinitionConfig {
	if err := s.validateBindParameters(_scope, _sfnPrincipal, _sfnProps); err != nil {
		panic(err)
	}
	var returns *DefinitionConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{_scope, _sfnPrincipal, _sfnProps, _graph},
		&returns,
	)

	return returns
}

