package awsroute53recoverycontrol

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53recoverycontrol/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ControlPanel.
// Experimental.
type IControlPanelRef interface {
	constructs.IConstruct
	// A reference to a ControlPanel resource.
	// Experimental.
	ControlPanelRef() *ControlPanelReference
}

// The jsii proxy for IControlPanelRef
type jsiiProxy_IControlPanelRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IControlPanelRef) ControlPanelRef() *ControlPanelReference {
	var returns *ControlPanelReference
	_jsii_.Get(
		j,
		"controlPanelRef",
		&returns,
	)
	return returns
}

