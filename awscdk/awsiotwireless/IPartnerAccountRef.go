package awsiotwireless

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiotwireless/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PartnerAccount.
// Experimental.
type IPartnerAccountRef interface {
	constructs.IConstruct
}

// The jsii proxy for IPartnerAccountRef
type jsiiProxy_IPartnerAccountRef struct {
	internal.Type__constructsIConstruct
}

