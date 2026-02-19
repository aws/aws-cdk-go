package interfacesawsiotwireless

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsiotwireless/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DeviceProfile.
// Experimental.
type IDeviceProfileRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a DeviceProfile resource.
	// Experimental.
	DeviceProfileRef() *DeviceProfileReference
}

// The jsii proxy for IDeviceProfileRef
type jsiiProxy_IDeviceProfileRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IDeviceProfileRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IDeviceProfileRef) DeviceProfileRef() *DeviceProfileReference {
	var returns *DeviceProfileReference
	_jsii_.Get(
		j,
		"deviceProfileRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDeviceProfileRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDeviceProfileRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

