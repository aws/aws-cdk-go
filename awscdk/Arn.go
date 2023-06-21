package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

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
func Arn_ExtractResourceName(arn *string, resourceType *string) *string {
	_init_.Initialize()

	if err := validateArn_ExtractResourceNameParameters(arn, resourceType); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Arn",
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
//   arn:{partition}:{service}:{region}:{account}:{resource}{sep}{resource-name}
//
// The required ARN pieces that are omitted will be taken from the stack that
// the 'scope' is attached to. If all ARN pieces are supplied, the supplied scope
// can be 'undefined'.
func Arn_Format(components *ArnComponents, stack Stack) *string {
	_init_.Initialize()

	if err := validateArn_FormatParameters(components); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Arn",
		"format",
		[]interface{}{components, stack},
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
func Arn_Split(arn *string, arnFormat ArnFormat) *ArnComponents {
	_init_.Initialize()

	if err := validateArn_SplitParameters(arn, arnFormat); err != nil {
		panic(err)
	}
	var returns *ArnComponents

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Arn",
		"split",
		[]interface{}{arn, arnFormat},
		&returns,
	)

	return returns
}

