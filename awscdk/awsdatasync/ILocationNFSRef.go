package awsdatasync

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdatasync/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LocationNFS.
// Experimental.
type ILocationNFSRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a LocationNFS resource.
	// Experimental.
	LocationNfsRef() *LocationNFSReference
}

// The jsii proxy for ILocationNFSRef
type jsiiProxy_ILocationNFSRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_ILocationNFSRef) LocationNfsRef() *LocationNFSReference {
	var returns *LocationNFSReference
	_jsii_.Get(
		j,
		"locationNfsRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILocationNFSRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILocationNFSRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

