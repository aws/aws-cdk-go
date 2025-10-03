package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VerifiedAccessTrustProvider.
// Experimental.
type IVerifiedAccessTrustProviderRef interface {
	constructs.IConstruct
}

// The jsii proxy for IVerifiedAccessTrustProviderRef
type jsiiProxy_IVerifiedAccessTrustProviderRef struct {
	internal.Type__constructsIConstruct
}

