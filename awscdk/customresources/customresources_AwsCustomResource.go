package customresources

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/customresources/internal"
	"github.com/aws/constructs-go/constructs/v3"
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
//   awsCustom := cr.NewAwsCustomResource(this, jsii.String("aws-custom"), &awsCustomResourceProps{
//   	onCreate: &awsSdkCall{
//   		service: jsii.String("..."),
//   		action: jsii.String("..."),
//   		parameters: map[string]*string{
//   			"text": jsii.String("..."),
//   		},
//   		physicalResourceId: cr.physicalResourceId.of(jsii.String("...")),
//   	},
//   	onUpdate: &awsSdkCall{
//   		service: jsii.String("..."),
//   		action: jsii.String("..."),
//   		parameters: map[string]interface{}{
//   			"text": jsii.String("..."),
//   			"resourceId": cr.NewPhysicalResourceIdReference(),
//   		},
//   	},
//   	policy: cr.awsCustomResourcePolicy.fromSdkCalls(&sdkCallsPolicyOptions{
//   		resources: cr.*awsCustomResourcePolicy_ANY_RESOURCE(),
//   	}),
//   })
//
// Experimental.
type AwsCustomResource interface {
	awscdk.Construct
	awsiam.IGrantable
	// The principal to grant permissions to.
	// Experimental.
	GrantPrincipal() awsiam.IPrincipal
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Returns response data for the AWS SDK call as string.
	//
	// Example for S3 / listBucket : 'Buckets.0.Name'
	//
	// Note that you cannot use this method if `ignoreErrorCodesMatching`
	// is configured for any of the SDK calls. This is because in such a case,
	// the response data might not exist, and will cause a CloudFormation deploy time error.
	// Experimental.
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
	// Experimental.
	GetResponseFieldReference(dataPath *string) awscdk.Reference
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
}

// The jsii proxy struct for AwsCustomResource
type jsiiProxy_AwsCustomResource struct {
	internal.Type__awscdkConstruct
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

func (j *jsiiProxy_AwsCustomResource) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsCustomResource(scope constructs.Construct, id *string, props *AwsCustomResourceProps) AwsCustomResource {
	_init_.Initialize()

	if err := validateNewAwsCustomResourceParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsCustomResource{}

	_jsii_.Create(
		"monocdk.custom_resources.AwsCustomResource",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsCustomResource_Override(a AwsCustomResource, scope constructs.Construct, id *string, props *AwsCustomResourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.custom_resources.AwsCustomResource",
		[]interface{}{scope, id, props},
		a,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func AwsCustomResource_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsCustomResource_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.custom_resources.AwsCustomResource",
		"isConstruct",
		[]interface{}{x},
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

func (a *jsiiProxy_AwsCustomResource) OnPrepare() {
	_jsii_.InvokeVoid(
		a,
		"onPrepare",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCustomResource) OnSynthesize(session constructs.ISynthesisSession) {
	if err := a.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (a *jsiiProxy_AwsCustomResource) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCustomResource) Prepare() {
	_jsii_.InvokeVoid(
		a,
		"prepare",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCustomResource) Synthesize(session awscdk.ISynthesisSession) {
	if err := a.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"synthesize",
		[]interface{}{session},
	)
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

func (a *jsiiProxy_AwsCustomResource) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

