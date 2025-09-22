package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VPCCidrBlock.
// Experimental.
type IVPCCidrBlockRef interface {
	constructs.IConstruct
	// A reference to a VPCCidrBlock resource.
	// Experimental.
	VpcCidrBlockRef() *VPCCidrBlockReference
}

// The jsii proxy for IVPCCidrBlockRef
type jsiiProxy_IVPCCidrBlockRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IVPCCidrBlockRef) VpcCidrBlockRef() *VPCCidrBlockReference {
	var returns *VPCCidrBlockReference
	_jsii_.Get(
		j,
		"vpcCidrBlockRef",
		&returns,
	)
	return returns
}

