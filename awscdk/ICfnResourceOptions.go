// Version 2 of the AWS Cloud Development Kit library
package awscdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type ICfnResourceOptions interface {
	// A condition to associate with this resource.
	//
	// This means that only if the condition evaluates to 'true' when the stack
	// is deployed, the resource will be included. This is provided to allow CDK projects to produce legacy templates, but normally
	// there is no need to use it in CDK projects.
	Condition() CfnCondition
	SetCondition(c CfnCondition)
	// Associate the CreationPolicy attribute with a resource to prevent its status from reaching create complete until AWS CloudFormation receives a specified number of success signals or the timeout period is exceeded.
	//
	// To signal a
	// resource, you can use the cfn-signal helper script or SignalResource API. AWS CloudFormation publishes valid signals
	// to the stack events so that you track the number of signals sent.
	CreationPolicy() *CfnCreationPolicy
	SetCreationPolicy(c *CfnCreationPolicy)
	// With the DeletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted.
	//
	// You specify a DeletionPolicy attribute for each resource that you want to control. If a resource has no DeletionPolicy
	// attribute, AWS CloudFormation deletes the resource by default. Note that this capability also applies to update operations
	// that lead to resources being removed.
	DeletionPolicy() CfnDeletionPolicy
	SetDeletionPolicy(d CfnDeletionPolicy)
	// The description of this resource.
	//
	// Used for informational purposes only, is not processed in any way
	// (and stays with the CloudFormation template, is not passed to the underlying resource,
	// even if it does have a 'description' property).
	Description() *string
	SetDescription(d *string)
	// Metadata associated with the CloudFormation resource.
	//
	// This is not the same as the construct metadata which can be added
	// using construct.addMetadata(), but would not appear in the CloudFormation template automatically.
	Metadata() *map[string]interface{}
	SetMetadata(m *map[string]interface{})
	// Use the UpdatePolicy attribute to specify how AWS CloudFormation handles updates to the AWS::AutoScaling::AutoScalingGroup resource.
	//
	// AWS CloudFormation invokes one of three update policies depending on the type of change you make or whether a
	// scheduled action is associated with the Auto Scaling group.
	UpdatePolicy() *CfnUpdatePolicy
	SetUpdatePolicy(u *CfnUpdatePolicy)
	// Use the UpdateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy() CfnDeletionPolicy
	SetUpdateReplacePolicy(u CfnDeletionPolicy)
	// The version of this resource.
	//
	// Used only for custom CloudFormation resources.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cfn-customresource.html
	//
	Version() *string
	SetVersion(v *string)
}

// The jsii proxy for ICfnResourceOptions
type jsiiProxy_ICfnResourceOptions struct {
	_ byte // padding
}

func (j *jsiiProxy_ICfnResourceOptions) Condition() CfnCondition {
	var returns CfnCondition
	_jsii_.Get(
		j,
		"condition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICfnResourceOptions)SetCondition(val CfnCondition) {
	_jsii_.Set(
		j,
		"condition",
		val,
	)
}

func (j *jsiiProxy_ICfnResourceOptions) CreationPolicy() *CfnCreationPolicy {
	var returns *CfnCreationPolicy
	_jsii_.Get(
		j,
		"creationPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICfnResourceOptions)SetCreationPolicy(val *CfnCreationPolicy) {
	if err := j.validateSetCreationPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"creationPolicy",
		val,
	)
}

func (j *jsiiProxy_ICfnResourceOptions) DeletionPolicy() CfnDeletionPolicy {
	var returns CfnDeletionPolicy
	_jsii_.Get(
		j,
		"deletionPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICfnResourceOptions)SetDeletionPolicy(val CfnDeletionPolicy) {
	_jsii_.Set(
		j,
		"deletionPolicy",
		val,
	)
}

func (j *jsiiProxy_ICfnResourceOptions) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICfnResourceOptions)SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ICfnResourceOptions) Metadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICfnResourceOptions)SetMetadata(val *map[string]interface{}) {
	_jsii_.Set(
		j,
		"metadata",
		val,
	)
}

func (j *jsiiProxy_ICfnResourceOptions) UpdatePolicy() *CfnUpdatePolicy {
	var returns *CfnUpdatePolicy
	_jsii_.Get(
		j,
		"updatePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICfnResourceOptions)SetUpdatePolicy(val *CfnUpdatePolicy) {
	if err := j.validateSetUpdatePolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updatePolicy",
		val,
	)
}

func (j *jsiiProxy_ICfnResourceOptions) UpdateReplacePolicy() CfnDeletionPolicy {
	var returns CfnDeletionPolicy
	_jsii_.Get(
		j,
		"updateReplacePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICfnResourceOptions)SetUpdateReplacePolicy(val CfnDeletionPolicy) {
	_jsii_.Set(
		j,
		"updateReplacePolicy",
		val,
	)
}

func (j *jsiiProxy_ICfnResourceOptions) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICfnResourceOptions)SetVersion(val *string) {
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

