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
//   chainDefinitionBody := awscdk.Aws_stepfunctions.ChainDefinitionBody_FromChainable(chainable)
//
type ChainDefinitionBody interface {
	DefinitionBody
	Chainable() IChainable
	Bind(scope constructs.Construct, _sfnPrincipal awsiam.IPrincipal, sfnProps *StateMachineProps, graph StateGraph) *DefinitionConfig
}

// The jsii proxy struct for ChainDefinitionBody
type jsiiProxy_ChainDefinitionBody struct {
	jsiiProxy_DefinitionBody
}

func (j *jsiiProxy_ChainDefinitionBody) Chainable() IChainable {
	var returns IChainable
	_jsii_.Get(
		j,
		"chainable",
		&returns,
	)
	return returns
}


func NewChainDefinitionBody(chainable IChainable) ChainDefinitionBody {
	_init_.Initialize()

	if err := validateNewChainDefinitionBodyParameters(chainable); err != nil {
		panic(err)
	}
	j := jsiiProxy_ChainDefinitionBody{}

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions.ChainDefinitionBody",
		[]interface{}{chainable},
		&j,
	)

	return &j
}

func NewChainDefinitionBody_Override(c ChainDefinitionBody, chainable IChainable) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions.ChainDefinitionBody",
		[]interface{}{chainable},
		c,
	)
}

func ChainDefinitionBody_FromChainable(chainable IChainable) DefinitionBody {
	_init_.Initialize()

	if err := validateChainDefinitionBody_FromChainableParameters(chainable); err != nil {
		panic(err)
	}
	var returns DefinitionBody

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.ChainDefinitionBody",
		"fromChainable",
		[]interface{}{chainable},
		&returns,
	)

	return returns
}

func ChainDefinitionBody_FromFile(path *string, options *awss3assets.AssetOptions) DefinitionBody {
	_init_.Initialize()

	if err := validateChainDefinitionBody_FromFileParameters(path, options); err != nil {
		panic(err)
	}
	var returns DefinitionBody

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.ChainDefinitionBody",
		"fromFile",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

func ChainDefinitionBody_FromString(definition *string) DefinitionBody {
	_init_.Initialize()

	if err := validateChainDefinitionBody_FromStringParameters(definition); err != nil {
		panic(err)
	}
	var returns DefinitionBody

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.ChainDefinitionBody",
		"fromString",
		[]interface{}{definition},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChainDefinitionBody) Bind(scope constructs.Construct, _sfnPrincipal awsiam.IPrincipal, sfnProps *StateMachineProps, graph StateGraph) *DefinitionConfig {
	if err := c.validateBindParameters(scope, _sfnPrincipal, sfnProps); err != nil {
		panic(err)
	}
	var returns *DefinitionConfig

	_jsii_.Invoke(
		c,
		"bind",
		[]interface{}{scope, _sfnPrincipal, sfnProps, graph},
		&returns,
	)

	return returns
}

