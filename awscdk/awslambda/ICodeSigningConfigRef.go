package awslambda

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CodeSigningConfig.
// Experimental.
type ICodeSigningConfigRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a CodeSigningConfig resource.
	// Experimental.
	CodeSigningConfigRef() *CodeSigningConfigReference
}

// The jsii proxy for ICodeSigningConfigRef
type jsiiProxy_ICodeSigningConfigRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_ICodeSigningConfigRef) CodeSigningConfigRef() *CodeSigningConfigReference {
	var returns *CodeSigningConfigReference
	_jsii_.Get(
		j,
		"codeSigningConfigRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICodeSigningConfigRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICodeSigningConfigRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

