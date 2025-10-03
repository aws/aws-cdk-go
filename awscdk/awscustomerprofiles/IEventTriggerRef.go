package awscustomerprofiles

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscustomerprofiles/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EventTrigger.
// Experimental.
type IEventTriggerRef interface {
	constructs.IConstruct
}

// The jsii proxy for IEventTriggerRef
type jsiiProxy_IEventTriggerRef struct {
	internal.Type__constructsIConstruct
}

