package interfacesawsiotwireless

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsiotwireless/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a WirelessDeviceImportTask.
// Experimental.
type IWirelessDeviceImportTaskRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a WirelessDeviceImportTask resource.
	// Experimental.
	WirelessDeviceImportTaskRef() *WirelessDeviceImportTaskReference
}

// The jsii proxy for IWirelessDeviceImportTaskRef
type jsiiProxy_IWirelessDeviceImportTaskRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IWirelessDeviceImportTaskRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IWirelessDeviceImportTaskRef) WirelessDeviceImportTaskRef() *WirelessDeviceImportTaskReference {
	var returns *WirelessDeviceImportTaskReference
	_jsii_.Get(
		j,
		"wirelessDeviceImportTaskRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWirelessDeviceImportTaskRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWirelessDeviceImportTaskRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

