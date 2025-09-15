package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VPCEndpointConnectionNotification.
// Experimental.
type IVPCEndpointConnectionNotificationRef interface {
	constructs.IConstruct
	// A reference to a VPCEndpointConnectionNotification resource.
	// Experimental.
	VpcEndpointConnectionNotificationRef() *VPCEndpointConnectionNotificationReference
}

// The jsii proxy for IVPCEndpointConnectionNotificationRef
type jsiiProxy_IVPCEndpointConnectionNotificationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IVPCEndpointConnectionNotificationRef) VpcEndpointConnectionNotificationRef() *VPCEndpointConnectionNotificationReference {
	var returns *VPCEndpointConnectionNotificationReference
	_jsii_.Get(
		j,
		"vpcEndpointConnectionNotificationRef",
		&returns,
	)
	return returns
}

