package awsconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ContactFlowModule.
// Experimental.
type IContactFlowModuleRef interface {
	constructs.IConstruct
	// A reference to a ContactFlowModule resource.
	// Experimental.
	ContactFlowModuleRef() *ContactFlowModuleReference
}

// The jsii proxy for IContactFlowModuleRef
type jsiiProxy_IContactFlowModuleRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IContactFlowModuleRef) ContactFlowModuleRef() *ContactFlowModuleReference {
	var returns *ContactFlowModuleReference
	_jsii_.Get(
		j,
		"contactFlowModuleRef",
		&returns,
	)
	return returns
}

