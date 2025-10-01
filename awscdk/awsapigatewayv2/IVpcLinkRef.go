package awsapigatewayv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VpcLink.
// Experimental.
type IVpcLinkRef interface {
	constructs.IConstruct
	// A reference to a VpcLink resource.
	// Experimental.
	VpcLinkRef() *VpcLinkReference
}

// The jsii proxy for IVpcLinkRef
type jsiiProxy_IVpcLinkRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IVpcLinkRef) VpcLinkRef() *VpcLinkReference {
	var returns *VpcLinkReference
	_jsii_.Get(
		j,
		"vpcLinkRef",
		&returns,
	)
	return returns
}

