package awsiam

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SAMLProvider.
// Experimental.
type ISAMLProviderRef interface {
	constructs.IConstruct
}

// The jsii proxy for ISAMLProviderRef
type jsiiProxy_ISAMLProviderRef struct {
	internal.Type__constructsIConstruct
}

