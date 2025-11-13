package interfacesawsdatasync

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsdatasync/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LocationObjectStorage.
// Experimental.
type ILocationObjectStorageRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a LocationObjectStorage resource.
	// Experimental.
	LocationObjectStorageRef() *LocationObjectStorageReference
}

// The jsii proxy for ILocationObjectStorageRef
type jsiiProxy_ILocationObjectStorageRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_ILocationObjectStorageRef) LocationObjectStorageRef() *LocationObjectStorageReference {
	var returns *LocationObjectStorageReference
	_jsii_.Get(
		j,
		"locationObjectStorageRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILocationObjectStorageRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILocationObjectStorageRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

