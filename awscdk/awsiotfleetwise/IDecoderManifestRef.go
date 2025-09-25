package awsiotfleetwise

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiotfleetwise/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DecoderManifest.
// Experimental.
type IDecoderManifestRef interface {
	constructs.IConstruct
	// A reference to a DecoderManifest resource.
	// Experimental.
	DecoderManifestRef() *DecoderManifestReference
}

// The jsii proxy for IDecoderManifestRef
type jsiiProxy_IDecoderManifestRef struct {
	internal.Type__constructsIConstruct
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

