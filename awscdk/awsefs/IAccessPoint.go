package awsefs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsefs/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsefs"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents an EFS AccessPoint.
type IAccessPoint interface {
	interfacesawsefs.IAccessPointRef
	awscdk.IResource
	// The ARN of the AccessPoint.
	AccessPointArn() *string
	// The ID of the AccessPoint.
	AccessPointId() *string
	// The EFS file system.
	FileSystem() IFileSystem
}

// The jsii proxy for IAccessPoint
type jsiiProxy_IAccessPoint struct {
	internal.Type__interfacesawsefsIAccessPointRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IAccessPoint) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (j *jsiiProxy_IAccessPoint) AccessPointArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessPointArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAccessPoint) AccessPointId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessPointId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAccessPoint) FileSystem() IFileSystem {
	var returns IFileSystem
	_jsii_.Get(
		j,
		"fileSystem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAccessPoint) AccessPointRef() *interfacesawsefs.AccessPointReference {
	var returns *interfacesawsefs.AccessPointReference
	_jsii_.Get(
		j,
		"accessPointRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAccessPoint) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAccessPoint) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAccessPoint) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

