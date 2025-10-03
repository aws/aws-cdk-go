package awsgreengrass

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsgreengrass/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DeviceDefinition.
// Experimental.
type IDeviceDefinitionRef interface {
	constructs.IConstruct
}

// The jsii proxy for IDeviceDefinitionRef
type jsiiProxy_IDeviceDefinitionRef struct {
	internal.Type__constructsIConstruct
}

