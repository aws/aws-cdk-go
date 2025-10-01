package awsentityresolution

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsentityresolution/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PolicyStatement.
// Experimental.
type IPolicyStatementRef interface {
	constructs.IConstruct
	// A reference to a PolicyStatement resource.
	// Experimental.
	PolicyStatementRef() *PolicyStatementReference
}

// The jsii proxy for IPolicyStatementRef
type jsiiProxy_IPolicyStatementRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IPolicyStatementRef) PolicyStatementRef() *PolicyStatementReference {
	var returns *PolicyStatementReference
	_jsii_.Get(
		j,
		"policyStatementRef",
		&returns,
	)
	return returns
}

