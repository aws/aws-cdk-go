package awscomprehend

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscomprehend/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Flywheel.
// Experimental.
type IFlywheelRef interface {
	constructs.IConstruct
}

// The jsii proxy for IFlywheelRef
type jsiiProxy_IFlywheelRef struct {
	internal.Type__constructsIConstruct
}

