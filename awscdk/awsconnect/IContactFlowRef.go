package awsconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ContactFlow.
// Experimental.
type IContactFlowRef interface {
	constructs.IConstruct
	// A reference to a ContactFlow resource.
	// Experimental.
	ContactFlowRef() *ContactFlowReference
}

// The jsii proxy for IContactFlowRef
type jsiiProxy_IContactFlowRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IContactFlowRef) ContactFlowRef() *ContactFlowReference {
	var returns *ContactFlowReference
	_jsii_.Get(
		j,
		"contactFlowRef",
		&returns,
	)
	return returns
}

