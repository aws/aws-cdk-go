package awscloudformation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudformation/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ModuleDefaultVersion.
// Experimental.
type IModuleDefaultVersionRef interface {
	constructs.IConstruct
	// A reference to a ModuleDefaultVersion resource.
	// Experimental.
	ModuleDefaultVersionRef() *ModuleDefaultVersionReference
}

// The jsii proxy for IModuleDefaultVersionRef
type jsiiProxy_IModuleDefaultVersionRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IModuleDefaultVersionRef) ModuleDefaultVersionRef() *ModuleDefaultVersionReference {
	var returns *ModuleDefaultVersionReference
	_jsii_.Get(
		j,
		"moduleDefaultVersionRef",
		&returns,
	)
	return returns
}

