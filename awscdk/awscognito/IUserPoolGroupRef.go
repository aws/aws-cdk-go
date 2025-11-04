package awscognito

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a UserPoolGroup.
// Experimental.
type IUserPoolGroupRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a UserPoolGroup resource.
	// Experimental.
	UserPoolGroupRef() *UserPoolGroupReference
}

// The jsii proxy for IUserPoolGroupRef
type jsiiProxy_IUserPoolGroupRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IUserPoolGroupRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUserPoolGroupRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

