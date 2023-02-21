package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
)

// A NetworkAcl.
type INetworkAcl interface {
	awscdk.IResource
	// Add a new entry to the ACL.
	AddEntry(id *string, options *CommonNetworkAclEntryOptions) NetworkAclEntry
	// ID for the current Network ACL.
	NetworkAclId() *string
}

// The jsii proxy for INetworkAcl
type jsiiProxy_INetworkAcl struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_INetworkAcl) AddEntry(id *string, options *CommonNetworkAclEntryOptions) NetworkAclEntry {
	if err := i.validateAddEntryParameters(id, options); err != nil {
		panic(err)
	}
	var returns NetworkAclEntry

	_jsii_.Invoke(
		i,
		"addEntry",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_INetworkAcl) NetworkAclId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkAclId",
		&returns,
	)
	return returns
}

