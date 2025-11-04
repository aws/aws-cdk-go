package awsdatasync

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdatasync/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LocationHDFS.
// Experimental.
type ILocationHDFSRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a LocationHDFS resource.
	// Experimental.
	LocationHdfsRef() *LocationHDFSReference
}

// The jsii proxy for ILocationHDFSRef
type jsiiProxy_ILocationHDFSRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_ILocationHDFSRef) LocationHdfsRef() *LocationHDFSReference {
	var returns *LocationHDFSReference
	_jsii_.Get(
		j,
		"locationHdfsRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILocationHDFSRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILocationHDFSRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

