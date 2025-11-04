package awsiotfleetwise

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiotfleetwise/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Campaign.
// Experimental.
type ICampaignRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a Campaign resource.
	// Experimental.
	CampaignRef() *CampaignReference
}

// The jsii proxy for ICampaignRef
type jsiiProxy_ICampaignRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_ICampaignRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICampaignRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

