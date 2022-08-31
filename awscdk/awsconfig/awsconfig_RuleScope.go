package awsconfig

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
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
// Experimental.
type RuleScope interface {
	// tag key applied to resources that will trigger evaluation of a rule.
	// Experimental.
	Key() *string
	// ID of the only AWS resource that will trigger evaluation of a rule.
	// Experimental.
	ResourceId() *string
	// Resource types that will trigger evaluation of a rule.
	// Experimental.
	ResourceTypes() *[]ResourceType
	// tag value applied to resources that will trigger evaluation of a rule.
	// Experimental.
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
// Experimental.
func RuleScope_FromResource(resourceType ResourceType, resourceId *string) RuleScope {
	_init_.Initialize()

	if err := validateRuleScope_FromResourceParameters(resourceType); err != nil {
		panic(err)
	}
	var returns RuleScope

	_jsii_.StaticInvoke(
		"monocdk.aws_config.RuleScope",
		"fromResource",
		[]interface{}{resourceType, resourceId},
		&returns,
	)

	return returns
}

// restricts scope of changes to specific resource types.
// Experimental.
func RuleScope_FromResources(resourceTypes *[]ResourceType) RuleScope {
	_init_.Initialize()

	if err := validateRuleScope_FromResourcesParameters(resourceTypes); err != nil {
		panic(err)
	}
	var returns RuleScope

	_jsii_.StaticInvoke(
		"monocdk.aws_config.RuleScope",
		"fromResources",
		[]interface{}{resourceTypes},
		&returns,
	)

	return returns
}

// restricts scope of changes to a specific tag.
// Experimental.
func RuleScope_FromTag(key *string, value *string) RuleScope {
	_init_.Initialize()

	if err := validateRuleScope_FromTagParameters(key); err != nil {
		panic(err)
	}
	var returns RuleScope

	_jsii_.StaticInvoke(
		"monocdk.aws_config.RuleScope",
		"fromTag",
		[]interface{}{key, value},
		&returns,
	)

	return returns
}

