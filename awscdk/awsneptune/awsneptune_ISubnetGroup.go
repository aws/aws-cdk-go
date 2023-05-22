package awsneptune

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsneptune/internal"
)

// Interface for a subnet group.
// Experimental.
type ISubnetGroup interface {
	awscdk.IResource
	// The name of the subnet group.
	// Experimental.
	SubnetGroupName() *string
}

// The jsii proxy for ISubnetGroup
type jsiiProxy_ISubnetGroup struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_ISubnetGroup) SubnetGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetGroupName",
		&returns,
	)
	return returns
}

