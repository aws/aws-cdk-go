package awscognito

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ManagedLoginBranding.
// Experimental.
type IManagedLoginBrandingRef interface {
	constructs.IConstruct
	// A reference to a ManagedLoginBranding resource.
	// Experimental.
	ManagedLoginBrandingRef() *ManagedLoginBrandingReference
}

// The jsii proxy for IManagedLoginBrandingRef
type jsiiProxy_IManagedLoginBrandingRef struct {
	internal.Type__constructsIConstruct
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

