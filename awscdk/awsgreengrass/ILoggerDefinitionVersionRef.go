package awsgreengrass

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsgreengrass/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LoggerDefinitionVersion.
// Experimental.
type ILoggerDefinitionVersionRef interface {
	constructs.IConstruct
}

// The jsii proxy for ILoggerDefinitionVersionRef
type jsiiProxy_ILoggerDefinitionVersionRef struct {
	internal.Type__constructsIConstruct
}

