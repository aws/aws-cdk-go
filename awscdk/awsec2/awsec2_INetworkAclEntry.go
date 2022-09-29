package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsec2/internal"
)

// A NetworkAclEntry.
// Experimental.
type INetworkAclEntry interface {
	awscdk.IResource
	// The network ACL.
	// Experimental.
	NetworkAcl() INetworkAcl
}

// The jsii proxy for INetworkAclEntry
type jsiiProxy_INetworkAclEntry struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_INetworkAclEntry) NetworkAcl() INetworkAcl {
	var returns INetworkAcl
	_jsii_.Get(
		j,
		"networkAcl",
		&returns,
	)
	return returns
}

