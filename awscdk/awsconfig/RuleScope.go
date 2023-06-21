package awsconfig

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Determines which resources trigger an evaluation of an AWS Config rule.
//
// Example:
//   // Lambda function containing logic that evaluates compliance with the rule.
//   evalComplianceFn := lambda.NewFunction(this, jsii.String("CustomFunction"), &FunctionProps{
//   	Code: lambda.AssetCode_FromInline(jsii.String("exports.handler = (event) => console.log(event);")),
//   	Handler: jsii.String("index.handler"),
//   	Runtime: lambda.Runtime_NODEJS_18_X(),
//   })
//
//   // A custom rule that runs on configuration changes of EC2 instances
//   customRule := config.NewCustomRule(this, jsii.String("Custom"), &CustomRuleProps{
//   	ConfigurationChanges: jsii.Boolean(true),
//   	LambdaFunction: evalComplianceFn,
//   	RuleScope: config.RuleScope_FromResource(config.ResourceType_EC2_INSTANCE()),
//   })
//
type RuleScope interface {
	// tag key applied to resources that will trigger evaluation of a rule.
	Key() *string
	// ID of the only AWS resource that will trigger evaluation of a rule.
	ResourceId() *string
	// Resource types that will trigger evaluation of a rule.
	ResourceTypes() *[]ResourceType
	// tag value applied to resources that will trigger evaluation of a rule.
	Value() *string
}

// The jsii proxy struct for RuleScope
type jsiiProxy_RuleScope struct {
	_ byte // padding
}

func (j *jsiiProxy_RuleScope) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RuleScope) ResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RuleScope) ResourceTypes() *[]ResourceType {
	var returns *[]ResourceType
	_jsii_.Get(
		j,
		"resourceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RuleScope) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// restricts scope of changes to a specific resource type or resource identifier.
func RuleScope_FromResource(resourceType ResourceType, resourceId *string) RuleScope {
	_init_.Initialize()

	if err := validateRuleScope_FromResourceParameters(resourceType); err != nil {
		panic(err)
	}
	var returns RuleScope

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_config.RuleScope",
		"fromResource",
		[]interface{}{resourceType, resourceId},
		&returns,
	)

	return returns
}

// restricts scope of changes to specific resource types.
func RuleScope_FromResources(resourceTypes *[]ResourceType) RuleScope {
	_init_.Initialize()

	if err := validateRuleScope_FromResourcesParameters(resourceTypes); err != nil {
		panic(err)
	}
	var returns RuleScope

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_config.RuleScope",
		"fromResources",
		[]interface{}{resourceTypes},
		&returns,
	)

	return returns
}

// restricts scope of changes to a specific tag.
func RuleScope_FromTag(key *string, value *string) RuleScope {
	_init_.Initialize()

	if err := validateRuleScope_FromTagParameters(key); err != nil {
		panic(err)
	}
	var returns RuleScope

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_config.RuleScope",
		"fromTag",
		[]interface{}{key, value},
		&returns,
	)

	return returns
}

