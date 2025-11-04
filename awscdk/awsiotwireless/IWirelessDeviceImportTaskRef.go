package awsiotwireless

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiotwireless/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a WirelessDeviceImportTask.
// Experimental.
type IWirelessDeviceImportTaskRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a WirelessDeviceImportTask resource.
	// Experimental.
	WirelessDeviceImportTaskRef() *WirelessDeviceImportTaskReference
}

// The jsii proxy for IWirelessDeviceImportTaskRef
type jsiiProxy_IWirelessDeviceImportTaskRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IWirelessDeviceImportTaskRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
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

