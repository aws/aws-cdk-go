package interfacesawsrds

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsrds/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DBSnapshot.
// Experimental.
type IDBSnapshotRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a DBSnapshot resource.
	// Experimental.
	DbSnapshotRef() *DBSnapshotReference
}

// The jsii proxy for IDBSnapshotRef
type jsiiProxy_IDBSnapshotRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IDBSnapshotRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IDBSnapshotRef) DbSnapshotRef() *DBSnapshotReference {
	var returns *DBSnapshotReference
	_jsii_.Get(
		j,
		"dbSnapshotRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDBSnapshotRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDBSnapshotRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

