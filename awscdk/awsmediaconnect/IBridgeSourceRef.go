package awsmediaconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsmediaconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a BridgeSource.
// Experimental.
type IBridgeSourceRef interface {
	constructs.IConstruct
	// A reference to a BridgeSource resource.
	// Experimental.
	BridgeSourceRef() *BridgeSourceReference
}

// The jsii proxy for IBridgeSourceRef
type jsiiProxy_IBridgeSourceRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IBridgeSourceRef) BridgeSourceRef() *BridgeSourceReference {
	var returns *BridgeSourceReference
	_jsii_.Get(
		j,
		"bridgeSourceRef",
		&returns,
	)
	return returns
}

