package awsiottwinmaker

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiottwinmaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SyncJob.
// Experimental.
type ISyncJobRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a SyncJob resource.
	// Experimental.
	SyncJobRef() *SyncJobReference
}

// The jsii proxy for ISyncJobRef
type jsiiProxy_ISyncJobRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_ISyncJobRef) SyncJobRef() *SyncJobReference {
	var returns *SyncJobReference
	_jsii_.Get(
		j,
		"syncJobRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISyncJobRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISyncJobRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

