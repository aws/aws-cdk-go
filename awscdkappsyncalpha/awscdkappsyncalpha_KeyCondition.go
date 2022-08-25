// The CDK Construct Library for AWS::AppSync
package awscdkappsyncalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkappsyncalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Factory class for DynamoDB key conditions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import appsync_alpha "github.com/aws/aws-cdk-go/awscdkappsyncalpha"
//
//   keyCondition := appsync_alpha.keyCondition.beginsWith(jsii.String("keyName"), jsii.String("arg"))
//
// Experimental.
type KeyCondition interface {
	// Conjunction between two conditions.
	// Experimental.
	And(keyCond KeyCondition) KeyCondition
	// Renders the key condition to a VTL string.
	// Experimental.
	RenderTemplate() *string
}

// The jsii proxy struct for KeyCondition
type jsiiProxy_KeyCondition struct {
	_ byte // padding
}

// Condition (k, arg).
//
// True if the key attribute k begins with the Query argument.
// Experimental.
func KeyCondition_BeginsWith(keyName *string, arg *string) KeyCondition {
	_init_.Initialize()

	var returns KeyCondition

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-appsync-alpha.KeyCondition",
		"beginsWith",
		[]interface{}{keyName, arg},
		&returns,
	)

	return returns
}

// Condition k BETWEEN arg1 AND arg2, true if k >= arg1 and k <= arg2.
// Experimental.
func KeyCondition_Between(keyName *string, arg1 *string, arg2 *string) KeyCondition {
	_init_.Initialize()

	var returns KeyCondition

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-appsync-alpha.KeyCondition",
		"between",
		[]interface{}{keyName, arg1, arg2},
		&returns,
	)

	return returns
}

// Condition k = arg, true if the key attribute k is equal to the Query argument.
// Experimental.
func KeyCondition_Eq(keyName *string, arg *string) KeyCondition {
	_init_.Initialize()

	var returns KeyCondition

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-appsync-alpha.KeyCondition",
		"eq",
		[]interface{}{keyName, arg},
		&returns,
	)

	return returns
}

// Condition k >= arg, true if the key attribute k is greater or equal to the Query argument.
// Experimental.
func KeyCondition_Ge(keyName *string, arg *string) KeyCondition {
	_init_.Initialize()

	var returns KeyCondition

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-appsync-alpha.KeyCondition",
		"ge",
		[]interface{}{keyName, arg},
		&returns,
	)

	return returns
}

// Condition k > arg, true if the key attribute k is greater than the the Query argument.
// Experimental.
func KeyCondition_Gt(keyName *string, arg *string) KeyCondition {
	_init_.Initialize()

	var returns KeyCondition

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-appsync-alpha.KeyCondition",
		"gt",
		[]interface{}{keyName, arg},
		&returns,
	)

	return returns
}

// Condition k <= arg, true if the key attribute k is less than or equal to the Query argument.
// Experimental.
func KeyCondition_Le(keyName *string, arg *string) KeyCondition {
	_init_.Initialize()

	var returns KeyCondition

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-appsync-alpha.KeyCondition",
		"le",
		[]interface{}{keyName, arg},
		&returns,
	)

	return returns
}

// Condition k < arg, true if the key attribute k is less than the Query argument.
// Experimental.
func KeyCondition_Lt(keyName *string, arg *string) KeyCondition {
	_init_.Initialize()

	var returns KeyCondition

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-appsync-alpha.KeyCondition",
		"lt",
		[]interface{}{keyName, arg},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyCondition) And(keyCond KeyCondition) KeyCondition {
	var returns KeyCondition

	_jsii_.Invoke(
		k,
		"and",
		[]interface{}{keyCond},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyCondition) RenderTemplate() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"renderTemplate",
		nil, // no parameters
		&returns,
	)

	return returns
}

