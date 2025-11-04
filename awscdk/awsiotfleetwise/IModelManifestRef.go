package awsiotfleetwise

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiotfleetwise/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ModelManifest.
// Experimental.
type IModelManifestRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a ModelManifest resource.
	// Experimental.
	ModelManifestRef() *ModelManifestReference
}

// The jsii proxy for IModelManifestRef
type jsiiProxy_IModelManifestRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IModelManifestRef) ModelManifestRef() *ModelManifestReference {
	var returns *ModelManifestReference
	_jsii_.Get(
		j,
		"modelManifestRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IModelManifestRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IModelManifestRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

