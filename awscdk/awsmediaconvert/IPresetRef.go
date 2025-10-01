package awsmediaconvert

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsmediaconvert/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Preset.
// Experimental.
type IPresetRef interface {
	constructs.IConstruct
	// A reference to a Preset resource.
	// Experimental.
	PresetRef() *PresetReference
}

// The jsii proxy for IPresetRef
type jsiiProxy_IPresetRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IPresetRef) PresetRef() *PresetReference {
	var returns *PresetReference
	_jsii_.Get(
		j,
		"presetRef",
		&returns,
	)
	return returns
}

