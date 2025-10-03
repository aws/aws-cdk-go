package awscloudformation

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudformation/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ModuleVersion.
// Experimental.
type IModuleVersionRef interface {
	constructs.IConstruct
}

// The jsii proxy for IModuleVersionRef
type jsiiProxy_IModuleVersionRef struct {
	internal.Type__constructsIConstruct
}

