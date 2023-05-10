package awseks

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awseks/internal"
)

// NodeGroup interface.
// Experimental.
type INodegroup interface {
	awscdk.IResource
	// Name of the nodegroup.
	// Experimental.
	NodegroupName() *string
}

// The jsii proxy for INodegroup
type jsiiProxy_INodegroup struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_INodegroup) NodegroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodegroupName",
		&returns,
	)
	return returns
}

