package awsdevopsguru

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdevopsguru/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ResourceCollection.
// Experimental.
type IResourceCollectionRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a ResourceCollection resource.
	// Experimental.
	ResourceCollectionRef() *ResourceCollectionReference
}

// The jsii proxy for IResourceCollectionRef
type jsiiProxy_IResourceCollectionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IResourceCollectionRef) ResourceCollectionRef() *ResourceCollectionReference {
	var returns *ResourceCollectionReference
	_jsii_.Get(
		j,
		"resourceCollectionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IResourceCollectionRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IResourceCollectionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

