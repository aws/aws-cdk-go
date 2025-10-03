package awssagemaker

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssagemaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Device.
// Experimental.
type IDeviceRef interface {
	constructs.IConstruct
}

// The jsii proxy for IDeviceRef
type jsiiProxy_IDeviceRef struct {
	internal.Type__constructsIConstruct
}

