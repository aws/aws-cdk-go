package awsconfig

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Determines which resources trigger an evaluation of an AWS Config rule.
//
// Example:
//   var evalComplianceFn function
//   sshRule := config.NewManagedRule(this, jsii.String("SSH"), &managedRuleProps{
//   	identifier: config.managedRuleIdentifiers_EC2_SECURITY_GROUPS_INCOMING_SSH_DISABLED(),
//   	ruleScope: config.ruleScope.fromResource(config.resourceType_EC2_SECURITY_GROUP(), jsii.String("sg-1234567890abcdefgh")),
//   })
//   customRule := config.NewCustomRule(this, jsii.String("Lambda"), &customRuleProps{
//   	lambdaFunction: evalComplianceFn,
//   	configurationChanges: jsii.Boolean(true),
//   	ruleScope: config.*ruleScope.fromResources([]*resourceType{
//   		config.*resourceType_CLOUDFORMATION_STACK(),
//   		config.*resourceType_S3_BUCKET(),
//   	}),
//   })
//
//   tagRule := config.NewCustomRule(this, jsii.String("CostCenterTagRule"), &customRuleProps{
//   	lambdaFunction: evalComplianceFn,
//   	configurationChanges: jsii.Boolean(true),
//   	ruleScope: config.*ruleScope.fromTag(jsii.String("Cost Center"), jsii.String("MyApp")),
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

	var returns RuleScope

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_config.RuleScope",
		"fromTag",
		[]interface{}{key, value},
		&returns,
	)

	return returns
}

