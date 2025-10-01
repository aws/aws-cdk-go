package awscloudformation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudformation/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ModuleVersion.
// Experimental.
type IModuleVersionRef interface {
	constructs.IConstruct
	// A reference to a ModuleVersion resource.
	// Experimental.
	ModuleVersionRef() *ModuleVersionReference
}

// The jsii proxy for IModuleVersionRef
type jsiiProxy_IModuleVersionRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IModuleVersionRef) ModuleVersionRef() *ModuleVersionReference {
	var returns *ModuleVersionReference
	_jsii_.Get(
		j,
		"moduleVersionRef",
		&returns,
	)
	return returns
}

