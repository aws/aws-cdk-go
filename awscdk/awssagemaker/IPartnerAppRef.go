package awssagemaker

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssagemaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PartnerApp.
// Experimental.
type IPartnerAppRef interface {
	constructs.IConstruct
}

// The jsii proxy for IPartnerAppRef
type jsiiProxy_IPartnerAppRef struct {
	internal.Type__constructsIConstruct
}

