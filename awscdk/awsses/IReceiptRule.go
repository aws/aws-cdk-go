package awsses

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsses/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsses"
	"github.com/aws/constructs-go/constructs/v10"
)

// A receipt rule.
type IReceiptRule interface {
	interfacesawsses.IReceiptRuleRef
	awscdk.IResource
	// The name of the receipt rule.
	ReceiptRuleName() *string
}

// The jsii proxy for IReceiptRule
type jsiiProxy_IReceiptRule struct {
	internal.Type__interfacesawssesIReceiptRuleRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IReceiptRule) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (j *jsiiProxy_IReceiptRule) ReceiptRuleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"receiptRuleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IReceiptRule) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IReceiptRule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IReceiptRule) ReceiptRuleRef() *interfacesawsses.ReceiptRuleReference {
	var returns *interfacesawsses.ReceiptRuleReference
	_jsii_.Get(
		j,
		"receiptRuleRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IReceiptRule) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

