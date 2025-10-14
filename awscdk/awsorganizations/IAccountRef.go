package awsorganizations

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsorganizations/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Account.
// Experimental.
type IAccountRef interface {
	constructs.IConstruct
	// A reference to a Account resource.
	// Experimental.
	AccountRef() *AccountReference
}

// The jsii proxy for IAccountRef
type jsiiProxy_IAccountRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IAccountRef) AccountRef() *AccountReference {
	var returns *AccountReference
	_jsii_.Get(
		j,
		"accountRef",
		&returns,
	)
	return returns
}

