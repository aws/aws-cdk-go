package awsses

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsses/internal"
)

// A dedicated IP pool.
type IDedicatedIpPool interface {
	awscdk.IResource
	// The name of the dedicated IP pool.
	DedicatedIpPoolName() *string
}

// The jsii proxy for IDedicatedIpPool
type jsiiProxy_IDedicatedIpPool struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IDedicatedIpPool) DedicatedIpPoolName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dedicatedIpPoolName",
		&returns,
	)
	return returns
}

