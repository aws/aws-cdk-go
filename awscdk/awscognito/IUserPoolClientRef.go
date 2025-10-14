package awscognito

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a UserPoolClient.
// Experimental.
type IUserPoolClientRef interface {
	constructs.IConstruct
	// A reference to a UserPoolClient resource.
	// Experimental.
	UserPoolClientRef() *UserPoolClientReference
}

// The jsii proxy for IUserPoolClientRef
type jsiiProxy_IUserPoolClientRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IUserPoolClientRef) UserPoolClientRef() *UserPoolClientReference {
	var returns *UserPoolClientReference
	_jsii_.Get(
		j,
		"userPoolClientRef",
		&returns,
	)
	return returns
}

