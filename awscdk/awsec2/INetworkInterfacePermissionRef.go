package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a NetworkInterfacePermission.
// Experimental.
type INetworkInterfacePermissionRef interface {
	constructs.IConstruct
}

// The jsii proxy for INetworkInterfacePermissionRef
type jsiiProxy_INetworkInterfacePermissionRef struct {
	internal.Type__constructsIConstruct
}

