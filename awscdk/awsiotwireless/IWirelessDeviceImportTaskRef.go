package awsiotwireless

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiotwireless/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a WirelessDeviceImportTask.
// Experimental.
type IWirelessDeviceImportTaskRef interface {
	constructs.IConstruct
	// A reference to a WirelessDeviceImportTask resource.
	// Experimental.
	WirelessDeviceImportTaskRef() *WirelessDeviceImportTaskReference
}

// The jsii proxy for IWirelessDeviceImportTaskRef
type jsiiProxy_IWirelessDeviceImportTaskRef struct {
	internal.Type__constructsIConstruct
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

