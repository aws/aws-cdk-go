package awslambda

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslambda"
	"github.com/aws/constructs-go/constructs/v10"
)

// A Code Signing Config.
type ICodeSigningConfig interface {
	interfacesawslambda.ICodeSigningConfigRef
	awscdk.IResource
	// The ARN of Code Signing Config.
	CodeSigningConfigArn() *string
	// The id of Code Signing Config.
	CodeSigningConfigId() *string
}

// The jsii proxy for ICodeSigningConfig
type jsiiProxy_ICodeSigningConfig struct {
	internal.Type__interfacesawslambdaICodeSigningConfigRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_ICodeSigningConfig) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (j *jsiiProxy_ICodeSigningConfig) CodeSigningConfigArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codeSigningConfigArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICodeSigningConfig) CodeSigningConfigId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codeSigningConfigId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICodeSigningConfig) CodeSigningConfigRef() *interfacesawslambda.CodeSigningConfigReference {
	var returns *interfacesawslambda.CodeSigningConfigReference
	_jsii_.Get(
		j,
		"codeSigningConfigRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICodeSigningConfig) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICodeSigningConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICodeSigningConfig) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

