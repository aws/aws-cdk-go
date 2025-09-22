package awsconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a QuickConnect.
// Experimental.
type IQuickConnectRef interface {
	constructs.IConstruct
	// A reference to a QuickConnect resource.
	// Experimental.
	QuickConnectRef() *QuickConnectReference
}

// The jsii proxy for IQuickConnectRef
type jsiiProxy_IQuickConnectRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IQuickConnectRef) QuickConnectRef() *QuickConnectReference {
	var returns *QuickConnectReference
	_jsii_.Get(
		j,
		"quickConnectRef",
		&returns,
	)
	return returns
}

