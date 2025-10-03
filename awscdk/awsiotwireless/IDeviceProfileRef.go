package awsiotwireless

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiotwireless/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DeviceProfile.
// Experimental.
type IDeviceProfileRef interface {
	constructs.IConstruct
}

// The jsii proxy for IDeviceProfileRef
type jsiiProxy_IDeviceProfileRef struct {
	internal.Type__constructsIConstruct
}

