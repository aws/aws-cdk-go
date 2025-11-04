package awsssm

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsssm/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ResourceDataSync.
// Experimental.
type IResourceDataSyncRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a ResourceDataSync resource.
	// Experimental.
	ResourceDataSyncRef() *ResourceDataSyncReference
}

// The jsii proxy for IResourceDataSyncRef
type jsiiProxy_IResourceDataSyncRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IResourceDataSyncRef) ResourceDataSyncRef() *ResourceDataSyncReference {
	var returns *ResourceDataSyncReference
	_jsii_.Get(
		j,
		"resourceDataSyncRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IResourceDataSyncRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IResourceDataSyncRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

