package awswafregional

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awswafregional/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::WAFRegional::ByteMatchSet`.
//
// > This is *AWS WAF Classic* documentation. For more information, see [AWS WAF Classic](https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html) in the developer guide.
// >
// > *For the latest version of AWS WAF* , use the AWS WAF V2 API and see the [AWS WAF Developer Guide](https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html) . With the latest version, AWS WAF has a single set of endpoints for regional and global use.
//
// The `AWS::WAFRegional::ByteMatchSet` resource creates an AWS WAF `ByteMatchSet` that identifies a part of a web request that you want to inspect.
//
// TODO: EXAMPLE
//
type CfnByteMatchSet interface {
	awscdk.CfnResource
	awscdk.IInspectable
	ByteMatchTuples() interface{}
	SetByteMatchTuples(val interface{})
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() constructs.Node
	Ref() *string
	Stack() awscdk.Stack
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

// The jsii proxy struct for CfnByteMatchSet
type jsiiProxy_CfnByteMatchSet struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnByteMatchSet) ByteMatchTuples() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"byteMatchTuples",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnByteMatchSet) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnByteMatchSet) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnByteMatchSet) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnByteMatchSet) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnByteMatchSet) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnByteMatchSet) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnByteMatchSet) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnByteMatchSet) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnByteMatchSet) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnByteMatchSet) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::WAFRegional::ByteMatchSet`.
func NewCfnByteMatchSet(scope constructs.Construct, id *string, props *CfnByteMatchSetProps) CfnByteMatchSet {
	_init_.Initialize()

	j := jsiiProxy_CfnByteMatchSet{}

	_jsii_.Create(
		"aws-cdk-lib.aws_wafregional.CfnByteMatchSet",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::WAFRegional::ByteMatchSet`.
func NewCfnByteMatchSet_Override(c CfnByteMatchSet, scope constructs.Construct, id *string, props *CfnByteMatchSetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_wafregional.CfnByteMatchSet",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnByteMatchSet) SetByteMatchTuples(val interface{}) {
	_jsii_.Set(
		j,
		"byteMatchTuples",
		val,
	)
}

func (j *jsiiProxy_CfnByteMatchSet) SetName(val *string) {
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
func CfnByteMatchSet_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_wafregional.CfnByteMatchSet",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnByteMatchSet_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_wafregional.CfnByteMatchSet",
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
func CfnByteMatchSet_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_wafregional.CfnByteMatchSet",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnByteMatchSet_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_wafregional.CfnByteMatchSet",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnByteMatchSet) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnByteMatchSet) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnByteMatchSet) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnByteMatchSet) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnByteMatchSet) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnByteMatchSet) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnByteMatchSet) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnByteMatchSet) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnByteMatchSet) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnByteMatchSet) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnByteMatchSet) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnByteMatchSet) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnByteMatchSet) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnByteMatchSet) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnByteMatchSet) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// > This is *AWS WAF Classic* documentation.
//
// For more information, see [AWS WAF Classic](https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html) in the developer guide.
// >
// > *For the latest version of AWS WAF* , use the AWS WAF V2 API and see the [AWS WAF Developer Guide](https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html) . With the latest version, AWS WAF has a single set of endpoints for regional and global use.
//
// The bytes (typically a string that corresponds with ASCII characters) that you want AWS WAF to search for in web requests, the location in requests that you want AWS WAF to search, and other settings.
//
// TODO: EXAMPLE
//
type CfnByteMatchSet_ByteMatchTupleProperty struct {
	// The part of a web request that you want AWS WAF to inspect, such as a specific header or a query string.
	FieldToMatch interface{} `json:"fieldToMatch" yaml:"fieldToMatch"`
	// Within the portion of a web request that you want to search (for example, in the query string, if any), specify where you want AWS WAF to search.
	//
	// Valid values include the following:
	//
	// *CONTAINS*
	//
	// The specified part of the web request must include the value of `TargetString` , but the location doesn't matter.
	//
	// *CONTAINS_WORD*
	//
	// The specified part of the web request must include the value of `TargetString` , and `TargetString` must contain only alphanumeric characters or underscore (A-Z, a-z, 0-9, or _). In addition, `TargetString` must be a word, which means one of the following:
	//
	// - `TargetString` exactly matches the value of the specified part of the web request, such as the value of a header.
	// - `TargetString` is at the beginning of the specified part of the web request and is followed by a character other than an alphanumeric character or underscore (_), for example, `BadBot;` .
	// - `TargetString` is at the end of the specified part of the web request and is preceded by a character other than an alphanumeric character or underscore (_), for example, `;BadBot` .
	// - `TargetString` is in the middle of the specified part of the web request and is preceded and followed by characters other than alphanumeric characters or underscore (_), for example, `-BadBot;` .
	//
	// *EXACTLY*
	//
	// The value of the specified part of the web request must exactly match the value of `TargetString` .
	//
	// *STARTS_WITH*
	//
	// The value of `TargetString` must appear at the beginning of the specified part of the web request.
	//
	// *ENDS_WITH*
	//
	// The value of `TargetString` must appear at the end of the specified part of the web request.
	PositionalConstraint *string `json:"positionalConstraint" yaml:"positionalConstraint"`
	// Text transformations eliminate some of the unusual formatting that attackers use in web requests in an effort to bypass AWS WAF .
	//
	// If you specify a transformation, AWS WAF performs the transformation on `FieldToMatch` before inspecting it for a match.
	//
	// You can only specify a single type of TextTransformation.
	//
	// *CMD_LINE*
	//
	// When you're concerned that attackers are injecting an operating system command line command and using unusual formatting to disguise some or all of the command, use this option to perform the following transformations:
	//
	// - Delete the following characters: \ " ' ^
	// - Delete spaces before the following characters: / (
	// - Replace the following characters with a space: , ;
	// - Replace multiple spaces with one space
	// - Convert uppercase letters (A-Z) to lowercase (a-z)
	//
	// *COMPRESS_WHITE_SPACE*
	//
	// Use this option to replace the following characters with a space character (decimal 32):
	//
	// - \f, formfeed, decimal 12
	// - \t, tab, decimal 9
	// - \n, newline, decimal 10
	// - \r, carriage return, decimal 13
	// - \v, vertical tab, decimal 11
	// - non-breaking space, decimal 160
	//
	// `COMPRESS_WHITE_SPACE` also replaces multiple spaces with one space.
	//
	// *HTML_ENTITY_DECODE*
	//
	// Use this option to replace HTML-encoded characters with unencoded characters. `HTML_ENTITY_DECODE` performs the following operations:
	//
	// - Replaces `(ampersand)quot;` with `"`
	// - Replaces `(ampersand)nbsp;` with a non-breaking space, decimal 160
	// - Replaces `(ampersand)lt;` with a "less than" symbol
	// - Replaces `(ampersand)gt;` with `>`
	// - Replaces characters that are represented in hexadecimal format, `(ampersand)#xhhhh;` , with the corresponding characters
	// - Replaces characters that are represented in decimal format, `(ampersand)#nnnn;` , with the corresponding characters
	//
	// *LOWERCASE*
	//
	// Use this option to convert uppercase letters (A-Z) to lowercase (a-z).
	//
	// *URL_DECODE*
	//
	// Use this option to decode a URL-encoded value.
	//
	// *NONE*
	//
	// Specify `NONE` if you don't want to perform any text transformations.
	TextTransformation *string `json:"textTransformation" yaml:"textTransformation"`
	// The value that you want AWS WAF to search for.
	//
	// AWS WAF searches for the specified string in the part of web requests that you specified in `FieldToMatch` . The maximum length of the value is 50 bytes.
	//
	// You must specify this property or the `TargetStringBase64` property.
	//
	// Valid values depend on the values that you specified for `FieldToMatch` :
	//
	// - `HEADER` : The value that you want AWS WAF to search for in the request header that you specified in `FieldToMatch` , for example, the value of the `User-Agent` or `Referer` header.
	// - `METHOD` : The HTTP method, which indicates the type of operation specified in the request.
	// - `QUERY_STRING` : The value that you want AWS WAF to search for in the query string, which is the part of a URL that appears after a `?` character.
	// - `URI` : The value that you want AWS WAF to search for in the part of a URL that identifies a resource, for example, `/images/daily-ad.jpg` .
	// - `BODY` : The part of a request that contains any additional data that you want to send to your web server as the HTTP request body, such as data from a form. The request body immediately follows the request headers. Note that only the first `8192` bytes of the request body are forwarded to AWS WAF for inspection. To allow or block requests based on the length of the body, you can create a size constraint set.
	// - `SINGLE_QUERY_ARG` : The parameter in the query string that you will inspect, such as *UserName* or *SalesRegion* . The maximum length for `SINGLE_QUERY_ARG` is 30 characters.
	// - `ALL_QUERY_ARGS` : Similar to `SINGLE_QUERY_ARG` , but instead of inspecting a single parameter, AWS WAF inspects all parameters within the query string for the value or regex pattern that you specify in `TargetString` .
	//
	// If `TargetString` includes alphabetic characters A-Z and a-z, note that the value is case sensitive.
	TargetString *string `json:"targetString" yaml:"targetString"`
	// The base64-encoded value that AWS WAF searches for. AWS CloudFormation sends this value to AWS WAF without encoding it.
	//
	// You must specify this property or the `TargetString` property.
	//
	// AWS WAF searches for this value in a specific part of web requests, which you define in the `FieldToMatch` property.
	//
	// Valid values depend on the Type value in the `FieldToMatch` property. For example, for a `METHOD` type, you must specify HTTP methods such as `DELETE, GET, HEAD, OPTIONS, PATCH, POST` , and `PUT` .
	TargetStringBase64 *string `json:"targetStringBase64" yaml:"targetStringBase64"`
}

// > This is *AWS WAF Classic* documentation.
//
// For more information, see [AWS WAF Classic](https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html) in the developer guide.
// >
// > *For the latest version of AWS WAF* , use the AWS WAF V2 API and see the [AWS WAF Developer Guide](https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html) . With the latest version, AWS WAF has a single set of endpoints for regional and global use.
//
// Specifies where in a web request to look for `TargetString` .
//
// TODO: EXAMPLE
//
type CfnByteMatchSet_FieldToMatchProperty struct {
	// The part of the web request that you want AWS WAF to search for a specified string.
	//
	// Parts of a request that you can search include the following:
	//
	// - `HEADER` : A specified request header, for example, the value of the `User-Agent` or `Referer` header. If you choose `HEADER` for the type, specify the name of the header in `Data` .
	// - `METHOD` : The HTTP method, which indicated the type of operation that the request is asking the origin to perform.
	// - `QUERY_STRING` : A query string, which is the part of a URL that appears after a `?` character, if any.
	// - `URI` : The part of a web request that identifies a resource, for example, `/images/daily-ad.jpg` .
	// - `BODY` : The part of a request that contains any additional data that you want to send to your web server as the HTTP request body, such as data from a form. The request body immediately follows the request headers. Note that only the first `8192` bytes of the request body are forwarded to AWS WAF for inspection. To allow or block requests based on the length of the body, you can create a size constraint set.
	// - `SINGLE_QUERY_ARG` : The parameter in the query string that you will inspect, such as *UserName* or *SalesRegion* . The maximum length for `SINGLE_QUERY_ARG` is 30 characters.
	// - `ALL_QUERY_ARGS` : Similar to `SINGLE_QUERY_ARG` , but rather than inspecting a single parameter, AWS WAF will inspect all parameters within the query for the value or regex pattern that you specify in `TargetString` .
	Type *string `json:"type" yaml:"type"`
	// When the value of `Type` is `HEADER` , enter the name of the header that you want AWS WAF to search, for example, `User-Agent` or `Referer` .
	//
	// The name of the header is not case sensitive.
	//
	// When the value of `Type` is `SINGLE_QUERY_ARG` , enter the name of the parameter that you want AWS WAF to search, for example, `UserName` or `SalesRegion` . The parameter name is not case sensitive.
	//
	// If the value of `Type` is any other value, omit `Data` .
	Data *string `json:"data" yaml:"data"`
}

// Properties for defining a `CfnByteMatchSet`.
//
// TODO: EXAMPLE
//
type CfnByteMatchSetProps struct {
	// A friendly name or description of the `ByteMatchSet` .
	//
	// You can't change `Name` after you create a `ByteMatchSet` .
	Name *string `json:"name" yaml:"name"`
	// Specifies the bytes (typically a string that corresponds with ASCII characters) that you want AWS WAF to search for in web requests, the location in requests that you want AWS WAF to search, and other settings.
	ByteMatchTuples interface{} `json:"byteMatchTuples" yaml:"byteMatchTuples"`
}

// A CloudFormation `AWS::WAFRegional::GeoMatchSet`.
//
// > This is *AWS WAF Classic* documentation. For more information, see [AWS WAF Classic](https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html) in the developer guide.
// >
// > *For the latest version of AWS WAF* , use the AWS WAF V2 API and see the [AWS WAF Developer Guide](https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html) . With the latest version, AWS WAF has a single set of endpoints for regional and global use.
//
// Contains one or more countries that AWS WAF will search for.
//
// TODO: EXAMPLE
//
type CfnGeoMatchSet interface {
	awscdk.CfnResource
	awscdk.IInspectable
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	GeoMatchConstraints() interface{}
	SetGeoMatchConstraints(val interface{})
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() constructs.Node
	Ref() *string
	Stack() awscdk.Stack
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

// The jsii proxy struct for CfnGeoMatchSet
type jsiiProxy_CfnGeoMatchSet struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnGeoMatchSet) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGeoMatchSet) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGeoMatchSet) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGeoMatchSet) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGeoMatchSet) GeoMatchConstraints() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"geoMatchConstraints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGeoMatchSet) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGeoMatchSet) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGeoMatchSet) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGeoMatchSet) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGeoMatchSet) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGeoMatchSet) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::WAFRegional::GeoMatchSet`.
func NewCfnGeoMatchSet(scope constructs.Construct, id *string, props *CfnGeoMatchSetProps) CfnGeoMatchSet {
	_init_.Initialize()

	j := jsiiProxy_CfnGeoMatchSet{}

	_jsii_.Create(
		"aws-cdk-lib.aws_wafregional.CfnGeoMatchSet",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::WAFRegional::GeoMatchSet`.
func NewCfnGeoMatchSet_Override(c CfnGeoMatchSet, scope constructs.Construct, id *string, props *CfnGeoMatchSetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_wafregional.CfnGeoMatchSet",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnGeoMatchSet) SetGeoMatchConstraints(val interface{}) {
	_jsii_.Set(
		j,
		"geoMatchConstraints",
		val,
	)
}

func (j *jsiiProxy_CfnGeoMatchSet) SetName(val *string) {
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
func CfnGeoMatchSet_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_wafregional.CfnGeoMatchSet",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnGeoMatchSet_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_wafregional.CfnGeoMatchSet",
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
func CfnGeoMatchSet_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_wafregional.CfnGeoMatchSet",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnGeoMatchSet_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_wafregional.CfnGeoMatchSet",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnGeoMatchSet) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnGeoMatchSet) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnGeoMatchSet) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnGeoMatchSet) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnGeoMatchSet) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnGeoMatchSet) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnGeoMatchSet) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnGeoMatchSet) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnGeoMatchSet) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnGeoMatchSet) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnGeoMatchSet) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnGeoMatchSet) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnGeoMatchSet) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnGeoMatchSet) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnGeoMatchSet) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// > This is *AWS WAF Classic* documentation.
//
// For more information, see [AWS WAF Classic](https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html) in the developer guide.
// >
// > *For the latest version of AWS WAF* , use the AWS WAF V2 API and see the [AWS WAF Developer Guide](https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html) . With the latest version, AWS WAF has a single set of endpoints for regional and global use.
//
// The country from which web requests originate that you want AWS WAF to search for.
//
// TODO: EXAMPLE
//
type CfnGeoMatchSet_GeoMatchConstraintProperty struct {
	// The type of geographical area you want AWS WAF to search for.
	//
	// Currently `Country` is the only valid value.
	Type *string `json:"type" yaml:"type"`
	// The country that you want AWS WAF to search for.
	Value *string `json:"value" yaml:"value"`
}

// Properties for defining a `CfnGeoMatchSet`.
//
// TODO: EXAMPLE
//
type CfnGeoMatchSetProps struct {
	// A friendly name or description of the `GeoMatchSet` .
	//
	// You can't change the name of an `GeoMatchSet` after you create it.
	Name *string `json:"name" yaml:"name"`
	// An array of `GeoMatchConstraint` objects, which contain the country that you want AWS WAF to search for.
	GeoMatchConstraints interface{} `json:"geoMatchConstraints" yaml:"geoMatchConstraints"`
}

// A CloudFormation `AWS::WAFRegional::IPSet`.
//
// > This is *AWS WAF Classic* documentation. For more information, see [AWS WAF Classic](https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html) in the developer guide.
// >
// > *For the latest version of AWS WAF* , use the AWS WAF V2 API and see the [AWS WAF Developer Guide](https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html) . With the latest version, AWS WAF has a single set of endpoints for regional and global use.
//
// Contains one or more IP addresses or blocks of IP addresses specified in Classless Inter-Domain Routing (CIDR) notation. AWS WAF supports IPv4 address ranges: /8 and any range between /16 through /32. AWS WAF supports IPv6 address ranges: /24, /32, /48, /56, /64, and /128.
//
// To specify an individual IP address, you specify the four-part IP address followed by a `/32` , for example, 192.0.2.0/32. To block a range of IP addresses, you can specify /8 or any range between /16 through /32 (for IPv4) or /24, /32, /48, /56, /64, or /128 (for IPv6). For more information about CIDR notation, see the Wikipedia entry [Classless Inter-Domain Routing](https://docs.aws.amazon.com/https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing) .
//
// TODO: EXAMPLE
//
type CfnIPSet interface {
	awscdk.CfnResource
	awscdk.IInspectable
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	IpSetDescriptors() interface{}
	SetIpSetDescriptors(val interface{})
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() constructs.Node
	Ref() *string
	Stack() awscdk.Stack
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

// The jsii proxy struct for CfnIPSet
type jsiiProxy_CfnIPSet struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnIPSet) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIPSet) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIPSet) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIPSet) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIPSet) IpSetDescriptors() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipSetDescriptors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIPSet) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIPSet) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIPSet) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIPSet) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIPSet) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIPSet) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::WAFRegional::IPSet`.
func NewCfnIPSet(scope constructs.Construct, id *string, props *CfnIPSetProps) CfnIPSet {
	_init_.Initialize()

	j := jsiiProxy_CfnIPSet{}

	_jsii_.Create(
		"aws-cdk-lib.aws_wafregional.CfnIPSet",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::WAFRegional::IPSet`.
func NewCfnIPSet_Override(c CfnIPSet, scope constructs.Construct, id *string, props *CfnIPSetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_wafregional.CfnIPSet",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnIPSet) SetIpSetDescriptors(val interface{}) {
	_jsii_.Set(
		j,
		"ipSetDescriptors",
		val,
	)
}

func (j *jsiiProxy_CfnIPSet) SetName(val *string) {
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
func CfnIPSet_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_wafregional.CfnIPSet",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnIPSet_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_wafregional.CfnIPSet",
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
func CfnIPSet_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_wafregional.CfnIPSet",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnIPSet_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_wafregional.CfnIPSet",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnIPSet) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnIPSet) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnIPSet) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnIPSet) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnIPSet) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnIPSet) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnIPSet) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnIPSet) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnIPSet) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnIPSet) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnIPSet) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnIPSet) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnIPSet) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnIPSet) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnIPSet) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// > This is *AWS WAF Classic* documentation.
//
// For more information, see [AWS WAF Classic](https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html) in the developer guide.
// >
// > *For the latest version of AWS WAF* , use the AWS WAF V2 API and see the [AWS WAF Developer Guide](https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html) . With the latest version, AWS WAF has a single set of endpoints for regional and global use.
//
// Specifies the IP address type ( `IPV4` or `IPV6` ) and the IP address range (in CIDR format) that web requests originate from.
//
// TODO: EXAMPLE
//
type CfnIPSet_IPSetDescriptorProperty struct {
	// Specify `IPV4` or `IPV6` .
	Type *string `json:"type" yaml:"type"`
	// Specify an IPv4 address by using CIDR notation. For example:.
	//
	// - To configure AWS WAF to allow, block, or count requests that originated from the IP address 192.0.2.44, specify `192.0.2.44/32` .
	// - To configure AWS WAF to allow, block, or count requests that originated from IP addresses from 192.0.2.0 to 192.0.2.255, specify `192.0.2.0/24` .
	//
	// For more information about CIDR notation, see the Wikipedia entry [Classless Inter-Domain Routing](https://docs.aws.amazon.com/https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing) .
	//
	// Specify an IPv6 address by using CIDR notation. For example:
	//
	// - To configure AWS WAF to allow, block, or count requests that originated from the IP address 1111:0000:0000:0000:0000:0000:0000:0111, specify `1111:0000:0000:0000:0000:0000:0000:0111/128` .
	// - To configure AWS WAF to allow, block, or count requests that originated from IP addresses 1111:0000:0000:0000:0000:0000:0000:0000 to 1111:0000:0000:0000:ffff:ffff:ffff:ffff, specify `1111:0000:0000:0000:0000:0000:0000:0000/64` .
	Value *string `json:"value" yaml:"value"`
}

// Properties for defining a `CfnIPSet`.
//
// TODO: EXAMPLE
//
type CfnIPSetProps struct {
	// A friendly name or description of the `IPSet` .
	//
	// You can't change the name of an `IPSet` after you create it.
	Name *string `json:"name" yaml:"name"`
	// The IP address type ( `IPV4` or `IPV6` ) and the IP address range (in CIDR notation) that web requests originate from.
	IpSetDescriptors interface{} `json:"ipSetDescriptors" yaml:"ipSetDescriptors"`
}

// A CloudFormation `AWS::WAFRegional::RateBasedRule`.
//
// > This is *AWS WAF Classic* documentation. For more information, see [AWS WAF Classic](https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html) in the developer guide.
// >
// > *For the latest version of AWS WAF* , use the AWS WAF V2 API and see the [AWS WAF Developer Guide](https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html) . With the latest version, AWS WAF has a single set of endpoints for regional and global use.
//
// A `RateBasedRule` is identical to a regular `Rule` , with one addition: a `RateBasedRule` counts the number of requests that arrive from a specified IP address every five minutes. For example, based on recent requests that you've seen from an attacker, you might create a `RateBasedRule` that includes the following conditions:
//
// - The requests come from 192.0.2.44.
// - They contain the value `BadBot` in the `User-Agent` header.
//
// In the rule, you also define the rate limit as 15,000.
//
// Requests that meet both of these conditions and exceed 15,000 requests every five minutes trigger the rule's action (block or count), which is defined in the web ACL.
//
// Note you can only create rate-based rules using an AWS CloudFormation template. To add the rate-based rules created through AWS CloudFormation to a web ACL, use the AWS WAF console, API, or command line interface (CLI). For more information, see [UpdateWebACL](https://docs.aws.amazon.com/waf/latest/APIReference/API_regional_UpdateWebACL.html) .
//
// TODO: EXAMPLE
//
type CfnRateBasedRule interface {
	awscdk.CfnResource
	awscdk.IInspectable
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	LogicalId() *string
	MatchPredicates() interface{}
	SetMatchPredicates(val interface{})
	MetricName() *string
	SetMetricName(val *string)
	Name() *string
	SetName(val *string)
	Node() constructs.Node
	RateKey() *string
	SetRateKey(val *string)
	RateLimit() *float64
	SetRateLimit(val *float64)
	Ref() *string
	Stack() awscdk.Stack
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

// The jsii proxy struct for CfnRateBasedRule
type jsiiProxy_CfnRateBasedRule struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnRateBasedRule) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRateBasedRule) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRateBasedRule) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRateBasedRule) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRateBasedRule) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRateBasedRule) MatchPredicates() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"matchPredicates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRateBasedRule) MetricName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRateBasedRule) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRateBasedRule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRateBasedRule) RateKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rateKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRateBasedRule) RateLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rateLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRateBasedRule) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRateBasedRule) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRateBasedRule) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::WAFRegional::RateBasedRule`.
func NewCfnRateBasedRule(scope constructs.Construct, id *string, props *CfnRateBasedRuleProps) CfnRateBasedRule {
	_init_.Initialize()

	j := jsiiProxy_CfnRateBasedRule{}

	_jsii_.Create(
		"aws-cdk-lib.aws_wafregional.CfnRateBasedRule",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::WAFRegional::RateBasedRule`.
func NewCfnRateBasedRule_Override(c CfnRateBasedRule, scope constructs.Construct, id *string, props *CfnRateBasedRuleProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_wafregional.CfnRateBasedRule",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnRateBasedRule) SetMatchPredicates(val interface{}) {
	_jsii_.Set(
		j,
		"matchPredicates",
		val,
	)
}

func (j *jsiiProxy_CfnRateBasedRule) SetMetricName(val *string) {
	_jsii_.Set(
		j,
		"metricName",
		val,
	)
}

func (j *jsiiProxy_CfnRateBasedRule) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnRateBasedRule) SetRateKey(val *string) {
	_jsii_.Set(
		j,
		"rateKey",
		val,
	)
}

func (j *jsiiProxy_CfnRateBasedRule) SetRateLimit(val *float64) {
	_jsii_.Set(
		j,
		"rateLimit",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnRateBasedRule_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_wafregional.CfnRateBasedRule",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnRateBasedRule_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_wafregional.CfnRateBasedRule",
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
func CfnRateBasedRule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_wafregional.CfnRateBasedRule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnRateBasedRule_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_wafregional.CfnRateBasedRule",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnRateBasedRule) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnRateBasedRule) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnRateBasedRule) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnRateBasedRule) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnRateBasedRule) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnRateBasedRule) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnRateBasedRule) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnRateBasedRule) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnRateBasedRule) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnRateBasedRule) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnRateBasedRule) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnRateBasedRule) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnRateBasedRule) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnRateBasedRule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnRateBasedRule) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Specifies the `ByteMatchSet` , `IPSet` , `SqlInjectionMatchSet` , `XssMatchSet` , `RegexMatchSet` , `GeoMatchSet` , and `SizeConstraintSet` objects that you want to add to a `Rule` and, for each object, indicates whether you want to negate the settings, for example, requests that do NOT originate from the IP address 192.0.2.44.
//
// TODO: EXAMPLE
//
type CfnRateBasedRule_PredicateProperty struct {
	// A unique identifier for a predicate in a `Rule` , such as `ByteMatchSetId` or `IPSetId` .
	//
	// The ID is returned by the corresponding `Create` or `List` command.
	DataId *string `json:"dataId" yaml:"dataId"`
	// Set `Negated` to `False` if you want AWS WAF to allow, block, or count requests based on the settings in the specified `ByteMatchSet` , `IPSet` , `SqlInjectionMatchSet` , `XssMatchSet` , `RegexMatchSet` , `GeoMatchSet` , or `SizeConstraintSet` .
	//
	// For example, if an `IPSet` includes the IP address `192.0.2.44` , AWS WAF will allow or block requests based on that IP address.
	//
	// Set `Negated` to `True` if you want AWS WAF to allow or block a request based on the negation of the settings in the `ByteMatchSet` , `IPSet` , `SqlInjectionMatchSet` , `XssMatchSet` , `RegexMatchSet` , `GeoMatchSet` , or `SizeConstraintSet` >. For example, if an `IPSet` includes the IP address `192.0.2.44` , AWS WAF will allow, block, or count requests based on all IP addresses *except* `192.0.2.44` .
	Negated interface{} `json:"negated" yaml:"negated"`
	// The type of predicate in a `Rule` , such as `ByteMatch` or `IPSet` .
	Type *string `json:"type" yaml:"type"`
}

// Properties for defining a `CfnRateBasedRule`.
//
// TODO: EXAMPLE
//
type CfnRateBasedRuleProps struct {
	// A name for the metrics for a `RateBasedRule` .
	//
	// The name can contain only alphanumeric characters (A-Z, a-z, 0-9), with maximum length 128 and minimum length one. It can't contain whitespace or metric names reserved for AWS WAF , including "All" and "Default_Action." You can't change the name of the metric after you create the `RateBasedRule` .
	MetricName *string `json:"metricName" yaml:"metricName"`
	// A friendly name or description for a `RateBasedRule` .
	//
	// You can't change the name of a `RateBasedRule` after you create it.
	Name *string `json:"name" yaml:"name"`
	// The field that AWS WAF uses to determine if requests are likely arriving from single source and thus subject to rate monitoring.
	//
	// The only valid value for `RateKey` is `IP` . `IP` indicates that requests arriving from the same IP address are subject to the `RateLimit` that is specified in the `RateBasedRule` .
	RateKey *string `json:"rateKey" yaml:"rateKey"`
	// The maximum number of requests, which have an identical value in the field specified by the `RateKey` , allowed in a five-minute period.
	//
	// If the number of requests exceeds the `RateLimit` and the other predicates specified in the rule are also met, AWS WAF triggers the action that is specified for this rule.
	RateLimit *float64 `json:"rateLimit" yaml:"rateLimit"`
	// The `Predicates` object contains one `Predicate` element for each `ByteMatchSet` , `IPSet` , or `SqlInjectionMatchSet>` object that you want to include in a `RateBasedRule` .
	MatchPredicates interface{} `json:"matchPredicates" yaml:"matchPredicates"`
}

// A CloudFormation `AWS::WAFRegional::RegexPatternSet`.
//
// The `RegexPatternSet` specifies the regular expression (regex) pattern that you want AWS WAF to search for, such as `B[a@]dB[o0]t` . You can then configure AWS WAF to reject those requests.
//
// Note that you can only create regex pattern sets using a AWS CloudFormation template. To add the regex pattern sets created through AWS CloudFormation to a RegexMatchSet, use the AWS WAF console, API, or command line interface (CLI). For more information, see [UpdateRegexMatchSet](https://docs.aws.amazon.com/waf/latest/APIReference/API_regional_UpdateRegexMatchSet.html) .
//
// TODO: EXAMPLE
//
type CfnRegexPatternSet interface {
	awscdk.CfnResource
	awscdk.IInspectable
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() constructs.Node
	Ref() *string
	RegexPatternStrings() *[]*string
	SetRegexPatternStrings(val *[]*string)
	Stack() awscdk.Stack
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

// The jsii proxy struct for CfnRegexPatternSet
type jsiiProxy_CfnRegexPatternSet struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnRegexPatternSet) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRegexPatternSet) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRegexPatternSet) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRegexPatternSet) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRegexPatternSet) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRegexPatternSet) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRegexPatternSet) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRegexPatternSet) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRegexPatternSet) RegexPatternStrings() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"regexPatternStrings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRegexPatternSet) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRegexPatternSet) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::WAFRegional::RegexPatternSet`.
func NewCfnRegexPatternSet(scope constructs.Construct, id *string, props *CfnRegexPatternSetProps) CfnRegexPatternSet {
	_init_.Initialize()

	j := jsiiProxy_CfnRegexPatternSet{}

	_jsii_.Create(
		"aws-cdk-lib.aws_wafregional.CfnRegexPatternSet",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::WAFRegional::RegexPatternSet`.
func NewCfnRegexPatternSet_Override(c CfnRegexPatternSet, scope constructs.Construct, id *string, props *CfnRegexPatternSetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_wafregional.CfnRegexPatternSet",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnRegexPatternSet) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnRegexPatternSet) SetRegexPatternStrings(val *[]*string) {
	_jsii_.Set(
		j,
		"regexPatternStrings",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnRegexPatternSet_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_wafregional.CfnRegexPatternSet",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnRegexPatternSet_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_wafregional.CfnRegexPatternSet",
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
func CfnRegexPatternSet_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_wafregional.CfnRegexPatternSet",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnRegexPatternSet_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_wafregional.CfnRegexPatternSet",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnRegexPatternSet) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnRegexPatternSet) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnRegexPatternSet) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnRegexPatternSet) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnRegexPatternSet) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnRegexPatternSet) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnRegexPatternSet) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnRegexPatternSet) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnRegexPatternSet) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnRegexPatternSet) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnRegexPatternSet) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnRegexPatternSet) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnRegexPatternSet) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnRegexPatternSet) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnRegexPatternSet) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnRegexPatternSet`.
//
// TODO: EXAMPLE
//
type CfnRegexPatternSetProps struct {
	// A friendly name or description of the `RegexPatternSet` .
	//
	// You can't change `Name` after you create a `RegexPatternSet` .
	Name *string `json:"name" yaml:"name"`
	// Specifies the regular expression (regex) patterns that you want AWS WAF to search for, such as `B[a@]dB[o0]t` .
	RegexPatternStrings *[]*string `json:"regexPatternStrings" yaml:"regexPatternStrings"`
}

// A CloudFormation `AWS::WAFRegional::Rule`.
//
// > This is *AWS WAF Classic* documentation. For more information, see [AWS WAF Classic](https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html) in the developer guide.
// >
// > *For the latest version of AWS WAF* , use the AWS WAF V2 API and see the [AWS WAF Developer Guide](https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html) . With the latest version, AWS WAF has a single set of endpoints for regional and global use.
//
// A combination of `ByteMatchSet` , `IPSet` , and/or `SqlInjectionMatchSet` objects that identify the web requests that you want to allow, block, or count. For example, you might create a `Rule` that includes the following predicates:
//
// - An `IPSet` that causes AWS WAF to search for web requests that originate from the IP address `192.0.2.44`
// - A `ByteMatchSet` that causes AWS WAF to search for web requests for which the value of the `User-Agent` header is `BadBot` .
//
// To match the settings in this `Rule` , a request must originate from `192.0.2.44` AND include a `User-Agent` header for which the value is `BadBot` .
//
// TODO: EXAMPLE
//
type CfnRule interface {
	awscdk.CfnResource
	awscdk.IInspectable
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	LogicalId() *string
	MetricName() *string
	SetMetricName(val *string)
	Name() *string
	SetName(val *string)
	Node() constructs.Node
	Predicates() interface{}
	SetPredicates(val interface{})
	Ref() *string
	Stack() awscdk.Stack
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

// The jsii proxy struct for CfnRule
type jsiiProxy_CfnRule struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnRule) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRule) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRule) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRule) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRule) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRule) MetricName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRule) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRule) Predicates() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"predicates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRule) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRule) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRule) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::WAFRegional::Rule`.
func NewCfnRule(scope constructs.Construct, id *string, props *CfnRuleProps) CfnRule {
	_init_.Initialize()

	j := jsiiProxy_CfnRule{}

	_jsii_.Create(
		"aws-cdk-lib.aws_wafregional.CfnRule",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::WAFRegional::Rule`.
func NewCfnRule_Override(c CfnRule, scope constructs.Construct, id *string, props *CfnRuleProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_wafregional.CfnRule",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnRule) SetMetricName(val *string) {
	_jsii_.Set(
		j,
		"metricName",
		val,
	)
}

func (j *jsiiProxy_CfnRule) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnRule) SetPredicates(val interface{}) {
	_jsii_.Set(
		j,
		"predicates",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnRule_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_wafregional.CfnRule",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnRule_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_wafregional.CfnRule",
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
func CfnRule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_wafregional.CfnRule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnRule_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_wafregional.CfnRule",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnRule) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnRule) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnRule) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnRule) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnRule) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnRule) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnRule) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnRule) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnRule) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnRule) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnRule) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnRule) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnRule) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnRule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnRule) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Specifies the `ByteMatchSet` , `IPSet` , `SqlInjectionMatchSet` , `XssMatchSet` , `RegexMatchSet` , `GeoMatchSet` , and `SizeConstraintSet` objects that you want to add to a `Rule` and, for each object, indicates whether you want to negate the settings, for example, requests that do NOT originate from the IP address 192.0.2.44.
//
// TODO: EXAMPLE
//
type CfnRule_PredicateProperty struct {
	// A unique identifier for a predicate in a `Rule` , such as `ByteMatchSetId` or `IPSetId` .
	//
	// The ID is returned by the corresponding `Create` or `List` command.
	DataId *string `json:"dataId" yaml:"dataId"`
	// Set `Negated` to `False` if you want AWS WAF to allow, block, or count requests based on the settings in the specified `ByteMatchSet` , `IPSet` , `SqlInjectionMatchSet` , `XssMatchSet` , `RegexMatchSet` , `GeoMatchSet` , or `SizeConstraintSet` .
	//
	// For example, if an `IPSet` includes the IP address `192.0.2.44` , AWS WAF will allow or block requests based on that IP address.
	//
	// Set `Negated` to `True` if you want AWS WAF to allow or block a request based on the negation of the settings in the `ByteMatchSet` , `IPSet` , `SqlInjectionMatchSet` , `XssMatchSet` , `RegexMatchSet` , `GeoMatchSet` , or `SizeConstraintSet` . For example, if an `IPSet` includes the IP address `192.0.2.44` , AWS WAF will allow, block, or count requests based on all IP addresses *except* `192.0.2.44` .
	Negated interface{} `json:"negated" yaml:"negated"`
	// The type of predicate in a `Rule` , such as `ByteMatch` or `IPSet` .
	Type *string `json:"type" yaml:"type"`
}

// Properties for defining a `CfnRule`.
//
// TODO: EXAMPLE
//
type CfnRuleProps struct {
	// A name for the metrics for this `Rule` .
	//
	// The name can contain only alphanumeric characters (A-Z, a-z, 0-9), with maximum length 128 and minimum length one. It can't contain whitespace or metric names reserved for AWS WAF, including "All" and "Default_Action." You can't change `MetricName` after you create the `Rule` .
	MetricName *string `json:"metricName" yaml:"metricName"`
	// The friendly name or description for the `Rule` .
	//
	// You can't change the name of a `Rule` after you create it.
	Name *string `json:"name" yaml:"name"`
	// The `Predicates` object contains one `Predicate` element for each `ByteMatchSet` , `IPSet` , or `SqlInjectionMatchSet` object that you want to include in a `Rule` .
	Predicates interface{} `json:"predicates" yaml:"predicates"`
}

// A CloudFormation `AWS::WAFRegional::SizeConstraintSet`.
//
// > This is *AWS WAF Classic* documentation. For more information, see [AWS WAF Classic](https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html) in the developer guide.
// >
// > *For the latest version of AWS WAF* , use the AWS WAF V2 API and see the [AWS WAF Developer Guide](https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html) . With the latest version, AWS WAF has a single set of endpoints for regional and global use.
//
// A complex type that contains `SizeConstraint` objects, which specify the parts of web requests that you want AWS WAF to inspect the size of. If a `SizeConstraintSet` contains more than one `SizeConstraint` object, a request only needs to match one constraint to be considered a match.
//
// TODO: EXAMPLE
//
type CfnSizeConstraintSet interface {
	awscdk.CfnResource
	awscdk.IInspectable
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() constructs.Node
	Ref() *string
	SizeConstraints() interface{}
	SetSizeConstraints(val interface{})
	Stack() awscdk.Stack
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

// The jsii proxy struct for CfnSizeConstraintSet
type jsiiProxy_CfnSizeConstraintSet struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnSizeConstraintSet) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSizeConstraintSet) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSizeConstraintSet) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSizeConstraintSet) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSizeConstraintSet) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSizeConstraintSet) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSizeConstraintSet) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSizeConstraintSet) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSizeConstraintSet) SizeConstraints() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sizeConstraints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSizeConstraintSet) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSizeConstraintSet) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::WAFRegional::SizeConstraintSet`.
func NewCfnSizeConstraintSet(scope constructs.Construct, id *string, props *CfnSizeConstraintSetProps) CfnSizeConstraintSet {
	_init_.Initialize()

	j := jsiiProxy_CfnSizeConstraintSet{}

	_jsii_.Create(
		"aws-cdk-lib.aws_wafregional.CfnSizeConstraintSet",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::WAFRegional::SizeConstraintSet`.
func NewCfnSizeConstraintSet_Override(c CfnSizeConstraintSet, scope constructs.Construct, id *string, props *CfnSizeConstraintSetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_wafregional.CfnSizeConstraintSet",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnSizeConstraintSet) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnSizeConstraintSet) SetSizeConstraints(val interface{}) {
	_jsii_.Set(
		j,
		"sizeConstraints",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnSizeConstraintSet_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_wafregional.CfnSizeConstraintSet",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnSizeConstraintSet_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_wafregional.CfnSizeConstraintSet",
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
func CfnSizeConstraintSet_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_wafregional.CfnSizeConstraintSet",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSizeConstraintSet_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_wafregional.CfnSizeConstraintSet",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnSizeConstraintSet) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnSizeConstraintSet) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnSizeConstraintSet) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnSizeConstraintSet) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnSizeConstraintSet) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnSizeConstraintSet) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnSizeConstraintSet) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnSizeConstraintSet) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnSizeConstraintSet) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnSizeConstraintSet) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnSizeConstraintSet) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnSizeConstraintSet) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnSizeConstraintSet) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnSizeConstraintSet) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSizeConstraintSet) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The part of a web request that you want AWS WAF to inspect, such as a specific header or a query string.
//
// TODO: EXAMPLE
//
type CfnSizeConstraintSet_FieldToMatchProperty struct {
	// The part of the web request that you want AWS WAF to search for a specified string.
	//
	// Parts of a request that you can search include the following:
	//
	// - `HEADER` : A specified request header, for example, the value of the `User-Agent` or `Referer` header. If you choose `HEADER` for the type, specify the name of the header in `Data` .
	// - `METHOD` : The HTTP method, which indicates the type of operation that the request is asking the origin to perform.
	// - `QUERY_STRING` : A query string, which is the part of a URL that appears after a `?` character, if any.
	// - `URI` : The part of a web request that identifies a resource, for example, `/images/daily-ad.jpg` .
	// - `BODY` : The part of a request that contains any additional data that you want to send to your web server as the HTTP request body, such as data from a form. The request body immediately follows the request headers. Note that only the first `8192` bytes of the request body are forwarded to AWS WAF for inspection. To allow or block requests based on the length of the body, you can create a size constraint set.
	// - `SINGLE_QUERY_ARG` : The parameter in the query string that you will inspect, such as *UserName* or *SalesRegion* . The maximum length for `SINGLE_QUERY_ARG` is 30 characters.
	// - `ALL_QUERY_ARGS` : Similar to `SINGLE_QUERY_ARG` , but rather than inspecting a single parameter, AWS WAF will inspect all parameters within the query for the value or regex pattern that you specify in `TargetString` .
	Type *string `json:"type" yaml:"type"`
	// When the value of `Type` is `HEADER` , enter the name of the header that you want AWS WAF to search, for example, `User-Agent` or `Referer` .
	//
	// The name of the header is not case sensitive.
	//
	// When the value of `Type` is `SINGLE_QUERY_ARG` , enter the name of the parameter that you want AWS WAF to search, for example, `UserName` or `SalesRegion` . The parameter name is not case sensitive.
	//
	// If the value of `Type` is any other value, omit `Data` .
	Data *string `json:"data" yaml:"data"`
}

// > This is *AWS WAF Classic* documentation.
//
// For more information, see [AWS WAF Classic](https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html) in the developer guide.
// >
// > *For the latest version of AWS WAF* , use the AWS WAF V2 API and see the [AWS WAF Developer Guide](https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html) . With the latest version, AWS WAF has a single set of endpoints for regional and global use.
//
// Specifies a constraint on the size of a part of the web request. AWS WAF uses the `Size` , `ComparisonOperator` , and `FieldToMatch` to build an expression in the form of " `Size` `ComparisonOperator` size in bytes of `FieldToMatch` ". If that expression is true, the `SizeConstraint` is considered to match.
//
// TODO: EXAMPLE
//
type CfnSizeConstraintSet_SizeConstraintProperty struct {
	// The type of comparison you want AWS WAF to perform.
	//
	// AWS WAF uses this in combination with the provided `Size` and `FieldToMatch` to build an expression in the form of " `Size` `ComparisonOperator` size in bytes of `FieldToMatch` ". If that expression is true, the `SizeConstraint` is considered to match.
	//
	// *EQ* : Used to test if the `Size` is equal to the size of the `FieldToMatch`
	//
	// *NE* : Used to test if the `Size` is not equal to the size of the `FieldToMatch`
	//
	// *LE* : Used to test if the `Size` is less than or equal to the size of the `FieldToMatch`
	//
	// *LT* : Used to test if the `Size` is strictly less than the size of the `FieldToMatch`
	//
	// *GE* : Used to test if the `Size` is greater than or equal to the size of the `FieldToMatch`
	//
	// *GT* : Used to test if the `Size` is strictly greater than the size of the `FieldToMatch`
	ComparisonOperator *string `json:"comparisonOperator" yaml:"comparisonOperator"`
	// The part of a web request that you want AWS WAF to inspect, such as a specific header or a query string.
	FieldToMatch interface{} `json:"fieldToMatch" yaml:"fieldToMatch"`
	// The size in bytes that you want AWS WAF to compare against the size of the specified `FieldToMatch` .
	//
	// AWS WAF uses this in combination with `ComparisonOperator` and `FieldToMatch` to build an expression in the form of " `Size` `ComparisonOperator` size in bytes of `FieldToMatch` ". If that expression is true, the `SizeConstraint` is considered to match.
	//
	// Valid values for size are 0 - 21474836480 bytes (0 - 20 GB).
	//
	// If you specify `URI` for the value of `Type` , the / in the URI path that you specify counts as one character. For example, the URI `/logo.jpg` is nine characters long.
	Size *float64 `json:"size" yaml:"size"`
	// Text transformations eliminate some of the unusual formatting that attackers use in web requests in an effort to bypass AWS WAF .
	//
	// If you specify a transformation, AWS WAF performs the transformation on `FieldToMatch` before inspecting a request for a match.
	//
	// You can only specify a single type of TextTransformation.
	//
	// Note that if you choose `BODY` for the value of `Type` , you must choose `NONE` for `TextTransformation` because the API Gateway API or Application Load Balancer forward only the first 8192 bytes for inspection.
	//
	// *NONE*
	//
	// Specify `NONE` if you don't want to perform any text transformations.
	//
	// *CMD_LINE*
	//
	// When you're concerned that attackers are injecting an operating system command line command and using unusual formatting to disguise some or all of the command, use this option to perform the following transformations:
	//
	// - Delete the following characters: \ " ' ^
	// - Delete spaces before the following characters: / (
	// - Replace the following characters with a space: , ;
	// - Replace multiple spaces with one space
	// - Convert uppercase letters (A-Z) to lowercase (a-z)
	//
	// *COMPRESS_WHITE_SPACE*
	//
	// Use this option to replace the following characters with a space character (decimal 32):
	//
	// - \f, formfeed, decimal 12
	// - \t, tab, decimal 9
	// - \n, newline, decimal 10
	// - \r, carriage return, decimal 13
	// - \v, vertical tab, decimal 11
	// - non-breaking space, decimal 160
	//
	// `COMPRESS_WHITE_SPACE` also replaces multiple spaces with one space.
	//
	// *HTML_ENTITY_DECODE*
	//
	// Use this option to replace HTML-encoded characters with unencoded characters. `HTML_ENTITY_DECODE` performs the following operations:
	//
	// - Replaces `(ampersand)quot;` with `"`
	// - Replaces `(ampersand)nbsp;` with a non-breaking space, decimal 160
	// - Replaces `(ampersand)lt;` with a "less than" symbol
	// - Replaces `(ampersand)gt;` with `>`
	// - Replaces characters that are represented in hexadecimal format, `(ampersand)#xhhhh;` , with the corresponding characters
	// - Replaces characters that are represented in decimal format, `(ampersand)#nnnn;` , with the corresponding characters
	//
	// *LOWERCASE*
	//
	// Use this option to convert uppercase letters (A-Z) to lowercase (a-z).
	//
	// *URL_DECODE*
	//
	// Use this option to decode a URL-encoded value.
	TextTransformation *string `json:"textTransformation" yaml:"textTransformation"`
}

// Properties for defining a `CfnSizeConstraintSet`.
//
// TODO: EXAMPLE
//
type CfnSizeConstraintSetProps struct {
	// The name, if any, of the `SizeConstraintSet` .
	Name *string `json:"name" yaml:"name"`
	// The size constraint and the part of the web request to check.
	SizeConstraints interface{} `json:"sizeConstraints" yaml:"sizeConstraints"`
}

// A CloudFormation `AWS::WAFRegional::SqlInjectionMatchSet`.
//
// > This is *AWS WAF Classic* documentation. For more information, see [AWS WAF Classic](https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html) in the developer guide.
// >
// > *For the latest version of AWS WAF* , use the AWS WAF V2 API and see the [AWS WAF Developer Guide](https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html) . With the latest version, AWS WAF has a single set of endpoints for regional and global use.
//
// A complex type that contains `SqlInjectionMatchTuple` objects, which specify the parts of web requests that you want AWS WAF to inspect for snippets of malicious SQL code and, if you want AWS WAF to inspect a header, the name of the header. If a `SqlInjectionMatchSet` contains more than one `SqlInjectionMatchTuple` object, a request needs to include snippets of SQL code in only one of the specified parts of the request to be considered a match.
//
// TODO: EXAMPLE
//
type CfnSqlInjectionMatchSet interface {
	awscdk.CfnResource
	awscdk.IInspectable
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() constructs.Node
	Ref() *string
	SqlInjectionMatchTuples() interface{}
	SetSqlInjectionMatchTuples(val interface{})
	Stack() awscdk.Stack
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

// The jsii proxy struct for CfnSqlInjectionMatchSet
type jsiiProxy_CfnSqlInjectionMatchSet struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnSqlInjectionMatchSet) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSqlInjectionMatchSet) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSqlInjectionMatchSet) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSqlInjectionMatchSet) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSqlInjectionMatchSet) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSqlInjectionMatchSet) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSqlInjectionMatchSet) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSqlInjectionMatchSet) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSqlInjectionMatchSet) SqlInjectionMatchTuples() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sqlInjectionMatchTuples",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSqlInjectionMatchSet) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSqlInjectionMatchSet) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::WAFRegional::SqlInjectionMatchSet`.
func NewCfnSqlInjectionMatchSet(scope constructs.Construct, id *string, props *CfnSqlInjectionMatchSetProps) CfnSqlInjectionMatchSet {
	_init_.Initialize()

	j := jsiiProxy_CfnSqlInjectionMatchSet{}

	_jsii_.Create(
		"aws-cdk-lib.aws_wafregional.CfnSqlInjectionMatchSet",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::WAFRegional::SqlInjectionMatchSet`.
func NewCfnSqlInjectionMatchSet_Override(c CfnSqlInjectionMatchSet, scope constructs.Construct, id *string, props *CfnSqlInjectionMatchSetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_wafregional.CfnSqlInjectionMatchSet",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnSqlInjectionMatchSet) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnSqlInjectionMatchSet) SetSqlInjectionMatchTuples(val interface{}) {
	_jsii_.Set(
		j,
		"sqlInjectionMatchTuples",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnSqlInjectionMatchSet_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_wafregional.CfnSqlInjectionMatchSet",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnSqlInjectionMatchSet_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_wafregional.CfnSqlInjectionMatchSet",
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
func CfnSqlInjectionMatchSet_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_wafregional.CfnSqlInjectionMatchSet",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSqlInjectionMatchSet_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_wafregional.CfnSqlInjectionMatchSet",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnSqlInjectionMatchSet) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnSqlInjectionMatchSet) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnSqlInjectionMatchSet) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnSqlInjectionMatchSet) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnSqlInjectionMatchSet) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnSqlInjectionMatchSet) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnSqlInjectionMatchSet) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnSqlInjectionMatchSet) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnSqlInjectionMatchSet) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnSqlInjectionMatchSet) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnSqlInjectionMatchSet) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnSqlInjectionMatchSet) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnSqlInjectionMatchSet) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnSqlInjectionMatchSet) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSqlInjectionMatchSet) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The part of a web request that you want AWS WAF to inspect, such as a specific header or a query string.
//
// TODO: EXAMPLE
//
type CfnSqlInjectionMatchSet_FieldToMatchProperty struct {
	// The part of the web request that you want AWS WAF to search for a specified string.
	//
	// Parts of a request that you can search include the following:
	//
	// - `HEADER` : A specified request header, for example, the value of the `User-Agent` or `Referer` header. If you choose `HEADER` for the type, specify the name of the header in `Data` .
	// - `METHOD` : The HTTP method, which indicates the type of operation that the request is asking the origin to perform.
	// - `QUERY_STRING` : A query string, which is the part of a URL that appears after a `?` character, if any.
	// - `URI` : The part of a web request that identifies a resource, for example, `/images/daily-ad.jpg` .
	// - `BODY` : The part of a request that contains any additional data that you want to send to your web server as the HTTP request body, such as data from a form. The request body immediately follows the request headers. Note that only the first `8192` bytes of the request body are forwarded to AWS WAF for inspection. To allow or block requests based on the length of the body, you can create a size constraint set.
	// - `SINGLE_QUERY_ARG` : The parameter in the query string that you will inspect, such as *UserName* or *SalesRegion* . The maximum length for `SINGLE_QUERY_ARG` is 30 characters.
	// - `ALL_QUERY_ARGS` : Similar to `SINGLE_QUERY_ARG` , but rather than inspecting a single parameter, AWS WAF will inspect all parameters within the query for the value or regex pattern that you specify in `TargetString` .
	Type *string `json:"type" yaml:"type"`
	// When the value of `Type` is `HEADER` , enter the name of the header that you want AWS WAF to search, for example, `User-Agent` or `Referer` .
	//
	// The name of the header is not case sensitive.
	//
	// When the value of `Type` is `SINGLE_QUERY_ARG` , enter the name of the parameter that you want AWS WAF to search, for example, `UserName` or `SalesRegion` . The parameter name is not case sensitive.
	//
	// If the value of `Type` is any other value, omit `Data` .
	Data *string `json:"data" yaml:"data"`
}

// > This is *AWS WAF Classic* documentation.
//
// For more information, see [AWS WAF Classic](https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html) in the developer guide.
// >
// > *For the latest version of AWS WAF* , use the AWS WAF V2 API and see the [AWS WAF Developer Guide](https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html) . With the latest version, AWS WAF has a single set of endpoints for regional and global use.
//
// Specifies the part of a web request that you want AWS WAF to inspect for snippets of malicious SQL code and, if you want AWS WAF to inspect a header, the name of the header.
//
// TODO: EXAMPLE
//
type CfnSqlInjectionMatchSet_SqlInjectionMatchTupleProperty struct {
	// The part of a web request that you want AWS WAF to inspect, such as a specific header or a query string.
	FieldToMatch interface{} `json:"fieldToMatch" yaml:"fieldToMatch"`
	// Text transformations eliminate some of the unusual formatting that attackers use in web requests in an effort to bypass AWS WAF .
	//
	// If you specify a transformation, AWS WAF performs the transformation on `FieldToMatch` before inspecting it for a match.
	//
	// You can only specify a single type of TextTransformation.
	//
	// *CMD_LINE*
	//
	// When you're concerned that attackers are injecting an operating system command line command and using unusual formatting to disguise some or all of the command, use this option to perform the following transformations:
	//
	// - Delete the following characters: \ " ' ^
	// - Delete spaces before the following characters: / (
	// - Replace the following characters with a space: , ;
	// - Replace multiple spaces with one space
	// - Convert uppercase letters (A-Z) to lowercase (a-z)
	//
	// *COMPRESS_WHITE_SPACE*
	//
	// Use this option to replace the following characters with a space character (decimal 32):
	//
	// - \f, formfeed, decimal 12
	// - \t, tab, decimal 9
	// - \n, newline, decimal 10
	// - \r, carriage return, decimal 13
	// - \v, vertical tab, decimal 11
	// - non-breaking space, decimal 160
	//
	// `COMPRESS_WHITE_SPACE` also replaces multiple spaces with one space.
	//
	// *HTML_ENTITY_DECODE*
	//
	// Use this option to replace HTML-encoded characters with unencoded characters. `HTML_ENTITY_DECODE` performs the following operations:
	//
	// - Replaces `(ampersand)quot;` with `"`
	// - Replaces `(ampersand)nbsp;` with a non-breaking space, decimal 160
	// - Replaces `(ampersand)lt;` with a "less than" symbol
	// - Replaces `(ampersand)gt;` with `>`
	// - Replaces characters that are represented in hexadecimal format, `(ampersand)#xhhhh;` , with the corresponding characters
	// - Replaces characters that are represented in decimal format, `(ampersand)#nnnn;` , with the corresponding characters
	//
	// *LOWERCASE*
	//
	// Use this option to convert uppercase letters (A-Z) to lowercase (a-z).
	//
	// *URL_DECODE*
	//
	// Use this option to decode a URL-encoded value.
	//
	// *NONE*
	//
	// Specify `NONE` if you don't want to perform any text transformations.
	TextTransformation *string `json:"textTransformation" yaml:"textTransformation"`
}

// Properties for defining a `CfnSqlInjectionMatchSet`.
//
// TODO: EXAMPLE
//
type CfnSqlInjectionMatchSetProps struct {
	// The name, if any, of the `SqlInjectionMatchSet` .
	Name *string `json:"name" yaml:"name"`
	// Specifies the parts of web requests that you want to inspect for snippets of malicious SQL code.
	SqlInjectionMatchTuples interface{} `json:"sqlInjectionMatchTuples" yaml:"sqlInjectionMatchTuples"`
}

// A CloudFormation `AWS::WAFRegional::WebACL`.
//
// > This is *AWS WAF Classic* documentation. For more information, see [AWS WAF Classic](https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html) in the developer guide.
// >
// > *For the latest version of AWS WAF* , use the AWS WAF V2 API and see the [AWS WAF Developer Guide](https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html) . With the latest version, AWS WAF has a single set of endpoints for regional and global use.
//
// Contains the `Rules` that identify the requests that you want to allow, block, or count. In a `WebACL` , you also specify a default action ( `ALLOW` or `BLOCK` ), and the action for each `Rule` that you add to a `WebACL` , for example, block requests from specified IP addresses or block requests from specified referrers. If you add more than one `Rule` to a `WebACL` , a request needs to match only one of the specifications to be allowed, blocked, or counted.
//
// To identify the requests that you want AWS WAF to filter, you associate the `WebACL` with an API Gateway API or an Application Load Balancer.
//
// TODO: EXAMPLE
//
type CfnWebACL interface {
	awscdk.CfnResource
	awscdk.IInspectable
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	DefaultAction() interface{}
	SetDefaultAction(val interface{})
	LogicalId() *string
	MetricName() *string
	SetMetricName(val *string)
	Name() *string
	SetName(val *string)
	Node() constructs.Node
	Ref() *string
	Rules() interface{}
	SetRules(val interface{})
	Stack() awscdk.Stack
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

// The jsii proxy struct for CfnWebACL
type jsiiProxy_CfnWebACL struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnWebACL) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebACL) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebACL) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebACL) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebACL) DefaultAction() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebACL) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebACL) MetricName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebACL) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebACL) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebACL) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebACL) Rules() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebACL) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebACL) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::WAFRegional::WebACL`.
func NewCfnWebACL(scope constructs.Construct, id *string, props *CfnWebACLProps) CfnWebACL {
	_init_.Initialize()

	j := jsiiProxy_CfnWebACL{}

	_jsii_.Create(
		"aws-cdk-lib.aws_wafregional.CfnWebACL",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::WAFRegional::WebACL`.
func NewCfnWebACL_Override(c CfnWebACL, scope constructs.Construct, id *string, props *CfnWebACLProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_wafregional.CfnWebACL",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnWebACL) SetDefaultAction(val interface{}) {
	_jsii_.Set(
		j,
		"defaultAction",
		val,
	)
}

func (j *jsiiProxy_CfnWebACL) SetMetricName(val *string) {
	_jsii_.Set(
		j,
		"metricName",
		val,
	)
}

func (j *jsiiProxy_CfnWebACL) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnWebACL) SetRules(val interface{}) {
	_jsii_.Set(
		j,
		"rules",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnWebACL_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_wafregional.CfnWebACL",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnWebACL_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_wafregional.CfnWebACL",
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
func CfnWebACL_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_wafregional.CfnWebACL",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnWebACL_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_wafregional.CfnWebACL",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnWebACL) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnWebACL) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnWebACL) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnWebACL) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnWebACL) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnWebACL) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnWebACL) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnWebACL) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnWebACL) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnWebACL) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnWebACL) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnWebACL) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnWebACL) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnWebACL) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnWebACL) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Specifies the action AWS WAF takes when a web request matches or doesn't match all rule conditions.
//
// TODO: EXAMPLE
//
type CfnWebACL_ActionProperty struct {
	// For actions that are associated with a rule, the action that AWS WAF takes when a web request matches all conditions in a rule.
	//
	// For the default action of a web access control list (ACL), the action that AWS WAF takes when a web request doesn't match all conditions in any rule.
	//
	// Valid settings include the following:
	//
	// - `ALLOW` : AWS WAF allows requests
	// - `BLOCK` : AWS WAF blocks requests
	// - `COUNT` : AWS WAF increments a counter of the requests that match all of the conditions in the rule. AWS WAF then continues to inspect the web request based on the remaining rules in the web ACL. You can't specify `COUNT` for the default action for a WebACL.
	Type *string `json:"type" yaml:"type"`
}

// A combination of `ByteMatchSet` , `IPSet` , and/or `SqlInjectionMatchSet` objects that identify the web requests that you want to allow, block, or count.
//
// For example, you might create a `Rule` that includes the following predicates:
//
// - An `IPSet` that causes AWS WAF to search for web requests that originate from the IP address `192.0.2.44`
// - A `ByteMatchSet` that causes AWS WAF to search for web requests for which the value of the `User-Agent` header is `BadBot` .
//
// To match the settings in this `Rule` , a request must originate from `192.0.2.44` AND include a `User-Agent` header for which the value is `BadBot` .
//
// TODO: EXAMPLE
//
type CfnWebACL_RuleProperty struct {
	// The action that AWS WAF takes when a web request matches all conditions in the rule, such as allow, block, or count the request.
	Action interface{} `json:"action" yaml:"action"`
	// The order in which AWS WAF evaluates the rules in a web ACL.
	//
	// AWS WAF evaluates rules with a lower value before rules with a higher value. The value must be a unique integer. If you have multiple rules in a web ACL, the priority numbers do not need to be consecutive.
	Priority *float64 `json:"priority" yaml:"priority"`
	// The ID of an AWS WAF Regional rule to associate with a web ACL.
	RuleId *string `json:"ruleId" yaml:"ruleId"`
}

// A CloudFormation `AWS::WAFRegional::WebACLAssociation`.
//
// > This is *AWS WAF Classic* documentation. For more information, see [AWS WAF Classic](https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html) in the developer guide.
// >
// > *For the latest version of AWS WAF* , use the AWS WAF V2 API and see the [AWS WAF Developer Guide](https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html) . With the latest version, AWS WAF has a single set of endpoints for regional and global use.
//
// The AWS::WAFRegional::WebACLAssociation resource associates an AWS WAF Regional web access control group (ACL) with a resource.
//
// TODO: EXAMPLE
//
type CfnWebACLAssociation interface {
	awscdk.CfnResource
	awscdk.IInspectable
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	LogicalId() *string
	Node() constructs.Node
	Ref() *string
	ResourceArn() *string
	SetResourceArn(val *string)
	Stack() awscdk.Stack
	UpdatedProperites() *map[string]interface{}
	WebAclId() *string
	SetWebAclId(val *string)
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

// The jsii proxy struct for CfnWebACLAssociation
type jsiiProxy_CfnWebACLAssociation struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnWebACLAssociation) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebACLAssociation) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebACLAssociation) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebACLAssociation) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebACLAssociation) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebACLAssociation) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebACLAssociation) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebACLAssociation) ResourceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebACLAssociation) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebACLAssociation) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebACLAssociation) WebAclId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webAclId",
		&returns,
	)
	return returns
}


// Create a new `AWS::WAFRegional::WebACLAssociation`.
func NewCfnWebACLAssociation(scope constructs.Construct, id *string, props *CfnWebACLAssociationProps) CfnWebACLAssociation {
	_init_.Initialize()

	j := jsiiProxy_CfnWebACLAssociation{}

	_jsii_.Create(
		"aws-cdk-lib.aws_wafregional.CfnWebACLAssociation",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::WAFRegional::WebACLAssociation`.
func NewCfnWebACLAssociation_Override(c CfnWebACLAssociation, scope constructs.Construct, id *string, props *CfnWebACLAssociationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_wafregional.CfnWebACLAssociation",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnWebACLAssociation) SetResourceArn(val *string) {
	_jsii_.Set(
		j,
		"resourceArn",
		val,
	)
}

func (j *jsiiProxy_CfnWebACLAssociation) SetWebAclId(val *string) {
	_jsii_.Set(
		j,
		"webAclId",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnWebACLAssociation_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_wafregional.CfnWebACLAssociation",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnWebACLAssociation_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_wafregional.CfnWebACLAssociation",
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
func CfnWebACLAssociation_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_wafregional.CfnWebACLAssociation",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnWebACLAssociation_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_wafregional.CfnWebACLAssociation",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnWebACLAssociation) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnWebACLAssociation) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnWebACLAssociation) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnWebACLAssociation) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnWebACLAssociation) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnWebACLAssociation) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnWebACLAssociation) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnWebACLAssociation) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnWebACLAssociation) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnWebACLAssociation) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnWebACLAssociation) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnWebACLAssociation) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnWebACLAssociation) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnWebACLAssociation) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnWebACLAssociation) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnWebACLAssociation`.
//
// TODO: EXAMPLE
//
type CfnWebACLAssociationProps struct {
	// The Amazon Resource Name (ARN) of the resource to protect with the web ACL.
	ResourceArn *string `json:"resourceArn" yaml:"resourceArn"`
	// A unique identifier (ID) for the web ACL.
	WebAclId *string `json:"webAclId" yaml:"webAclId"`
}

// Properties for defining a `CfnWebACL`.
//
// TODO: EXAMPLE
//
type CfnWebACLProps struct {
	// The action to perform if none of the `Rules` contained in the `WebACL` match.
	//
	// The action is specified by the `WafAction` object.
	DefaultAction interface{} `json:"defaultAction" yaml:"defaultAction"`
	// A name for the metrics for this `WebACL` .
	//
	// The name can contain only alphanumeric characters (A-Z, a-z, 0-9), with maximum length 128 and minimum length one. It can't contain whitespace or metric names reserved for AWS WAF, including "All" and "Default_Action." You can't change `MetricName` after you create the `WebACL` .
	MetricName *string `json:"metricName" yaml:"metricName"`
	// A friendly name or description of the `WebACL` .
	//
	// You can't change the name of a `WebACL` after you create it.
	Name *string `json:"name" yaml:"name"`
	// An array that contains the action for each `Rule` in a `WebACL` , the priority of the `Rule` , and the ID of the `Rule` .
	Rules interface{} `json:"rules" yaml:"rules"`
}

// A CloudFormation `AWS::WAFRegional::XssMatchSet`.
//
// > This is *AWS WAF Classic* documentation. For more information, see [AWS WAF Classic](https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html) in the developer guide.
// >
// > *For the latest version of AWS WAF* , use the AWS WAF V2 API and see the [AWS WAF Developer Guide](https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html) . With the latest version, AWS WAF has a single set of endpoints for regional and global use.
//
// A complex type that contains `XssMatchTuple` objects, which specify the parts of web requests that you want AWS WAF to inspect for cross-site scripting attacks and, if you want AWS WAF to inspect a header, the name of the header. If a `XssMatchSet` contains more than one `XssMatchTuple` object, a request needs to include cross-site scripting attacks in only one of the specified parts of the request to be considered a match.
//
// TODO: EXAMPLE
//
type CfnXssMatchSet interface {
	awscdk.CfnResource
	awscdk.IInspectable
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() constructs.Node
	Ref() *string
	Stack() awscdk.Stack
	UpdatedProperites() *map[string]interface{}
	XssMatchTuples() interface{}
	SetXssMatchTuples(val interface{})
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

// The jsii proxy struct for CfnXssMatchSet
type jsiiProxy_CfnXssMatchSet struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnXssMatchSet) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnXssMatchSet) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnXssMatchSet) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnXssMatchSet) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnXssMatchSet) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnXssMatchSet) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnXssMatchSet) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnXssMatchSet) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnXssMatchSet) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnXssMatchSet) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnXssMatchSet) XssMatchTuples() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"xssMatchTuples",
		&returns,
	)
	return returns
}


// Create a new `AWS::WAFRegional::XssMatchSet`.
func NewCfnXssMatchSet(scope constructs.Construct, id *string, props *CfnXssMatchSetProps) CfnXssMatchSet {
	_init_.Initialize()

	j := jsiiProxy_CfnXssMatchSet{}

	_jsii_.Create(
		"aws-cdk-lib.aws_wafregional.CfnXssMatchSet",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::WAFRegional::XssMatchSet`.
func NewCfnXssMatchSet_Override(c CfnXssMatchSet, scope constructs.Construct, id *string, props *CfnXssMatchSetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_wafregional.CfnXssMatchSet",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnXssMatchSet) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnXssMatchSet) SetXssMatchTuples(val interface{}) {
	_jsii_.Set(
		j,
		"xssMatchTuples",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnXssMatchSet_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_wafregional.CfnXssMatchSet",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnXssMatchSet_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_wafregional.CfnXssMatchSet",
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
func CfnXssMatchSet_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_wafregional.CfnXssMatchSet",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnXssMatchSet_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_wafregional.CfnXssMatchSet",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnXssMatchSet) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnXssMatchSet) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnXssMatchSet) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnXssMatchSet) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnXssMatchSet) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnXssMatchSet) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnXssMatchSet) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnXssMatchSet) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnXssMatchSet) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnXssMatchSet) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnXssMatchSet) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnXssMatchSet) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnXssMatchSet) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnXssMatchSet) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnXssMatchSet) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The part of a web request that you want AWS WAF to inspect, such as a specific header or a query string.
//
// TODO: EXAMPLE
//
type CfnXssMatchSet_FieldToMatchProperty struct {
	// The part of the web request that you want AWS WAF to search for a specified string.
	//
	// Parts of a request that you can search include the following:
	//
	// - `HEADER` : A specified request header, for example, the value of the `User-Agent` or `Referer` header. If you choose `HEADER` for the type, specify the name of the header in `Data` .
	// - `METHOD` : The HTTP method, which indicates the type of operation that the request is asking the origin to perform.
	// - `QUERY_STRING` : A query string, which is the part of a URL that appears after a `?` character, if any.
	// - `URI` : The part of a web request that identifies a resource, for example, `/images/daily-ad.jpg` .
	// - `BODY` : The part of a request that contains any additional data that you want to send to your web server as the HTTP request body, such as data from a form. The request body immediately follows the request headers. Note that only the first `8192` bytes of the request body are forwarded to AWS WAF for inspection. To allow or block requests based on the length of the body, you can create a size constraint set.
	// - `SINGLE_QUERY_ARG` : The parameter in the query string that you will inspect, such as *UserName* or *SalesRegion* . The maximum length for `SINGLE_QUERY_ARG` is 30 characters.
	// - `ALL_QUERY_ARGS` : Similar to `SINGLE_QUERY_ARG` , but rather than inspecting a single parameter, AWS WAF will inspect all parameters within the query for the value or regex pattern that you specify in `TargetString` .
	Type *string `json:"type" yaml:"type"`
	// When the value of `Type` is `HEADER` , enter the name of the header that you want AWS WAF to search, for example, `User-Agent` or `Referer` .
	//
	// The name of the header is not case sensitive.
	//
	// When the value of `Type` is `SINGLE_QUERY_ARG` , enter the name of the parameter that you want AWS WAF to search, for example, `UserName` or `SalesRegion` . The parameter name is not case sensitive.
	//
	// If the value of `Type` is any other value, omit `Data` .
	Data *string `json:"data" yaml:"data"`
}

// > This is *AWS WAF Classic* documentation.
//
// For more information, see [AWS WAF Classic](https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html) in the developer guide.
// >
// > *For the latest version of AWS WAF* , use the AWS WAF V2 API and see the [AWS WAF Developer Guide](https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html) . With the latest version, AWS WAF has a single set of endpoints for regional and global use.
//
// Specifies the part of a web request that you want AWS WAF to inspect for cross-site scripting attacks and, if you want AWS WAF to inspect a header, the name of the header.
//
// TODO: EXAMPLE
//
type CfnXssMatchSet_XssMatchTupleProperty struct {
	// The part of a web request that you want AWS WAF to inspect, such as a specified header or a query string.
	FieldToMatch interface{} `json:"fieldToMatch" yaml:"fieldToMatch"`
	// Text transformations eliminate some of the unusual formatting that attackers use in web requests in an effort to bypass AWS WAF .
	//
	// If you specify a transformation, AWS WAF performs the transformation on `FieldToMatch` before inspecting it for a match.
	//
	// You can only specify a single type of TextTransformation.
	//
	// *CMD_LINE*
	//
	// When you're concerned that attackers are injecting an operating system command line command and using unusual formatting to disguise some or all of the command, use this option to perform the following transformations:
	//
	// - Delete the following characters: \ " ' ^
	// - Delete spaces before the following characters: / (
	// - Replace the following characters with a space: , ;
	// - Replace multiple spaces with one space
	// - Convert uppercase letters (A-Z) to lowercase (a-z)
	//
	// *COMPRESS_WHITE_SPACE*
	//
	// Use this option to replace the following characters with a space character (decimal 32):
	//
	// - \f, formfeed, decimal 12
	// - \t, tab, decimal 9
	// - \n, newline, decimal 10
	// - \r, carriage return, decimal 13
	// - \v, vertical tab, decimal 11
	// - non-breaking space, decimal 160
	//
	// `COMPRESS_WHITE_SPACE` also replaces multiple spaces with one space.
	//
	// *HTML_ENTITY_DECODE*
	//
	// Use this option to replace HTML-encoded characters with unencoded characters. `HTML_ENTITY_DECODE` performs the following operations:
	//
	// - Replaces `(ampersand)quot;` with `"`
	// - Replaces `(ampersand)nbsp;` with a non-breaking space, decimal 160
	// - Replaces `(ampersand)lt;` with a "less than" symbol
	// - Replaces `(ampersand)gt;` with `>`
	// - Replaces characters that are represented in hexadecimal format, `(ampersand)#xhhhh;` , with the corresponding characters
	// - Replaces characters that are represented in decimal format, `(ampersand)#nnnn;` , with the corresponding characters
	//
	// *LOWERCASE*
	//
	// Use this option to convert uppercase letters (A-Z) to lowercase (a-z).
	//
	// *URL_DECODE*
	//
	// Use this option to decode a URL-encoded value.
	//
	// *NONE*
	//
	// Specify `NONE` if you don't want to perform any text transformations.
	TextTransformation *string `json:"textTransformation" yaml:"textTransformation"`
}

// Properties for defining a `CfnXssMatchSet`.
//
// TODO: EXAMPLE
//
type CfnXssMatchSetProps struct {
	// The name, if any, of the `XssMatchSet` .
	Name *string `json:"name" yaml:"name"`
	// Specifies the parts of web requests that you want to inspect for cross-site scripting attacks.
	XssMatchTuples interface{} `json:"xssMatchTuples" yaml:"xssMatchTuples"`
}

