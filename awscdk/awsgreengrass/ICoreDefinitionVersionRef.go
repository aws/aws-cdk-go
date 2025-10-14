package awsgreengrass

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsgreengrass/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CoreDefinitionVersion.
// Experimental.
type ICoreDefinitionVersionRef interface {
	constructs.IConstruct
	// A reference to a CoreDefinitionVersion resource.
	// Experimental.
	CoreDefinitionVersionRef() *CoreDefinitionVersionReference
}

// The jsii proxy for ICoreDefinitionVersionRef
type jsiiProxy_ICoreDefinitionVersionRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ICoreDefinitionVersionRef) CoreDefinitionVersionRef() *CoreDefinitionVersionReference {
	var returns *CoreDefinitionVersionReference
	_jsii_.Get(
		j,
		"coreDefinitionVersionRef",
		&returns,
	)
	return returns
}

