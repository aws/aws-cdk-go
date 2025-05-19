package awsappsync

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Factory class for DynamoDB key conditions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   keyCondition := awscdk.Aws_appsync.KeyCondition_BeginsWith(jsii.String("keyName"), jsii.String("arg"))
//
type KeyCondition interface {
	// Conjunction between two conditions.
	And(keyCond KeyCondition) KeyCondition
	// Renders the key condition to a VTL string.
	RenderTemplate() *string
}

// The jsii proxy struct for KeyCondition
type jsiiProxy_KeyCondition struct {
	_ byte // padding
}

// Condition (k, arg).
//
// True if the key attribute k begins with the Query argument.
func KeyCondition_BeginsWith(keyName *string, arg *string) KeyCondition {
	_init_.Initialize()

	if err := validateKeyCondition_BeginsWithParameters(keyName, arg); err != nil {
		panic(err)
	}
	var returns KeyCondition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.KeyCondition",
		"beginsWith",
		[]interface{}{keyName, arg},
		&returns,
	)

	return returns
}

// Condition k BETWEEN arg1 AND arg2, true if k >= arg1 and k <= arg2.
func KeyCondition_Between(keyName *string, arg1 *string, arg2 *string) KeyCondition {
	_init_.Initialize()

	if err := validateKeyCondition_BetweenParameters(keyName, arg1, arg2); err != nil {
		panic(err)
	}
	var returns KeyCondition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.KeyCondition",
		"between",
		[]interface{}{keyName, arg1, arg2},
		&returns,
	)

	return returns
}

// Condition k = arg, true if the key attribute k is equal to the Query argument.
func KeyCondition_Eq(keyName *string, arg *string) KeyCondition {
	_init_.Initialize()

	if err := validateKeyCondition_EqParameters(keyName, arg); err != nil {
		panic(err)
	}
	var returns KeyCondition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.KeyCondition",
		"eq",
		[]interface{}{keyName, arg},
		&returns,
	)

	return returns
}

// Condition k >= arg, true if the key attribute k is greater or equal to the Query argument.
func KeyCondition_Ge(keyName *string, arg *string) KeyCondition {
	_init_.Initialize()

	if err := validateKeyCondition_GeParameters(keyName, arg); err != nil {
		panic(err)
	}
	var returns KeyCondition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.KeyCondition",
		"ge",
		[]interface{}{keyName, arg},
		&returns,
	)

	return returns
}

// Condition k > arg, true if the key attribute k is greater than the the Query argument.
func KeyCondition_Gt(keyName *string, arg *string) KeyCondition {
	_init_.Initialize()

	if err := validateKeyCondition_GtParameters(keyName, arg); err != nil {
		panic(err)
	}
	var returns KeyCondition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.KeyCondition",
		"gt",
		[]interface{}{keyName, arg},
		&returns,
	)

	return returns
}

// Condition k <= arg, true if the key attribute k is less than or equal to the Query argument.
func KeyCondition_Le(keyName *string, arg *string) KeyCondition {
	_init_.Initialize()

	if err := validateKeyCondition_LeParameters(keyName, arg); err != nil {
		panic(err)
	}
	var returns KeyCondition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.KeyCondition",
		"le",
		[]interface{}{keyName, arg},
		&returns,
	)

	return returns
}

// Condition k < arg, true if the key attribute k is less than the Query argument.
func KeyCondition_Lt(keyName *string, arg *string) KeyCondition {
	_init_.Initialize()

	if err := validateKeyCondition_LtParameters(keyName, arg); err != nil {
		panic(err)
	}
	var returns KeyCondition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.KeyCondition",
		"lt",
		[]interface{}{keyName, arg},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyCondition) And(keyCond KeyCondition) KeyCondition {
	if err := k.validateAndParameters(keyCond); err != nil {
		panic(err)
	}
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

