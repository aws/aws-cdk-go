package awsses

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsses/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ReceiptRuleSet.
// Experimental.
type IReceiptRuleSetRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a ReceiptRuleSet resource.
	// Experimental.
	ReceiptRuleSetRef() *ReceiptRuleSetReference
}

// The jsii proxy for IReceiptRuleSetRef
type jsiiProxy_IReceiptRuleSetRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IReceiptRuleSetRef) ReceiptRuleSetRef() *ReceiptRuleSetReference {
	var returns *ReceiptRuleSetReference
	_jsii_.Get(
		j,
		"receiptRuleSetRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IReceiptRuleSetRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IReceiptRuleSetRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

