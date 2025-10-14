package awscognito

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a UserPoolGroup.
// Experimental.
type IUserPoolGroupRef interface {
	constructs.IConstruct
	// A reference to a UserPoolGroup resource.
	// Experimental.
	UserPoolGroupRef() *UserPoolGroupReference
}

// The jsii proxy for IUserPoolGroupRef
type jsiiProxy_IUserPoolGroupRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IUserPoolGroupRef) UserPoolGroupRef() *UserPoolGroupReference {
	var returns *UserPoolGroupReference
	_jsii_.Get(
		j,
		"userPoolGroupRef",
		&returns,
	)
	return returns
}

