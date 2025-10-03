package awsmediaconnect

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsmediaconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a BridgeSource.
// Experimental.
type IBridgeSourceRef interface {
	constructs.IConstruct
}

// The jsii proxy for IBridgeSourceRef
type jsiiProxy_IBridgeSourceRef struct {
	internal.Type__constructsIConstruct
}

