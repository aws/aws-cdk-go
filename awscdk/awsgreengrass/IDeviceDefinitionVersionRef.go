package awsgreengrass

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsgreengrass/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DeviceDefinitionVersion.
// Experimental.
type IDeviceDefinitionVersionRef interface {
	constructs.IConstruct
}

// The jsii proxy for IDeviceDefinitionVersionRef
type jsiiProxy_IDeviceDefinitionVersionRef struct {
	internal.Type__constructsIConstruct
}

