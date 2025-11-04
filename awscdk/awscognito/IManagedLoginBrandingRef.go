package awscognito

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ManagedLoginBranding.
// Experimental.
type IManagedLoginBrandingRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a ManagedLoginBranding resource.
	// Experimental.
	ManagedLoginBrandingRef() *ManagedLoginBrandingReference
}

// The jsii proxy for IManagedLoginBrandingRef
type jsiiProxy_IManagedLoginBrandingRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IManagedLoginBrandingRef) ManagedLoginBrandingRef() *ManagedLoginBrandingReference {
	var returns *ManagedLoginBrandingReference
	_jsii_.Get(
		j,
		"managedLoginBrandingRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IManagedLoginBrandingRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IManagedLoginBrandingRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

