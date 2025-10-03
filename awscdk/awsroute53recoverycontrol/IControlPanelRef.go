package awsroute53recoverycontrol

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53recoverycontrol/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ControlPanel.
// Experimental.
type IControlPanelRef interface {
	constructs.IConstruct
}

// The jsii proxy for IControlPanelRef
type jsiiProxy_IControlPanelRef struct {
	internal.Type__constructsIConstruct
}

