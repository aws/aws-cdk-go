package awsb2bi

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsb2bi/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Partnership.
// Experimental.
type IPartnershipRef interface {
	constructs.IConstruct
}

// The jsii proxy for IPartnershipRef
type jsiiProxy_IPartnershipRef struct {
	internal.Type__constructsIConstruct
}

