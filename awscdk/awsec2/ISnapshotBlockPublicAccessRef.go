package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SnapshotBlockPublicAccess.
// Experimental.
type ISnapshotBlockPublicAccessRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a SnapshotBlockPublicAccess resource.
	// Experimental.
	SnapshotBlockPublicAccessRef() *SnapshotBlockPublicAccessReference
}

// The jsii proxy for ISnapshotBlockPublicAccessRef
type jsiiProxy_ISnapshotBlockPublicAccessRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_ISnapshotBlockPublicAccessRef) SnapshotBlockPublicAccessRef() *SnapshotBlockPublicAccessReference {
	var returns *SnapshotBlockPublicAccessReference
	_jsii_.Get(
		j,
		"snapshotBlockPublicAccessRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISnapshotBlockPublicAccessRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISnapshotBlockPublicAccessRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

