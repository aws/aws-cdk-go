package awslambda

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CodeSigningConfig.
// Experimental.
type ICodeSigningConfigRef interface {
	constructs.IConstruct
}

// The jsii proxy for ICodeSigningConfigRef
type jsiiProxy_ICodeSigningConfigRef struct {
	internal.Type__constructsIConstruct
}

