package awsses

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsses/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsses"
	"github.com/aws/constructs-go/constructs/v10"
)

// Virtual Deliverability Manager (VDM) attributes.
type IVdmAttributes interface {
	awscdk.IResource
	interfacesawsses.IVdmAttributesRef
	// The name of the resource behind the Virtual Deliverability Manager attributes.
	VdmAttributesName() *string
}

// The jsii proxy for IVdmAttributes
type jsiiProxy_IVdmAttributes struct {
	internal.Type__awscdkIResource
	internal.Type__interfacesawssesIVdmAttributesRef
}

func (i *jsiiProxy_IVdmAttributes) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
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

func (j *jsiiProxy_IVdmAttributes) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVdmAttributes) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVdmAttributes) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVdmAttributes) VdmAttributesRef() *interfacesawsses.VdmAttributesReference {
	var returns *interfacesawsses.VdmAttributesReference
	_jsii_.Get(
		j,
		"vdmAttributesRef",
		&returns,
	)
	return returns
}

