package awsecr

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Repository.
// Experimental.
type IRepositoryRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a Repository resource.
	// Experimental.
	RepositoryRef() *RepositoryReference
}

// The jsii proxy for IRepositoryRef
type jsiiProxy_IRepositoryRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IRepositoryRef) RepositoryRef() *RepositoryReference {
	var returns *RepositoryReference
	_jsii_.Get(
		j,
		"repositoryRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRepositoryRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRepositoryRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

