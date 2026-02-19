package interfacesawsses

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsses/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ReceiptRule.
// Experimental.
type IReceiptRuleRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ReceiptRule resource.
	// Experimental.
	ReceiptRuleRef() *ReceiptRuleReference
}

// The jsii proxy for IReceiptRuleRef
type jsiiProxy_IReceiptRuleRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IReceiptRuleRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IReceiptRuleRef) ReceiptRuleRef() *ReceiptRuleReference {
	var returns *ReceiptRuleReference
	_jsii_.Get(
		j,
		"receiptRuleRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IReceiptRuleRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IReceiptRuleRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

