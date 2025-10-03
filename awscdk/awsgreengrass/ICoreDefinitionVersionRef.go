package awsgreengrass

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsgreengrass/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CoreDefinitionVersion.
// Experimental.
type ICoreDefinitionVersionRef interface {
	constructs.IConstruct
}

// The jsii proxy for ICoreDefinitionVersionRef
type jsiiProxy_ICoreDefinitionVersionRef struct {
	internal.Type__constructsIConstruct
}

