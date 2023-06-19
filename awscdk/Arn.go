package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Experimental.
type Arn interface {
}

// The jsii proxy struct for Arn
type jsiiProxy_Arn struct {
	_ byte // padding
}

// Extract the full resource name from an ARN.
//
// Necessary for resource names (paths) that may contain the separator, like
// `arn:aws:iam::111111111111:role/path/to/role/name`.
//
// Only works if we statically know the expected `resourceType` beforehand, since we're going
// to use that to split the string on ':<resourceType>/' (and take the right-hand side).
//
// We can't extract the 'resourceType' from the ARN at hand, because CloudFormation Expressions
// only allow literals in the 'separator' argument to `{ Fn::Split }`, and so it can't be
// `{ Fn::Select: [5, { Fn::Split: [':', ARN] }}`.
//
// Only necessary for ARN formats for which the type-name separator is `/`.
// Experimental.
func Arn_ExtractResourceName(arn *string, resourceType *string) *string {
	_init_.Initialize()

	if err := validateArn_ExtractResourceNameParameters(arn, resourceType); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"monocdk.Arn",
		"extractResourceName",
		[]interface{}{arn, resourceType},
		&returns,
	)

	return returns
}

// Creates an ARN from components.
//
// If `partition`, `region` or `account` are not specified, the stack's
// partition, region and account will be used.
//
// If any component is the empty string, an empty string will be inserted
// into the generated ARN at the location that component corresponds to.
//
// The ARN will be formatted as follows:
//
//    arn:{partition}:{service}:{region}:{account}:{resource}{sep}{resource-name}
//
// The required ARN pieces that are omitted will be taken from the stack that
// the 'scope' is attached to. If all ARN pieces are supplied, the supplied scope
// can be 'undefined'.
// Experimental.
func Arn_Format(components *ArnComponents, stack Stack) *string {
	_init_.Initialize()

	if err := validateArn_FormatParameters(components); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"monocdk.Arn",
		"format",
		[]interface{}{components, stack},
		&returns,
	)

	return returns
}

// Given an ARN, parses it and returns components.
//
// IF THE ARN IS A CONCRETE STRING...
//
// ...it will be parsed and validated. The separator (`sep`) will be set to '/'
// if the 6th component includes a '/', in which case, `resource` will be set
// to the value before the '/' and `resourceName` will be the rest. In case
// there is no '/', `resource` will be set to the 6th components and
// `resourceName` will be set to the rest of the string.
//
// IF THE ARN IS A TOKEN...
//
// ...it cannot be validated, since we don't have the actual value yet at the
// time of this function call. You will have to supply `sepIfToken` and
// whether or not ARNs of the expected format usually have resource names
// in order to parse it properly. The resulting `ArnComponents` object will
// contain tokens for the subexpressions of the ARN, not string literals.
//
// If the resource name could possibly contain the separator char, the actual
// resource name cannot be properly parsed. This only occurs if the separator
// char is '/', and happens for example for S3 object ARNs, IAM Role ARNs,
// IAM OIDC Provider ARNs, etc. To properly extract the resource name from a
// Tokenized ARN, you must know the resource type and call
// `Arn.extractResourceName`.
//
// Returns: an ArnComponents object which allows access to the various
// components of the ARN.
// Deprecated: use split instead.
func Arn_Parse(arn *string, sepIfToken *string, hasName *bool) *ArnComponents {
	_init_.Initialize()

	if err := validateArn_ParseParameters(arn); err != nil {
		panic(err)
	}
	var returns *ArnComponents

	_jsii_.StaticInvoke(
		"monocdk.Arn",
		"parse",
		[]interface{}{arn, sepIfToken, hasName},
		&returns,
	)

	return returns
}

// Splits the provided ARN into its components.
//
// Works both if 'arn' is a string like 'arn:aws:s3:::bucket',
// and a Token representing a dynamic CloudFormation expression
// (in which case the returned components will also be dynamic CloudFormation expressions,
// encoded as Tokens).
// Experimental.
func Arn_Split(arn *string, arnFormat ArnFormat) *ArnComponents {
	_init_.Initialize()

	if err := validateArn_SplitParameters(arn, arnFormat); err != nil {
		panic(err)
	}
	var returns *ArnComponents

	_jsii_.StaticInvoke(
		"monocdk.Arn",
		"split",
		[]interface{}{arn, arnFormat},
		&returns,
	)

	return returns
}

