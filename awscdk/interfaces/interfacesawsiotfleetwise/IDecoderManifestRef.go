package interfacesawsiotfleetwise

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsiotfleetwise/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DecoderManifest.
// Experimental.
type IDecoderManifestRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a DecoderManifest resource.
	// Experimental.
	DecoderManifestRef() *DecoderManifestReference
}

// The jsii proxy for IDecoderManifestRef
type jsiiProxy_IDecoderManifestRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IDecoderManifestRef) DecoderManifestRef() *DecoderManifestReference {
	var returns *DecoderManifestReference
	_jsii_.Get(
		j,
		"decoderManifestRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDecoderManifestRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDecoderManifestRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

