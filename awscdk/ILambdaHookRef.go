package awscdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LambdaHook.
// Experimental.
type ILambdaHookRef interface {
	constructs.IConstruct
	// A reference to a LambdaHook resource.
	// Experimental.
	LambdaHookRef() *LambdaHookReference
}

// The jsii proxy for ILambdaHookRef
type jsiiProxy_ILambdaHookRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ILambdaHookRef) LambdaHookRef() *LambdaHookReference {
	var returns *LambdaHookReference
	_jsii_.Get(
		j,
		"lambdaHookRef",
		&returns,
	)
	return returns
}

