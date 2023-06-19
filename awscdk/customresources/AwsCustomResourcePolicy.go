package customresources

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

// The IAM Policy that will be applied to the different calls.
//
// Example:
//   awsCustom := cr.NewAwsCustomResource(this, jsii.String("aws-custom"), &AwsCustomResourceProps{
//   	OnCreate: &AwsSdkCall{
//   		Service: jsii.String("..."),
//   		Action: jsii.String("..."),
//   		Parameters: map[string]*string{
//   			"text": jsii.String("..."),
//   		},
//   		PhysicalResourceId: cr.PhysicalResourceId_Of(jsii.String("...")),
//   	},
//   	OnUpdate: &AwsSdkCall{
//   		Service: jsii.String("..."),
//   		Action: jsii.String("..."),
//   		Parameters: map[string]interface{}{
//   			"text": jsii.String("..."),
//   			"resourceId": cr.NewPhysicalResourceIdReference(),
//   		},
//   	},
//   	Policy: cr.AwsCustomResourcePolicy_FromSdkCalls(&SdkCallsPolicyOptions{
//   		Resources: cr.AwsCustomResourcePolicy_ANY_RESOURCE(),
//   	}),
//   })
//
// Experimental.
type AwsCustomResourcePolicy interface {
	// resources for auto-generated from SDK calls.
	// Experimental.
	Resources() *[]*string
	// statements for explicit policy.
	// Experimental.
	Statements() *[]awsiam.PolicyStatement
}

// The jsii proxy struct for AwsCustomResourcePolicy
type jsiiProxy_AwsCustomResourcePolicy struct {
	_ byte // padding
}

func (j *jsiiProxy_AwsCustomResourcePolicy) Resources() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomResourcePolicy) Statements() *[]awsiam.PolicyStatement {
	var returns *[]awsiam.PolicyStatement
	_jsii_.Get(
		j,
		"statements",
		&returns,
	)
	return returns
}


// Generate IAM Policy Statements from the configured SDK calls.
//
// Each SDK call with be translated to an IAM Policy Statement in the form of: `call.service:call.action` (e.g `s3:PutObject`).
//
// This policy generator assumes the IAM policy name has the same name as the API
// call. This is true in 99% of cases, but there are exceptions (for example,
// S3's `PutBucketLifecycleConfiguration` requires
// `s3:PutLifecycleConfiguration` permissions, Lambda's `Invoke` requires
// `lambda:InvokeFunction` permissions). Use `fromStatements` if you want to
// do a call that requires different IAM action names.
// Experimental.
func AwsCustomResourcePolicy_FromSdkCalls(options *SdkCallsPolicyOptions) AwsCustomResourcePolicy {
	_init_.Initialize()

	if err := validateAwsCustomResourcePolicy_FromSdkCallsParameters(options); err != nil {
		panic(err)
	}
	var returns AwsCustomResourcePolicy

	_jsii_.StaticInvoke(
		"monocdk.custom_resources.AwsCustomResourcePolicy",
		"fromSdkCalls",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Explicit IAM Policy Statements.
// Experimental.
func AwsCustomResourcePolicy_FromStatements(statements *[]awsiam.PolicyStatement) AwsCustomResourcePolicy {
	_init_.Initialize()

	if err := validateAwsCustomResourcePolicy_FromStatementsParameters(statements); err != nil {
		panic(err)
	}
	var returns AwsCustomResourcePolicy

	_jsii_.StaticInvoke(
		"monocdk.custom_resources.AwsCustomResourcePolicy",
		"fromStatements",
		[]interface{}{statements},
		&returns,
	)

	return returns
}

func AwsCustomResourcePolicy_ANY_RESOURCE() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"monocdk.custom_resources.AwsCustomResourcePolicy",
		"ANY_RESOURCE",
		&returns,
	)
	return returns
}

