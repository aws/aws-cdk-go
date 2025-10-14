package awsiot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiot/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Command.
// Experimental.
type ICommandRef interface {
	constructs.IConstruct
	// A reference to a Command resource.
	// Experimental.
	CommandRef() *CommandReference
}

// The jsii proxy for ICommandRef
type jsiiProxy_ICommandRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ICommandRef) CommandRef() *CommandReference {
	var returns *CommandReference
	_jsii_.Get(
		j,
		"commandRef",
		&returns,
	)
	return returns
}

