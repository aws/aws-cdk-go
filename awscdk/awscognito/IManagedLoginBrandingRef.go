package awscognito

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ManagedLoginBranding.
// Experimental.
type IManagedLoginBrandingRef interface {
	constructs.IConstruct
}

// The jsii proxy for IManagedLoginBrandingRef
type jsiiProxy_IManagedLoginBrandingRef struct {
	internal.Type__constructsIConstruct
}

