package awsgreengrass

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsgreengrass/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CoreDefinition.
// Experimental.
type ICoreDefinitionRef interface {
	constructs.IConstruct
}

// The jsii proxy for ICoreDefinitionRef
type jsiiProxy_ICoreDefinitionRef struct {
	internal.Type__constructsIConstruct
}

