package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VPCBlockPublicAccessOptions.
// Experimental.
type IVPCBlockPublicAccessOptionsRef interface {
	constructs.IConstruct
	// A reference to a VPCBlockPublicAccessOptions resource.
	// Experimental.
	VpcBlockPublicAccessOptionsRef() *VPCBlockPublicAccessOptionsReference
}

// The jsii proxy for IVPCBlockPublicAccessOptionsRef
type jsiiProxy_IVPCBlockPublicAccessOptionsRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IVPCBlockPublicAccessOptionsRef) VpcBlockPublicAccessOptionsRef() *VPCBlockPublicAccessOptionsReference {
	var returns *VPCBlockPublicAccessOptionsReference
	_jsii_.Get(
		j,
		"vpcBlockPublicAccessOptionsRef",
		&returns,
	)
	return returns
}

