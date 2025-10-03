package awsiotfleetwise

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiotfleetwise/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Campaign.
// Experimental.
type ICampaignRef interface {
	constructs.IConstruct
}

// The jsii proxy for ICampaignRef
type jsiiProxy_ICampaignRef struct {
	internal.Type__constructsIConstruct
}

