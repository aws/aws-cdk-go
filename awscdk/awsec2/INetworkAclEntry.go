package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
)

// A NetworkAclEntry.
type INetworkAclEntry interface {
	awscdk.IResource
	// The network ACL.
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

