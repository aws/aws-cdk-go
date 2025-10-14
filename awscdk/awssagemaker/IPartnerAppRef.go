package awssagemaker

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssagemaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PartnerApp.
// Experimental.
type IPartnerAppRef interface {
	constructs.IConstruct
	// A reference to a PartnerApp resource.
	// Experimental.
	PartnerAppRef() *PartnerAppReference
}

// The jsii proxy for IPartnerAppRef
type jsiiProxy_IPartnerAppRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IPartnerAppRef) PartnerAppRef() *PartnerAppReference {
	var returns *PartnerAppReference
	_jsii_.Get(
		j,
		"partnerAppRef",
		&returns,
	)
	return returns
}

