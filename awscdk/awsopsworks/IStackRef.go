package awsopsworks

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsopsworks/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Stack.
// Experimental.
type IStackRef interface {
	constructs.IConstruct
	// A reference to a Stack resource.
	// Experimental.
	StackRef() *StackReference
}

// The jsii proxy for IStackRef
type jsiiProxy_IStackRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IStackRef) StackRef() *StackReference {
	var returns *StackReference
	_jsii_.Get(
		j,
		"stackRef",
		&returns,
	)
	return returns
}

