package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VPCBlockPublicAccessExclusion.
// Experimental.
type IVPCBlockPublicAccessExclusionRef interface {
	constructs.IConstruct
	// A reference to a VPCBlockPublicAccessExclusion resource.
	// Experimental.
	VpcBlockPublicAccessExclusionRef() *VPCBlockPublicAccessExclusionReference
}

// The jsii proxy for IVPCBlockPublicAccessExclusionRef
type jsiiProxy_IVPCBlockPublicAccessExclusionRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IVPCBlockPublicAccessExclusionRef) VpcBlockPublicAccessExclusionRef() *VPCBlockPublicAccessExclusionReference {
	var returns *VPCBlockPublicAccessExclusionReference
	_jsii_.Get(
		j,
		"vpcBlockPublicAccessExclusionRef",
		&returns,
	)
	return returns
}

