package customresources

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/customresources/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Defines a custom resource that is materialized using specific AWS API calls.
//
// These calls are created using
// a singleton Lambda function.
//
// Use this to bridge any gap that might exist in the CloudFormation Coverage.
// You can specify exactly which calls are invoked for the 'CREATE', 'UPDATE' and 'DELETE' life cycle events.
//
// Example:
//   crossAccountRoleArn := "arn:aws:iam::OTHERACCOUNT:role/CrossAccountRoleName" // arn of role deployed in separate account
//
//   callRegion := "us-west-1" // sdk call to be made in specified region (optional)
//
//    // sdk call to be made in specified region (optional)
//   cr.NewAwsCustomResource(this, jsii.String("CrossAccount"), &AwsCustomResourceProps{
//   	OnCreate: &AwsSdkCall{
//   		AssumedRoleArn: crossAccountRoleArn,
//   		Region: callRegion,
//   		 // optional
//   		Service: jsii.String("sts"),
//   		Action: jsii.String("GetCallerIdentity"),
//   		PhysicalResourceId: cr.PhysicalResourceId_Of(jsii.String("id")),
//   	},
//   	Policy: cr.AwsCustomResourcePolicy_FromStatements([]PolicyStatement{
//   		iam.PolicyStatement_FromJson(map[string]*string{
//   			"Effect": jsii.String("Allow"),
//   			"Action": jsii.String("sts:AssumeRole"),
//   			"Resource": crossAccountRoleArn,
//   		}),
//   	}),
//   })
//
type AwsCustomResource interface {
	constructs.Construct
	awsiam.IGrantable
	// The principal to grant permissions to.
	GrantPrincipal() awsiam.IPrincipal
	// The tree node.
	Node() constructs.Node
	// Returns response data for the AWS SDK call as string.
	//
	// Example for S3 / listBucket : 'Buckets.0.Name'
	//
	// Note that you cannot use this method if `ignoreErrorCodesMatching`
	// is configured for any of the SDK calls. This is because in such a case,
	// the response data might not exist, and will cause a CloudFormation deploy time error.
	GetResponseField(dataPath *string) *string
	// Returns response data for the AWS SDK call.
	//
	// Example for S3 / listBucket : 'Buckets.0.Name'
	//
	// Use `Token.asXxx` to encode the returned `Reference` as a specific type or
	// use the convenience `getDataString` for string attributes.
	//
	// Note that you cannot use this method if `ignoreErrorCodesMatching`
	// is configured for any of the SDK calls. This is because in such a case,
	// the response data might not exist, and will cause a CloudFormation deploy time error.
	GetResponseFieldReference(dataPath *string) awscdk.Reference
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for AwsCustomResource
type jsiiProxy_AwsCustomResource struct {
	internal.Type__constructsConstruct
	internal.Type__awsiamIGrantable
}

func (j *jsiiProxy_AwsCustomResource) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomResource) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


func NewAwsCustomResource(scope constructs.Construct, id *string, props *AwsCustomResourceProps) AwsCustomResource {
	_init_.Initialize()

	if err := validateNewAwsCustomResourceParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsCustomResource{}

	_jsii_.Create(
		"aws-cdk-lib.custom_resources.AwsCustomResource",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewAwsCustomResource_Override(a AwsCustomResource, scope constructs.Construct, id *string, props *AwsCustomResourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.custom_resources.AwsCustomResource",
		[]interface{}{scope, id, props},
		a,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func AwsCustomResource_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsCustomResource_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.custom_resources.AwsCustomResource",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsCustomResource_PROPERTY_INJECTION_ID() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.custom_resources.AwsCustomResource",
		"PROPERTY_INJECTION_ID",
		&returns,
	)
	return returns
}

func AwsCustomResource_PROVIDER_FUNCTION_UUID() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.custom_resources.AwsCustomResource",
		"PROVIDER_FUNCTION_UUID",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsCustomResource) GetResponseField(dataPath *string) *string {
	if err := a.validateGetResponseFieldParameters(dataPath); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getResponseField",
		[]interface{}{dataPath},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCustomResource) GetResponseFieldReference(dataPath *string) awscdk.Reference {
	if err := a.validateGetResponseFieldReferenceParameters(dataPath); err != nil {
		panic(err)
	}
	var returns awscdk.Reference

	_jsii_.Invoke(
		a,
		"getResponseFieldReference",
		[]interface{}{dataPath},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCustomResource) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

