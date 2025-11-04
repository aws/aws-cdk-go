package awssagemaker

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssagemaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CodeRepository.
// Experimental.
type ICodeRepositoryRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a CodeRepository resource.
	// Experimental.
	CodeRepositoryRef() *CodeRepositoryReference
}

// The jsii proxy for ICodeRepositoryRef
type jsiiProxy_ICodeRepositoryRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_ICodeRepositoryRef) CodeRepositoryRef() *CodeRepositoryReference {
	var returns *CodeRepositoryReference
	_jsii_.Get(
		j,
		"codeRepositoryRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICodeRepositoryRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICodeRepositoryRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

