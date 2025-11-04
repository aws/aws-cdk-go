package awsses

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsses/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EmailIdentity.
// Experimental.
type IEmailIdentityRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a EmailIdentity resource.
	// Experimental.
	EmailIdentityRef() *EmailIdentityReference
}

// The jsii proxy for IEmailIdentityRef
type jsiiProxy_IEmailIdentityRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IEmailIdentityRef) EmailIdentityRef() *EmailIdentityReference {
	var returns *EmailIdentityReference
	_jsii_.Get(
		j,
		"emailIdentityRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEmailIdentityRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEmailIdentityRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

