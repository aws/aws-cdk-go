package awscdk

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ModuleDefaultVersion.
// Experimental.
type IModuleDefaultVersionRef interface {
	constructs.IConstruct
}

// The jsii proxy for IModuleDefaultVersionRef
type jsiiProxy_IModuleDefaultVersionRef struct {
	internal.Type__constructsIConstruct
}

