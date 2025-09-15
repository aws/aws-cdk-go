package awsconnectcampaigns

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsconnectcampaigns/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Campaign.
// Experimental.
type ICampaignRef interface {
	constructs.IConstruct
	// A reference to a Campaign resource.
	// Experimental.
	CampaignRef() *CampaignReference
}

// The jsii proxy for ICampaignRef
type jsiiProxy_ICampaignRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ICampaignRef) CampaignRef() *CampaignReference {
	var returns *CampaignReference
	_jsii_.Get(
		j,
		"campaignRef",
		&returns,
	)
	return returns
}

