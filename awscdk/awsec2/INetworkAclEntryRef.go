package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a NetworkAclEntry.
// Experimental.
type INetworkAclEntryRef interface {
	constructs.IConstruct
	// A reference to a NetworkAclEntry resource.
	// Experimental.
	NetworkAclEntryRef() *NetworkAclEntryReference
}

// The jsii proxy for INetworkAclEntryRef
type jsiiProxy_INetworkAclEntryRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_INetworkAclEntryRef) NetworkAclEntryRef() *NetworkAclEntryReference {
	var returns *NetworkAclEntryReference
	_jsii_.Get(
		j,
		"networkAclEntryRef",
		&returns,
	)
	return returns
}

