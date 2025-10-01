package awsopensearchserverless

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsopensearchserverless/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VpcEndpoint.
// Experimental.
type IVpcEndpointRef interface {
	constructs.IConstruct
	// A reference to a VpcEndpoint resource.
	// Experimental.
	VpcEndpointRef() *VpcEndpointReference
}

// The jsii proxy for IVpcEndpointRef
type jsiiProxy_IVpcEndpointRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IVpcEndpointRef) VpcEndpointRef() *VpcEndpointReference {
	var returns *VpcEndpointReference
	_jsii_.Get(
		j,
		"vpcEndpointRef",
		&returns,
	)
	return returns
}

