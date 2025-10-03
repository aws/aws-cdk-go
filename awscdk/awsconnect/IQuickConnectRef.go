package awsconnect

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a QuickConnect.
// Experimental.
type IQuickConnectRef interface {
	constructs.IConstruct
}

// The jsii proxy for IQuickConnectRef
type jsiiProxy_IQuickConnectRef struct {
	internal.Type__constructsIConstruct
}

