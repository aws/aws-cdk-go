package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// CloudFormation intrinsic functions.
//
// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference.html
//
// Example:
//   var cfnTemplate cfnInclude
//
//   // mutating the rule
//   var myParameter cfnParameter
//
//   rule := cfnTemplate.GetRule(jsii.String("MyRule"))
//   rule.AddAssertion(core.Fn_ConditionContains([]*string{
//   	jsii.String("m1.small"),
//   }, myParameter.valueAsString), jsii.String("MyParameter has to be m1.small"))
//
type Fn interface {
}

// The jsii proxy struct for Fn
type jsiiProxy_Fn struct {
	_ byte // padding
}

// The intrinsic function ``Fn::Base64`` returns the Base64 representation of the input string.
//
// This function is typically used to pass encoded data to
// Amazon EC2 instances by way of the UserData property.
//
// Returns: a token represented as a string.
func Fn_Base64(data *string) *string {
	_init_.Initialize()

	if err := validateFn_Base64Parameters(data); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Fn",
		"base64",
		[]interface{}{data},
		&returns,
	)

	return returns
}

// The intrinsic function ``Fn::Cidr`` returns the specified Cidr address block.
//
// Returns: a token represented as a string.
func Fn_Cidr(ipBlock *string, count *float64, sizeMask *string) *[]*string {
	_init_.Initialize()

	if err := validateFn_CidrParameters(ipBlock, count); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Fn",
		"cidr",
		[]interface{}{ipBlock, count, sizeMask},
		&returns,
	)

	return returns
}

// Returns true if all the specified conditions evaluate to true, or returns false if any one of the conditions evaluates to false.
//
// ``Fn::And`` acts as
// an AND operator. The minimum number of conditions that you can include is
// 1.
//
// Returns: an FnCondition token.
func Fn_ConditionAnd(conditions ...ICfnConditionExpression) ICfnRuleConditionExpression {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range conditions {
		args = append(args, a)
	}

	var returns ICfnRuleConditionExpression

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Fn",
		"conditionAnd",
		args,
		&returns,
	)

	return returns
}

// Returns true if a specified string matches at least one value in a list of strings.
//
// Returns: an FnCondition token.
func Fn_ConditionContains(listOfStrings *[]*string, value *string) ICfnRuleConditionExpression {
	_init_.Initialize()

	if err := validateFn_ConditionContainsParameters(listOfStrings, value); err != nil {
		panic(err)
	}
	var returns ICfnRuleConditionExpression

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Fn",
		"conditionContains",
		[]interface{}{listOfStrings, value},
		&returns,
	)

	return returns
}

// Returns true if a specified string matches all values in a list.
//
// Returns: an FnCondition token.
func Fn_ConditionEachMemberEquals(listOfStrings *[]*string, value *string) ICfnRuleConditionExpression {
	_init_.Initialize()

	if err := validateFn_ConditionEachMemberEqualsParameters(listOfStrings, value); err != nil {
		panic(err)
	}
	var returns ICfnRuleConditionExpression

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Fn",
		"conditionEachMemberEquals",
		[]interface{}{listOfStrings, value},
		&returns,
	)

	return returns
}

// Returns true if each member in a list of strings matches at least one value in a second list of strings.
//
// Returns: an FnCondition token.
func Fn_ConditionEachMemberIn(stringsToCheck *[]*string, stringsToMatch *[]*string) ICfnRuleConditionExpression {
	_init_.Initialize()

	if err := validateFn_ConditionEachMemberInParameters(stringsToCheck, stringsToMatch); err != nil {
		panic(err)
	}
	var returns ICfnRuleConditionExpression

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Fn",
		"conditionEachMemberIn",
		[]interface{}{stringsToCheck, stringsToMatch},
		&returns,
	)

	return returns
}

// Compares if two values are equal.
//
// Returns true if the two values are equal
// or false if they aren't.
//
// Returns: an FnCondition token.
func Fn_ConditionEquals(lhs interface{}, rhs interface{}) ICfnRuleConditionExpression {
	_init_.Initialize()

	if err := validateFn_ConditionEqualsParameters(lhs, rhs); err != nil {
		panic(err)
	}
	var returns ICfnRuleConditionExpression

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Fn",
		"conditionEquals",
		[]interface{}{lhs, rhs},
		&returns,
	)

	return returns
}

// Returns one value if the specified condition evaluates to true and another value if the specified condition evaluates to false.
//
// Currently, AWS
// CloudFormation supports the ``Fn::If`` intrinsic function in the metadata
// attribute, update policy attribute, and property values in the Resources
// section and Outputs sections of a template. You can use the AWS::NoValue
// pseudo parameter as a return value to remove the corresponding property.
//
// Returns: an FnCondition token.
func Fn_ConditionIf(conditionId *string, valueIfTrue interface{}, valueIfFalse interface{}) ICfnRuleConditionExpression {
	_init_.Initialize()

	if err := validateFn_ConditionIfParameters(conditionId, valueIfTrue, valueIfFalse); err != nil {
		panic(err)
	}
	var returns ICfnRuleConditionExpression

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Fn",
		"conditionIf",
		[]interface{}{conditionId, valueIfTrue, valueIfFalse},
		&returns,
	)

	return returns
}

// Returns true for a condition that evaluates to false or returns false for a condition that evaluates to true.
//
// ``Fn::Not`` acts as a NOT operator.
//
// Returns: an FnCondition token.
func Fn_ConditionNot(condition ICfnConditionExpression) ICfnRuleConditionExpression {
	_init_.Initialize()

	if err := validateFn_ConditionNotParameters(condition); err != nil {
		panic(err)
	}
	var returns ICfnRuleConditionExpression

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Fn",
		"conditionNot",
		[]interface{}{condition},
		&returns,
	)

	return returns
}

// Returns true if any one of the specified conditions evaluate to true, or returns false if all of the conditions evaluates to false.
//
// ``Fn::Or`` acts
// as an OR operator. The minimum number of conditions that you can include is
// 1.
//
// Returns: an FnCondition token.
func Fn_ConditionOr(conditions ...ICfnConditionExpression) ICfnRuleConditionExpression {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range conditions {
		args = append(args, a)
	}

	var returns ICfnRuleConditionExpression

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Fn",
		"conditionOr",
		args,
		&returns,
	)

	return returns
}

// The intrinsic function ``Fn::FindInMap`` returns the value corresponding to keys in a two-level map that is declared in the Mappings section.
//
// Warning: do not use with lazy mappings as this function will not guarentee a lazy mapping to render in the template.
// Prefer to use `CfnMapping.findInMap` in general.
//
// Returns: a token represented as a string.
func Fn_FindInMap(mapName *string, topLevelKey *string, secondLevelKey *string, defaultValue *string) *string {
	_init_.Initialize()

	if err := validateFn_FindInMapParameters(mapName, topLevelKey, secondLevelKey); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Fn",
		"findInMap",
		[]interface{}{mapName, topLevelKey, secondLevelKey, defaultValue},
		&returns,
	)

	return returns
}

// The ``Fn::GetAtt`` intrinsic function returns the value of an attribute from a resource in the template.
//
// Returns: an IResolvable object.
func Fn_GetAtt(logicalNameOfResource *string, attributeName *string) IResolvable {
	_init_.Initialize()

	if err := validateFn_GetAttParameters(logicalNameOfResource, attributeName); err != nil {
		panic(err)
	}
	var returns IResolvable

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Fn",
		"getAtt",
		[]interface{}{logicalNameOfResource, attributeName},
		&returns,
	)

	return returns
}

// The intrinsic function ``Fn::GetAZs`` returns an array that lists Availability Zones for a specified region.
//
// Because customers have access to
// different Availability Zones, the intrinsic function ``Fn::GetAZs`` enables
// template authors to write templates that adapt to the calling user's
// access. That way you don't have to hard-code a full list of Availability
// Zones for a specified region.
//
// Returns: a token represented as a string array.
func Fn_GetAzs(region *string) *[]*string {
	_init_.Initialize()

	var returns *[]*string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Fn",
		"getAzs",
		[]interface{}{region},
		&returns,
	)

	return returns
}

// Like `Fn.importValue`, but import a list with a known length.
//
// If you explicitly want a list with an unknown length, call `Fn.split(',',
// Fn.importValue(exportName))`. See the documentation of `Fn.split` to read
// more about the limitations of using lists of unknown length.
//
// `Fn.importListValue(exportName, assumedLength)` is the same as
// `Fn.split(',', Fn.importValue(exportName), assumedLength)`,
// but easier to read and impossible to forget to pass `assumedLength`.
func Fn_ImportListValue(sharedValueToImport *string, assumedLength *float64, delimiter *string) *[]*string {
	_init_.Initialize()

	if err := validateFn_ImportListValueParameters(sharedValueToImport, assumedLength); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Fn",
		"importListValue",
		[]interface{}{sharedValueToImport, assumedLength, delimiter},
		&returns,
	)

	return returns
}

// The intrinsic function ``Fn::ImportValue`` returns the value of an output exported by another stack.
//
// You typically use this function to create
// cross-stack references. In the following example template snippets, Stack A
// exports VPC security group values and Stack B imports them.
//
// Returns: a token represented as a string.
func Fn_ImportValue(sharedValueToImport *string) *string {
	_init_.Initialize()

	if err := validateFn_ImportValueParameters(sharedValueToImport); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Fn",
		"importValue",
		[]interface{}{sharedValueToImport},
		&returns,
	)

	return returns
}

// The intrinsic function ``Fn::Join`` appends a set of values into a single value, separated by the specified delimiter.
//
// If a delimiter is the empty
// string, the set of values are concatenated with no delimiter.
//
// Returns: a token represented as a string.
func Fn_Join(delimiter *string, listOfValues *[]*string) *string {
	_init_.Initialize()

	if err := validateFn_JoinParameters(delimiter, listOfValues); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Fn",
		"join",
		[]interface{}{delimiter, listOfValues},
		&returns,
	)

	return returns
}

// The intrinsic function `Fn::Length` returns the number of elements within an array or an intrinsic function that returns an array.
func Fn_Len(array interface{}) *float64 {
	_init_.Initialize()

	if err := validateFn_LenParameters(array); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Fn",
		"len",
		[]interface{}{array},
		&returns,
	)

	return returns
}

// Given an url, parse the domain name.
func Fn_ParseDomainName(url *string) *string {
	_init_.Initialize()

	if err := validateFn_ParseDomainNameParameters(url); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Fn",
		"parseDomainName",
		[]interface{}{url},
		&returns,
	)

	return returns
}

// The ``Ref`` intrinsic function returns the value of the specified parameter or resource.
//
// Note that it doesn't validate the logicalName, it mainly serves parameter/resource reference defined in a ``CfnInclude`` template.
func Fn_Ref(logicalName *string) *string {
	_init_.Initialize()

	if err := validateFn_RefParameters(logicalName); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Fn",
		"ref",
		[]interface{}{logicalName},
		&returns,
	)

	return returns
}

// Returns all values for a specified parameter type.
//
// Returns: a token represented as a string array.
func Fn_RefAll(parameterType *string) *[]*string {
	_init_.Initialize()

	if err := validateFn_RefAllParameters(parameterType); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Fn",
		"refAll",
		[]interface{}{parameterType},
		&returns,
	)

	return returns
}

// The intrinsic function ``Fn::Select`` returns a single object from a list of objects by index.
//
// Returns: a token represented as a string.
func Fn_Select(index *float64, array *[]*string) *string {
	_init_.Initialize()

	if err := validateFn_SelectParameters(index, array); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Fn",
		"select",
		[]interface{}{index, array},
		&returns,
	)

	return returns
}

// Split a string token into a token list of string values.
//
// Specify the location of splits with a delimiter such as ',' (a comma).
// Renders to the `Fn::Split` intrinsic function.
//
// Lists with unknown lengths (default)
// -------------------------------------
//
// Since this function is used to work with deploy-time values, if `assumedLength`
// is not given the CDK cannot know the length of the resulting list at synthesis time.
// This brings the following restrictions:
//
// - You must use `Fn.select(i, list)` to pick elements out of the list (you must not use
//   `list[i]`).
// - You cannot add elements to the list, remove elements from the list,
//   combine two such lists together, or take a slice of the list.
// - You cannot pass the list to constructs that do any of the above.
//
// The only valid operation with such a tokenized list is to pass it unmodified to a
// CloudFormation Resource construct.
//
// Lists with assumed lengths
// --------------------------
//
// Pass `assumedLength` if you know the length of the list that will be
// produced by splitting. The actual list length at deploy time may be
// *longer* than the number you pass, but not *shorter*.
//
// The returned list will look like:
//
// ```
// [Fn.select(0, split), Fn.select(1, split), Fn.select(2, split), ...]
// ```
//
// The restrictions from the section "Lists with unknown lengths" will now be lifted,
// at the expense of having to know and fix the length of the list.
//
// Returns: a token represented as a string array.
func Fn_Split(delimiter *string, source *string, assumedLength *float64) *[]*string {
	_init_.Initialize()

	if err := validateFn_SplitParameters(delimiter, source); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Fn",
		"split",
		[]interface{}{delimiter, source, assumedLength},
		&returns,
	)

	return returns
}

// The intrinsic function ``Fn::Sub`` substitutes variables in an input string with values that you specify.
//
// In your templates, you can use this function
// to construct commands or outputs that include values that aren't available
// until you create or update a stack.
//
// Returns: a token represented as a string.
func Fn_Sub(body *string, variables *map[string]*string) *string {
	_init_.Initialize()

	if err := validateFn_SubParameters(body); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Fn",
		"sub",
		[]interface{}{body, variables},
		&returns,
	)

	return returns
}

// The `Fn::ToJsonString` intrinsic function converts an object or array to its corresponding JSON string.
func Fn_ToJsonString(object interface{}) *string {
	_init_.Initialize()

	if err := validateFn_ToJsonStringParameters(object); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Fn",
		"toJsonString",
		[]interface{}{object},
		&returns,
	)

	return returns
}

// Creates a token representing the ``Fn::Transform`` expression.
//
// Returns: a token representing the transform expression.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-transform.html
//
func Fn_Transform(macroName *string, parameters *map[string]interface{}) IResolvable {
	_init_.Initialize()

	if err := validateFn_TransformParameters(macroName, parameters); err != nil {
		panic(err)
	}
	var returns IResolvable

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Fn",
		"transform",
		[]interface{}{macroName, parameters},
		&returns,
	)

	return returns
}

// Returns an attribute value or list of values for a specific parameter and attribute.
//
// Returns: a token represented as a string.
func Fn_ValueOf(parameterOrLogicalId *string, attribute *string) *string {
	_init_.Initialize()

	if err := validateFn_ValueOfParameters(parameterOrLogicalId, attribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Fn",
		"valueOf",
		[]interface{}{parameterOrLogicalId, attribute},
		&returns,
	)

	return returns
}

// Returns a list of all attribute values for a given parameter type and attribute.
//
// Returns: a token represented as a string array.
func Fn_ValueOfAll(parameterType *string, attribute *string) *[]*string {
	_init_.Initialize()

	if err := validateFn_ValueOfAllParameters(parameterType, attribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Fn",
		"valueOfAll",
		[]interface{}{parameterType, attribute},
		&returns,
	)

	return returns
}

