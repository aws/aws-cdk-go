package awsnotificationscontacts

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsnotificationscontacts/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EmailContact.
// Experimental.
type IEmailContactRef interface {
	constructs.IConstruct
	// A reference to a EmailContact resource.
	// Experimental.
	EmailContactRef() *EmailContactReference
}

// The jsii proxy for IEmailContactRef
type jsiiProxy_IEmailContactRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IEmailContactRef) EmailContactRef() *EmailContactReference {
	var returns *EmailContactReference
	_jsii_.Get(
		j,
		"emailContactRef",
		&returns,
	)
	return returns
}

