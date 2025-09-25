package awsiotwireless

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiotwireless/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PartnerAccount.
// Experimental.
type IPartnerAccountRef interface {
	constructs.IConstruct
	// A reference to a PartnerAccount resource.
	// Experimental.
	PartnerAccountRef() *PartnerAccountReference
}

// The jsii proxy for IPartnerAccountRef
type jsiiProxy_IPartnerAccountRef struct {
	internal.Type__constructsIConstruct
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

