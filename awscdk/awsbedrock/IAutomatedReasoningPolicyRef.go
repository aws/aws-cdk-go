package awsbedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsbedrock/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AutomatedReasoningPolicy.
// Experimental.
type IAutomatedReasoningPolicyRef interface {
	constructs.IConstruct
	// A reference to a AutomatedReasoningPolicy resource.
	// Experimental.
	AutomatedReasoningPolicyRef() *AutomatedReasoningPolicyReference
}

// The jsii proxy for IAutomatedReasoningPolicyRef
type jsiiProxy_IAutomatedReasoningPolicyRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IAutomatedReasoningPolicyRef) AutomatedReasoningPolicyRef() *AutomatedReasoningPolicyReference {
	var returns *AutomatedReasoningPolicyReference
	_jsii_.Get(
		j,
		"automatedReasoningPolicyRef",
		&returns,
	)
	return returns
}

