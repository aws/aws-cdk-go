package awsamplifyuibuilder

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsamplifyuibuilder/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Form.
// Experimental.
type IFormRef interface {
	constructs.IConstruct
}

// The jsii proxy for IFormRef
type jsiiProxy_IFormRef struct {
	internal.Type__constructsIConstruct
}

