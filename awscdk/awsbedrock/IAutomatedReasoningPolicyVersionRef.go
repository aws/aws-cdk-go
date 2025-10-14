package awsbedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsbedrock/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AutomatedReasoningPolicyVersion.
// Experimental.
type IAutomatedReasoningPolicyVersionRef interface {
	constructs.IConstruct
	// A reference to a AutomatedReasoningPolicyVersion resource.
	// Experimental.
	AutomatedReasoningPolicyVersionRef() *AutomatedReasoningPolicyVersionReference
}

// The jsii proxy for IAutomatedReasoningPolicyVersionRef
type jsiiProxy_IAutomatedReasoningPolicyVersionRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IAutomatedReasoningPolicyVersionRef) AutomatedReasoningPolicyVersionRef() *AutomatedReasoningPolicyVersionReference {
	var returns *AutomatedReasoningPolicyVersionReference
	_jsii_.Get(
		j,
		"automatedReasoningPolicyVersionRef",
		&returns,
	)
	return returns
}

