package awsapprunner

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapprunner/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VpcIngressConnection.
// Experimental.
type IVpcIngressConnectionRef interface {
	constructs.IConstruct
	// A reference to a VpcIngressConnection resource.
	// Experimental.
	VpcIngressConnectionRef() *VpcIngressConnectionReference
}

// The jsii proxy for IVpcIngressConnectionRef
type jsiiProxy_IVpcIngressConnectionRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IVpcIngressConnectionRef) VpcIngressConnectionRef() *VpcIngressConnectionReference {
	var returns *VpcIngressConnectionReference
	_jsii_.Get(
		j,
		"vpcIngressConnectionRef",
		&returns,
	)
	return returns
}

