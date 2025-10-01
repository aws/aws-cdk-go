package awsiot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiot/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Authorizer.
// Experimental.
type IAuthorizerRef interface {
	constructs.IConstruct
	// A reference to a Authorizer resource.
	// Experimental.
	AuthorizerRef() *AuthorizerReference
}

// The jsii proxy for IAuthorizerRef
type jsiiProxy_IAuthorizerRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IAuthorizerRef) AuthorizerRef() *AuthorizerReference {
	var returns *AuthorizerReference
	_jsii_.Get(
		j,
		"authorizerRef",
		&returns,
	)
	return returns
}

