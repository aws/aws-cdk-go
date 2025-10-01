package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a NetworkInterfacePermission.
// Experimental.
type INetworkInterfacePermissionRef interface {
	constructs.IConstruct
	// A reference to a NetworkInterfacePermission resource.
	// Experimental.
	NetworkInterfacePermissionRef() *NetworkInterfacePermissionReference
}

// The jsii proxy for INetworkInterfacePermissionRef
type jsiiProxy_INetworkInterfacePermissionRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_INetworkInterfacePermissionRef) NetworkInterfacePermissionRef() *NetworkInterfacePermissionReference {
	var returns *NetworkInterfacePermissionReference
	_jsii_.Get(
		j,
		"networkInterfacePermissionRef",
		&returns,
	)
	return returns
}

