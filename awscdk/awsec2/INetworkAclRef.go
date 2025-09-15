package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a NetworkAcl.
// Experimental.
type INetworkAclRef interface {
	constructs.IConstruct
	// A reference to a NetworkAcl resource.
	// Experimental.
	NetworkAclRef() *NetworkAclReference
}

// The jsii proxy for INetworkAclRef
type jsiiProxy_INetworkAclRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_INetworkAclRef) NetworkAclRef() *NetworkAclReference {
	var returns *NetworkAclReference
	_jsii_.Get(
		j,
		"networkAclRef",
		&returns,
	)
	return returns
}

