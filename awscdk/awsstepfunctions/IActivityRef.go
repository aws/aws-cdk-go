package awsstepfunctions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Activity.
// Experimental.
type IActivityRef interface {
	constructs.IConstruct
	// A reference to a Activity resource.
	// Experimental.
	ActivityRef() *ActivityReference
}

// The jsii proxy for IActivityRef
type jsiiProxy_IActivityRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IActivityRef) ActivityRef() *ActivityReference {
	var returns *ActivityReference
	_jsii_.Get(
		j,
		"activityRef",
		&returns,
	)
	return returns
}

