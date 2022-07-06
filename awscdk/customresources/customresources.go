package customresources

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscloudformation"
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/awslogs"
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

// The IAM Policy that will be applied to the different calls.
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

// Properties for AwsCustomResource.
//
// Note that at least onCreate, onUpdate or onDelete must be specified.
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
type AwsCustomResourceProps struct {
	// The policy that will be added to the execution role of the Lambda function implementing this custom resource provider.
	//
	// The custom resource also implements `iam.IGrantable`, making it possible
	// to use the `grantXxx()` methods.
	//
	// As this custom resource uses a singleton Lambda function, it's important
	// to note the that function's role will eventually accumulate the
	// permissions/grants from all resources.
	// See: Policy.fromSdkCalls
	//
	// Experimental.
	Policy AwsCustomResourcePolicy `field:"required" json:"policy" yaml:"policy"`
	// A name for the singleton Lambda function implementing this custom resource.
	//
	// The function name will remain the same after the first AwsCustomResource is created in a stack.
	// Experimental.
	FunctionName *string `field:"optional" json:"functionName" yaml:"functionName"`
	// Whether to install the latest AWS SDK v2. Allows to use the latest API calls documented at https://docs.aws.amazon.com/AWSJavaScriptSDK/latest/index.html.
	//
	// The installation takes around 60 seconds.
	// Experimental.
	InstallLatestAwsSdk *bool `field:"optional" json:"installLatestAwsSdk" yaml:"installLatestAwsSdk"`
	// The number of days log events of the singleton Lambda function implementing this custom resource are kept in CloudWatch Logs.
	// Experimental.
	LogRetention awslogs.RetentionDays `field:"optional" json:"logRetention" yaml:"logRetention"`
	// The AWS SDK call to make when the resource is created.
	// Experimental.
	OnCreate *AwsSdkCall `field:"optional" json:"onCreate" yaml:"onCreate"`
	// The AWS SDK call to make when the resource is deleted.
	// Experimental.
	OnDelete *AwsSdkCall `field:"optional" json:"onDelete" yaml:"onDelete"`
	// The AWS SDK call to make when the resource is updated.
	// Experimental.
	OnUpdate *AwsSdkCall `field:"optional" json:"onUpdate" yaml:"onUpdate"`
	// Cloudformation Resource type.
	// Experimental.
	ResourceType *string `field:"optional" json:"resourceType" yaml:"resourceType"`
	// The execution role for the singleton Lambda function implementing this custom resource provider.
	//
	// This role will apply to all `AwsCustomResource`
	// instances in the stack. The role must be assumable by the
	// `lambda.amazonaws.com` service principal.
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// The timeout for the singleton Lambda function implementing this custom resource.
	// Experimental.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
}

// An AWS SDK call.
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
type AwsSdkCall struct {
	// The service action to call.
	// See: https://docs.aws.amazon.com/AWSJavaScriptSDK/latest/index.html
	//
	// Experimental.
	Action *string `field:"required" json:"action" yaml:"action"`
	// The service to call.
	// See: https://docs.aws.amazon.com/AWSJavaScriptSDK/latest/index.html
	//
	// Experimental.
	Service *string `field:"required" json:"service" yaml:"service"`
	// API version to use for the service.
	// See: https://docs.aws.amazon.com/sdk-for-javascript/v2/developer-guide/locking-api-versions.html
	//
	// Experimental.
	ApiVersion *string `field:"optional" json:"apiVersion" yaml:"apiVersion"`
	// Used for running the SDK calls in underlying lambda with a different role Can be used primarily for cross-account requests to for example connect hostedzone with a shared vpc.
	//
	// Example for Route53 / associateVPCWithHostedZone.
	// Experimental.
	AssumedRoleArn *string `field:"optional" json:"assumedRoleArn" yaml:"assumedRoleArn"`
	// The regex pattern to use to catch API errors.
	//
	// The `code` property of the
	// `Error` object will be tested against this pattern. If there is a match an
	// error will not be thrown.
	// Experimental.
	IgnoreErrorCodesMatching *string `field:"optional" json:"ignoreErrorCodesMatching" yaml:"ignoreErrorCodesMatching"`
	// Restrict the data returned by the custom resource to a specific path in the API response.
	//
	// Use this to limit the data returned by the custom
	// resource if working with API calls that could potentially result in custom
	// response objects exceeding the hard limit of 4096 bytes.
	//
	// Example for ECS / updateService: 'service.deploymentConfiguration.maximumPercent'
	// Deprecated: use outputPaths instead.
	OutputPath *string `field:"optional" json:"outputPath" yaml:"outputPath"`
	// Restrict the data returned by the custom resource to specific paths in the API response.
	//
	// Use this to limit the data returned by the custom
	// resource if working with API calls that could potentially result in custom
	// response objects exceeding the hard limit of 4096 bytes.
	//
	// Example for ECS / updateService: ['service.deploymentConfiguration.maximumPercent']
	// Experimental.
	OutputPaths *[]*string `field:"optional" json:"outputPaths" yaml:"outputPaths"`
	// The parameters for the service action.
	// See: https://docs.aws.amazon.com/AWSJavaScriptSDK/latest/index.html
	//
	// Experimental.
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// The physical resource id of the custom resource for this call.
	//
	// Mandatory for onCreate or onUpdate calls.
	// Experimental.
	PhysicalResourceId PhysicalResourceId `field:"optional" json:"physicalResourceId" yaml:"physicalResourceId"`
	// The region to send service requests to.
	//
	// **Note: Cross-region operations are generally considered an anti-pattern.**
	// **Consider first deploying a stack in that region.**
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
}

// Physical ID of the custom resource.
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
type PhysicalResourceId interface {
	// Literal string to be used as the physical id.
	// Experimental.
	Id() *string
	// Path to a response data element to be used as the physical id.
	// Experimental.
	ResponsePath() *string
}

// The jsii proxy struct for PhysicalResourceId
type jsiiProxy_PhysicalResourceId struct {
	_ byte // padding
}

func (j *jsiiProxy_PhysicalResourceId) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PhysicalResourceId) ResponsePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responsePath",
		&returns,
	)
	return returns
}


// Extract the physical resource id from the path (dot notation) to the data in the API call response.
// Experimental.
func PhysicalResourceId_FromResponse(responsePath *string) PhysicalResourceId {
	_init_.Initialize()

	var returns PhysicalResourceId

	_jsii_.StaticInvoke(
		"monocdk.custom_resources.PhysicalResourceId",
		"fromResponse",
		[]interface{}{responsePath},
		&returns,
	)

	return returns
}

// Explicit physical resource id.
// Experimental.
func PhysicalResourceId_Of(id *string) PhysicalResourceId {
	_init_.Initialize()

	var returns PhysicalResourceId

	_jsii_.StaticInvoke(
		"monocdk.custom_resources.PhysicalResourceId",
		"of",
		[]interface{}{id},
		&returns,
	)

	return returns
}

// Reference to the physical resource id that can be passed to the AWS operation as a parameter.
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
type PhysicalResourceIdReference interface {
	awscdk.IResolvable
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// This may return an array with a single informational element indicating how
	// to get this property populated, if it was skipped for performance reasons.
	// Experimental.
	CreationStack() *[]*string
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_arg awscdk.IResolveContext) interface{}
	// toJSON serialization to replace `PhysicalResourceIdReference` with a magic string.
	// Experimental.
	ToJSON() *string
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PhysicalResourceIdReference
type jsiiProxy_PhysicalResourceIdReference struct {
	internal.Type__awscdkIResolvable
}

func (j *jsiiProxy_PhysicalResourceIdReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}


// Experimental.
func NewPhysicalResourceIdReference() PhysicalResourceIdReference {
	_init_.Initialize()

	j := jsiiProxy_PhysicalResourceIdReference{}

	_jsii_.Create(
		"monocdk.custom_resources.PhysicalResourceIdReference",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewPhysicalResourceIdReference_Override(p PhysicalResourceIdReference) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.custom_resources.PhysicalResourceIdReference",
		nil, // no parameters
		p,
	)
}

func (p *jsiiProxy_PhysicalResourceIdReference) Resolve(_arg awscdk.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{_arg},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PhysicalResourceIdReference) ToJSON() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toJSON",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PhysicalResourceIdReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Defines an AWS CloudFormation custom resource provider.
//
// Example:
//   var onEvent function
//   var isComplete function
//   var myRole role
//
//   myProvider := cr.NewProvider(this, jsii.String("MyProvider"), &providerProps{
//   	onEventHandler: onEvent,
//   	isCompleteHandler: isComplete,
//   	logRetention: logs.retentionDays_ONE_DAY,
//   	role: myRole,
//   	providerFunctionName: jsii.String("the-lambda-name"),
//   })
//
// Experimental.
type Provider interface {
	awscdk.Construct
	awscloudformation.ICustomResourceProvider
	// The user-defined AWS Lambda function which is invoked asynchronously in order to determine if the operation is complete.
	// Experimental.
	IsCompleteHandler() awslambda.IFunction
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// The user-defined AWS Lambda function which is invoked for all resource lifecycle operations (CREATE/UPDATE/DELETE).
	// Experimental.
	OnEventHandler() awslambda.IFunction
	// The service token to use in order to define custom resources that are backed by this provider.
	// Experimental.
	ServiceToken() *string
	// Called by `CustomResource` which uses this provider.
	// Deprecated: use `provider.serviceToken` instead
	Bind(_scope awscdk.Construct) *awscloudformation.CustomResourceProviderConfig
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

// The jsii proxy struct for Provider
type jsiiProxy_Provider struct {
	internal.Type__awscdkConstruct
	internal.Type__awscloudformationICustomResourceProvider
}

func (j *jsiiProxy_Provider) IsCompleteHandler() awslambda.IFunction {
	var returns awslambda.IFunction
	_jsii_.Get(
		j,
		"isCompleteHandler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Provider) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Provider) OnEventHandler() awslambda.IFunction {
	var returns awslambda.IFunction
	_jsii_.Get(
		j,
		"onEventHandler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Provider) ServiceToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceToken",
		&returns,
	)
	return returns
}


// Experimental.
func NewProvider(scope constructs.Construct, id *string, props *ProviderProps) Provider {
	_init_.Initialize()

	j := jsiiProxy_Provider{}

	_jsii_.Create(
		"monocdk.custom_resources.Provider",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewProvider_Override(p Provider, scope constructs.Construct, id *string, props *ProviderProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.custom_resources.Provider",
		[]interface{}{scope, id, props},
		p,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func Provider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.custom_resources.Provider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Provider) Bind(_scope awscdk.Construct) *awscloudformation.CustomResourceProviderConfig {
	var returns *awscloudformation.CustomResourceProviderConfig

	_jsii_.Invoke(
		p,
		"bind",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Provider) OnPrepare() {
	_jsii_.InvokeVoid(
		p,
		"onPrepare",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Provider) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		p,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (p *jsiiProxy_Provider) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Provider) Prepare() {
	_jsii_.InvokeVoid(
		p,
		"prepare",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Provider) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		p,
		"synthesize",
		[]interface{}{session},
	)
}

func (p *jsiiProxy_Provider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Provider) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Initialization properties for the `Provider` construct.
//
// Example:
//   var onEvent function
//   var isComplete function
//   var myRole role
//
//   myProvider := cr.NewProvider(this, jsii.String("MyProvider"), &providerProps{
//   	onEventHandler: onEvent,
//   	isCompleteHandler: isComplete,
//   	logRetention: logs.retentionDays_ONE_DAY,
//   	role: myRole,
//   	providerFunctionName: jsii.String("the-lambda-name"),
//   })
//
// Experimental.
type ProviderProps struct {
	// The AWS Lambda function to invoke for all resource lifecycle operations (CREATE/UPDATE/DELETE).
	//
	// This function is responsible to begin the requested resource operation
	// (CREATE/UPDATE/DELETE) and return any additional properties to add to the
	// event, which will later be passed to `isComplete`. The `PhysicalResourceId`
	// property must be included in the response.
	// Experimental.
	OnEventHandler awslambda.IFunction `field:"required" json:"onEventHandler" yaml:"onEventHandler"`
	// The AWS Lambda function to invoke in order to determine if the operation is complete.
	//
	// This function will be called immediately after `onEvent` and then
	// periodically based on the configured query interval as long as it returns
	// `false`. If the function still returns `false` and the alloted timeout has
	// passed, the operation will fail.
	// Experimental.
	IsCompleteHandler awslambda.IFunction `field:"optional" json:"isCompleteHandler" yaml:"isCompleteHandler"`
	// The number of days framework log events are kept in CloudWatch Logs.
	//
	// When
	// updating this property, unsetting it doesn't remove the log retention policy.
	// To remove the retention policy, set the value to `INFINITE`.
	// Experimental.
	LogRetention awslogs.RetentionDays `field:"optional" json:"logRetention" yaml:"logRetention"`
	// Provider Lambda name.
	//
	// The provider lambda function name.
	// Experimental.
	ProviderFunctionName *string `field:"optional" json:"providerFunctionName" yaml:"providerFunctionName"`
	// Time between calls to the `isComplete` handler which determines if the resource has been stabilized.
	//
	// The first `isComplete` will be called immediately after `handler` and then
	// every `queryInterval` seconds, and until `timeout` has been reached or until
	// `isComplete` returns `true`.
	// Experimental.
	QueryInterval awscdk.Duration `field:"optional" json:"queryInterval" yaml:"queryInterval"`
	// AWS Lambda execution role.
	//
	// The role that will be assumed by the AWS Lambda.
	// Must be assumable by the 'lambda.amazonaws.com' service principal.
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// Security groups to attach to the provider functions.
	//
	// Only used if 'vpc' is supplied.
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Total timeout for the entire operation.
	//
	// The maximum timeout is 2 hours (yes, it can exceed the AWS Lambda 15 minutes).
	// Experimental.
	TotalTimeout awscdk.Duration `field:"optional" json:"totalTimeout" yaml:"totalTimeout"`
	// The vpc to provision the lambda functions in.
	// Experimental.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
	// Which subnets from the VPC to place the lambda functions in.
	//
	// Only used if 'vpc' is supplied. Note: internet access for Lambdas
	// requires a NAT gateway, so picking Public subnets is not allowed.
	// Experimental.
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
}

// Options for the auto-generation of policies based on the configured SDK calls.
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
type SdkCallsPolicyOptions struct {
	// The resources that the calls will have access to.
	//
	// It is best to use specific resource ARN's when possible. However, you can also use `AwsCustomResourcePolicy.ANY_RESOURCE`
	// to allow access to all resources. For example, when `onCreate` is used to create a resource which you don't
	// know the physical name of in advance.
	//
	// Note that will apply to ALL SDK calls.
	// Experimental.
	Resources *[]*string `field:"required" json:"resources" yaml:"resources"`
}

