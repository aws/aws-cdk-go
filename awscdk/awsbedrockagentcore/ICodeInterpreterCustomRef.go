package awsbedrockagentcore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsbedrockagentcore/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CodeInterpreterCustom.
// Experimental.
type ICodeInterpreterCustomRef interface {
	constructs.IConstruct
	// A reference to a CodeInterpreterCustom resource.
	// Experimental.
	CodeInterpreterCustomRef() *CodeInterpreterCustomReference
}

// The jsii proxy for ICodeInterpreterCustomRef
type jsiiProxy_ICodeInterpreterCustomRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ICodeInterpreterCustomRef) CodeInterpreterCustomRef() *CodeInterpreterCustomReference {
	var returns *CodeInterpreterCustomReference
	_jsii_.Get(
		j,
		"codeInterpreterCustomRef",
		&returns,
	)
	return returns
}

