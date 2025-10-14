package awsgreengrass

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsgreengrass/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CoreDefinition.
// Experimental.
type ICoreDefinitionRef interface {
	constructs.IConstruct
	// A reference to a CoreDefinition resource.
	// Experimental.
	CoreDefinitionRef() *CoreDefinitionReference
}

// The jsii proxy for ICoreDefinitionRef
type jsiiProxy_ICoreDefinitionRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ICoreDefinitionRef) CoreDefinitionRef() *CoreDefinitionReference {
	var returns *CoreDefinitionReference
	_jsii_.Get(
		j,
		"coreDefinitionRef",
		&returns,
	)
	return returns
}

