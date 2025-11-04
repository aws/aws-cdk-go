package awsiotwireless

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiotwireless/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PartnerAccount.
// Experimental.
type IPartnerAccountRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a PartnerAccount resource.
	// Experimental.
	PartnerAccountRef() *PartnerAccountReference
}

// The jsii proxy for IPartnerAccountRef
type jsiiProxy_IPartnerAccountRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IPartnerAccountRef) PartnerAccountRef() *PartnerAccountReference {
	var returns *PartnerAccountReference
	_jsii_.Get(
		j,
		"partnerAccountRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPartnerAccountRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPartnerAccountRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

