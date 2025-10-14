package awsb2bi

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsb2bi/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Partnership.
// Experimental.
type IPartnershipRef interface {
	constructs.IConstruct
	// A reference to a Partnership resource.
	// Experimental.
	PartnershipRef() *PartnershipReference
}

// The jsii proxy for IPartnershipRef
type jsiiProxy_IPartnershipRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IPartnershipRef) PartnershipRef() *PartnershipReference {
	var returns *PartnershipReference
	_jsii_.Get(
		j,
		"partnershipRef",
		&returns,
	)
	return returns
}

