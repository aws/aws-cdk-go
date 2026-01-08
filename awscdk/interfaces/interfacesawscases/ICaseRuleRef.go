package interfacesawscases

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscases/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CaseRule.
// Experimental.
type ICaseRuleRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a CaseRule resource.
	// Experimental.
	CaseRuleRef() *CaseRuleReference
}

// The jsii proxy for ICaseRuleRef
type jsiiProxy_ICaseRuleRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_ICaseRuleRef) CaseRuleRef() *CaseRuleReference {
	var returns *CaseRuleReference
	_jsii_.Get(
		j,
		"caseRuleRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICaseRuleRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICaseRuleRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

