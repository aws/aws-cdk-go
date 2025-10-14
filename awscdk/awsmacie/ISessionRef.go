package awsmacie

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsmacie/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Session.
// Experimental.
type ISessionRef interface {
	constructs.IConstruct
	// A reference to a Session resource.
	// Experimental.
	SessionRef() *SessionReference
}

// The jsii proxy for ISessionRef
type jsiiProxy_ISessionRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ISessionRef) SessionRef() *SessionReference {
	var returns *SessionReference
	_jsii_.Get(
		j,
		"sessionRef",
		&returns,
	)
	return returns
}

