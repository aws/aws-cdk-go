package awsrds

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsrds/internal"
)

// Interface for a subnet group.
type ISubnetGroup interface {
	awscdk.IResource
	// The name of the subnet group.
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

