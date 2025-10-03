package awsiotwireless

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiotwireless/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a WirelessDeviceImportTask.
// Experimental.
type IWirelessDeviceImportTaskRef interface {
	constructs.IConstruct
}

// The jsii proxy for IWirelessDeviceImportTaskRef
type jsiiProxy_IWirelessDeviceImportTaskRef struct {
	internal.Type__constructsIConstruct
}

