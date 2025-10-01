package awseks

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awseks/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Addon.
// Experimental.
type IAddonRef interface {
	constructs.IConstruct
	// A reference to a Addon resource.
	// Experimental.
	AddonRef() *AddonReference
}

// The jsii proxy for IAddonRef
type jsiiProxy_IAddonRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IAddonRef) AddonRef() *AddonReference {
	var returns *AddonReference
	_jsii_.Get(
		j,
		"addonRef",
		&returns,
	)
	return returns
}

