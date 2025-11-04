package awsverifiedpermissions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsverifiedpermissions/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a IdentitySource.
// Experimental.
type IIdentitySourceRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a IdentitySource resource.
	// Experimental.
	IdentitySourceRef() *IdentitySourceReference
}

// The jsii proxy for IIdentitySourceRef
type jsiiProxy_IIdentitySourceRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IIdentitySourceRef) IdentitySourceRef() *IdentitySourceReference {
	var returns *IdentitySourceReference
	_jsii_.Get(
		j,
		"identitySourceRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIdentitySourceRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIdentitySourceRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

