package awsmediaconnect

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsmediaconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Bridge.
// Experimental.
type IBridgeRef interface {
	constructs.IConstruct
}

// The jsii proxy for IBridgeRef
type jsiiProxy_IBridgeRef struct {
	internal.Type__constructsIConstruct
}

