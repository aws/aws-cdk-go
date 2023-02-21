package awsses

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsses/internal"
)

// Virtual Deliverablity Manager (VDM) attributes.
type IVdmAttributes interface {
	awscdk.IResource
	// The name of the resource behind the Virtual Deliverablity Manager attributes.
	VdmAttributesName() *string
}

// The jsii proxy for IVdmAttributes
type jsiiProxy_IVdmAttributes struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IVdmAttributes) VdmAttributesName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vdmAttributesName",
		&returns,
	)
	return returns
}

