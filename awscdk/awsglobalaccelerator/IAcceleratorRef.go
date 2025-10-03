package awsglobalaccelerator

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsglobalaccelerator/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Accelerator.
// Experimental.
type IAcceleratorRef interface {
	constructs.IConstruct
}

// The jsii proxy for IAcceleratorRef
type jsiiProxy_IAcceleratorRef struct {
	internal.Type__constructsIConstruct
}

