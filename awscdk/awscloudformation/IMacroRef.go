package awscloudformation

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudformation/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Macro.
// Experimental.
type IMacroRef interface {
	constructs.IConstruct
}

// The jsii proxy for IMacroRef
type jsiiProxy_IMacroRef struct {
	internal.Type__constructsIConstruct
}

