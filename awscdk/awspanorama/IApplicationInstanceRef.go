package awspanorama

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awspanorama/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ApplicationInstance.
// Experimental.
type IApplicationInstanceRef interface {
	constructs.IConstruct
	// A reference to a ApplicationInstance resource.
	// Experimental.
	ApplicationInstanceRef() *ApplicationInstanceReference
}

// The jsii proxy for IApplicationInstanceRef
type jsiiProxy_IApplicationInstanceRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IApplicationInstanceRef) ApplicationInstanceRef() *ApplicationInstanceReference {
	var returns *ApplicationInstanceReference
	_jsii_.Get(
		j,
		"applicationInstanceRef",
		&returns,
	)
	return returns
}

