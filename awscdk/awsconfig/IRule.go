package awsconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsconfig/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsconfig"
	"github.com/aws/constructs-go/constructs/v10"
)

// Interface representing an AWS Config rule.
type IRule interface {
	interfacesawsconfig.IConfigRuleRef
	awscdk.IResource
	// Defines a EventBridge event rule which triggers for rule compliance events.
	OnComplianceChange(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Defines an EventBridge event rule which triggers for rule events.
	//
	// Use
	// `rule.addEventPattern(pattern)` to specify a filter.
	OnEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Defines a EventBridge event rule which triggers for rule re-evaluation status events.
	OnReEvaluationStatus(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// The name of the rule.
	ConfigRuleName() *string
}

// The jsii proxy for IRule
type jsiiProxy_IRule struct {
	internal.Type__interfacesawsconfigIConfigRuleRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IRule) OnComplianceChange(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	if err := i.validateOnComplianceChangeParameters(id, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		i,
		"onComplianceChange",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRule) OnEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	if err := i.validateOnEventParameters(id, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		i,
		"onEvent",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRule) OnReEvaluationStatus(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	if err := i.validateOnReEvaluationStatusParameters(id, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		i,
		"onReEvaluationStatus",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRule) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (j *jsiiProxy_IRule) ConfigRuleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configRuleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRule) ConfigRuleRef() *interfacesawsconfig.ConfigRuleReference {
	var returns *interfacesawsconfig.ConfigRuleReference
	_jsii_.Get(
		j,
		"configRuleRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRule) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRule) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

