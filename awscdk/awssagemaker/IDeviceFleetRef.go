package awssagemaker

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssagemaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DeviceFleet.
// Experimental.
type IDeviceFleetRef interface {
	constructs.IConstruct
}

// The jsii proxy for IDeviceFleetRef
type jsiiProxy_IDeviceFleetRef struct {
	internal.Type__constructsIConstruct
}

