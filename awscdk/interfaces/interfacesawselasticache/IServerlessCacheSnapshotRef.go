package interfacesawselasticache

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawselasticache/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ServerlessCacheSnapshot.
// Experimental.
type IServerlessCacheSnapshotRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ServerlessCacheSnapshot resource.
	// Experimental.
	ServerlessCacheSnapshotRef() *ServerlessCacheSnapshotReference
}

// The jsii proxy for IServerlessCacheSnapshotRef
type jsiiProxy_IServerlessCacheSnapshotRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IServerlessCacheSnapshotRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IServerlessCacheSnapshotRef) ServerlessCacheSnapshotRef() *ServerlessCacheSnapshotReference {
	var returns *ServerlessCacheSnapshotReference
	_jsii_.Get(
		j,
		"serverlessCacheSnapshotRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServerlessCacheSnapshotRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServerlessCacheSnapshotRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

