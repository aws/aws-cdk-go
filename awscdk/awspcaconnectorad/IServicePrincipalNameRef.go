package awspcaconnectorad

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awspcaconnectorad/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ServicePrincipalName.
// Experimental.
type IServicePrincipalNameRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a ServicePrincipalName resource.
	// Experimental.
	ServicePrincipalNameRef() *ServicePrincipalNameReference
}

// The jsii proxy for IServicePrincipalNameRef
type jsiiProxy_IServicePrincipalNameRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IServicePrincipalNameRef) ServicePrincipalNameRef() *ServicePrincipalNameReference {
	var returns *ServicePrincipalNameReference
	_jsii_.Get(
		j,
		"servicePrincipalNameRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServicePrincipalNameRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServicePrincipalNameRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

