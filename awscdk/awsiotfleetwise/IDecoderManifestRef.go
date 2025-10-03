package awsiotfleetwise

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiotfleetwise/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DecoderManifest.
// Experimental.
type IDecoderManifestRef interface {
	constructs.IConstruct
}

// The jsii proxy for IDecoderManifestRef
type jsiiProxy_IDecoderManifestRef struct {
	internal.Type__constructsIConstruct
}

