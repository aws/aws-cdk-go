package appstagingsynthesizeralpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Staging Resource Factory interface.
//
// The function included in this class will be called by the synthesizer
// to create or reference an IStagingResources construct that has the necessary
// staging resources for the stack.
// Experimental.
type IStagingResourcesFactory interface {
	// Return an object that will manage staging resources for the given stack.
	//
	// This is called whenever the the `AppStagingSynthesizer` binds to a specific
	// stack, and allows selecting where the staging resources go.
	//
	// This method can choose to either create a new construct (perhaps a stack)
	// and return it, or reference an existing construct.
	// Experimental.
	ObtainStagingResources(stack awscdk.Stack, context *ObtainStagingResourcesContext) IStagingResources
}

// The jsii proxy for IStagingResourcesFactory
type jsiiProxy_IStagingResourcesFactory struct {
	_ byte // padding
}

func (i *jsiiProxy_IStagingResourcesFactory) ObtainStagingResources(stack awscdk.Stack, context *ObtainStagingResourcesContext) IStagingResources {
	if err := i.validateObtainStagingResourcesParameters(stack, context); err != nil {
		panic(err)
	}
	var returns IStagingResources

	_jsii_.Invoke(
		i,
		"obtainStagingResources",
		[]interface{}{stack, context},
		&returns,
	)

	return returns
}

