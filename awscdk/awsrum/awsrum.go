package awsrum

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsrum/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::RUM::AppMonitor`.
//
// Creates a CloudWatch RUM app monitor, which you can use to collect telemetry data from your application and send it to CloudWatch RUM. The data includes performance and reliability information such as page load time, client-side errors, and user behavior.
//
// After you create an app monitor, sign in to the CloudWatch RUM console to get the JavaScript code snippet to add to your web application. For more information, see [How do I find a code snippet that I've already generated?](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-RUM-find-code-snippet.html)
//
// TODO: EXAMPLE
//
type CfnAppMonitor interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AppMonitorConfiguration() interface{}
	SetAppMonitorConfiguration(val interface{})
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	CwLogEnabled() interface{}
	SetCwLogEnabled(val interface{})
	Domain() *string
	SetDomain(val *string)
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() constructs.Node
	Ref() *string
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
	UpdatedProperites() *map[string]interface{}
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnAppMonitor
type jsiiProxy_CfnAppMonitor struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnAppMonitor) AppMonitorConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"appMonitorConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAppMonitor) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAppMonitor) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAppMonitor) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAppMonitor) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAppMonitor) CwLogEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cwLogEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAppMonitor) Domain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAppMonitor) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAppMonitor) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAppMonitor) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAppMonitor) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAppMonitor) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAppMonitor) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAppMonitor) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::RUM::AppMonitor`.
func NewCfnAppMonitor(scope constructs.Construct, id *string, props *CfnAppMonitorProps) CfnAppMonitor {
	_init_.Initialize()

	j := jsiiProxy_CfnAppMonitor{}

	_jsii_.Create(
		"aws-cdk-lib.aws_rum.CfnAppMonitor",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::RUM::AppMonitor`.
func NewCfnAppMonitor_Override(c CfnAppMonitor, scope constructs.Construct, id *string, props *CfnAppMonitorProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_rum.CfnAppMonitor",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnAppMonitor) SetAppMonitorConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"appMonitorConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnAppMonitor) SetCwLogEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"cwLogEnabled",
		val,
	)
}

func (j *jsiiProxy_CfnAppMonitor) SetDomain(val *string) {
	_jsii_.Set(
		j,
		"domain",
		val,
	)
}

func (j *jsiiProxy_CfnAppMonitor) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnAppMonitor_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rum.CfnAppMonitor",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnAppMonitor_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rum.CfnAppMonitor",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnAppMonitor_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rum.CfnAppMonitor",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAppMonitor_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rum.CfnAppMonitor",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnAppMonitor) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
func (c *jsiiProxy_CfnAppMonitor) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnAppMonitor) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

// Adds an override to the synthesized CloudFormation resource.
//
// To add a
// property override, either use `addPropertyOverride` or prefix `path` with
// "Properties." (i.e. `Properties.TopicName`).
//
// If the override is nested, separate each nested level using a dot (.) in the path parameter.
// If there is an array as part of the nesting, specify the index in the path.
//
// To include a literal `.` in the property name, prefix with a `\`. In most
// programming languages you will need to write this as `"\\."` because the
// `\` itself will need to be escaped.
//
// For example,
// ```typescript
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
// ```
// would add the overrides
// ```json
// "Properties": {
//    "GlobalSecondaryIndexes": [
//      {
//        "Projection": {
//          "NonKeyAttributes": [ "myattribute" ]
//          ...
//        }
//        ...
//      },
//      {
//        "ProjectionType": "INCLUDE"
//        ...
//      },
//    ]
//    ...
// }
// ```
func (c *jsiiProxy_CfnAppMonitor) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnAppMonitor) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnAppMonitor) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (c *jsiiProxy_CfnAppMonitor) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
func (c *jsiiProxy_CfnAppMonitor) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnAppMonitor) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnAppMonitor) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnAppMonitor) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnAppMonitor) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
func (c *jsiiProxy_CfnAppMonitor) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
func (c *jsiiProxy_CfnAppMonitor) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAppMonitor) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// This structure contains much of the configuration data for the app monitor.
//
// TODO: EXAMPLE
//
type CfnAppMonitor_AppMonitorConfigurationProperty struct {
	// If you set this to `true` , the CloudWatch RUM web client sets two cookies, a session cookie and a user cookie.
	//
	// The cookies allow the CloudWatch RUM web client to collect data relating to the number of users an application has and the behavior of the application across a sequence of events. Cookies are stored in the top-level domain of the current page.
	AllowCookies interface{} `json:"allowCookies" yaml:"allowCookies"`
	// If you set this to `true` , CloudWatch RUM sends client-side traces to X-Ray for each sampled session.
	//
	// You can then see traces and segments from these user sessions in the RUM dashboard and the CloudWatch ServiceLens console. For more information, see [What is AWS X-Ray ?](https://docs.aws.amazon.com/xray/latest/devguide/aws-xray.html)
	EnableXRay interface{} `json:"enableXRay" yaml:"enableXRay"`
	// A list of URLs in your website or application to exclude from RUM data collection.
	//
	// You can't include both `ExcludedPages` and `IncludedPages` in the same app monitor.
	ExcludedPages *[]*string `json:"excludedPages" yaml:"excludedPages"`
	// A list of pages in your application that are to be displayed with a "favorite" icon in the CloudWatch RUM console.
	FavoritePages *[]*string `json:"favoritePages" yaml:"favoritePages"`
	// The ARN of the guest IAM role that is attached to the Amazon Cognito identity pool that is used to authorize the sending of data to CloudWatch RUM.
	GuestRoleArn *string `json:"guestRoleArn" yaml:"guestRoleArn"`
	// The ID of the Amazon Cognito identity pool that is used to authorize the sending of data to CloudWatch RUM.
	IdentityPoolId *string `json:"identityPoolId" yaml:"identityPoolId"`
	// If this app monitor is to collect data from only certain pages in your application, this structure lists those pages.
	//
	// You can't include both `ExcludedPages` and `IncludedPages` in the same app monitor.
	IncludedPages *[]*string `json:"includedPages" yaml:"includedPages"`
	// Specifies the portion of user sessions to use for CloudWatch RUM data collection.
	//
	// Choosing a higher portion gives you more data but also incurs more costs.
	//
	// The range for this value is 0 to 1 inclusive. Setting this to 1 means that 100% of user sessions are sampled, and setting it to 0.1 means that 10% of user sessions are sampled.
	//
	// If you omit this parameter, the default of 0.1 is used, and 10% of sessions will be sampled.
	SessionSampleRate *float64 `json:"sessionSampleRate" yaml:"sessionSampleRate"`
	// An array that lists the types of telemetry data that this app monitor is to collect.
	//
	// - `errors` indicates that RUM collects data about unhandled JavaScript errors raised by your application.
	// - `performance` indicates that RUM collects performance data about how your application and its resources are loaded and rendered. This includes Core Web Vitals.
	// - `http` indicates that RUM collects data about HTTP errors thrown by your application.
	Telemetries *[]*string `json:"telemetries" yaml:"telemetries"`
}

// Properties for defining a `CfnAppMonitor`.
//
// TODO: EXAMPLE
//
type CfnAppMonitorProps struct {
	// A structure that contains much of the configuration data for the app monitor.
	//
	// If you are using Amazon Cognito for authorization, you must include this structure in your request, and it must include the ID of the Amazon Cognito identity pool to use for authorization. If you don't include `AppMonitorConfiguration` , you must set up your own authorization method. For more information, see [Authorize your application to send data to AWS](https://docs.aws.amazon.com/monitoring/CloudWatch-RUM-get-started-authorization.html) .
	//
	// If you omit this argument, the sample rate used for CloudWatch RUM is set to 10% of the user sessions.
	AppMonitorConfiguration interface{} `json:"appMonitorConfiguration" yaml:"appMonitorConfiguration"`
	// Data collected by CloudWatch RUM is kept by RUM for 30 days and then deleted.
	//
	// This parameter specifies whether CloudWatch RUM sends a copy of this telemetry data to Amazon CloudWatch Logs in your account. This enables you to keep the telemetry data for more than 30 days, but it does incur Amazon CloudWatch Logs charges.
	//
	// If you omit this parameter, the default is `false` .
	CwLogEnabled interface{} `json:"cwLogEnabled" yaml:"cwLogEnabled"`
	// The top-level internet domain name for which your application has administrative authority.
	//
	// This parameter is required.
	Domain *string `json:"domain" yaml:"domain"`
	// A name for the app monitor.
	//
	// This parameter is required.
	Name *string `json:"name" yaml:"name"`
	// Assigns one or more tags (key-value pairs) to the app monitor.
	//
	// Tags can help you organize and categorize your resources. You can also use them to scope user permissions by granting a user permission to access or change only resources with certain tag values.
	//
	// Tags don't have any semantic meaning to AWS and are interpreted strictly as strings of characters.
	//
	// You can associate as many as 50 tags with an app monitor.
	//
	// For more information, see [Tagging AWS resources](https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) .
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
}

