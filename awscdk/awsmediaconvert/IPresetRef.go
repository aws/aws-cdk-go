package awsmediaconvert

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsmediaconvert/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Preset.
// Experimental.
type IPresetRef interface {
	constructs.IConstruct
}

// The jsii proxy for IPresetRef
type jsiiProxy_IPresetRef struct {
	internal.Type__constructsIConstruct
}

