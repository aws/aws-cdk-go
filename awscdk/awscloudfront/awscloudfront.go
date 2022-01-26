package awscloudfront

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscertificatemanager"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
)

// Options for adding a new behavior to a Distribution.
//
// TODO: EXAMPLE
//
type AddBehaviorOptions struct {
	// HTTP methods to allow for this behavior.
	AllowedMethods AllowedMethods `json:"allowedMethods" yaml:"allowedMethods"`
	// HTTP methods to cache for this behavior.
	CachedMethods CachedMethods `json:"cachedMethods" yaml:"cachedMethods"`
	// The cache policy for this behavior.
	//
	// The cache policy determines what values are included in the cache key,
	// and the time-to-live (TTL) values for the cache.
	// See: https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-the-cache-key.html.
	//
	CachePolicy ICachePolicy `json:"cachePolicy" yaml:"cachePolicy"`
	// Whether you want CloudFront to automatically compress certain files for this cache behavior.
	//
	// See https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/ServingCompressedFiles.html#compressed-content-cloudfront-file-types
	// for file types CloudFront will compress.
	Compress *bool `json:"compress" yaml:"compress"`
	// The Lambda@Edge functions to invoke before serving the contents.
	// See: https://aws.amazon.com/lambda/edge
	//
	EdgeLambdas *[]*EdgeLambda `json:"edgeLambdas" yaml:"edgeLambdas"`
	// The CloudFront functions to invoke before serving the contents.
	FunctionAssociations *[]*FunctionAssociation `json:"functionAssociations" yaml:"functionAssociations"`
	// The origin request policy for this behavior.
	//
	// The origin request policy determines which values (e.g., headers, cookies)
	// are included in requests that CloudFront sends to the origin.
	OriginRequestPolicy IOriginRequestPolicy `json:"originRequestPolicy" yaml:"originRequestPolicy"`
	// The response headers policy for this behavior.
	//
	// The response headers policy determines which headers are included in responses
	ResponseHeadersPolicy IResponseHeadersPolicy `json:"responseHeadersPolicy" yaml:"responseHeadersPolicy"`
	// Set this to true to indicate you want to distribute media files in the Microsoft Smooth Streaming format using this behavior.
	SmoothStreaming *bool `json:"smoothStreaming" yaml:"smoothStreaming"`
	// A list of Key Groups that CloudFront can use to validate signed URLs or signed cookies.
	// See: https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/PrivateContent.html
	//
	TrustedKeyGroups *[]IKeyGroup `json:"trustedKeyGroups" yaml:"trustedKeyGroups"`
	// The protocol that viewers can use to access the files controlled by this behavior.
	ViewerProtocolPolicy ViewerProtocolPolicy `json:"viewerProtocolPolicy" yaml:"viewerProtocolPolicy"`
}

// The HTTP methods that the Behavior will accept requests on.
//
// TODO: EXAMPLE
//
type AllowedMethods interface {
	Methods() *[]*string
}

// The jsii proxy struct for AllowedMethods
type jsiiProxy_AllowedMethods struct {
	_ byte // padding
}

func (j *jsiiProxy_AllowedMethods) Methods() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"methods",
		&returns,
	)
	return returns
}


func AllowedMethods_ALLOW_ALL() AllowedMethods {
	_init_.Initialize()
	var returns AllowedMethods
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cloudfront.AllowedMethods",
		"ALLOW_ALL",
		&returns,
	)
	return returns
}

func AllowedMethods_ALLOW_GET_HEAD() AllowedMethods {
	_init_.Initialize()
	var returns AllowedMethods
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cloudfront.AllowedMethods",
		"ALLOW_GET_HEAD",
		&returns,
	)
	return returns
}

func AllowedMethods_ALLOW_GET_HEAD_OPTIONS() AllowedMethods {
	_init_.Initialize()
	var returns AllowedMethods
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cloudfront.AllowedMethods",
		"ALLOW_GET_HEAD_OPTIONS",
		&returns,
	)
	return returns
}

// A CloudFront behavior wrapper.
//
// TODO: EXAMPLE
//
type Behavior struct {
	// The method this CloudFront distribution responds do.
	AllowedMethods CloudFrontAllowedMethods `json:"allowedMethods" yaml:"allowedMethods"`
	// Which methods are cached by CloudFront by default.
	CachedMethods CloudFrontAllowedCachedMethods `json:"cachedMethods" yaml:"cachedMethods"`
	// If CloudFront should automatically compress some content types.
	Compress *bool `json:"compress" yaml:"compress"`
	// The default amount of time CloudFront will cache an object.
	//
	// This value applies only when your custom origin does not add HTTP headers,
	// such as Cache-Control max-age, Cache-Control s-maxage, and Expires to objects.
	DefaultTtl awscdk.Duration `json:"defaultTtl" yaml:"defaultTtl"`
	// The values CloudFront will forward to the origin when making a request.
	ForwardedValues *CfnDistribution_ForwardedValuesProperty `json:"forwardedValues" yaml:"forwardedValues"`
	// The CloudFront functions to invoke before serving the contents.
	FunctionAssociations *[]*FunctionAssociation `json:"functionAssociations" yaml:"functionAssociations"`
	// If this behavior is the default behavior for the distribution.
	//
	// You must specify exactly one default distribution per CloudFront distribution.
	// The default behavior is allowed to omit the "path" property.
	IsDefaultBehavior *bool `json:"isDefaultBehavior" yaml:"isDefaultBehavior"`
	// Declares associated lambda@edge functions for this distribution behaviour.
	LambdaFunctionAssociations *[]*LambdaFunctionAssociation `json:"lambdaFunctionAssociations" yaml:"lambdaFunctionAssociations"`
	// The max amount of time you want objects to stay in the cache before CloudFront queries your origin.
	MaxTtl awscdk.Duration `json:"maxTtl" yaml:"maxTtl"`
	// The minimum amount of time that you want objects to stay in the cache before CloudFront queries your origin.
	MinTtl awscdk.Duration `json:"minTtl" yaml:"minTtl"`
	// The path this behavior responds to.
	//
	// Required for all non-default behaviors. (The default behavior implicitly has "*" as the path pattern. )
	PathPattern *string `json:"pathPattern" yaml:"pathPattern"`
	// A list of Key Groups that CloudFront can use to validate signed URLs or signed cookies.
	// See: https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/PrivateContent.html
	//
	TrustedKeyGroups *[]IKeyGroup `json:"trustedKeyGroups" yaml:"trustedKeyGroups"`
	// The viewer policy for this behavior.
	ViewerProtocolPolicy ViewerProtocolPolicy `json:"viewerProtocolPolicy" yaml:"viewerProtocolPolicy"`
}

// Options for creating a new behavior.
//
// TODO: EXAMPLE
//
type BehaviorOptions struct {
	// HTTP methods to allow for this behavior.
	AllowedMethods AllowedMethods `json:"allowedMethods" yaml:"allowedMethods"`
	// HTTP methods to cache for this behavior.
	CachedMethods CachedMethods `json:"cachedMethods" yaml:"cachedMethods"`
	// The cache policy for this behavior.
	//
	// The cache policy determines what values are included in the cache key,
	// and the time-to-live (TTL) values for the cache.
	// See: https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-the-cache-key.html.
	//
	CachePolicy ICachePolicy `json:"cachePolicy" yaml:"cachePolicy"`
	// Whether you want CloudFront to automatically compress certain files for this cache behavior.
	//
	// See https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/ServingCompressedFiles.html#compressed-content-cloudfront-file-types
	// for file types CloudFront will compress.
	Compress *bool `json:"compress" yaml:"compress"`
	// The Lambda@Edge functions to invoke before serving the contents.
	// See: https://aws.amazon.com/lambda/edge
	//
	EdgeLambdas *[]*EdgeLambda `json:"edgeLambdas" yaml:"edgeLambdas"`
	// The CloudFront functions to invoke before serving the contents.
	FunctionAssociations *[]*FunctionAssociation `json:"functionAssociations" yaml:"functionAssociations"`
	// The origin request policy for this behavior.
	//
	// The origin request policy determines which values (e.g., headers, cookies)
	// are included in requests that CloudFront sends to the origin.
	OriginRequestPolicy IOriginRequestPolicy `json:"originRequestPolicy" yaml:"originRequestPolicy"`
	// The response headers policy for this behavior.
	//
	// The response headers policy determines which headers are included in responses
	ResponseHeadersPolicy IResponseHeadersPolicy `json:"responseHeadersPolicy" yaml:"responseHeadersPolicy"`
	// Set this to true to indicate you want to distribute media files in the Microsoft Smooth Streaming format using this behavior.
	SmoothStreaming *bool `json:"smoothStreaming" yaml:"smoothStreaming"`
	// A list of Key Groups that CloudFront can use to validate signed URLs or signed cookies.
	// See: https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/PrivateContent.html
	//
	TrustedKeyGroups *[]IKeyGroup `json:"trustedKeyGroups" yaml:"trustedKeyGroups"`
	// The protocol that viewers can use to access the files controlled by this behavior.
	ViewerProtocolPolicy ViewerProtocolPolicy `json:"viewerProtocolPolicy" yaml:"viewerProtocolPolicy"`
	// The origin that you want CloudFront to route requests to when they match this behavior.
	Origin IOrigin `json:"origin" yaml:"origin"`
}

// Determines whether any cookies in viewer requests are included in the cache key and automatically included in requests that CloudFront sends to the origin.
//
// TODO: EXAMPLE
//
type CacheCookieBehavior interface {
	Behavior() *string
	Cookies() *[]*string
}

// The jsii proxy struct for CacheCookieBehavior
type jsiiProxy_CacheCookieBehavior struct {
	_ byte // padding
}

func (j *jsiiProxy_CacheCookieBehavior) Behavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"behavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CacheCookieBehavior) Cookies() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cookies",
		&returns,
	)
	return returns
}


// All cookies in viewer requests are included in the cache key and are automatically included in requests that CloudFront sends to the origin.
func CacheCookieBehavior_All() CacheCookieBehavior {
	_init_.Initialize()

	var returns CacheCookieBehavior

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.CacheCookieBehavior",
		"all",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Only the provided `cookies` are included in the cache key and automatically included in requests that CloudFront sends to the origin.
func CacheCookieBehavior_AllowList(cookies ...*string) CacheCookieBehavior {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range cookies {
		args = append(args, a)
	}

	var returns CacheCookieBehavior

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.CacheCookieBehavior",
		"allowList",
		args,
		&returns,
	)

	return returns
}

// All cookies except the provided `cookies` are included in the cache key and automatically included in requests that CloudFront sends to the origin.
func CacheCookieBehavior_DenyList(cookies ...*string) CacheCookieBehavior {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range cookies {
		args = append(args, a)
	}

	var returns CacheCookieBehavior

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.CacheCookieBehavior",
		"denyList",
		args,
		&returns,
	)

	return returns
}

// Cookies in viewer requests are not included in the cache key and are not automatically included in requests that CloudFront sends to the origin.
func CacheCookieBehavior_None() CacheCookieBehavior {
	_init_.Initialize()

	var returns CacheCookieBehavior

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.CacheCookieBehavior",
		"none",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Determines whether any HTTP headers are included in the cache key and automatically included in requests that CloudFront sends to the origin.
//
// TODO: EXAMPLE
//
type CacheHeaderBehavior interface {
	Behavior() *string
	Headers() *[]*string
}

// The jsii proxy struct for CacheHeaderBehavior
type jsiiProxy_CacheHeaderBehavior struct {
	_ byte // padding
}

func (j *jsiiProxy_CacheHeaderBehavior) Behavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"behavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CacheHeaderBehavior) Headers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"headers",
		&returns,
	)
	return returns
}


// Listed headers are included in the cache key and are automatically included in requests that CloudFront sends to the origin.
func CacheHeaderBehavior_AllowList(headers ...*string) CacheHeaderBehavior {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range headers {
		args = append(args, a)
	}

	var returns CacheHeaderBehavior

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.CacheHeaderBehavior",
		"allowList",
		args,
		&returns,
	)

	return returns
}

// HTTP headers are not included in the cache key and are not automatically included in requests that CloudFront sends to the origin.
func CacheHeaderBehavior_None() CacheHeaderBehavior {
	_init_.Initialize()

	var returns CacheHeaderBehavior

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.CacheHeaderBehavior",
		"none",
		nil, // no parameters
		&returns,
	)

	return returns
}

// A Cache Policy configuration.
//
// TODO: EXAMPLE
//
type CachePolicy interface {
	awscdk.Resource
	ICachePolicy
	CachePolicyId() *string
	Env() *awscdk.ResourceEnvironment
	Node() constructs.Node
	PhysicalName() *string
	Stack() awscdk.Stack
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ToString() *string
}

// The jsii proxy struct for CachePolicy
type jsiiProxy_CachePolicy struct {
	internal.Type__awscdkResource
	jsiiProxy_ICachePolicy
}

func (j *jsiiProxy_CachePolicy) CachePolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cachePolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CachePolicy) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CachePolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CachePolicy) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CachePolicy) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewCachePolicy(scope constructs.Construct, id *string, props *CachePolicyProps) CachePolicy {
	_init_.Initialize()

	j := jsiiProxy_CachePolicy{}

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront.CachePolicy",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCachePolicy_Override(c CachePolicy, scope constructs.Construct, id *string, props *CachePolicyProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront.CachePolicy",
		[]interface{}{scope, id, props},
		c,
	)
}

// Imports a Cache Policy from its id.
func CachePolicy_FromCachePolicyId(scope constructs.Construct, id *string, cachePolicyId *string) ICachePolicy {
	_init_.Initialize()

	var returns ICachePolicy

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.CachePolicy",
		"fromCachePolicyId",
		[]interface{}{scope, id, cachePolicyId},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CachePolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.CachePolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func CachePolicy_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.CachePolicy",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func CachePolicy_AMPLIFY() ICachePolicy {
	_init_.Initialize()
	var returns ICachePolicy
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cloudfront.CachePolicy",
		"AMPLIFY",
		&returns,
	)
	return returns
}

func CachePolicy_CACHING_DISABLED() ICachePolicy {
	_init_.Initialize()
	var returns ICachePolicy
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cloudfront.CachePolicy",
		"CACHING_DISABLED",
		&returns,
	)
	return returns
}

func CachePolicy_CACHING_OPTIMIZED() ICachePolicy {
	_init_.Initialize()
	var returns ICachePolicy
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cloudfront.CachePolicy",
		"CACHING_OPTIMIZED",
		&returns,
	)
	return returns
}

func CachePolicy_CACHING_OPTIMIZED_FOR_UNCOMPRESSED_OBJECTS() ICachePolicy {
	_init_.Initialize()
	var returns ICachePolicy
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cloudfront.CachePolicy",
		"CACHING_OPTIMIZED_FOR_UNCOMPRESSED_OBJECTS",
		&returns,
	)
	return returns
}

func CachePolicy_ELEMENTAL_MEDIA_PACKAGE() ICachePolicy {
	_init_.Initialize()
	var returns ICachePolicy
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cloudfront.CachePolicy",
		"ELEMENTAL_MEDIA_PACKAGE",
		&returns,
	)
	return returns
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (c *jsiiProxy_CachePolicy) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (c *jsiiProxy_CachePolicy) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
func (c *jsiiProxy_CachePolicy) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
func (c *jsiiProxy_CachePolicy) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (c *jsiiProxy_CachePolicy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for creating a Cache Policy.
//
// TODO: EXAMPLE
//
type CachePolicyProps struct {
	// A unique name to identify the cache policy.
	//
	// The name must only include '-', '_', or alphanumeric characters.
	CachePolicyName *string `json:"cachePolicyName" yaml:"cachePolicyName"`
	// A comment to describe the cache policy.
	Comment *string `json:"comment" yaml:"comment"`
	// Determines whether any cookies in viewer requests are included in the cache key and automatically included in requests that CloudFront sends to the origin.
	CookieBehavior CacheCookieBehavior `json:"cookieBehavior" yaml:"cookieBehavior"`
	// The default amount of time for objects to stay in the CloudFront cache.
	//
	// Only used when the origin does not send Cache-Control or Expires headers with the object.
	DefaultTtl awscdk.Duration `json:"defaultTtl" yaml:"defaultTtl"`
	// Whether to normalize and include the `Accept-Encoding` header in the cache key when the `Accept-Encoding` header is 'br'.
	EnableAcceptEncodingBrotli *bool `json:"enableAcceptEncodingBrotli" yaml:"enableAcceptEncodingBrotli"`
	// Whether to normalize and include the `Accept-Encoding` header in the cache key when the `Accept-Encoding` header is 'gzip'.
	EnableAcceptEncodingGzip *bool `json:"enableAcceptEncodingGzip" yaml:"enableAcceptEncodingGzip"`
	// Determines whether any HTTP headers are included in the cache key and automatically included in requests that CloudFront sends to the origin.
	HeaderBehavior CacheHeaderBehavior `json:"headerBehavior" yaml:"headerBehavior"`
	// The maximum amount of time for objects to stay in the CloudFront cache.
	//
	// CloudFront uses this value only when the origin sends Cache-Control or Expires headers with the object.
	MaxTtl awscdk.Duration `json:"maxTtl" yaml:"maxTtl"`
	// The minimum amount of time for objects to stay in the CloudFront cache.
	MinTtl awscdk.Duration `json:"minTtl" yaml:"minTtl"`
	// Determines whether any query strings are included in the cache key and automatically included in requests that CloudFront sends to the origin.
	QueryStringBehavior CacheQueryStringBehavior `json:"queryStringBehavior" yaml:"queryStringBehavior"`
}

// Determines whether any URL query strings in viewer requests are included in the cache key and automatically included in requests that CloudFront sends to the origin.
//
// TODO: EXAMPLE
//
type CacheQueryStringBehavior interface {
	Behavior() *string
	QueryStrings() *[]*string
}

// The jsii proxy struct for CacheQueryStringBehavior
type jsiiProxy_CacheQueryStringBehavior struct {
	_ byte // padding
}

func (j *jsiiProxy_CacheQueryStringBehavior) Behavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"behavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CacheQueryStringBehavior) QueryStrings() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"queryStrings",
		&returns,
	)
	return returns
}


// All query strings in viewer requests are included in the cache key and are automatically included in requests that CloudFront sends to the origin.
func CacheQueryStringBehavior_All() CacheQueryStringBehavior {
	_init_.Initialize()

	var returns CacheQueryStringBehavior

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.CacheQueryStringBehavior",
		"all",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Only the provided `queryStrings` are included in the cache key and automatically included in requests that CloudFront sends to the origin.
func CacheQueryStringBehavior_AllowList(queryStrings ...*string) CacheQueryStringBehavior {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range queryStrings {
		args = append(args, a)
	}

	var returns CacheQueryStringBehavior

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.CacheQueryStringBehavior",
		"allowList",
		args,
		&returns,
	)

	return returns
}

// All query strings except the provided `queryStrings` are included in the cache key and automatically included in requests that CloudFront sends to the origin.
func CacheQueryStringBehavior_DenyList(queryStrings ...*string) CacheQueryStringBehavior {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range queryStrings {
		args = append(args, a)
	}

	var returns CacheQueryStringBehavior

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.CacheQueryStringBehavior",
		"denyList",
		args,
		&returns,
	)

	return returns
}

// Query strings in viewer requests are not included in the cache key and are not automatically included in requests that CloudFront sends to the origin.
func CacheQueryStringBehavior_None() CacheQueryStringBehavior {
	_init_.Initialize()

	var returns CacheQueryStringBehavior

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.CacheQueryStringBehavior",
		"none",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The HTTP methods that the Behavior will cache requests on.
//
// TODO: EXAMPLE
//
type CachedMethods interface {
	Methods() *[]*string
}

// The jsii proxy struct for CachedMethods
type jsiiProxy_CachedMethods struct {
	_ byte // padding
}

func (j *jsiiProxy_CachedMethods) Methods() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"methods",
		&returns,
	)
	return returns
}


func CachedMethods_CACHE_GET_HEAD() CachedMethods {
	_init_.Initialize()
	var returns CachedMethods
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cloudfront.CachedMethods",
		"CACHE_GET_HEAD",
		&returns,
	)
	return returns
}

func CachedMethods_CACHE_GET_HEAD_OPTIONS() CachedMethods {
	_init_.Initialize()
	var returns CachedMethods
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cloudfront.CachedMethods",
		"CACHE_GET_HEAD_OPTIONS",
		&returns,
	)
	return returns
}

// A CloudFormation `AWS::CloudFront::CachePolicy`.
//
// A cache policy.
//
// When it’s attached to a cache behavior, the cache policy determines the following:
//
// - The values that CloudFront includes in the cache key. These values can include HTTP headers, cookies, and URL query strings. CloudFront uses the cache key to find an object in its cache that it can return to the viewer.
// - The default, minimum, and maximum time to live (TTL) values that you want objects to stay in the CloudFront cache.
//
// The headers, cookies, and query strings that are included in the cache key are automatically included in requests that CloudFront sends to the origin. CloudFront sends a request when it can’t find a valid object in its cache that matches the request’s cache key. If you want to send values to the origin but *not* include them in the cache key, use `OriginRequestPolicy` .
//
// TODO: EXAMPLE
//
type CfnCachePolicy interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrId() *string
	AttrLastModifiedTime() *string
	CachePolicyConfig() interface{}
	SetCachePolicyConfig(val interface{})
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	LogicalId() *string
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

// The jsii proxy struct for CfnCachePolicy
type jsiiProxy_CfnCachePolicy struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnCachePolicy) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCachePolicy) AttrLastModifiedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLastModifiedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCachePolicy) CachePolicyConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cachePolicyConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCachePolicy) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCachePolicy) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCachePolicy) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCachePolicy) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCachePolicy) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCachePolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCachePolicy) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCachePolicy) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCachePolicy) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::CloudFront::CachePolicy`.
func NewCfnCachePolicy(scope constructs.Construct, id *string, props *CfnCachePolicyProps) CfnCachePolicy {
	_init_.Initialize()

	j := jsiiProxy_CfnCachePolicy{}

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront.CfnCachePolicy",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::CloudFront::CachePolicy`.
func NewCfnCachePolicy_Override(c CfnCachePolicy, scope constructs.Construct, id *string, props *CfnCachePolicyProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront.CfnCachePolicy",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnCachePolicy) SetCachePolicyConfig(val interface{}) {
	_jsii_.Set(
		j,
		"cachePolicyConfig",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnCachePolicy_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.CfnCachePolicy",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnCachePolicy_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.CfnCachePolicy",
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
func CfnCachePolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.CfnCachePolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCachePolicy_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cloudfront.CfnCachePolicy",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnCachePolicy) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnCachePolicy) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnCachePolicy) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnCachePolicy) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnCachePolicy) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnCachePolicy) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnCachePolicy) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnCachePolicy) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnCachePolicy) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnCachePolicy) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnCachePolicy) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnCachePolicy) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnCachePolicy) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnCachePolicy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCachePolicy) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// A cache policy configuration.
//
// This configuration determines the following:
//
// - The values that CloudFront includes in the cache key. These values can include HTTP headers, cookies, and URL query strings. CloudFront uses the cache key to find an object in its cache that it can return to the viewer.
// - The default, minimum, and maximum time to live (TTL) values that you want objects to stay in the CloudFront cache.
//
// The headers, cookies, and query strings that are included in the cache key are automatically included in requests that CloudFront sends to the origin. CloudFront sends a request when it can’t find a valid object in its cache that matches the request’s cache key. If you want to send values to the origin but *not* include them in the cache key, use `OriginRequestPolicy` .
//
// TODO: EXAMPLE
//
type CfnCachePolicy_CachePolicyConfigProperty struct {
	// The default amount of time, in seconds, that you want objects to stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated.
	//
	// CloudFront uses this value as the object’s time to live (TTL) only when the origin does *not* send `Cache-Control` or `Expires` headers with the object. For more information, see [Managing How Long Content Stays in an Edge Cache (Expiration)](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/Expiration.html) in the *Amazon CloudFront Developer Guide* .
	//
	// The default value for this field is 86400 seconds (one day). If the value of `MinTTL` is more than 86400 seconds, then the default value for this field is the same as the value of `MinTTL` .
	DefaultTtl *float64 `json:"defaultTtl" yaml:"defaultTtl"`
	// The maximum amount of time, in seconds, that objects stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated.
	//
	// CloudFront uses this value only when the origin sends `Cache-Control` or `Expires` headers with the object. For more information, see [Managing How Long Content Stays in an Edge Cache (Expiration)](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/Expiration.html) in the *Amazon CloudFront Developer Guide* .
	//
	// The default value for this field is 31536000 seconds (one year). If the value of `MinTTL` or `DefaultTTL` is more than 31536000 seconds, then the default value for this field is the same as the value of `DefaultTTL` .
	MaxTtl *float64 `json:"maxTtl" yaml:"maxTtl"`
	// The minimum amount of time, in seconds, that you want objects to stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated.
	//
	// For more information, see [Managing How Long Content Stays in an Edge Cache (Expiration)](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/Expiration.html) in the *Amazon CloudFront Developer Guide* .
	MinTtl *float64 `json:"minTtl" yaml:"minTtl"`
	// A unique name to identify the cache policy.
	Name *string `json:"name" yaml:"name"`
	// The HTTP headers, cookies, and URL query strings to include in the cache key.
	//
	// The values included in the cache key are automatically included in requests that CloudFront sends to the origin.
	ParametersInCacheKeyAndForwardedToOrigin interface{} `json:"parametersInCacheKeyAndForwardedToOrigin" yaml:"parametersInCacheKeyAndForwardedToOrigin"`
	// A comment to describe the cache policy.
	//
	// The comment cannot be longer than 128 characters.
	Comment *string `json:"comment" yaml:"comment"`
}

// An object that determines whether any cookies in viewer requests (and if so, which cookies) are included in the cache key and automatically included in requests that CloudFront sends to the origin.
//
// TODO: EXAMPLE
//
type CfnCachePolicy_CookiesConfigProperty struct {
	// Determines whether any cookies in viewer requests are included in the cache key and automatically included in requests that CloudFront sends to the origin.
	//
	// Valid values are:
	//
	// - `none` – Cookies in viewer requests are not included in the cache key and are not automatically included in requests that CloudFront sends to the origin. Even when this field is set to `none` , any cookies that are listed in an `OriginRequestPolicy` *are* included in origin requests.
	// - `whitelist` – The cookies in viewer requests that are listed in the `CookieNames` type are included in the cache key and automatically included in requests that CloudFront sends to the origin.
	// - `allExcept` – All cookies in viewer requests that are **not** listed in the `CookieNames` type are included in the cache key and automatically included in requests that CloudFront sends to the origin.
	// - `all` – All cookies in viewer requests are included in the cache key and are automatically included in requests that CloudFront sends to the origin.
	CookieBehavior *string `json:"cookieBehavior" yaml:"cookieBehavior"`
	// Contains a list of cookie names.
	Cookies *[]*string `json:"cookies" yaml:"cookies"`
}

// An object that determines whether any HTTP headers (and if so, which headers) are included in the cache key and automatically included in requests that CloudFront sends to the origin.
//
// TODO: EXAMPLE
//
type CfnCachePolicy_HeadersConfigProperty struct {
	// Determines whether any HTTP headers are included in the cache key and automatically included in requests that CloudFront sends to the origin.
	//
	// Valid values are:
	//
	// - `none` – HTTP headers are not included in the cache key and are not automatically included in requests that CloudFront sends to the origin. Even when this field is set to `none` , any headers that are listed in an `OriginRequestPolicy` *are* included in origin requests.
	// - `whitelist` – The HTTP headers that are listed in the `Headers` type are included in the cache key and are automatically included in requests that CloudFront sends to the origin.
	HeaderBehavior *string `json:"headerBehavior" yaml:"headerBehavior"`
	// Contains a list of HTTP header names.
	Headers *[]*string `json:"headers" yaml:"headers"`
}

// This object determines the values that CloudFront includes in the cache key.
//
// These values can include HTTP headers, cookies, and URL query strings. CloudFront uses the cache key to find an object in its cache that it can return to the viewer.
//
// The headers, cookies, and query strings that are included in the cache key are automatically included in requests that CloudFront sends to the origin. CloudFront sends a request when it can’t find an object in its cache that matches the request’s cache key. If you want to send values to the origin but *not* include them in the cache key, use `OriginRequestPolicy` .
//
// TODO: EXAMPLE
//
type CfnCachePolicy_ParametersInCacheKeyAndForwardedToOriginProperty struct {
	// An object that determines whether any cookies in viewer requests (and if so, which cookies) are included in the cache key and automatically included in requests that CloudFront sends to the origin.
	CookiesConfig interface{} `json:"cookiesConfig" yaml:"cookiesConfig"`
	// A flag that can affect whether the `Accept-Encoding` HTTP header is included in the cache key and included in requests that CloudFront sends to the origin.
	//
	// This field is related to the `EnableAcceptEncodingBrotli` field. If one or both of these fields is `true` *and* the viewer request includes the `Accept-Encoding` header, then CloudFront does the following:
	//
	// - Normalizes the value of the viewer’s `Accept-Encoding` header
	// - Includes the normalized header in the cache key
	// - Includes the normalized header in the request to the origin, if a request is necessary
	//
	// For more information, see [Compression support](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-the-cache-key.html#cache-policy-compressed-objects) in the *Amazon CloudFront Developer Guide* .
	//
	// If you set this value to `true` , and this cache behavior also has an origin request policy attached, do not include the `Accept-Encoding` header in the origin request policy. CloudFront always includes the `Accept-Encoding` header in origin requests when the value of this field is `true` , so including this header in an origin request policy has no effect.
	//
	// If both of these fields are `false` , then CloudFront treats the `Accept-Encoding` header the same as any other HTTP header in the viewer request. By default, it’s not included in the cache key and it’s not included in origin requests. In this case, you can manually add `Accept-Encoding` to the headers whitelist like any other HTTP header.
	EnableAcceptEncodingGzip interface{} `json:"enableAcceptEncodingGzip" yaml:"enableAcceptEncodingGzip"`
	// An object that determines whether any HTTP headers (and if so, which headers) are included in the cache key and automatically included in requests that CloudFront sends to the origin.
	HeadersConfig interface{} `json:"headersConfig" yaml:"headersConfig"`
	// An object that determines whether any URL query strings in viewer requests (and if so, which query strings) are included in the cache key and automatically included in requests that CloudFront sends to the origin.
	QueryStringsConfig interface{} `json:"queryStringsConfig" yaml:"queryStringsConfig"`
	// A flag that can affect whether the `Accept-Encoding` HTTP header is included in the cache key and included in requests that CloudFront sends to the origin.
	//
	// This field is related to the `EnableAcceptEncodingGzip` field. If one or both of these fields is `true` *and* the viewer request includes the `Accept-Encoding` header, then CloudFront does the following:
	//
	// - Normalizes the value of the viewer’s `Accept-Encoding` header
	// - Includes the normalized header in the cache key
	// - Includes the normalized header in the request to the origin, if a request is necessary
	//
	// For more information, see [Compression support](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-the-cache-key.html#cache-policy-compressed-objects) in the *Amazon CloudFront Developer Guide* .
	//
	// If you set this value to `true` , and this cache behavior also has an origin request policy attached, do not include the `Accept-Encoding` header in the origin request policy. CloudFront always includes the `Accept-Encoding` header in origin requests when the value of this field is `true` , so including this header in an origin request policy has no effect.
	//
	// If both of these fields are `false` , then CloudFront treats the `Accept-Encoding` header the same as any other HTTP header in the viewer request. By default, it’s not included in the cache key and it’s not included in origin requests. In this case, you can manually add `Accept-Encoding` to the headers whitelist like any other HTTP header.
	EnableAcceptEncodingBrotli interface{} `json:"enableAcceptEncodingBrotli" yaml:"enableAcceptEncodingBrotli"`
}

// An object that determines whether any URL query strings in viewer requests (and if so, which query strings) are included in the cache key and automatically included in requests that CloudFront sends to the origin.
//
// TODO: EXAMPLE
//
type CfnCachePolicy_QueryStringsConfigProperty struct {
	// Determines whether any URL query strings in viewer requests are included in the cache key and automatically included in requests that CloudFront sends to the origin.
	//
	// Valid values are:
	//
	// - `none` – Query strings in viewer requests are not included in the cache key and are not automatically included in requests that CloudFront sends to the origin. Even when this field is set to `none` , any query strings that are listed in an `OriginRequestPolicy` *are* included in origin requests.
	// - `whitelist` – The query strings in viewer requests that are listed in the `QueryStringNames` type are included in the cache key and automatically included in requests that CloudFront sends to the origin.
	// - `allExcept` – All query strings in viewer requests that are **not** listed in the `QueryStringNames` type are included in the cache key and automatically included in requests that CloudFront sends to the origin.
	// - `all` – All query strings in viewer requests are included in the cache key and are automatically included in requests that CloudFront sends to the origin.
	QueryStringBehavior *string `json:"queryStringBehavior" yaml:"queryStringBehavior"`
	// Contains a list of query string names.
	QueryStrings *[]*string `json:"queryStrings" yaml:"queryStrings"`
}

// Properties for defining a `CfnCachePolicy`.
//
// TODO: EXAMPLE
//
type CfnCachePolicyProps struct {
	// The cache policy configuration.
	CachePolicyConfig interface{} `json:"cachePolicyConfig" yaml:"cachePolicyConfig"`
}

// A CloudFormation `AWS::CloudFront::CloudFrontOriginAccessIdentity`.
//
// The request to create a new origin access identity (OAI). An origin access identity is a special CloudFront user that you can associate with Amazon S3 origins, so that you can secure all or just some of your Amazon S3 content. For more information, see [Restricting Access to Amazon S3 Content by Using an Origin Access Identity](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/private-content-restricting-access-to-s3.html) in the *Amazon CloudFront Developer Guide* .
//
// TODO: EXAMPLE
//
type CfnCloudFrontOriginAccessIdentity interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrId() *string
	AttrS3CanonicalUserId() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CloudFrontOriginAccessIdentityConfig() interface{}
	SetCloudFrontOriginAccessIdentityConfig(val interface{})
	CreationStack() *[]*string
	LogicalId() *string
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

// The jsii proxy struct for CfnCloudFrontOriginAccessIdentity
type jsiiProxy_CfnCloudFrontOriginAccessIdentity struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnCloudFrontOriginAccessIdentity) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudFrontOriginAccessIdentity) AttrS3CanonicalUserId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrS3CanonicalUserId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudFrontOriginAccessIdentity) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudFrontOriginAccessIdentity) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudFrontOriginAccessIdentity) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudFrontOriginAccessIdentity) CloudFrontOriginAccessIdentityConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cloudFrontOriginAccessIdentityConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudFrontOriginAccessIdentity) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudFrontOriginAccessIdentity) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudFrontOriginAccessIdentity) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudFrontOriginAccessIdentity) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudFrontOriginAccessIdentity) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudFrontOriginAccessIdentity) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::CloudFront::CloudFrontOriginAccessIdentity`.
func NewCfnCloudFrontOriginAccessIdentity(scope constructs.Construct, id *string, props *CfnCloudFrontOriginAccessIdentityProps) CfnCloudFrontOriginAccessIdentity {
	_init_.Initialize()

	j := jsiiProxy_CfnCloudFrontOriginAccessIdentity{}

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront.CfnCloudFrontOriginAccessIdentity",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::CloudFront::CloudFrontOriginAccessIdentity`.
func NewCfnCloudFrontOriginAccessIdentity_Override(c CfnCloudFrontOriginAccessIdentity, scope constructs.Construct, id *string, props *CfnCloudFrontOriginAccessIdentityProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront.CfnCloudFrontOriginAccessIdentity",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnCloudFrontOriginAccessIdentity) SetCloudFrontOriginAccessIdentityConfig(val interface{}) {
	_jsii_.Set(
		j,
		"cloudFrontOriginAccessIdentityConfig",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnCloudFrontOriginAccessIdentity_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.CfnCloudFrontOriginAccessIdentity",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnCloudFrontOriginAccessIdentity_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.CfnCloudFrontOriginAccessIdentity",
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
func CfnCloudFrontOriginAccessIdentity_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.CfnCloudFrontOriginAccessIdentity",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCloudFrontOriginAccessIdentity_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cloudfront.CfnCloudFrontOriginAccessIdentity",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnCloudFrontOriginAccessIdentity) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnCloudFrontOriginAccessIdentity) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnCloudFrontOriginAccessIdentity) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnCloudFrontOriginAccessIdentity) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnCloudFrontOriginAccessIdentity) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnCloudFrontOriginAccessIdentity) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnCloudFrontOriginAccessIdentity) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnCloudFrontOriginAccessIdentity) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnCloudFrontOriginAccessIdentity) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnCloudFrontOriginAccessIdentity) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnCloudFrontOriginAccessIdentity) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnCloudFrontOriginAccessIdentity) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnCloudFrontOriginAccessIdentity) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnCloudFrontOriginAccessIdentity) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCloudFrontOriginAccessIdentity) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Origin access identity configuration.
//
// Send a `GET` request to the `/ *CloudFront API version* /CloudFront/identity ID/config` resource.
//
// TODO: EXAMPLE
//
type CfnCloudFrontOriginAccessIdentity_CloudFrontOriginAccessIdentityConfigProperty struct {
	// A comment to describe the origin access identity.
	//
	// The comment cannot be longer than 128 characters.
	Comment *string `json:"comment" yaml:"comment"`
}

// Properties for defining a `CfnCloudFrontOriginAccessIdentity`.
//
// TODO: EXAMPLE
//
type CfnCloudFrontOriginAccessIdentityProps struct {
	// The current configuration information for the identity.
	CloudFrontOriginAccessIdentityConfig interface{} `json:"cloudFrontOriginAccessIdentityConfig" yaml:"cloudFrontOriginAccessIdentityConfig"`
}

// A CloudFormation `AWS::CloudFront::Distribution`.
//
// A distribution tells CloudFront where you want content to be delivered from, and the details about how to track and manage content delivery.
//
// TODO: EXAMPLE
//
type CfnDistribution interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrDomainName() *string
	AttrId() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	DistributionConfig() interface{}
	SetDistributionConfig(val interface{})
	LogicalId() *string
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

// The jsii proxy struct for CfnDistribution
type jsiiProxy_CfnDistribution struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnDistribution) AttrDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDistribution) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDistribution) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDistribution) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDistribution) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDistribution) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDistribution) DistributionConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"distributionConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDistribution) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDistribution) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDistribution) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDistribution) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDistribution) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDistribution) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::CloudFront::Distribution`.
func NewCfnDistribution(scope constructs.Construct, id *string, props *CfnDistributionProps) CfnDistribution {
	_init_.Initialize()

	j := jsiiProxy_CfnDistribution{}

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront.CfnDistribution",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::CloudFront::Distribution`.
func NewCfnDistribution_Override(c CfnDistribution, scope constructs.Construct, id *string, props *CfnDistributionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront.CfnDistribution",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDistribution) SetDistributionConfig(val interface{}) {
	_jsii_.Set(
		j,
		"distributionConfig",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnDistribution_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.CfnDistribution",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnDistribution_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.CfnDistribution",
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
func CfnDistribution_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.CfnDistribution",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDistribution_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cloudfront.CfnDistribution",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnDistribution) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnDistribution) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnDistribution) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnDistribution) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnDistribution) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnDistribution) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnDistribution) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnDistribution) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnDistribution) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnDistribution) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnDistribution) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDistribution) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnDistribution) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnDistribution) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDistribution) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// A complex type that describes how CloudFront processes requests.
//
// You must create at least as many cache behaviors (including the default cache behavior) as you have origins if you want CloudFront to serve objects from all of the origins. Each cache behavior specifies the one origin from which you want CloudFront to get objects. If you have two origins and only the default cache behavior, the default cache behavior will cause CloudFront to get objects from one of the origins, but the other origin is never used.
//
// For the current quota (formerly known as limit) on the number of cache behaviors that you can add to a distribution, see [Quotas](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/cloudfront-limits.html) in the *Amazon CloudFront Developer Guide* .
//
// If you don’t want to specify any cache behaviors, include only an empty `CacheBehaviors` element. Don’t include an empty `CacheBehavior` element because this is invalid.
//
// To delete all cache behaviors in an existing distribution, update the distribution configuration and include only an empty `CacheBehaviors` element.
//
// To add, change, or remove one or more cache behaviors, update the distribution configuration and specify all of the cache behaviors that you want to include in the updated distribution.
//
// For more information about cache behaviors, see [Cache Behavior Settings](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/distribution-web-values-specify.html#DownloadDistValuesCacheBehavior) in the *Amazon CloudFront Developer Guide* .
//
// TODO: EXAMPLE
//
type CfnDistribution_CacheBehaviorProperty struct {
	// The pattern (for example, `images/*.jpg` ) that specifies which requests to apply the behavior to. When CloudFront receives a viewer request, the requested path is compared with path patterns in the order in which cache behaviors are listed in the distribution.
	//
	// > You can optionally include a slash ( `/` ) at the beginning of the path pattern. For example, `/images/*.jpg` . CloudFront behavior is the same with or without the leading `/` .
	//
	// The path pattern for the default cache behavior is `*` and cannot be changed. If the request for an object does not match the path pattern for any cache behaviors, CloudFront applies the behavior in the default cache behavior.
	//
	// For more information, see [Path Pattern](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/distribution-web-values-specify.html#DownloadDistValuesPathPattern) in the *Amazon CloudFront Developer Guide* .
	PathPattern *string `json:"pathPattern" yaml:"pathPattern"`
	// The value of `ID` for the origin that you want CloudFront to route requests to when they match this cache behavior.
	TargetOriginId *string `json:"targetOriginId" yaml:"targetOriginId"`
	// The protocol that viewers can use to access the files in the origin specified by `TargetOriginId` when a request matches the path pattern in `PathPattern` .
	//
	// You can specify the following options:
	//
	// - `allow-all` : Viewers can use HTTP or HTTPS.
	// - `redirect-to-https` : If a viewer submits an HTTP request, CloudFront returns an HTTP status code of 301 (Moved Permanently) to the viewer along with the HTTPS URL. The viewer then resubmits the request using the new URL.
	// - `https-only` : If a viewer sends an HTTP request, CloudFront returns an HTTP status code of 403 (Forbidden).
	//
	// For more information about requiring the HTTPS protocol, see [Requiring HTTPS Between Viewers and CloudFront](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/using-https-viewers-to-cloudfront.html) in the *Amazon CloudFront Developer Guide* .
	//
	// > The only way to guarantee that viewers retrieve an object that was fetched from the origin using HTTPS is never to use any other protocol to fetch the object. If you have recently changed from HTTP to HTTPS, we recommend that you clear your objects’ cache because cached objects are protocol agnostic. That means that an edge location will return an object from the cache regardless of whether the current request protocol matches the protocol used previously. For more information, see [Managing Cache Expiration](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/Expiration.html) in the *Amazon CloudFront Developer Guide* .
	ViewerProtocolPolicy *string `json:"viewerProtocolPolicy" yaml:"viewerProtocolPolicy"`
	// A complex type that controls which HTTP methods CloudFront processes and forwards to your Amazon S3 bucket or your custom origin.
	//
	// There are three choices:
	//
	// - CloudFront forwards only `GET` and `HEAD` requests.
	// - CloudFront forwards only `GET` , `HEAD` , and `OPTIONS` requests.
	// - CloudFront forwards `GET, HEAD, OPTIONS, PUT, PATCH, POST` , and `DELETE` requests.
	//
	// If you pick the third choice, you may need to restrict access to your Amazon S3 bucket or to your custom origin so users can't perform operations that you don't want them to. For example, you might not want users to have permissions to delete objects from your origin.
	AllowedMethods *[]*string `json:"allowedMethods" yaml:"allowedMethods"`
	// A complex type that controls whether CloudFront caches the response to requests using the specified HTTP methods.
	//
	// There are two choices:
	//
	// - CloudFront caches responses to `GET` and `HEAD` requests.
	// - CloudFront caches responses to `GET` , `HEAD` , and `OPTIONS` requests.
	//
	// If you pick the second choice for your Amazon S3 Origin, you may need to forward Access-Control-Request-Method, Access-Control-Request-Headers, and Origin headers for the responses to be cached correctly.
	CachedMethods *[]*string `json:"cachedMethods" yaml:"cachedMethods"`
	// The unique identifier of the cache policy that is attached to this cache behavior.
	//
	// For more information, see [Creating cache policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-the-cache-key.html#cache-key-create-cache-policy) or [Using the managed cache policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/using-managed-cache-policies.html) in the *Amazon CloudFront Developer Guide* .
	//
	// A `CacheBehavior` must include either a `CachePolicyId` or `ForwardedValues` . We recommend that you use a `CachePolicyId` .
	CachePolicyId *string `json:"cachePolicyId" yaml:"cachePolicyId"`
	// Whether you want CloudFront to automatically compress certain files for this cache behavior.
	//
	// If so, specify true; if not, specify false. For more information, see [Serving Compressed Files](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/ServingCompressedFiles.html) in the *Amazon CloudFront Developer Guide* .
	Compress interface{} `json:"compress" yaml:"compress"`
	// This field is deprecated.
	//
	// We recommend that you use the `DefaultTTL` field in a cache policy instead of this field. For more information, see [Creating cache policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-the-cache-key.html#cache-key-create-cache-policy) or [Using the managed cache policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/using-managed-cache-policies.html) in the *Amazon CloudFront Developer Guide* .
	//
	// The default amount of time that you want objects to stay in CloudFront caches before CloudFront forwards another request to your origin to determine whether the object has been updated. The value that you specify applies only when your origin does not add HTTP headers such as `Cache-Control max-age` , `Cache-Control s-maxage` , and `Expires` to objects. For more information, see [Managing How Long Content Stays in an Edge Cache (Expiration)](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/Expiration.html) in the *Amazon CloudFront Developer Guide* .
	DefaultTtl *float64 `json:"defaultTtl" yaml:"defaultTtl"`
	// The value of `ID` for the field-level encryption configuration that you want CloudFront to use for encrypting specific fields of data for this cache behavior.
	FieldLevelEncryptionId *string `json:"fieldLevelEncryptionId" yaml:"fieldLevelEncryptionId"`
	// This field is deprecated.
	//
	// We recommend that you use a cache policy or an origin request policy instead of this field. For more information, see [Working with policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/working-with-policies.html) in the *Amazon CloudFront Developer Guide* .
	//
	// If you want to include values in the cache key, use a cache policy. For more information, see [Creating cache policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-the-cache-key.html#cache-key-create-cache-policy) or [Using the managed cache policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/using-managed-cache-policies.html) in the *Amazon CloudFront Developer Guide* .
	//
	// If you want to send values to the origin but not include them in the cache key, use an origin request policy. For more information, see [Creating origin request policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-origin-requests.html#origin-request-create-origin-request-policy) or [Using the managed origin request policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/using-managed-origin-request-policies.html) in the *Amazon CloudFront Developer Guide* .
	//
	// A `CacheBehavior` must include either a `CachePolicyId` or `ForwardedValues` . We recommend that you use a `CachePolicyId` .
	//
	// A complex type that specifies how CloudFront handles query strings, cookies, and HTTP headers.
	ForwardedValues interface{} `json:"forwardedValues" yaml:"forwardedValues"`
	// A list of CloudFront functions that are associated with this cache behavior.
	//
	// CloudFront functions must be published to the `LIVE` stage to associate them with a cache behavior.
	FunctionAssociations interface{} `json:"functionAssociations" yaml:"functionAssociations"`
	// A complex type that contains zero or more Lambda@Edge function associations for a cache behavior.
	LambdaFunctionAssociations interface{} `json:"lambdaFunctionAssociations" yaml:"lambdaFunctionAssociations"`
	// This field is deprecated.
	//
	// We recommend that you use the `MaxTTL` field in a cache policy instead of this field. For more information, see [Creating cache policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-the-cache-key.html#cache-key-create-cache-policy) or [Using the managed cache policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/using-managed-cache-policies.html) in the *Amazon CloudFront Developer Guide* .
	//
	// The maximum amount of time that you want objects to stay in CloudFront caches before CloudFront forwards another request to your origin to determine whether the object has been updated. The value that you specify applies only when your origin adds HTTP headers such as `Cache-Control max-age` , `Cache-Control s-maxage` , and `Expires` to objects. For more information, see [Managing How Long Content Stays in an Edge Cache (Expiration)](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/Expiration.html) in the *Amazon CloudFront Developer Guide* .
	MaxTtl *float64 `json:"maxTtl" yaml:"maxTtl"`
	// This field is deprecated.
	//
	// We recommend that you use the `MinTTL` field in a cache policy instead of this field. For more information, see [Creating cache policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-the-cache-key.html#cache-key-create-cache-policy) or [Using the managed cache policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/using-managed-cache-policies.html) in the *Amazon CloudFront Developer Guide* .
	//
	// The minimum amount of time that you want objects to stay in CloudFront caches before CloudFront forwards another request to your origin to determine whether the object has been updated. For more information, see [Managing How Long Content Stays in an Edge Cache (Expiration)](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/Expiration.html) in the *Amazon CloudFront Developer Guide* .
	//
	// You must specify `0` for `MinTTL` if you configure CloudFront to forward all headers to your origin (under `Headers` , if you specify `1` for `Quantity` and `*` for `Name` ).
	MinTtl *float64 `json:"minTtl" yaml:"minTtl"`
	// The unique identifier of the origin request policy that is attached to this cache behavior.
	//
	// For more information, see [Creating origin request policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-origin-requests.html#origin-request-create-origin-request-policy) or [Using the managed origin request policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/using-managed-origin-request-policies.html) in the *Amazon CloudFront Developer Guide* .
	OriginRequestPolicyId *string `json:"originRequestPolicyId" yaml:"originRequestPolicyId"`
	// The Amazon Resource Name (ARN) of the real-time log configuration that is attached to this cache behavior.
	//
	// For more information, see [Real-time logs](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/real-time-logs.html) in the *Amazon CloudFront Developer Guide* .
	RealtimeLogConfigArn *string `json:"realtimeLogConfigArn" yaml:"realtimeLogConfigArn"`
	// The identifier for a response headers policy.
	ResponseHeadersPolicyId *string `json:"responseHeadersPolicyId" yaml:"responseHeadersPolicyId"`
	// Indicates whether you want to distribute media files in the Microsoft Smooth Streaming format using the origin that is associated with this cache behavior.
	//
	// If so, specify `true` ; if not, specify `false` . If you specify `true` for `SmoothStreaming` , you can still distribute other content using this cache behavior if the content matches the value of `PathPattern` .
	SmoothStreaming interface{} `json:"smoothStreaming" yaml:"smoothStreaming"`
	// A list of key groups that CloudFront can use to validate signed URLs or signed cookies.
	//
	// When a cache behavior contains trusted key groups, CloudFront requires signed URLs or signed cookies for all requests that match the cache behavior. The URLs or cookies must be signed with a private key whose corresponding public key is in the key group. The signed URL or cookie contains information about which public key CloudFront should use to verify the signature. For more information, see [Serving private content](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/PrivateContent.html) in the *Amazon CloudFront Developer Guide* .
	TrustedKeyGroups *[]*string `json:"trustedKeyGroups" yaml:"trustedKeyGroups"`
	// > We recommend using `TrustedKeyGroups` instead of `TrustedSigners` .
	//
	// A list of AWS account IDs whose public keys CloudFront can use to validate signed URLs or signed cookies.
	//
	// When a cache behavior contains trusted signers, CloudFront requires signed URLs or signed cookies for all requests that match the cache behavior. The URLs or cookies must be signed with the private key of a CloudFront key pair in the trusted signer’s AWS account . The signed URL or cookie contains information about which public key CloudFront should use to verify the signature. For more information, see [Serving private content](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/PrivateContent.html) in the *Amazon CloudFront Developer Guide* .
	TrustedSigners *[]*string `json:"trustedSigners" yaml:"trustedSigners"`
}

// This field is deprecated.
//
// We recommend that you use a cache policy or an origin request policy instead of this field.
//
// If you want to include cookies in the cache key, use a cache policy. For more information, see [Creating cache policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-the-cache-key.html#cache-key-create-cache-policy) in the *Amazon CloudFront Developer Guide* .
//
// If you want to send cookies to the origin but not include them in the cache key, use an origin request policy. For more information, see [Creating origin request policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-origin-requests.html#origin-request-create-origin-request-policy) in the *Amazon CloudFront Developer Guide* .
//
// A complex type that specifies whether you want CloudFront to forward cookies to the origin and, if so, which ones. For more information about forwarding cookies to the origin, see [How CloudFront Forwards, Caches, and Logs Cookies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/Cookies.html) in the *Amazon CloudFront Developer Guide* .
//
// TODO: EXAMPLE
//
type CfnDistribution_CookiesProperty struct {
	// This field is deprecated.
	//
	// We recommend that you use a cache policy or an origin request policy instead of this field.
	//
	// If you want to include cookies in the cache key, use a cache policy. For more information, see [Creating cache policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-the-cache-key.html#cache-key-create-cache-policy) in the *Amazon CloudFront Developer Guide* .
	//
	// If you want to send cookies to the origin but not include them in the cache key, use origin request policy. For more information, see [Creating origin request policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-origin-requests.html#origin-request-create-origin-request-policy) in the *Amazon CloudFront Developer Guide* .
	//
	// Specifies which cookies to forward to the origin for this cache behavior: all, none, or the list of cookies specified in the `WhitelistedNames` complex type.
	//
	// Amazon S3 doesn't process cookies. When the cache behavior is forwarding requests to an Amazon S3 origin, specify none for the `Forward` element.
	Forward *string `json:"forward" yaml:"forward"`
	// This field is deprecated.
	//
	// We recommend that you use a cache policy or an origin request policy instead of this field.
	//
	// If you want to include cookies in the cache key, use a cache policy. For more information, see [Creating cache policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-the-cache-key.html#cache-key-create-cache-policy) in the *Amazon CloudFront Developer Guide* .
	//
	// If you want to send cookies to the origin but not include them in the cache key, use an origin request policy. For more information, see [Creating origin request policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-origin-requests.html#origin-request-create-origin-request-policy) in the *Amazon CloudFront Developer Guide* .
	//
	// Required if you specify `whitelist` for the value of `Forward` . A complex type that specifies how many different cookies you want CloudFront to forward to the origin for this cache behavior and, if you want to forward selected cookies, the names of those cookies.
	//
	// If you specify `all` or `none` for the value of `Forward` , omit `WhitelistedNames` . If you change the value of `Forward` from `whitelist` to `all` or `none` and you don't delete the `WhitelistedNames` element and its child elements, CloudFront deletes them automatically.
	//
	// For the current limit on the number of cookie names that you can whitelist for each cache behavior, see [CloudFront Limits](https://docs.aws.amazon.com/general/latest/gr/xrefaws_service_limits.html#limits_cloudfront) in the *AWS General Reference* .
	WhitelistedNames *[]*string `json:"whitelistedNames" yaml:"whitelistedNames"`
}

// A complex type that controls:.
//
// - Whether CloudFront replaces HTTP status codes in the 4xx and 5xx range with custom error messages before returning the response to the viewer.
// - How long CloudFront caches HTTP status codes in the 4xx and 5xx range.
//
// For more information about custom error pages, see [Customizing Error Responses](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/custom-error-pages.html) in the *Amazon CloudFront Developer Guide* .
//
// TODO: EXAMPLE
//
type CfnDistribution_CustomErrorResponseProperty struct {
	// The HTTP status code for which you want to specify a custom error page and/or a caching duration.
	ErrorCode *float64 `json:"errorCode" yaml:"errorCode"`
	// The minimum amount of time, in seconds, that you want CloudFront to cache the HTTP status code specified in `ErrorCode` .
	//
	// When this time period has elapsed, CloudFront queries your origin to see whether the problem that caused the error has been resolved and the requested object is now available.
	//
	// For more information, see [Customizing Error Responses](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/custom-error-pages.html) in the *Amazon CloudFront Developer Guide* .
	ErrorCachingMinTtl *float64 `json:"errorCachingMinTtl" yaml:"errorCachingMinTtl"`
	// The HTTP status code that you want CloudFront to return to the viewer along with the custom error page.
	//
	// There are a variety of reasons that you might want CloudFront to return a status code different from the status code that your origin returned to CloudFront, for example:
	//
	// - Some Internet devices (some firewalls and corporate proxies, for example) intercept HTTP 4xx and 5xx and prevent the response from being returned to the viewer. If you substitute `200` , the response typically won't be intercepted.
	// - If you don't care about distinguishing among different client errors or server errors, you can specify `400` or `500` as the `ResponseCode` for all 4xx or 5xx errors.
	// - You might want to return a `200` status code (OK) and static website so your customers don't know that your website is down.
	//
	// If you specify a value for `ResponseCode` , you must also specify a value for `ResponsePagePath` .
	ResponseCode *float64 `json:"responseCode" yaml:"responseCode"`
	// The path to the custom error page that you want CloudFront to return to a viewer when your origin returns the HTTP status code specified by `ErrorCode` , for example, `/4xx-errors/403-forbidden.html` . If you want to store your objects and your custom error pages in different locations, your distribution must include a cache behavior for which the following is true:.
	//
	// - The value of `PathPattern` matches the path to your custom error messages. For example, suppose you saved custom error pages for 4xx errors in an Amazon S3 bucket in a directory named `/4xx-errors` . Your distribution must include a cache behavior for which the path pattern routes requests for your custom error pages to that location, for example, `/4xx-errors/*` .
	// - The value of `TargetOriginId` specifies the value of the `ID` element for the origin that contains your custom error pages.
	//
	// If you specify a value for `ResponsePagePath` , you must also specify a value for `ResponseCode` .
	//
	// We recommend that you store custom error pages in an Amazon S3 bucket. If you store custom error pages on an HTTP server and the server starts to return 5xx errors, CloudFront can't get the files that you want to return to viewers because the origin server is unavailable.
	ResponsePagePath *string `json:"responsePagePath" yaml:"responsePagePath"`
}

// A custom origin.
//
// A custom origin is any origin that is *not* an Amazon S3 bucket, with one exception. An Amazon S3 bucket that is [configured with static website hosting](https://docs.aws.amazon.com/AmazonS3/latest/dev/WebsiteHosting.html) *is* a custom origin.
//
// TODO: EXAMPLE
//
type CfnDistribution_CustomOriginConfigProperty struct {
	// Specifies the protocol (HTTP or HTTPS) that CloudFront uses to connect to the origin. Valid values are:.
	//
	// - `http-only` – CloudFront always uses HTTP to connect to the origin.
	// - `match-viewer` – CloudFront connects to the origin using the same protocol that the viewer used to connect to CloudFront.
	// - `https-only` – CloudFront always uses HTTPS to connect to the origin.
	OriginProtocolPolicy *string `json:"originProtocolPolicy" yaml:"originProtocolPolicy"`
	// The HTTP port that CloudFront uses to connect to the origin.
	//
	// Specify the HTTP port that the origin listens on.
	HttpPort *float64 `json:"httpPort" yaml:"httpPort"`
	// The HTTPS port that CloudFront uses to connect to the origin.
	//
	// Specify the HTTPS port that the origin listens on.
	HttpsPort *float64 `json:"httpsPort" yaml:"httpsPort"`
	// Specifies how long, in seconds, CloudFront persists its connection to the origin.
	//
	// The minimum timeout is 1 second, the maximum is 60 seconds, and the default (if you don’t specify otherwise) is 5 seconds.
	//
	// For more information, see [Origin Keep-alive Timeout](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/distribution-web-values-specify.html#DownloadDistValuesOriginKeepaliveTimeout) in the *Amazon CloudFront Developer Guide* .
	OriginKeepaliveTimeout *float64 `json:"originKeepaliveTimeout" yaml:"originKeepaliveTimeout"`
	// Specifies how long, in seconds, CloudFront waits for a response from the origin.
	//
	// This is also known as the *origin response timeout* . The minimum timeout is 1 second, the maximum is 60 seconds, and the default (if you don’t specify otherwise) is 30 seconds.
	//
	// For more information, see [Origin Response Timeout](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/distribution-web-values-specify.html#DownloadDistValuesOriginResponseTimeout) in the *Amazon CloudFront Developer Guide* .
	OriginReadTimeout *float64 `json:"originReadTimeout" yaml:"originReadTimeout"`
	// Specifies the minimum SSL/TLS protocol that CloudFront uses when connecting to your origin over HTTPS.
	//
	// Valid values include `SSLv3` , `TLSv1` , `TLSv1.1` , and `TLSv1.2` .
	//
	// For more information, see [Minimum Origin SSL Protocol](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/distribution-web-values-specify.html#DownloadDistValuesOriginSSLProtocols) in the *Amazon CloudFront Developer Guide* .
	OriginSslProtocols *[]*string `json:"originSslProtocols" yaml:"originSslProtocols"`
}

// A complex type that describes the default cache behavior if you don’t specify a `CacheBehavior` element or if request URLs don’t match any of the values of `PathPattern` in `CacheBehavior` elements.
//
// You must create exactly one default cache behavior.
//
// TODO: EXAMPLE
//
type CfnDistribution_DefaultCacheBehaviorProperty struct {
	// The value of `ID` for the origin that you want CloudFront to route requests to when they use the default cache behavior.
	TargetOriginId *string `json:"targetOriginId" yaml:"targetOriginId"`
	// The protocol that viewers can use to access the files in the origin specified by `TargetOriginId` when a request matches the path pattern in `PathPattern` .
	//
	// You can specify the following options:
	//
	// - `allow-all` : Viewers can use HTTP or HTTPS.
	// - `redirect-to-https` : If a viewer submits an HTTP request, CloudFront returns an HTTP status code of 301 (Moved Permanently) to the viewer along with the HTTPS URL. The viewer then resubmits the request using the new URL.
	// - `https-only` : If a viewer sends an HTTP request, CloudFront returns an HTTP status code of 403 (Forbidden).
	//
	// For more information about requiring the HTTPS protocol, see [Requiring HTTPS Between Viewers and CloudFront](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/using-https-viewers-to-cloudfront.html) in the *Amazon CloudFront Developer Guide* .
	//
	// > The only way to guarantee that viewers retrieve an object that was fetched from the origin using HTTPS is never to use any other protocol to fetch the object. If you have recently changed from HTTP to HTTPS, we recommend that you clear your objects’ cache because cached objects are protocol agnostic. That means that an edge location will return an object from the cache regardless of whether the current request protocol matches the protocol used previously. For more information, see [Managing Cache Expiration](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/Expiration.html) in the *Amazon CloudFront Developer Guide* .
	ViewerProtocolPolicy *string `json:"viewerProtocolPolicy" yaml:"viewerProtocolPolicy"`
	// A complex type that controls which HTTP methods CloudFront processes and forwards to your Amazon S3 bucket or your custom origin.
	//
	// There are three choices:
	//
	// - CloudFront forwards only `GET` and `HEAD` requests.
	// - CloudFront forwards only `GET` , `HEAD` , and `OPTIONS` requests.
	// - CloudFront forwards `GET, HEAD, OPTIONS, PUT, PATCH, POST` , and `DELETE` requests.
	//
	// If you pick the third choice, you may need to restrict access to your Amazon S3 bucket or to your custom origin so users can't perform operations that you don't want them to. For example, you might not want users to have permissions to delete objects from your origin.
	AllowedMethods *[]*string `json:"allowedMethods" yaml:"allowedMethods"`
	// A complex type that controls whether CloudFront caches the response to requests using the specified HTTP methods.
	//
	// There are two choices:
	//
	// - CloudFront caches responses to `GET` and `HEAD` requests.
	// - CloudFront caches responses to `GET` , `HEAD` , and `OPTIONS` requests.
	//
	// If you pick the second choice for your Amazon S3 Origin, you may need to forward Access-Control-Request-Method, Access-Control-Request-Headers, and Origin headers for the responses to be cached correctly.
	CachedMethods *[]*string `json:"cachedMethods" yaml:"cachedMethods"`
	// The unique identifier of the cache policy that is attached to the default cache behavior.
	//
	// For more information, see [Creating cache policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-the-cache-key.html#cache-key-create-cache-policy) or [Using the managed cache policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/using-managed-cache-policies.html) in the *Amazon CloudFront Developer Guide* .
	//
	// A `DefaultCacheBehavior` must include either a `CachePolicyId` or `ForwardedValues` . We recommend that you use a `CachePolicyId` .
	CachePolicyId *string `json:"cachePolicyId" yaml:"cachePolicyId"`
	// Whether you want CloudFront to automatically compress certain files for this cache behavior.
	//
	// If so, specify `true` ; if not, specify `false` . For more information, see [Serving Compressed Files](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/ServingCompressedFiles.html) in the *Amazon CloudFront Developer Guide* .
	Compress interface{} `json:"compress" yaml:"compress"`
	// This field is deprecated.
	//
	// We recommend that you use the `DefaultTTL` field in a cache policy instead of this field. For more information, see [Creating cache policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-the-cache-key.html#cache-key-create-cache-policy) or [Using the managed cache policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/using-managed-cache-policies.html) in the *Amazon CloudFront Developer Guide* .
	//
	// The default amount of time that you want objects to stay in CloudFront caches before CloudFront forwards another request to your origin to determine whether the object has been updated. The value that you specify applies only when your origin does not add HTTP headers such as `Cache-Control max-age` , `Cache-Control s-maxage` , and `Expires` to objects. For more information, see [Managing How Long Content Stays in an Edge Cache (Expiration)](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/Expiration.html) in the *Amazon CloudFront Developer Guide* .
	DefaultTtl *float64 `json:"defaultTtl" yaml:"defaultTtl"`
	// The value of `ID` for the field-level encryption configuration that you want CloudFront to use for encrypting specific fields of data for the default cache behavior.
	FieldLevelEncryptionId *string `json:"fieldLevelEncryptionId" yaml:"fieldLevelEncryptionId"`
	// This field is deprecated.
	//
	// We recommend that you use a cache policy or an origin request policy instead of this field. For more information, see [Working with policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/working-with-policies.html) in the *Amazon CloudFront Developer Guide* .
	//
	// If you want to include values in the cache key, use a cache policy. For more information, see [Creating cache policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-the-cache-key.html#cache-key-create-cache-policy) or [Using the managed cache policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/using-managed-cache-policies.html) in the *Amazon CloudFront Developer Guide* .
	//
	// If you want to send values to the origin but not include them in the cache key, use an origin request policy. For more information, see [Creating origin request policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-origin-requests.html#origin-request-create-origin-request-policy) or [Using the managed origin request policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/using-managed-origin-request-policies.html) in the *Amazon CloudFront Developer Guide* .
	//
	// A `DefaultCacheBehavior` must include either a `CachePolicyId` or `ForwardedValues` . We recommend that you use a `CachePolicyId` .
	//
	// A complex type that specifies how CloudFront handles query strings, cookies, and HTTP headers.
	ForwardedValues interface{} `json:"forwardedValues" yaml:"forwardedValues"`
	// A list of CloudFront functions that are associated with this cache behavior.
	//
	// CloudFront functions must be published to the `LIVE` stage to associate them with a cache behavior.
	FunctionAssociations interface{} `json:"functionAssociations" yaml:"functionAssociations"`
	// A complex type that contains zero or more Lambda@Edge function associations for a cache behavior.
	LambdaFunctionAssociations interface{} `json:"lambdaFunctionAssociations" yaml:"lambdaFunctionAssociations"`
	// This field is deprecated.
	//
	// We recommend that you use the `MaxTTL` field in a cache policy instead of this field. For more information, see [Creating cache policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-the-cache-key.html#cache-key-create-cache-policy) or [Using the managed cache policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/using-managed-cache-policies.html) in the *Amazon CloudFront Developer Guide* .
	//
	// The maximum amount of time that you want objects to stay in CloudFront caches before CloudFront forwards another request to your origin to determine whether the object has been updated. The value that you specify applies only when your origin adds HTTP headers such as `Cache-Control max-age` , `Cache-Control s-maxage` , and `Expires` to objects. For more information, see [Managing How Long Content Stays in an Edge Cache (Expiration)](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/Expiration.html) in the *Amazon CloudFront Developer Guide* .
	MaxTtl *float64 `json:"maxTtl" yaml:"maxTtl"`
	// This field is deprecated.
	//
	// We recommend that you use the `MinTTL` field in a cache policy instead of this field. For more information, see [Creating cache policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-the-cache-key.html#cache-key-create-cache-policy) or [Using the managed cache policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/using-managed-cache-policies.html) in the *Amazon CloudFront Developer Guide* .
	//
	// The minimum amount of time that you want objects to stay in CloudFront caches before CloudFront forwards another request to your origin to determine whether the object has been updated. For more information, see [Managing How Long Content Stays in an Edge Cache (Expiration)](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/Expiration.html) in the *Amazon CloudFront Developer Guide* .
	//
	// You must specify `0` for `MinTTL` if you configure CloudFront to forward all headers to your origin (under `Headers` , if you specify `1` for `Quantity` and `*` for `Name` ).
	MinTtl *float64 `json:"minTtl" yaml:"minTtl"`
	// The unique identifier of the origin request policy that is attached to the default cache behavior.
	//
	// For more information, see [Creating origin request policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-origin-requests.html#origin-request-create-origin-request-policy) or [Using the managed origin request policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/using-managed-origin-request-policies.html) in the *Amazon CloudFront Developer Guide* .
	OriginRequestPolicyId *string `json:"originRequestPolicyId" yaml:"originRequestPolicyId"`
	// The Amazon Resource Name (ARN) of the real-time log configuration that is attached to this cache behavior.
	//
	// For more information, see [Real-time logs](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/real-time-logs.html) in the *Amazon CloudFront Developer Guide* .
	RealtimeLogConfigArn *string `json:"realtimeLogConfigArn" yaml:"realtimeLogConfigArn"`
	// The identifier for a response headers policy.
	ResponseHeadersPolicyId *string `json:"responseHeadersPolicyId" yaml:"responseHeadersPolicyId"`
	// Indicates whether you want to distribute media files in the Microsoft Smooth Streaming format using the origin that is associated with this cache behavior.
	//
	// If so, specify `true` ; if not, specify `false` . If you specify `true` for `SmoothStreaming` , you can still distribute other content using this cache behavior if the content matches the value of `PathPattern` .
	SmoothStreaming interface{} `json:"smoothStreaming" yaml:"smoothStreaming"`
	// A list of key groups that CloudFront can use to validate signed URLs or signed cookies.
	//
	// When a cache behavior contains trusted key groups, CloudFront requires signed URLs or signed cookies for all requests that match the cache behavior. The URLs or cookies must be signed with a private key whose corresponding public key is in the key group. The signed URL or cookie contains information about which public key CloudFront should use to verify the signature. For more information, see [Serving private content](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/PrivateContent.html) in the *Amazon CloudFront Developer Guide* .
	TrustedKeyGroups *[]*string `json:"trustedKeyGroups" yaml:"trustedKeyGroups"`
	// > We recommend using `TrustedKeyGroups` instead of `TrustedSigners` .
	//
	// A list of AWS account IDs whose public keys CloudFront can use to validate signed URLs or signed cookies.
	//
	// When a cache behavior contains trusted signers, CloudFront requires signed URLs or signed cookies for all requests that match the cache behavior. The URLs or cookies must be signed with the private key of a CloudFront key pair in a trusted signer’s AWS account . The signed URL or cookie contains information about which public key CloudFront should use to verify the signature. For more information, see [Serving private content](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/PrivateContent.html) in the *Amazon CloudFront Developer Guide* .
	TrustedSigners *[]*string `json:"trustedSigners" yaml:"trustedSigners"`
}

// A distribution configuration.
//
// TODO: EXAMPLE
//
type CfnDistribution_DistributionConfigProperty struct {
	// From this field, you can enable or disable the selected distribution.
	Enabled interface{} `json:"enabled" yaml:"enabled"`
	// A complex type that contains information about CNAMEs (alternate domain names), if any, for this distribution.
	Aliases *[]*string `json:"aliases" yaml:"aliases"`
	// A complex type that contains zero or more `CacheBehavior` elements.
	CacheBehaviors interface{} `json:"cacheBehaviors" yaml:"cacheBehaviors"`
	// `CfnDistribution.DistributionConfigProperty.CNAMEs`.
	CnamEs *[]*string `json:"cnamEs" yaml:"cnamEs"`
	// An optional comment to describe the distribution.
	//
	// The comment cannot be longer than 128 characters.
	Comment *string `json:"comment" yaml:"comment"`
	// A complex type that controls the following:.
	//
	// - Whether CloudFront replaces HTTP status codes in the 4xx and 5xx range with custom error messages before returning the response to the viewer.
	// - How long CloudFront caches HTTP status codes in the 4xx and 5xx range.
	//
	// For more information about custom error pages, see [Customizing Error Responses](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/custom-error-pages.html) in the *Amazon CloudFront Developer Guide* .
	CustomErrorResponses interface{} `json:"customErrorResponses" yaml:"customErrorResponses"`
	// `CfnDistribution.DistributionConfigProperty.CustomOrigin`.
	CustomOrigin interface{} `json:"customOrigin" yaml:"customOrigin"`
	// A complex type that describes the default cache behavior if you don't specify a `CacheBehavior` element or if files don't match any of the values of `PathPattern` in `CacheBehavior` elements.
	//
	// You must create exactly one default cache behavior.
	DefaultCacheBehavior interface{} `json:"defaultCacheBehavior" yaml:"defaultCacheBehavior"`
	// The object that you want CloudFront to request from your origin (for example, `index.html` ) when a viewer requests the root URL for your distribution ( `http://www.example.com` ) instead of an object in your distribution ( `http://www.example.com/product-description.html` ). Specifying a default root object avoids exposing the contents of your distribution.
	//
	// Specify only the object name, for example, `index.html` . Don't add a `/` before the object name.
	//
	// If you don't want to specify a default root object when you create a distribution, include an empty `DefaultRootObject` element.
	//
	// To delete the default root object from an existing distribution, update the distribution configuration and include an empty `DefaultRootObject` element.
	//
	// To replace the default root object, update the distribution configuration and specify the new object.
	//
	// For more information about the default root object, see [Creating a Default Root Object](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/DefaultRootObject.html) in the *Amazon CloudFront Developer Guide* .
	DefaultRootObject *string `json:"defaultRootObject" yaml:"defaultRootObject"`
	// (Optional) Specify the maximum HTTP version that you want viewers to use to communicate with CloudFront .
	//
	// The default value for new web distributions is `http1.1` .
	//
	// For viewers and CloudFront to use HTTP/2, viewers must support TLS 1.2 or later, and must support server name identification (SNI).
	//
	// In general, configuring CloudFront to communicate with viewers using HTTP/2 reduces latency. You can improve performance by optimizing for HTTP/2.
	HttpVersion *string `json:"httpVersion" yaml:"httpVersion"`
	// If you want CloudFront to respond to IPv6 DNS requests with an IPv6 address for your distribution, specify `true` .
	//
	// If you specify `false` , CloudFront responds to IPv6 DNS requests with the DNS response code `NOERROR` and with no IP addresses. This allows viewers to submit a second request, for an IPv4 address for your distribution.
	//
	// In general, you should enable IPv6 if you have users on IPv6 networks who want to access your content. However, if you're using signed URLs or signed cookies to restrict access to your content, and if you're using a custom policy that includes the `IpAddress` parameter to restrict the IP addresses that can access your content, don't enable IPv6. If you want to restrict access to some content by IP address and not restrict access to other content (or restrict access but not by IP address), you can create two distributions. For more information, see [Creating a Signed URL Using a Custom Policy](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/private-content-creating-signed-url-custom-policy.html) in the *Amazon CloudFront Developer Guide* .
	//
	// If you're using an Amazon Route 53 AWS Integration alias resource record set to route traffic to your CloudFront distribution, you need to create a second alias resource record set when both of the following are true:
	//
	// - You enable IPv6 for the distribution
	// - You're using alternate domain names in the URLs for your objects
	//
	// For more information, see [Routing Traffic to an Amazon CloudFront Web Distribution by Using Your Domain Name](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/routing-to-cloudfront-distribution.html) in the *Amazon Route 53 AWS Integration Developer Guide* .
	//
	// If you created a CNAME resource record set, either with Amazon Route 53 AWS Integration or with another DNS service, you don't need to make any changes. A CNAME record will route traffic to your distribution regardless of the IP address format of the viewer request.
	Ipv6Enabled interface{} `json:"ipv6Enabled" yaml:"ipv6Enabled"`
	// A complex type that controls whether access logs are written for the distribution.
	//
	// For more information about logging, see [Access Logs](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/AccessLogs.html) in the *Amazon CloudFront Developer Guide* .
	Logging interface{} `json:"logging" yaml:"logging"`
	// A complex type that contains information about origin groups for this distribution.
	OriginGroups interface{} `json:"originGroups" yaml:"originGroups"`
	// A complex type that contains information about origins for this distribution.
	Origins interface{} `json:"origins" yaml:"origins"`
	// The price class that corresponds with the maximum price that you want to pay for CloudFront service.
	//
	// If you specify `PriceClass_All` , CloudFront responds to requests for your objects from all CloudFront edge locations.
	//
	// If you specify a price class other than `PriceClass_All` , CloudFront serves your objects from the CloudFront edge location that has the lowest latency among the edge locations in your price class. Viewers who are in or near regions that are excluded from your specified price class may encounter slower performance.
	//
	// For more information about price classes, see [Choosing the Price Class for a CloudFront Distribution](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/PriceClass.html) in the *Amazon CloudFront Developer Guide* . For information about CloudFront pricing, including how price classes (such as Price Class 100) map to CloudFront regions, see [Amazon CloudFront Pricing](https://docs.aws.amazon.com/cloudfront/pricing/) .
	PriceClass *string `json:"priceClass" yaml:"priceClass"`
	// A complex type that identifies ways in which you want to restrict distribution of your content.
	Restrictions interface{} `json:"restrictions" yaml:"restrictions"`
	// `CfnDistribution.DistributionConfigProperty.S3Origin`.
	S3Origin interface{} `json:"s3Origin" yaml:"s3Origin"`
	// A complex type that determines the distribution’s SSL/TLS configuration for communicating with viewers.
	ViewerCertificate interface{} `json:"viewerCertificate" yaml:"viewerCertificate"`
	// A unique identifier that specifies the AWS WAF web ACL, if any, to associate with this distribution.
	//
	// To specify a web ACL created using the latest version of AWS WAF , use the ACL ARN, for example `arn:aws:wafv2:us-east-1:123456789012:global/webacl/ExampleWebACL/473e64fd-f30b-4765-81a0-62ad96dd167a` . To specify a web ACL created using AWS WAF Classic, use the ACL ID, for example `473e64fd-f30b-4765-81a0-62ad96dd167a` .
	//
	// AWS WAF is a web application firewall that lets you monitor the HTTP and HTTPS requests that are forwarded to CloudFront, and lets you control access to your content. Based on conditions that you specify, such as the IP addresses that requests originate from or the values of query strings, CloudFront responds to requests either with the requested content or with an HTTP 403 status code (Forbidden). You can also configure CloudFront to return a custom error page when a request is blocked. For more information about AWS WAF , see the [AWS WAF Developer Guide](https://docs.aws.amazon.com/waf/latest/developerguide/what-is-aws-waf.html) .
	WebAclId *string `json:"webAclId" yaml:"webAclId"`
}

// This field is deprecated.
//
// We recommend that you use a cache policy or an origin request policy instead of this field.
//
// If you want to include values in the cache key, use a cache policy. For more information, see [Creating cache policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-the-cache-key.html#cache-key-create-cache-policy) in the *Amazon CloudFront Developer Guide* .
//
// If you want to send values to the origin but not include them in the cache key, use an origin request policy. For more information, see [Creating origin request policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-origin-requests.html#origin-request-create-origin-request-policy) in the *Amazon CloudFront Developer Guide* .
//
// A complex type that specifies how CloudFront handles query strings, cookies, and HTTP headers.
//
// TODO: EXAMPLE
//
type CfnDistribution_ForwardedValuesProperty struct {
	// This field is deprecated.
	//
	// We recommend that you use a cache policy or an origin request policy instead of this field.
	//
	// If you want to include query strings in the cache key, use a cache policy. For more information, see [Creating cache policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-the-cache-key.html#cache-key-create-cache-policy) in the *Amazon CloudFront Developer Guide* .
	//
	// If you want to send query strings to the origin but not include them in the cache key, use an origin request policy. For more information, see [Creating origin request policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-origin-requests.html#origin-request-create-origin-request-policy) in the *Amazon CloudFront Developer Guide* .
	//
	// Indicates whether you want CloudFront to forward query strings to the origin that is associated with this cache behavior and cache based on the query string parameters. CloudFront behavior depends on the value of `QueryString` and on the values that you specify for `QueryStringCacheKeys` , if any:
	//
	// If you specify true for `QueryString` and you don't specify any values for `QueryStringCacheKeys` , CloudFront forwards all query string parameters to the origin and caches based on all query string parameters. Depending on how many query string parameters and values you have, this can adversely affect performance because CloudFront must forward more requests to the origin.
	//
	// If you specify true for `QueryString` and you specify one or more values for `QueryStringCacheKeys` , CloudFront forwards all query string parameters to the origin, but it only caches based on the query string parameters that you specify.
	//
	// If you specify false for `QueryString` , CloudFront doesn't forward any query string parameters to the origin, and doesn't cache based on query string parameters.
	//
	// For more information, see [Configuring CloudFront to Cache Based on Query String Parameters](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/QueryStringParameters.html) in the *Amazon CloudFront Developer Guide* .
	QueryString interface{} `json:"queryString" yaml:"queryString"`
	// This field is deprecated.
	//
	// We recommend that you use a cache policy or an origin request policy instead of this field.
	//
	// If you want to include cookies in the cache key, use a cache policy. For more information, see [Creating cache policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-the-cache-key.html#cache-key-create-cache-policy) in the *Amazon CloudFront Developer Guide* .
	//
	// If you want to send cookies to the origin but not include them in the cache key, use an origin request policy. For more information, see [Creating origin request policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-origin-requests.html#origin-request-create-origin-request-policy) in the *Amazon CloudFront Developer Guide* .
	//
	// A complex type that specifies whether you want CloudFront to forward cookies to the origin and, if so, which ones. For more information about forwarding cookies to the origin, see [How CloudFront Forwards, Caches, and Logs Cookies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/Cookies.html) in the *Amazon CloudFront Developer Guide* .
	Cookies interface{} `json:"cookies" yaml:"cookies"`
	// This field is deprecated.
	//
	// We recommend that you use a cache policy or an origin request policy instead of this field.
	//
	// If you want to include headers in the cache key, use a cache policy. For more information, see [Creating cache policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-the-cache-key.html#cache-key-create-cache-policy) in the *Amazon CloudFront Developer Guide* .
	//
	// If you want to send headers to the origin but not include them in the cache key, use an origin request policy. For more information, see [Creating origin request policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-origin-requests.html#origin-request-create-origin-request-policy) in the *Amazon CloudFront Developer Guide* .
	//
	// A complex type that specifies the `Headers` , if any, that you want CloudFront to forward to the origin for this cache behavior (whitelisted headers). For the headers that you specify, CloudFront also caches separate versions of a specified object that is based on the header values in viewer requests.
	//
	// For more information, see [Caching Content Based on Request Headers](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/header-caching.html) in the *Amazon CloudFront Developer Guide* .
	Headers *[]*string `json:"headers" yaml:"headers"`
	// This field is deprecated.
	//
	// We recommend that you use a cache policy or an origin request policy instead of this field.
	//
	// If you want to include query strings in the cache key, use a cache policy. For more information, see [Creating cache policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-the-cache-key.html#cache-key-create-cache-policy) in the *Amazon CloudFront Developer Guide* .
	//
	// If you want to send query strings to the origin but not include them in the cache key, use an origin request policy. For more information, see [Creating origin request policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-origin-requests.html#origin-request-create-origin-request-policy) in the *Amazon CloudFront Developer Guide* .
	//
	// A complex type that contains information about the query string parameters that you want CloudFront to use for caching for this cache behavior.
	QueryStringCacheKeys *[]*string `json:"queryStringCacheKeys" yaml:"queryStringCacheKeys"`
}

// A CloudFront function that is associated with a cache behavior in a CloudFront distribution.
//
// TODO: EXAMPLE
//
type CfnDistribution_FunctionAssociationProperty struct {
	// The event type of the function, either `viewer-request` or `viewer-response` .
	//
	// You cannot use origin-facing event types ( `origin-request` and `origin-response` ) with a CloudFront function.
	EventType *string `json:"eventType" yaml:"eventType"`
	// The Amazon Resource Name (ARN) of the function.
	FunctionArn *string `json:"functionArn" yaml:"functionArn"`
}

// A complex type that controls the countries in which your content is distributed.
//
// CloudFront determines the location of your users using `MaxMind` GeoIP databases. To disable geo restriction, remove the [Restrictions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-distributionconfig.html#cfn-cloudfront-distribution-distributionconfig-restrictions) property from your stack template.
//
// TODO: EXAMPLE
//
type CfnDistribution_GeoRestrictionProperty struct {
	// The method that you want to use to restrict distribution of your content by country:.
	//
	// - `none` : No geo restriction is enabled, meaning access to content is not restricted by client geo location.
	// - `blacklist` : The `Location` elements specify the countries in which you don't want CloudFront to distribute your content.
	// - `whitelist` : The `Location` elements specify the countries in which you want CloudFront to distribute your content.
	RestrictionType *string `json:"restrictionType" yaml:"restrictionType"`
	// A complex type that contains a `Location` element for each country in which you want CloudFront either to distribute your content ( `whitelist` ) or not distribute your content ( `blacklist` ).
	//
	// The `Location` element is a two-letter, uppercase country code for a country that you want to include in your `blacklist` or `whitelist` . Include one `Location` element for each country.
	//
	// CloudFront and `MaxMind` both use `ISO 3166` country codes. For the current list of countries and the corresponding codes, see `ISO 3166-1-alpha-2` code on the *International Organization for Standardization* website. You can also refer to the country list on the CloudFront console, which includes both country names and codes.
	Locations *[]*string `json:"locations" yaml:"locations"`
}

// A complex type that contains a Lambda@Edge function association.
//
// TODO: EXAMPLE
//
type CfnDistribution_LambdaFunctionAssociationProperty struct {
	// Specifies the event type that triggers a Lambda@Edge function invocation. You can specify the following values:.
	//
	// - `viewer-request` : The function executes when CloudFront receives a request from a viewer and before it checks to see whether the requested object is in the edge cache.
	// - `origin-request` : The function executes only when CloudFront sends a request to your origin. When the requested object is in the edge cache, the function doesn't execute.
	// - `origin-response` : The function executes after CloudFront receives a response from the origin and before it caches the object in the response. When the requested object is in the edge cache, the function doesn't execute.
	// - `viewer-response` : The function executes before CloudFront returns the requested object to the viewer. The function executes regardless of whether the object was already in the edge cache.
	//
	// If the origin returns an HTTP status code other than HTTP 200 (OK), the function doesn't execute.
	EventType *string `json:"eventType" yaml:"eventType"`
	// A flag that allows a Lambda@Edge function to have read access to the body content.
	//
	// For more information, see [Accessing the Request Body by Choosing the Include Body Option](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/lambda-include-body-access.html) in the Amazon CloudFront Developer Guide.
	IncludeBody interface{} `json:"includeBody" yaml:"includeBody"`
	// The ARN of the Lambda@Edge function.
	//
	// You must specify the ARN of a function version; you can't specify an alias or $LATEST.
	LambdaFunctionArn *string `json:"lambdaFunctionArn" yaml:"lambdaFunctionArn"`
}

// TODO: EXAMPLE
//
type CfnDistribution_LegacyCustomOriginProperty struct {
	// `CfnDistribution.LegacyCustomOriginProperty.DNSName`.
	DnsName *string `json:"dnsName" yaml:"dnsName"`
	// `CfnDistribution.LegacyCustomOriginProperty.OriginProtocolPolicy`.
	OriginProtocolPolicy *string `json:"originProtocolPolicy" yaml:"originProtocolPolicy"`
	// `CfnDistribution.LegacyCustomOriginProperty.OriginSSLProtocols`.
	OriginSslProtocols *[]*string `json:"originSslProtocols" yaml:"originSslProtocols"`
	// `CfnDistribution.LegacyCustomOriginProperty.HTTPPort`.
	HttpPort *float64 `json:"httpPort" yaml:"httpPort"`
	// `CfnDistribution.LegacyCustomOriginProperty.HTTPSPort`.
	HttpsPort *float64 `json:"httpsPort" yaml:"httpsPort"`
}

// TODO: EXAMPLE
//
type CfnDistribution_LegacyS3OriginProperty struct {
	// `CfnDistribution.LegacyS3OriginProperty.DNSName`.
	DnsName *string `json:"dnsName" yaml:"dnsName"`
	// `CfnDistribution.LegacyS3OriginProperty.OriginAccessIdentity`.
	OriginAccessIdentity *string `json:"originAccessIdentity" yaml:"originAccessIdentity"`
}

// A complex type that controls whether access logs are written for the distribution.
//
// TODO: EXAMPLE
//
type CfnDistribution_LoggingProperty struct {
	// The Amazon S3 bucket to store the access logs in, for example, `myawslogbucket.s3.amazonaws.com` .
	Bucket *string `json:"bucket" yaml:"bucket"`
	// Specifies whether you want CloudFront to include cookies in access logs, specify `true` for `IncludeCookies` .
	//
	// If you choose to include cookies in logs, CloudFront logs all cookies regardless of how you configure the cache behaviors for this distribution. If you don't want to include cookies when you create a distribution or if you want to disable include cookies for an existing distribution, specify `false` for `IncludeCookies` .
	IncludeCookies interface{} `json:"includeCookies" yaml:"includeCookies"`
	// An optional string that you want CloudFront to prefix to the access log `filenames` for this distribution, for example, `myprefix/` .
	//
	// If you want to enable logging, but you don't want to specify a prefix, you still must include an empty `Prefix` element in the `Logging` element.
	Prefix *string `json:"prefix" yaml:"prefix"`
}

// A complex type that contains `HeaderName` and `HeaderValue` elements, if any, for this distribution.
//
// TODO: EXAMPLE
//
type CfnDistribution_OriginCustomHeaderProperty struct {
	// The name of a header that you want CloudFront to send to your origin.
	//
	// For more information, see [Adding Custom Headers to Origin Requests](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/forward-custom-headers.html) in the *Amazon CloudFront Developer Guide* .
	HeaderName *string `json:"headerName" yaml:"headerName"`
	// The value for the header that you specified in the `HeaderName` field.
	HeaderValue *string `json:"headerValue" yaml:"headerValue"`
}

// A complex data type that includes information about the failover criteria for an origin group, including the status codes for which CloudFront will failover from the primary origin to the second origin.
//
// TODO: EXAMPLE
//
type CfnDistribution_OriginGroupFailoverCriteriaProperty struct {
	// The status codes that, when returned from the primary origin, will trigger CloudFront to failover to the second origin.
	StatusCodes interface{} `json:"statusCodes" yaml:"statusCodes"`
}

// An origin in an origin group.
//
// TODO: EXAMPLE
//
type CfnDistribution_OriginGroupMemberProperty struct {
	// The ID for an origin in an origin group.
	OriginId *string `json:"originId" yaml:"originId"`
}

// A complex data type for the origins included in an origin group.
//
// TODO: EXAMPLE
//
type CfnDistribution_OriginGroupMembersProperty struct {
	// Items (origins) in an origin group.
	Items interface{} `json:"items" yaml:"items"`
	// The number of origins in an origin group.
	Quantity *float64 `json:"quantity" yaml:"quantity"`
}

// An origin group includes two origins (a primary origin and a second origin to failover to) and a failover criteria that you specify.
//
// You create an origin group to support origin failover in CloudFront. When you create or update a distribution, you can specifiy the origin group instead of a single origin, and CloudFront will failover from the primary origin to the second origin under the failover conditions that you've chosen.
//
// TODO: EXAMPLE
//
type CfnDistribution_OriginGroupProperty struct {
	// A complex type that contains information about the failover criteria for an origin group.
	FailoverCriteria interface{} `json:"failoverCriteria" yaml:"failoverCriteria"`
	// The origin group's ID.
	Id *string `json:"id" yaml:"id"`
	// A complex type that contains information about the origins in an origin group.
	Members interface{} `json:"members" yaml:"members"`
}

// A complex data type for the origin groups specified for a distribution.
//
// TODO: EXAMPLE
//
type CfnDistribution_OriginGroupsProperty struct {
	// The number of origin groups.
	Quantity *float64 `json:"quantity" yaml:"quantity"`
	// The items (origin groups) in a distribution.
	Items interface{} `json:"items" yaml:"items"`
}

// An origin.
//
// An origin is the location where content is stored, and from which CloudFront gets content to serve to viewers. To specify an origin:
//
// - Use `S3OriginConfig` to specify an Amazon S3 bucket that is not configured with static website hosting.
// - Use `CustomOriginConfig` to specify all other kinds of origins, including:
//
// - An Amazon S3 bucket that is configured with static website hosting
// - An Elastic Load Balancing load balancer
// - An AWS Elemental MediaPackage endpoint
// - An AWS Elemental MediaStore container
// - Any other HTTP server, running on an Amazon EC2 instance or any other kind of host
//
// For the current maximum number of origins that you can specify per distribution, see [General Quotas on Web Distributions](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/cloudfront-limits.html#limits-web-distributions) in the *Amazon CloudFront Developer Guide* (quotas were formerly referred to as limits).
//
// TODO: EXAMPLE
//
type CfnDistribution_OriginProperty struct {
	// The domain name for the origin.
	//
	// For more information, see [Origin Domain Name](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/distribution-web-values-specify.html#DownloadDistValuesDomainName) in the *Amazon CloudFront Developer Guide* .
	DomainName *string `json:"domainName" yaml:"domainName"`
	// A unique identifier for the origin. This value must be unique within the distribution.
	//
	// Use this value to specify the `TargetOriginId` in a `CacheBehavior` or `DefaultCacheBehavior` .
	Id *string `json:"id" yaml:"id"`
	// The number of times that CloudFront attempts to connect to the origin.
	//
	// The minimum number is 1, the maximum is 3, and the default (if you don’t specify otherwise) is 3.
	//
	// For a custom origin (including an Amazon S3 bucket that’s configured with static website hosting), this value also specifies the number of times that CloudFront attempts to get a response from the origin, in the case of an [Origin Response Timeout](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/distribution-web-values-specify.html#DownloadDistValuesOriginResponseTimeout) .
	//
	// For more information, see [Origin Connection Attempts](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/distribution-web-values-specify.html#origin-connection-attempts) in the *Amazon CloudFront Developer Guide* .
	ConnectionAttempts *float64 `json:"connectionAttempts" yaml:"connectionAttempts"`
	// The number of seconds that CloudFront waits when trying to establish a connection to the origin.
	//
	// The minimum timeout is 1 second, the maximum is 10 seconds, and the default (if you don’t specify otherwise) is 10 seconds.
	//
	// For more information, see [Origin Connection Timeout](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/distribution-web-values-specify.html#origin-connection-timeout) in the *Amazon CloudFront Developer Guide* .
	ConnectionTimeout *float64 `json:"connectionTimeout" yaml:"connectionTimeout"`
	// Use this type to specify an origin that is not an Amazon S3 bucket, with one exception.
	//
	// If the Amazon S3 bucket is configured with static website hosting, use this type. If the Amazon S3 bucket is not configured with static website hosting, use the `S3OriginConfig` type instead.
	CustomOriginConfig interface{} `json:"customOriginConfig" yaml:"customOriginConfig"`
	// A list of HTTP header names and values that CloudFront adds to the requests that it sends to the origin.
	//
	// For more information, see [Adding Custom Headers to Origin Requests](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/add-origin-custom-headers.html) in the *Amazon CloudFront Developer Guide* .
	OriginCustomHeaders interface{} `json:"originCustomHeaders" yaml:"originCustomHeaders"`
	// An optional path that CloudFront appends to the origin domain name when CloudFront requests content from the origin.
	//
	// For more information, see [Origin Path](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/distribution-web-values-specify.html#DownloadDistValuesOriginPath) in the *Amazon CloudFront Developer Guide* .
	OriginPath *string `json:"originPath" yaml:"originPath"`
	// CloudFront Origin Shield. Using Origin Shield can help reduce the load on your origin.
	//
	// For more information, see [Using Origin Shield](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/origin-shield.html) in the *Amazon CloudFront Developer Guide* .
	OriginShield interface{} `json:"originShield" yaml:"originShield"`
	// Use this type to specify an origin that is an Amazon S3 bucket that is not configured with static website hosting.
	//
	// To specify any other type of origin, including an Amazon S3 bucket that is configured with static website hosting, use the `CustomOriginConfig` type instead.
	S3OriginConfig interface{} `json:"s3OriginConfig" yaml:"s3OriginConfig"`
}

// CloudFront Origin Shield.
//
// Using Origin Shield can help reduce the load on your origin. For more information, see [Using Origin Shield](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/origin-shield.html) in the *Amazon CloudFront Developer Guide* .
//
// TODO: EXAMPLE
//
type CfnDistribution_OriginShieldProperty struct {
	// A flag that specifies whether Origin Shield is enabled.
	//
	// When it’s enabled, CloudFront routes all requests through Origin Shield, which can help protect your origin. When it’s disabled, CloudFront might send requests directly to your origin from multiple edge locations or regional edge caches.
	Enabled interface{} `json:"enabled" yaml:"enabled"`
	// The AWS Region for Origin Shield.
	//
	// Specify the AWS Region that has the lowest latency to your origin. To specify a region, use the region code, not the region name. For example, specify the US East (Ohio) region as `us-east-2` .
	//
	// When you enable CloudFront Origin Shield, you must specify the AWS Region for Origin Shield. For the list of AWS Regions that you can specify, and for help choosing the best Region for your origin, see [Choosing the AWS Region for Origin Shield](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/origin-shield.html#choose-origin-shield-region) in the *Amazon CloudFront Developer Guide* .
	OriginShieldRegion *string `json:"originShieldRegion" yaml:"originShieldRegion"`
}

// A complex type that identifies ways in which you want to restrict distribution of your content.
//
// TODO: EXAMPLE
//
type CfnDistribution_RestrictionsProperty struct {
	// A complex type that controls the countries in which your content is distributed.
	//
	// CloudFront determines the location of your users using `MaxMind` GeoIP databases. To disable geo restriction, remove the [Restrictions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-distributionconfig.html#cfn-cloudfront-distribution-distributionconfig-restrictions) property from your stack template.
	GeoRestriction interface{} `json:"geoRestriction" yaml:"geoRestriction"`
}

// A complex type that contains information about the Amazon S3 origin.
//
// If the origin is a custom origin or an S3 bucket that is configured as a website endpoint, use the `CustomOriginConfig` element instead.
//
// TODO: EXAMPLE
//
type CfnDistribution_S3OriginConfigProperty struct {
	// The CloudFront origin access identity to associate with the origin.
	//
	// Use an origin access identity to configure the origin so that viewers can *only* access objects in an Amazon S3 bucket through CloudFront. The format of the value is:
	//
	// origin-access-identity/cloudfront/ *ID-of-origin-access-identity*
	//
	// where `*ID-of-origin-access-identity*` is the value that CloudFront returned in the `ID` element when you created the origin access identity.
	//
	// If you want viewers to be able to access objects using either the CloudFront URL or the Amazon S3 URL, specify an empty `OriginAccessIdentity` element.
	//
	// To delete the origin access identity from an existing distribution, update the distribution configuration and include an empty `OriginAccessIdentity` element.
	//
	// To replace the origin access identity, update the distribution configuration and specify the new origin access identity.
	//
	// For more information about the origin access identity, see [Serving Private Content through CloudFront](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/PrivateContent.html) in the *Amazon CloudFront Developer Guide* .
	OriginAccessIdentity *string `json:"originAccessIdentity" yaml:"originAccessIdentity"`
}

// A complex data type for the status codes that you specify that, when returned by a primary origin, trigger CloudFront to failover to a second origin.
//
// TODO: EXAMPLE
//
type CfnDistribution_StatusCodesProperty struct {
	// The items (status codes) for an origin group.
	Items interface{} `json:"items" yaml:"items"`
	// The number of status codes.
	Quantity *float64 `json:"quantity" yaml:"quantity"`
}

// A complex type that determines the distribution’s SSL/TLS configuration for communicating with viewers.
//
// If the distribution doesn’t use `Aliases` (also known as alternate domain names or CNAMEs)—that is, if the distribution uses the CloudFront domain name such as `d111111abcdef8.cloudfront.net` —set `CloudFrontDefaultCertificate` to `true` and leave all other fields empty.
//
// If the distribution uses `Aliases` (alternate domain names or CNAMEs), use the fields in this type to specify the following settings:
//
// - Which viewers the distribution accepts HTTPS connections from: only viewers that support [server name indication (SNI)](https://docs.aws.amazon.com/https://en.wikipedia.org/wiki/Server_Name_Indication) (recommended), or all viewers including those that don’t support SNI.
//
// - To accept HTTPS connections from only viewers that support SNI, set `SSLSupportMethod` to `sni-only` . This is recommended. Most browsers and clients support SNI. (In CloudFormation, the field name is `SslSupportMethod` . Note the different capitalization.)
// - To accept HTTPS connections from all viewers, including those that don’t support SNI, set `SSLSupportMethod` to `vip` . This is not recommended, and results in additional monthly charges from CloudFront. (In CloudFormation, the field name is `SslSupportMethod` . Note the different capitalization.)
// - The minimum SSL/TLS protocol version that the distribution can use to communicate with viewers. To specify a minimum version, choose a value for `MinimumProtocolVersion` . For more information, see [Security Policy](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/distribution-web-values-specify.html#DownloadDistValues-security-policy) in the *Amazon CloudFront Developer Guide* .
// - The location of the SSL/TLS certificate, [AWS Certificate Manager (ACM)](https://docs.aws.amazon.com/acm/latest/userguide/acm-overview.html) (recommended) or [AWS Identity and Access Management (IAM)](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_server-certs.html) . You specify the location by setting a value in one of the following fields (not both):
//
// - `ACMCertificateArn` (In CloudFormation, this field name is `AcmCertificateArn` . Note the different capitalization.)
// - `IAMCertificateId` (In CloudFormation, this field name is `IamCertificateId` . Note the different capitalization.)
//
// All distributions support HTTPS connections from viewers. To require viewers to use HTTPS only, or to redirect them from HTTP to HTTPS, use `ViewerProtocolPolicy` in the `CacheBehavior` or `DefaultCacheBehavior` . To specify how CloudFront should use SSL/TLS to communicate with your custom origin, use `CustomOriginConfig` .
//
// For more information, see [Using HTTPS with CloudFront](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/using-https.html) and [Using Alternate Domain Names and HTTPS](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/using-https-alternate-domain-names.html) in the *Amazon CloudFront Developer Guide* .
//
// TODO: EXAMPLE
//
type CfnDistribution_ViewerCertificateProperty struct {
	// > In CloudFormation, this field name is `AcmCertificateArn` . Note the different capitalization.
	//
	// If the distribution uses `Aliases` (alternate domain names or CNAMEs) and the SSL/TLS certificate is stored in [AWS Certificate Manager (ACM)](https://docs.aws.amazon.com/acm/latest/userguide/acm-overview.html) , provide the Amazon Resource Name (ARN) of the ACM certificate. CloudFront only supports ACM certificates in the US East (N. Virginia) Region ( `us-east-1` ).
	//
	// If you specify an ACM certificate ARN, you must also specify values for `MinimumProtocolVersion` and `SSLSupportMethod` . (In CloudFormation, the field name is `SslSupportMethod` . Note the different capitalization.)
	AcmCertificateArn *string `json:"acmCertificateArn" yaml:"acmCertificateArn"`
	// If the distribution uses the CloudFront domain name such as `d111111abcdef8.cloudfront.net` , set this field to `true` .
	//
	// If the distribution uses `Aliases` (alternate domain names or CNAMEs), set this field to `false` and specify values for the following fields:
	//
	// - `ACMCertificateArn` or `IAMCertificateId` (specify a value for one, not both)
	//
	// In CloudFormation, these field names are `AcmCertificateArn` and `IamCertificateId` . Note the different capitalization.
	// - `MinimumProtocolVersion`
	// - `SSLSupportMethod` (In CloudFormation, this field name is `SslSupportMethod` . Note the different capitalization.)
	CloudFrontDefaultCertificate interface{} `json:"cloudFrontDefaultCertificate" yaml:"cloudFrontDefaultCertificate"`
	// > In CloudFormation, this field name is `IamCertificateId` . Note the different capitalization.
	//
	// If the distribution uses `Aliases` (alternate domain names or CNAMEs) and the SSL/TLS certificate is stored in [AWS Identity and Access Management (IAM)](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_server-certs.html) , provide the ID of the IAM certificate.
	//
	// If you specify an IAM certificate ID, you must also specify values for `MinimumProtocolVersion` and `SSLSupportMethod` . (In CloudFormation, the field name is `SslSupportMethod` . Note the different capitalization.)
	IamCertificateId *string `json:"iamCertificateId" yaml:"iamCertificateId"`
	// If the distribution uses `Aliases` (alternate domain names or CNAMEs), specify the security policy that you want CloudFront to use for HTTPS connections with viewers.
	//
	// The security policy determines two settings:
	//
	// - The minimum SSL/TLS protocol that CloudFront can use to communicate with viewers.
	// - The ciphers that CloudFront can use to encrypt the content that it returns to viewers.
	//
	// For more information, see [Security Policy](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/distribution-web-values-specify.html#DownloadDistValues-security-policy) and [Supported Protocols and Ciphers Between Viewers and CloudFront](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/secure-connections-supported-viewer-protocols-ciphers.html#secure-connections-supported-ciphers) in the *Amazon CloudFront Developer Guide* .
	//
	// > On the CloudFront console, this setting is called *Security Policy* .
	//
	// When you’re using SNI only (you set `SSLSupportMethod` to `sni-only` ), you must specify `TLSv1` or higher. (In CloudFormation, the field name is `SslSupportMethod` . Note the different capitalization.)
	//
	// If the distribution uses the CloudFront domain name such as `d111111abcdef8.cloudfront.net` (you set `CloudFrontDefaultCertificate` to `true` ), CloudFront automatically sets the security policy to `TLSv1` regardless of the value that you set here.
	MinimumProtocolVersion *string `json:"minimumProtocolVersion" yaml:"minimumProtocolVersion"`
	// > In CloudFormation, this field name is `SslSupportMethod` . Note the different capitalization.
	//
	// If the distribution uses `Aliases` (alternate domain names or CNAMEs), specify which viewers the distribution accepts HTTPS connections from.
	//
	// - `sni-only` – The distribution accepts HTTPS connections from only viewers that support [server name indication (SNI)](https://docs.aws.amazon.com/https://en.wikipedia.org/wiki/Server_Name_Indication) . This is recommended. Most browsers and clients support SNI.
	// - `vip` – The distribution accepts HTTPS connections from all viewers including those that don’t support SNI. This is not recommended, and results in additional monthly charges from CloudFront.
	// - `static-ip` - Do not specify this value unless your distribution has been enabled for this feature by the CloudFront team. If you have a use case that requires static IP addresses for a distribution, contact CloudFront through the [AWS Support Center](https://docs.aws.amazon.com/support/home) .
	//
	// If the distribution uses the CloudFront domain name such as `d111111abcdef8.cloudfront.net` , don’t set a value for this field.
	SslSupportMethod *string `json:"sslSupportMethod" yaml:"sslSupportMethod"`
}

// Properties for defining a `CfnDistribution`.
//
// TODO: EXAMPLE
//
type CfnDistributionProps struct {
	// The current configuration information for the distribution.
	//
	// Send a `GET` request to the `/ *CloudFront API version* /distribution ID/config` resource.
	DistributionConfig interface{} `json:"distributionConfig" yaml:"distributionConfig"`
	// A complex type that contains zero or more `Tag` elements.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::CloudFront::Function`.
//
// Creates a CloudFront function.
//
// To create a function, you provide the function code and some configuration information about the function. The response contains an Amazon Resource Name (ARN) that uniquely identifies the function, and the function’s stage.
//
// By default, when you create a function, it’s in the `DEVELOPMENT` stage. In this stage, you can [test the function](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/test-function.html) in the CloudFront console (or with `TestFunction` in the CloudFront API).
//
// When you’re ready to use your function with a CloudFront distribution, publish the function to the `LIVE` stage. You can do this in the CloudFront console, with `PublishFunction` in the CloudFront API, or by updating the `AWS::CloudFront::Function` resource with the `AutoPublish` property set to `true` . When the function is published to the `LIVE` stage, you can attach it to a distribution’s cache behavior, using the function’s ARN.
//
// To automatically publish the function to the `LIVE` stage when it’s created, set the `AutoPublish` property to `true` .
//
// TODO: EXAMPLE
//
type CfnFunction interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrFunctionArn() *string
	AttrFunctionMetadataFunctionArn() *string
	AttrStage() *string
	AutoPublish() interface{}
	SetAutoPublish(val interface{})
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	FunctionCode() *string
	SetFunctionCode(val *string)
	FunctionConfig() interface{}
	SetFunctionConfig(val interface{})
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

// The jsii proxy struct for CfnFunction
type jsiiProxy_CfnFunction struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnFunction) AttrFunctionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrFunctionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) AttrFunctionMetadataFunctionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrFunctionMetadataFunctionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) AttrStage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrStage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) AutoPublish() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoPublish",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) FunctionCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) FunctionConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"functionConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::CloudFront::Function`.
func NewCfnFunction(scope constructs.Construct, id *string, props *CfnFunctionProps) CfnFunction {
	_init_.Initialize()

	j := jsiiProxy_CfnFunction{}

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront.CfnFunction",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::CloudFront::Function`.
func NewCfnFunction_Override(c CfnFunction, scope constructs.Construct, id *string, props *CfnFunctionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront.CfnFunction",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnFunction) SetAutoPublish(val interface{}) {
	_jsii_.Set(
		j,
		"autoPublish",
		val,
	)
}

func (j *jsiiProxy_CfnFunction) SetFunctionCode(val *string) {
	_jsii_.Set(
		j,
		"functionCode",
		val,
	)
}

func (j *jsiiProxy_CfnFunction) SetFunctionConfig(val interface{}) {
	_jsii_.Set(
		j,
		"functionConfig",
		val,
	)
}

func (j *jsiiProxy_CfnFunction) SetName(val *string) {
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
func CfnFunction_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.CfnFunction",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnFunction_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.CfnFunction",
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
func CfnFunction_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.CfnFunction",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnFunction_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cloudfront.CfnFunction",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnFunction) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnFunction) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnFunction) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnFunction) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnFunction) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnFunction) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnFunction) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnFunction) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnFunction) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnFunction) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnFunction) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnFunction) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnFunction) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnFunction) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFunction) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Contains configuration information about a CloudFront function.
//
// TODO: EXAMPLE
//
type CfnFunction_FunctionConfigProperty struct {
	// A comment to describe the function.
	Comment *string `json:"comment" yaml:"comment"`
	// The function’s runtime environment.
	//
	// The only valid value is `cloudfront-js-1.0` .
	Runtime *string `json:"runtime" yaml:"runtime"`
}

// Contains metadata about a CloudFront function.
//
// TODO: EXAMPLE
//
type CfnFunction_FunctionMetadataProperty struct {
	// The Amazon Resource Name (ARN) of the function.
	//
	// The ARN uniquely identifies the function.
	FunctionArn *string `json:"functionArn" yaml:"functionArn"`
}

// Properties for defining a `CfnFunction`.
//
// TODO: EXAMPLE
//
type CfnFunctionProps struct {
	// A name to identify the function.
	Name *string `json:"name" yaml:"name"`
	// A flag that determines whether to automatically publish the function to the `LIVE` stage when it’s created.
	//
	// To automatically publish to the `LIVE` stage, set this property to `true` .
	AutoPublish interface{} `json:"autoPublish" yaml:"autoPublish"`
	// The function code.
	//
	// For more information about writing a CloudFront function, see [Writing function code for CloudFront Functions](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/writing-function-code.html) in the *Amazon CloudFront Developer Guide* .
	FunctionCode *string `json:"functionCode" yaml:"functionCode"`
	// Contains configuration information about a CloudFront function.
	FunctionConfig interface{} `json:"functionConfig" yaml:"functionConfig"`
}

// A CloudFormation `AWS::CloudFront::KeyGroup`.
//
// A key group.
//
// A key group contains a list of public keys that you can use with [CloudFront signed URLs and signed cookies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/PrivateContent.html) .
//
// TODO: EXAMPLE
//
type CfnKeyGroup interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrId() *string
	AttrLastModifiedTime() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	KeyGroupConfig() interface{}
	SetKeyGroupConfig(val interface{})
	LogicalId() *string
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

// The jsii proxy struct for CfnKeyGroup
type jsiiProxy_CfnKeyGroup struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnKeyGroup) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnKeyGroup) AttrLastModifiedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLastModifiedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnKeyGroup) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnKeyGroup) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnKeyGroup) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnKeyGroup) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnKeyGroup) KeyGroupConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keyGroupConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnKeyGroup) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnKeyGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnKeyGroup) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnKeyGroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnKeyGroup) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::CloudFront::KeyGroup`.
func NewCfnKeyGroup(scope constructs.Construct, id *string, props *CfnKeyGroupProps) CfnKeyGroup {
	_init_.Initialize()

	j := jsiiProxy_CfnKeyGroup{}

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront.CfnKeyGroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::CloudFront::KeyGroup`.
func NewCfnKeyGroup_Override(c CfnKeyGroup, scope constructs.Construct, id *string, props *CfnKeyGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront.CfnKeyGroup",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnKeyGroup) SetKeyGroupConfig(val interface{}) {
	_jsii_.Set(
		j,
		"keyGroupConfig",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnKeyGroup_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.CfnKeyGroup",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnKeyGroup_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.CfnKeyGroup",
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
func CfnKeyGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.CfnKeyGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnKeyGroup_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cloudfront.CfnKeyGroup",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnKeyGroup) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnKeyGroup) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnKeyGroup) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnKeyGroup) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnKeyGroup) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnKeyGroup) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnKeyGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnKeyGroup) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnKeyGroup) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnKeyGroup) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnKeyGroup) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnKeyGroup) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnKeyGroup) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnKeyGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnKeyGroup) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// A key group configuration.
//
// A key group contains a list of public keys that you can use with [CloudFront signed URLs and signed cookies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/PrivateContent.html) .
//
// TODO: EXAMPLE
//
type CfnKeyGroup_KeyGroupConfigProperty struct {
	// A list of the identifiers of the public keys in the key group.
	Items *[]*string `json:"items" yaml:"items"`
	// A name to identify the key group.
	Name *string `json:"name" yaml:"name"`
	// A comment to describe the key group.
	//
	// The comment cannot be longer than 128 characters.
	Comment *string `json:"comment" yaml:"comment"`
}

// Properties for defining a `CfnKeyGroup`.
//
// TODO: EXAMPLE
//
type CfnKeyGroupProps struct {
	// The key group configuration.
	KeyGroupConfig interface{} `json:"keyGroupConfig" yaml:"keyGroupConfig"`
}

// A CloudFormation `AWS::CloudFront::OriginRequestPolicy`.
//
// An origin request policy.
//
// When it’s attached to a cache behavior, the origin request policy determines the values that CloudFront includes in requests that it sends to the origin. Each request that CloudFront sends to the origin includes the following:
//
// - The request body and the URL path (without the domain name) from the viewer request.
// - The headers that CloudFront automatically includes in every origin request, including `Host` , `User-Agent` , and `X-Amz-Cf-Id` .
// - All HTTP headers, cookies, and URL query strings that are specified in the cache policy or the origin request policy. These can include items from the viewer request and, in the case of headers, additional ones that are added by CloudFront.
//
// CloudFront sends a request when it can’t find an object in its cache that matches the request. If you want to send values to the origin and also include them in the cache key, use `CachePolicy` .
//
// TODO: EXAMPLE
//
type CfnOriginRequestPolicy interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrId() *string
	AttrLastModifiedTime() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	LogicalId() *string
	Node() constructs.Node
	OriginRequestPolicyConfig() interface{}
	SetOriginRequestPolicyConfig(val interface{})
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

// The jsii proxy struct for CfnOriginRequestPolicy
type jsiiProxy_CfnOriginRequestPolicy struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnOriginRequestPolicy) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginRequestPolicy) AttrLastModifiedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLastModifiedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginRequestPolicy) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginRequestPolicy) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginRequestPolicy) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginRequestPolicy) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginRequestPolicy) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginRequestPolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginRequestPolicy) OriginRequestPolicyConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"originRequestPolicyConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginRequestPolicy) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginRequestPolicy) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginRequestPolicy) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::CloudFront::OriginRequestPolicy`.
func NewCfnOriginRequestPolicy(scope constructs.Construct, id *string, props *CfnOriginRequestPolicyProps) CfnOriginRequestPolicy {
	_init_.Initialize()

	j := jsiiProxy_CfnOriginRequestPolicy{}

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront.CfnOriginRequestPolicy",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::CloudFront::OriginRequestPolicy`.
func NewCfnOriginRequestPolicy_Override(c CfnOriginRequestPolicy, scope constructs.Construct, id *string, props *CfnOriginRequestPolicyProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront.CfnOriginRequestPolicy",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnOriginRequestPolicy) SetOriginRequestPolicyConfig(val interface{}) {
	_jsii_.Set(
		j,
		"originRequestPolicyConfig",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnOriginRequestPolicy_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.CfnOriginRequestPolicy",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnOriginRequestPolicy_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.CfnOriginRequestPolicy",
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
func CfnOriginRequestPolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.CfnOriginRequestPolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnOriginRequestPolicy_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cloudfront.CfnOriginRequestPolicy",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnOriginRequestPolicy) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnOriginRequestPolicy) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnOriginRequestPolicy) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnOriginRequestPolicy) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnOriginRequestPolicy) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnOriginRequestPolicy) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnOriginRequestPolicy) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnOriginRequestPolicy) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnOriginRequestPolicy) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnOriginRequestPolicy) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnOriginRequestPolicy) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnOriginRequestPolicy) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnOriginRequestPolicy) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnOriginRequestPolicy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnOriginRequestPolicy) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// An object that determines whether any cookies in viewer requests (and if so, which cookies) are included in requests that CloudFront sends to the origin.
//
// TODO: EXAMPLE
//
type CfnOriginRequestPolicy_CookiesConfigProperty struct {
	// Determines whether cookies in viewer requests are included in requests that CloudFront sends to the origin. Valid values are:.
	//
	// - `none` – Cookies in viewer requests are not included in requests that CloudFront sends to the origin. Even when this field is set to `none` , any cookies that are listed in a `CachePolicy` *are* included in origin requests.
	// - `whitelist` – The cookies in viewer requests that are listed in the `CookieNames` type are included in requests that CloudFront sends to the origin.
	// - `all` – All cookies in viewer requests are included in requests that CloudFront sends to the origin.
	CookieBehavior *string `json:"cookieBehavior" yaml:"cookieBehavior"`
	// Contains a list of cookie names.
	Cookies *[]*string `json:"cookies" yaml:"cookies"`
}

// An object that determines whether any HTTP headers (and if so, which headers) are included in requests that CloudFront sends to the origin.
//
// TODO: EXAMPLE
//
type CfnOriginRequestPolicy_HeadersConfigProperty struct {
	// Determines whether any HTTP headers are included in requests that CloudFront sends to the origin. Valid values are:.
	//
	// - `none` – HTTP headers are not included in requests that CloudFront sends to the origin. Even when this field is set to `none` , any headers that are listed in a `CachePolicy` *are* included in origin requests.
	// - `whitelist` – The HTTP headers that are listed in the `Headers` type are included in requests that CloudFront sends to the origin.
	// - `allViewer` – All HTTP headers in viewer requests are included in requests that CloudFront sends to the origin.
	// - `allViewerAndWhitelistCloudFront` – All HTTP headers in viewer requests and the additional CloudFront headers that are listed in the `Headers` type are included in requests that CloudFront sends to the origin. The additional headers are added by CloudFront.
	HeaderBehavior *string `json:"headerBehavior" yaml:"headerBehavior"`
	// Contains a list of HTTP header names.
	Headers *[]*string `json:"headers" yaml:"headers"`
}

// An origin request policy configuration.
//
// This configuration determines the values that CloudFront includes in requests that it sends to the origin. Each request that CloudFront sends to the origin includes the following:
//
// - The request body and the URL path (without the domain name) from the viewer request.
// - The headers that CloudFront automatically includes in every origin request, including `Host` , `User-Agent` , and `X-Amz-Cf-Id` .
// - All HTTP headers, cookies, and URL query strings that are specified in the cache policy or the origin request policy. These can include items from the viewer request and, in the case of headers, additional ones that are added by CloudFront.
//
// CloudFront sends a request when it can’t find an object in its cache that matches the request. If you want to send values to the origin and also include them in the cache key, use `CachePolicy` .
//
// TODO: EXAMPLE
//
type CfnOriginRequestPolicy_OriginRequestPolicyConfigProperty struct {
	// The cookies from viewer requests to include in origin requests.
	CookiesConfig interface{} `json:"cookiesConfig" yaml:"cookiesConfig"`
	// The HTTP headers to include in origin requests.
	//
	// These can include headers from viewer requests and additional headers added by CloudFront.
	HeadersConfig interface{} `json:"headersConfig" yaml:"headersConfig"`
	// A unique name to identify the origin request policy.
	Name *string `json:"name" yaml:"name"`
	// The URL query strings from viewer requests to include in origin requests.
	QueryStringsConfig interface{} `json:"queryStringsConfig" yaml:"queryStringsConfig"`
	// A comment to describe the origin request policy.
	//
	// The comment cannot be longer than 128 characters.
	Comment *string `json:"comment" yaml:"comment"`
}

// An object that determines whether any URL query strings in viewer requests (and if so, which query strings) are included in requests that CloudFront sends to the origin.
//
// TODO: EXAMPLE
//
type CfnOriginRequestPolicy_QueryStringsConfigProperty struct {
	// Determines whether any URL query strings in viewer requests are included in requests that CloudFront sends to the origin.
	//
	// Valid values are:
	//
	// - `none` – Query strings in viewer requests are not included in requests that CloudFront sends to the origin. Even when this field is set to `none` , any query strings that are listed in a `CachePolicy` *are* included in origin requests.
	// - `whitelist` – The query strings in viewer requests that are listed in the `QueryStringNames` type are included in requests that CloudFront sends to the origin.
	// - `all` – All query strings in viewer requests are included in requests that CloudFront sends to the origin.
	QueryStringBehavior *string `json:"queryStringBehavior" yaml:"queryStringBehavior"`
	// Contains a list of query string names.
	QueryStrings *[]*string `json:"queryStrings" yaml:"queryStrings"`
}

// Properties for defining a `CfnOriginRequestPolicy`.
//
// TODO: EXAMPLE
//
type CfnOriginRequestPolicyProps struct {
	// The origin request policy configuration.
	OriginRequestPolicyConfig interface{} `json:"originRequestPolicyConfig" yaml:"originRequestPolicyConfig"`
}

// A CloudFormation `AWS::CloudFront::PublicKey`.
//
// A public key that you can use with [signed URLs and signed cookies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/PrivateContent.html) , or with [field-level encryption](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/field-level-encryption.html) .
//
// TODO: EXAMPLE
//
type CfnPublicKey interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrCreatedTime() *string
	AttrId() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	LogicalId() *string
	Node() constructs.Node
	PublicKeyConfig() interface{}
	SetPublicKeyConfig(val interface{})
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

// The jsii proxy struct for CfnPublicKey
type jsiiProxy_CfnPublicKey struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnPublicKey) AttrCreatedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCreatedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublicKey) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublicKey) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublicKey) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublicKey) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublicKey) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublicKey) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublicKey) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublicKey) PublicKeyConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicKeyConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublicKey) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublicKey) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublicKey) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::CloudFront::PublicKey`.
func NewCfnPublicKey(scope constructs.Construct, id *string, props *CfnPublicKeyProps) CfnPublicKey {
	_init_.Initialize()

	j := jsiiProxy_CfnPublicKey{}

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront.CfnPublicKey",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::CloudFront::PublicKey`.
func NewCfnPublicKey_Override(c CfnPublicKey, scope constructs.Construct, id *string, props *CfnPublicKeyProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront.CfnPublicKey",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnPublicKey) SetPublicKeyConfig(val interface{}) {
	_jsii_.Set(
		j,
		"publicKeyConfig",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnPublicKey_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.CfnPublicKey",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnPublicKey_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.CfnPublicKey",
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
func CfnPublicKey_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.CfnPublicKey",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPublicKey_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cloudfront.CfnPublicKey",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnPublicKey) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnPublicKey) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnPublicKey) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnPublicKey) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnPublicKey) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnPublicKey) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnPublicKey) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnPublicKey) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnPublicKey) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnPublicKey) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnPublicKey) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnPublicKey) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnPublicKey) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnPublicKey) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPublicKey) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Configuration information about a public key that you can use with [signed URLs and signed cookies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/PrivateContent.html) , or with [field-level encryption](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/field-level-encryption.html) .
//
// TODO: EXAMPLE
//
type CfnPublicKey_PublicKeyConfigProperty struct {
	// A string included in the request to help make sure that the request can’t be replayed.
	CallerReference *string `json:"callerReference" yaml:"callerReference"`
	// The public key that you can use with [signed URLs and signed cookies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/PrivateContent.html) , or with [field-level encryption](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/field-level-encryption.html) .
	EncodedKey *string `json:"encodedKey" yaml:"encodedKey"`
	// A name to help identify the public key.
	Name *string `json:"name" yaml:"name"`
	// A comment to describe the public key.
	//
	// The comment cannot be longer than 128 characters.
	Comment *string `json:"comment" yaml:"comment"`
}

// Properties for defining a `CfnPublicKey`.
//
// TODO: EXAMPLE
//
type CfnPublicKeyProps struct {
	// Configuration information about a public key that you can use with [signed URLs and signed cookies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/PrivateContent.html) , or with [field-level encryption](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/field-level-encryption.html) .
	PublicKeyConfig interface{} `json:"publicKeyConfig" yaml:"publicKeyConfig"`
}

// A CloudFormation `AWS::CloudFront::RealtimeLogConfig`.
//
// A real-time log configuration.
//
// TODO: EXAMPLE
//
type CfnRealtimeLogConfig interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrArn() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	EndPoints() interface{}
	SetEndPoints(val interface{})
	Fields() *[]*string
	SetFields(val *[]*string)
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() constructs.Node
	Ref() *string
	SamplingRate() *float64
	SetSamplingRate(val *float64)
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

// The jsii proxy struct for CfnRealtimeLogConfig
type jsiiProxy_CfnRealtimeLogConfig struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnRealtimeLogConfig) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRealtimeLogConfig) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRealtimeLogConfig) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRealtimeLogConfig) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRealtimeLogConfig) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRealtimeLogConfig) EndPoints() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"endPoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRealtimeLogConfig) Fields() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"fields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRealtimeLogConfig) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRealtimeLogConfig) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRealtimeLogConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRealtimeLogConfig) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRealtimeLogConfig) SamplingRate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"samplingRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRealtimeLogConfig) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRealtimeLogConfig) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::CloudFront::RealtimeLogConfig`.
func NewCfnRealtimeLogConfig(scope constructs.Construct, id *string, props *CfnRealtimeLogConfigProps) CfnRealtimeLogConfig {
	_init_.Initialize()

	j := jsiiProxy_CfnRealtimeLogConfig{}

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront.CfnRealtimeLogConfig",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::CloudFront::RealtimeLogConfig`.
func NewCfnRealtimeLogConfig_Override(c CfnRealtimeLogConfig, scope constructs.Construct, id *string, props *CfnRealtimeLogConfigProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront.CfnRealtimeLogConfig",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnRealtimeLogConfig) SetEndPoints(val interface{}) {
	_jsii_.Set(
		j,
		"endPoints",
		val,
	)
}

func (j *jsiiProxy_CfnRealtimeLogConfig) SetFields(val *[]*string) {
	_jsii_.Set(
		j,
		"fields",
		val,
	)
}

func (j *jsiiProxy_CfnRealtimeLogConfig) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnRealtimeLogConfig) SetSamplingRate(val *float64) {
	_jsii_.Set(
		j,
		"samplingRate",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnRealtimeLogConfig_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.CfnRealtimeLogConfig",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnRealtimeLogConfig_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.CfnRealtimeLogConfig",
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
func CfnRealtimeLogConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.CfnRealtimeLogConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnRealtimeLogConfig_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cloudfront.CfnRealtimeLogConfig",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnRealtimeLogConfig) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnRealtimeLogConfig) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnRealtimeLogConfig) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnRealtimeLogConfig) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnRealtimeLogConfig) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnRealtimeLogConfig) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnRealtimeLogConfig) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnRealtimeLogConfig) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnRealtimeLogConfig) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnRealtimeLogConfig) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnRealtimeLogConfig) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnRealtimeLogConfig) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnRealtimeLogConfig) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnRealtimeLogConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnRealtimeLogConfig) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Contains information about the Amazon Kinesis data stream where you are sending real-time log data in a real-time log configuration.
//
// TODO: EXAMPLE
//
type CfnRealtimeLogConfig_EndPointProperty struct {
	// Contains information about the Amazon Kinesis data stream where you are sending real-time log data.
	KinesisStreamConfig interface{} `json:"kinesisStreamConfig" yaml:"kinesisStreamConfig"`
	// The type of data stream where you are sending real-time log data.
	//
	// The only valid value is `Kinesis` .
	StreamType *string `json:"streamType" yaml:"streamType"`
}

// Contains information about the Amazon Kinesis data stream where you are sending real-time log data.
//
// TODO: EXAMPLE
//
type CfnRealtimeLogConfig_KinesisStreamConfigProperty struct {
	// The Amazon Resource Name (ARN) of an AWS Identity and Access Management (IAM) role that CloudFront can use to send real-time log data to your Kinesis data stream.
	//
	// For more information the IAM role, see [Real-time log configuration IAM role](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/real-time-logs.html#understand-real-time-log-config-iam-role) in the *Amazon CloudFront Developer Guide* .
	RoleArn *string `json:"roleArn" yaml:"roleArn"`
	// The Amazon Resource Name (ARN) of the Kinesis data stream where you are sending real-time log data.
	StreamArn *string `json:"streamArn" yaml:"streamArn"`
}

// Properties for defining a `CfnRealtimeLogConfig`.
//
// TODO: EXAMPLE
//
type CfnRealtimeLogConfigProps struct {
	// Contains information about the Amazon Kinesis data stream where you are sending real-time log data for this real-time log configuration.
	EndPoints interface{} `json:"endPoints" yaml:"endPoints"`
	// A list of fields that are included in each real-time log record.
	//
	// In an API response, the fields are provided in the same order in which they are sent to the Amazon Kinesis data stream.
	//
	// For more information about fields, see [Real-time log configuration fields](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/real-time-logs.html#understand-real-time-log-config-fields) in the *Amazon CloudFront Developer Guide* .
	Fields *[]*string `json:"fields" yaml:"fields"`
	// The unique name of this real-time log configuration.
	Name *string `json:"name" yaml:"name"`
	// The sampling rate for this real-time log configuration.
	//
	// The sampling rate determines the percentage of viewer requests that are represented in the real-time log data. The sampling rate is an integer between 1 and 100, inclusive.
	SamplingRate *float64 `json:"samplingRate" yaml:"samplingRate"`
}

// A CloudFormation `AWS::CloudFront::ResponseHeadersPolicy`.
//
// A response headers policy.
//
// A response headers policy contains information about a set of HTTP response headers and their values.
//
// After you create a response headers policy, you can use its ID to attach it to one or more cache behaviors in a CloudFront distribution. When it’s attached to a cache behavior, CloudFront adds the headers in the policy to HTTP responses that it sends for requests that match the cache behavior.
//
// For more information, see [Adding HTTP headers to CloudFront responses](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/adding-response-headers.html) in the *Amazon CloudFront Developer Guide* .
//
// TODO: EXAMPLE
//
type CfnResponseHeadersPolicy interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrId() *string
	AttrLastModifiedTime() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	LogicalId() *string
	Node() constructs.Node
	Ref() *string
	ResponseHeadersPolicyConfig() interface{}
	SetResponseHeadersPolicyConfig(val interface{})
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

// The jsii proxy struct for CfnResponseHeadersPolicy
type jsiiProxy_CfnResponseHeadersPolicy struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnResponseHeadersPolicy) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResponseHeadersPolicy) AttrLastModifiedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLastModifiedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResponseHeadersPolicy) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResponseHeadersPolicy) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResponseHeadersPolicy) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResponseHeadersPolicy) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResponseHeadersPolicy) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResponseHeadersPolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResponseHeadersPolicy) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResponseHeadersPolicy) ResponseHeadersPolicyConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"responseHeadersPolicyConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResponseHeadersPolicy) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResponseHeadersPolicy) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::CloudFront::ResponseHeadersPolicy`.
func NewCfnResponseHeadersPolicy(scope constructs.Construct, id *string, props *CfnResponseHeadersPolicyProps) CfnResponseHeadersPolicy {
	_init_.Initialize()

	j := jsiiProxy_CfnResponseHeadersPolicy{}

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront.CfnResponseHeadersPolicy",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::CloudFront::ResponseHeadersPolicy`.
func NewCfnResponseHeadersPolicy_Override(c CfnResponseHeadersPolicy, scope constructs.Construct, id *string, props *CfnResponseHeadersPolicyProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront.CfnResponseHeadersPolicy",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnResponseHeadersPolicy) SetResponseHeadersPolicyConfig(val interface{}) {
	_jsii_.Set(
		j,
		"responseHeadersPolicyConfig",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnResponseHeadersPolicy_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.CfnResponseHeadersPolicy",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnResponseHeadersPolicy_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.CfnResponseHeadersPolicy",
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
func CfnResponseHeadersPolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.CfnResponseHeadersPolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnResponseHeadersPolicy_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cloudfront.CfnResponseHeadersPolicy",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnResponseHeadersPolicy) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnResponseHeadersPolicy) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnResponseHeadersPolicy) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnResponseHeadersPolicy) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnResponseHeadersPolicy) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnResponseHeadersPolicy) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnResponseHeadersPolicy) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnResponseHeadersPolicy) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnResponseHeadersPolicy) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnResponseHeadersPolicy) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnResponseHeadersPolicy) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnResponseHeadersPolicy) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnResponseHeadersPolicy) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnResponseHeadersPolicy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResponseHeadersPolicy) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// A list of HTTP header names that CloudFront includes as values for the `Access-Control-Allow-Headers` HTTP response header.
//
// For more information about the `Access-Control-Allow-Headers` HTTP response header, see [Access-Control-Allow-Headers](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Headers) in the MDN Web Docs.
//
// TODO: EXAMPLE
//
type CfnResponseHeadersPolicy_AccessControlAllowHeadersProperty struct {
	// The list of HTTP header names.
	//
	// You can specify `*` to allow all headers.
	Items *[]*string `json:"items" yaml:"items"`
}

// A list of HTTP methods that CloudFront includes as values for the `Access-Control-Allow-Methods` HTTP response header.
//
// For more information about the `Access-Control-Allow-Methods` HTTP response header, see [Access-Control-Allow-Methods](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Methods) in the MDN Web Docs.
//
// TODO: EXAMPLE
//
type CfnResponseHeadersPolicy_AccessControlAllowMethodsProperty struct {
	// The list of HTTP methods. Valid values are:.
	//
	// - `GET`
	// - `DELETE`
	// - `HEAD`
	// - `OPTIONS`
	// - `PATCH`
	// - `POST`
	// - `PUT`
	// - `ALL`
	//
	// `ALL` is a special value that includes all of the listed HTTP methods.
	Items *[]*string `json:"items" yaml:"items"`
}

// A list of origins (domain names) that CloudFront can use as the value for the `Access-Control-Allow-Origin` HTTP response header.
//
// For more information about the `Access-Control-Allow-Origin` HTTP response header, see [Access-Control-Allow-Origin](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Origin) in the MDN Web Docs.
//
// TODO: EXAMPLE
//
type CfnResponseHeadersPolicy_AccessControlAllowOriginsProperty struct {
	// The list of origins (domain names).
	//
	// You can specify `*` to allow all origins.
	Items *[]*string `json:"items" yaml:"items"`
}

// A list of HTTP headers that CloudFront includes as values for the `Access-Control-Expose-Headers` HTTP response header.
//
// For more information about the `Access-Control-Expose-Headers` HTTP response header, see [Access-Control-Expose-Headers](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Expose-Headers) in the MDN Web Docs.
//
// TODO: EXAMPLE
//
type CfnResponseHeadersPolicy_AccessControlExposeHeadersProperty struct {
	// The list of HTTP headers.
	//
	// You can specify `*` to expose all headers.
	Items *[]*string `json:"items" yaml:"items"`
}

// The policy directives and their values that CloudFront includes as values for the `Content-Security-Policy` HTTP response header.
//
// For more information about the `Content-Security-Policy` HTTP response header, see [Content-Security-Policy](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Security-Policy) in the MDN Web Docs.
//
// TODO: EXAMPLE
//
type CfnResponseHeadersPolicy_ContentSecurityPolicyProperty struct {
	// The policy directives and their values that CloudFront includes as values for the `Content-Security-Policy` HTTP response header.
	ContentSecurityPolicy *string `json:"contentSecurityPolicy" yaml:"contentSecurityPolicy"`
	// A Boolean that determines whether CloudFront overrides the `Content-Security-Policy` HTTP response header received from the origin with the one specified in this response headers policy.
	Override interface{} `json:"override" yaml:"override"`
}

// Determines whether CloudFront includes the `X-Content-Type-Options` HTTP response header with its value set to `nosniff` .
//
// For more information about the `X-Content-Type-Options` HTTP response header, see [X-Content-Type-Options](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Content-Type-Options) in the MDN Web Docs.
//
// TODO: EXAMPLE
//
type CfnResponseHeadersPolicy_ContentTypeOptionsProperty struct {
	// A Boolean that determines whether CloudFront overrides the `X-Content-Type-Options` HTTP response header received from the origin with the one specified in this response headers policy.
	Override interface{} `json:"override" yaml:"override"`
}

// A configuration for a set of HTTP response headers that are used for cross-origin resource sharing (CORS).
//
// CloudFront adds these headers to HTTP responses that it sends for CORS requests that match a cache behavior associated with this response headers policy.
//
// For more information about CORS, see [Cross-Origin Resource Sharing (CORS)](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS) in the MDN Web Docs.
//
// TODO: EXAMPLE
//
type CfnResponseHeadersPolicy_CorsConfigProperty struct {
	// A Boolean that CloudFront uses as the value for the `Access-Control-Allow-Credentials` HTTP response header.
	//
	// For more information about the `Access-Control-Allow-Credentials` HTTP response header, see [Access-Control-Allow-Credentials](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Credentials) in the MDN Web Docs.
	AccessControlAllowCredentials interface{} `json:"accessControlAllowCredentials" yaml:"accessControlAllowCredentials"`
	// A list of HTTP header names that CloudFront includes as values for the `Access-Control-Allow-Headers` HTTP response header.
	//
	// For more information about the `Access-Control-Allow-Headers` HTTP response header, see [Access-Control-Allow-Headers](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Headers) in the MDN Web Docs.
	AccessControlAllowHeaders interface{} `json:"accessControlAllowHeaders" yaml:"accessControlAllowHeaders"`
	// A list of HTTP methods that CloudFront includes as values for the `Access-Control-Allow-Methods` HTTP response header.
	//
	// For more information about the `Access-Control-Allow-Methods` HTTP response header, see [Access-Control-Allow-Methods](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Methods) in the MDN Web Docs.
	AccessControlAllowMethods interface{} `json:"accessControlAllowMethods" yaml:"accessControlAllowMethods"`
	// A list of origins (domain names) that CloudFront can use as the value for the `Access-Control-Allow-Origin` HTTP response header.
	//
	// For more information about the `Access-Control-Allow-Origin` HTTP response header, see [Access-Control-Allow-Origin](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Origin) in the MDN Web Docs.
	AccessControlAllowOrigins interface{} `json:"accessControlAllowOrigins" yaml:"accessControlAllowOrigins"`
	// A Boolean that determines whether CloudFront overrides HTTP response headers received from the origin with the ones specified in this response headers policy.
	OriginOverride interface{} `json:"originOverride" yaml:"originOverride"`
	// A list of HTTP headers that CloudFront includes as values for the `Access-Control-Expose-Headers` HTTP response header.
	//
	// For more information about the `Access-Control-Expose-Headers` HTTP response header, see [Access-Control-Expose-Headers](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Expose-Headers) in the MDN Web Docs.
	AccessControlExposeHeaders interface{} `json:"accessControlExposeHeaders" yaml:"accessControlExposeHeaders"`
	// A number that CloudFront uses as the value for the `Access-Control-Max-Age` HTTP response header.
	//
	// For more information about the `Access-Control-Max-Age` HTTP response header, see [Access-Control-Max-Age](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Max-Age) in the MDN Web Docs.
	AccessControlMaxAgeSec *float64 `json:"accessControlMaxAgeSec" yaml:"accessControlMaxAgeSec"`
}

// An HTTP response header name and its value.
//
// CloudFront includes this header in HTTP responses that it sends for requests that match a cache behavior that’s associated with this response headers policy.
//
// TODO: EXAMPLE
//
type CfnResponseHeadersPolicy_CustomHeaderProperty struct {
	// The HTTP response header name.
	Header *string `json:"header" yaml:"header"`
	// A Boolean that determines whether CloudFront overrides a response header with the same name received from the origin with the header specified here.
	Override interface{} `json:"override" yaml:"override"`
	// The value for the HTTP response header.
	Value *string `json:"value" yaml:"value"`
}

// A list of HTTP response header names and their values.
//
// CloudFront includes these headers in HTTP responses that it sends for requests that match a cache behavior that’s associated with this response headers policy.
//
// TODO: EXAMPLE
//
type CfnResponseHeadersPolicy_CustomHeadersConfigProperty struct {
	// The list of HTTP response headers and their values.
	Items interface{} `json:"items" yaml:"items"`
}

// Determines whether CloudFront includes the `X-Frame-Options` HTTP response header and the header’s value.
//
// For more information about the `X-Frame-Options` HTTP response header, see [X-Frame-Options](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Frame-Options) in the MDN Web Docs.
//
// TODO: EXAMPLE
//
type CfnResponseHeadersPolicy_FrameOptionsProperty struct {
	// The value of the `X-Frame-Options` HTTP response header. Valid values are `DENY` and `SAMEORIGIN` .
	//
	// For more information about these values, see [X-Frame-Options](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Frame-Options) in the MDN Web Docs.
	FrameOption *string `json:"frameOption" yaml:"frameOption"`
	// A Boolean that determines whether CloudFront overrides the `X-Frame-Options` HTTP response header received from the origin with the one specified in this response headers policy.
	Override interface{} `json:"override" yaml:"override"`
}

// Determines whether CloudFront includes the `Referrer-Policy` HTTP response header and the header’s value.
//
// For more information about the `Referrer-Policy` HTTP response header, see [Referrer-Policy](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Referrer-Policy) in the MDN Web Docs.
//
// TODO: EXAMPLE
//
type CfnResponseHeadersPolicy_ReferrerPolicyProperty struct {
	// A Boolean that determines whether CloudFront overrides the `Referrer-Policy` HTTP response header received from the origin with the one specified in this response headers policy.
	Override interface{} `json:"override" yaml:"override"`
	// The value of the `Referrer-Policy` HTTP response header. Valid values are:.
	//
	// - `no-referrer`
	// - `no-referrer-when-downgrade`
	// - `origin`
	// - `origin-when-cross-origin`
	// - `same-origin`
	// - `strict-origin`
	// - `strict-origin-when-cross-origin`
	// - `unsafe-url`
	//
	// For more information about these values, see [Referrer-Policy](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Referrer-Policy) in the MDN Web Docs.
	ReferrerPolicy *string `json:"referrerPolicy" yaml:"referrerPolicy"`
}

// A response headers policy configuration.
//
// A response headers policy configuration contains metadata about the response headers policy, and configurations for sets of HTTP response headers and their values. CloudFront adds the headers in the policy to HTTP responses that it sends for requests that match a cache behavior associated with the policy.
//
// TODO: EXAMPLE
//
type CfnResponseHeadersPolicy_ResponseHeadersPolicyConfigProperty struct {
	// A name to identify the response headers policy.
	//
	// The name must be unique for response headers policies in this AWS account .
	Name *string `json:"name" yaml:"name"`
	// A comment to describe the response headers policy.
	//
	// The comment cannot be longer than 128 characters.
	Comment *string `json:"comment" yaml:"comment"`
	// A configuration for a set of HTTP response headers that are used for cross-origin resource sharing (CORS).
	CorsConfig interface{} `json:"corsConfig" yaml:"corsConfig"`
	// A configuration for a set of custom HTTP response headers.
	CustomHeadersConfig interface{} `json:"customHeadersConfig" yaml:"customHeadersConfig"`
	// A configuration for a set of security-related HTTP response headers.
	SecurityHeadersConfig interface{} `json:"securityHeadersConfig" yaml:"securityHeadersConfig"`
}

// A configuration for a set of security-related HTTP response headers.
//
// CloudFront adds these headers to HTTP responses that it sends for requests that match a cache behavior associated with this response headers policy.
//
// TODO: EXAMPLE
//
type CfnResponseHeadersPolicy_SecurityHeadersConfigProperty struct {
	// The policy directives and their values that CloudFront includes as values for the `Content-Security-Policy` HTTP response header.
	//
	// For more information about the `Content-Security-Policy` HTTP response header, see [Content-Security-Policy](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Security-Policy) in the MDN Web Docs.
	ContentSecurityPolicy interface{} `json:"contentSecurityPolicy" yaml:"contentSecurityPolicy"`
	// Determines whether CloudFront includes the `X-Content-Type-Options` HTTP response header with its value set to `nosniff` .
	//
	// For more information about the `X-Content-Type-Options` HTTP response header, see [X-Content-Type-Options](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Content-Type-Options) in the MDN Web Docs.
	ContentTypeOptions interface{} `json:"contentTypeOptions" yaml:"contentTypeOptions"`
	// Determines whether CloudFront includes the `X-Frame-Options` HTTP response header and the header’s value.
	//
	// For more information about the `X-Frame-Options` HTTP response header, see [X-Frame-Options](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Frame-Options) in the MDN Web Docs.
	FrameOptions interface{} `json:"frameOptions" yaml:"frameOptions"`
	// Determines whether CloudFront includes the `Referrer-Policy` HTTP response header and the header’s value.
	//
	// For more information about the `Referrer-Policy` HTTP response header, see [Referrer-Policy](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Referrer-Policy) in the MDN Web Docs.
	ReferrerPolicy interface{} `json:"referrerPolicy" yaml:"referrerPolicy"`
	// Determines whether CloudFront includes the `Strict-Transport-Security` HTTP response header and the header’s value.
	//
	// For more information about the `Strict-Transport-Security` HTTP response header, see [Strict-Transport-Security](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Strict-Transport-Security) in the MDN Web Docs.
	StrictTransportSecurity interface{} `json:"strictTransportSecurity" yaml:"strictTransportSecurity"`
	// Determines whether CloudFront includes the `X-XSS-Protection` HTTP response header and the header’s value.
	//
	// For more information about the `X-XSS-Protection` HTTP response header, see [X-XSS-Protection](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-XSS-Protection) in the MDN Web Docs.
	XssProtection interface{} `json:"xssProtection" yaml:"xssProtection"`
}

// Determines whether CloudFront includes the `Strict-Transport-Security` HTTP response header and the header’s value.
//
// For more information about the `Strict-Transport-Security` HTTP response header, see [Strict-Transport-Security](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Strict-Transport-Security) in the MDN Web Docs.
//
// TODO: EXAMPLE
//
type CfnResponseHeadersPolicy_StrictTransportSecurityProperty struct {
	// A number that CloudFront uses as the value for the `max-age` directive in the `Strict-Transport-Security` HTTP response header.
	AccessControlMaxAgeSec *float64 `json:"accessControlMaxAgeSec" yaml:"accessControlMaxAgeSec"`
	// A Boolean that determines whether CloudFront overrides the `Strict-Transport-Security` HTTP response header received from the origin with the one specified in this response headers policy.
	Override interface{} `json:"override" yaml:"override"`
	// A Boolean that determines whether CloudFront includes the `includeSubDomains` directive in the `Strict-Transport-Security` HTTP response header.
	IncludeSubdomains interface{} `json:"includeSubdomains" yaml:"includeSubdomains"`
	// A Boolean that determines whether CloudFront includes the `preload` directive in the `Strict-Transport-Security` HTTP response header.
	Preload interface{} `json:"preload" yaml:"preload"`
}

// Determines whether CloudFront includes the `X-XSS-Protection` HTTP response header and the header’s value.
//
// For more information about the `X-XSS-Protection` HTTP response header, see [X-XSS-Protection](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-XSS-Protection) in the MDN Web Docs.
//
// TODO: EXAMPLE
//
type CfnResponseHeadersPolicy_XSSProtectionProperty struct {
	// A Boolean that determines whether CloudFront overrides the `X-XSS-Protection` HTTP response header received from the origin with the one specified in this response headers policy.
	Override interface{} `json:"override" yaml:"override"`
	// A Boolean that determines the value of the `X-XSS-Protection` HTTP response header.
	//
	// When this setting is `true` , the value of the `X-XSS-Protection` header is `1` . When this setting is `false` , the value of the `X-XSS-Protection` header is `0` .
	//
	// For more information about these settings, see [X-XSS-Protection](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-XSS-Protection) in the MDN Web Docs.
	Protection interface{} `json:"protection" yaml:"protection"`
	// A Boolean that determines whether CloudFront includes the `mode=block` directive in the `X-XSS-Protection` header.
	//
	// For more information about this directive, see [X-XSS-Protection](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-XSS-Protection) in the MDN Web Docs.
	ModeBlock interface{} `json:"modeBlock" yaml:"modeBlock"`
	// A reporting URI, which CloudFront uses as the value of the `report` directive in the `X-XSS-Protection` header.
	//
	// You cannot specify a `ReportUri` when `ModeBlock` is `true` .
	//
	// For more information about using a reporting URL, see [X-XSS-Protection](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-XSS-Protection) in the MDN Web Docs.
	ReportUri *string `json:"reportUri" yaml:"reportUri"`
}

// Properties for defining a `CfnResponseHeadersPolicy`.
//
// TODO: EXAMPLE
//
type CfnResponseHeadersPolicyProps struct {
	// A response headers policy configuration.
	//
	// A response headers policy contains information about a set of HTTP response headers and their values. CloudFront adds the headers in the policy to HTTP responses that it sends for requests that match a cache behavior that’s associated with the policy.
	ResponseHeadersPolicyConfig interface{} `json:"responseHeadersPolicyConfig" yaml:"responseHeadersPolicyConfig"`
}

// A CloudFormation `AWS::CloudFront::StreamingDistribution`.
//
// This resource is deprecated. Amazon CloudFront is deprecating real-time messaging protocol (RTMP) distributions on December 31, 2020. For more information, [read the announcement](https://docs.aws.amazon.com/ann.jspa?annID=7356) on the Amazon CloudFront discussion forum.
//
// TODO: EXAMPLE
//
type CfnStreamingDistribution interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrDomainName() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	LogicalId() *string
	Node() constructs.Node
	Ref() *string
	Stack() awscdk.Stack
	StreamingDistributionConfig() interface{}
	SetStreamingDistributionConfig(val interface{})
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

// The jsii proxy struct for CfnStreamingDistribution
type jsiiProxy_CfnStreamingDistribution struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnStreamingDistribution) AttrDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStreamingDistribution) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStreamingDistribution) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStreamingDistribution) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStreamingDistribution) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStreamingDistribution) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStreamingDistribution) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStreamingDistribution) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStreamingDistribution) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStreamingDistribution) StreamingDistributionConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"streamingDistributionConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStreamingDistribution) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStreamingDistribution) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::CloudFront::StreamingDistribution`.
func NewCfnStreamingDistribution(scope constructs.Construct, id *string, props *CfnStreamingDistributionProps) CfnStreamingDistribution {
	_init_.Initialize()

	j := jsiiProxy_CfnStreamingDistribution{}

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront.CfnStreamingDistribution",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::CloudFront::StreamingDistribution`.
func NewCfnStreamingDistribution_Override(c CfnStreamingDistribution, scope constructs.Construct, id *string, props *CfnStreamingDistributionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront.CfnStreamingDistribution",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnStreamingDistribution) SetStreamingDistributionConfig(val interface{}) {
	_jsii_.Set(
		j,
		"streamingDistributionConfig",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnStreamingDistribution_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.CfnStreamingDistribution",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnStreamingDistribution_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.CfnStreamingDistribution",
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
func CfnStreamingDistribution_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.CfnStreamingDistribution",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnStreamingDistribution_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cloudfront.CfnStreamingDistribution",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnStreamingDistribution) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnStreamingDistribution) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnStreamingDistribution) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnStreamingDistribution) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnStreamingDistribution) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnStreamingDistribution) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnStreamingDistribution) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnStreamingDistribution) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnStreamingDistribution) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnStreamingDistribution) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnStreamingDistribution) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnStreamingDistribution) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnStreamingDistribution) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnStreamingDistribution) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStreamingDistribution) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// A complex type that controls whether access logs are written for the streaming distribution.
//
// TODO: EXAMPLE
//
type CfnStreamingDistribution_LoggingProperty struct {
	// The Amazon S3 bucket to store the access logs in, for example, `myawslogbucket.s3.amazonaws.com` .
	Bucket *string `json:"bucket" yaml:"bucket"`
	// Specifies whether you want CloudFront to save access logs to an Amazon S3 bucket.
	//
	// If you don't want to enable logging when you create a streaming distribution or if you want to disable logging for an existing streaming distribution, specify `false` for `Enabled` , and specify `empty Bucket` and `Prefix` elements. If you specify `false` for `Enabled` but you specify values for `Bucket` and `Prefix` , the values are automatically deleted.
	Enabled interface{} `json:"enabled" yaml:"enabled"`
	// An optional string that you want CloudFront to prefix to the access log filenames for this streaming distribution, for example, `myprefix/` .
	//
	// If you want to enable logging, but you don't want to specify a prefix, you still must include an empty `Prefix` element in the `Logging` element.
	Prefix *string `json:"prefix" yaml:"prefix"`
}

// A complex type that contains information about the Amazon S3 bucket from which you want CloudFront to get your media files for distribution.
//
// TODO: EXAMPLE
//
type CfnStreamingDistribution_S3OriginProperty struct {
	// The DNS name of the Amazon S3 origin.
	DomainName *string `json:"domainName" yaml:"domainName"`
	// The CloudFront origin access identity to associate with the distribution.
	//
	// Use an origin access identity to configure the distribution so that end users can only access objects in an Amazon S3 bucket through CloudFront.
	//
	// If you want end users to be able to access objects using either the CloudFront URL or the Amazon S3 URL, specify an empty `OriginAccessIdentity` element.
	//
	// To delete the origin access identity from an existing distribution, update the distribution configuration and include an empty `OriginAccessIdentity` element.
	//
	// To replace the origin access identity, update the distribution configuration and specify the new origin access identity.
	//
	// For more information, see [Using an Origin Access Identity to Restrict Access to Your Amazon S3 Content](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/private-content-restricting-access-to-s3.html) in the *Amazon CloudFront Developer Guide* .
	OriginAccessIdentity *string `json:"originAccessIdentity" yaml:"originAccessIdentity"`
}

// The RTMP distribution's configuration information.
//
// TODO: EXAMPLE
//
type CfnStreamingDistribution_StreamingDistributionConfigProperty struct {
	// Any comments you want to include about the streaming distribution.
	Comment *string `json:"comment" yaml:"comment"`
	// Whether the streaming distribution is enabled to accept user requests for content.
	Enabled interface{} `json:"enabled" yaml:"enabled"`
	// A complex type that contains information about the Amazon S3 bucket from which you want CloudFront to get your media files for distribution.
	S3Origin interface{} `json:"s3Origin" yaml:"s3Origin"`
	// A complex type that specifies any AWS accounts that you want to permit to create signed URLs for private content.
	//
	// If you want the distribution to use signed URLs, include this element; if you want the distribution to use public URLs, remove this element. For more information, see [Serving Private Content through CloudFront](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/PrivateContent.html) in the *Amazon CloudFront Developer Guide* .
	TrustedSigners interface{} `json:"trustedSigners" yaml:"trustedSigners"`
	// A complex type that contains information about CNAMEs (alternate domain names), if any, for this streaming distribution.
	Aliases *[]*string `json:"aliases" yaml:"aliases"`
	// A complex type that controls whether access logs are written for the streaming distribution.
	Logging interface{} `json:"logging" yaml:"logging"`
	// A complex type that contains information about price class for this streaming distribution.
	PriceClass *string `json:"priceClass" yaml:"priceClass"`
}

// A list of AWS accounts whose public keys CloudFront can use to verify the signatures of signed URLs and signed cookies.
//
// TODO: EXAMPLE
//
type CfnStreamingDistribution_TrustedSignersProperty struct {
	// This field is `true` if any of the AWS accounts have public keys that CloudFront can use to verify the signatures of signed URLs and signed cookies.
	//
	// If not, this field is `false` .
	Enabled interface{} `json:"enabled" yaml:"enabled"`
	// An AWS account number that contains active CloudFront key pairs that CloudFront can use to verify the signatures of signed URLs and signed cookies.
	//
	// If the AWS account that owns the key pairs is the same account that owns the CloudFront distribution, the value of this field is `self` .
	AwsAccountNumbers *[]*string `json:"awsAccountNumbers" yaml:"awsAccountNumbers"`
}

// Properties for defining a `CfnStreamingDistribution`.
//
// TODO: EXAMPLE
//
type CfnStreamingDistributionProps struct {
	// The current configuration information for the RTMP distribution.
	StreamingDistributionConfig interface{} `json:"streamingDistributionConfig" yaml:"streamingDistributionConfig"`
	// A complex type that contains zero or more `Tag` elements.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
}

// Enums for the methods CloudFront can cache.
type CloudFrontAllowedCachedMethods string

const (
	CloudFrontAllowedCachedMethods_GET_HEAD CloudFrontAllowedCachedMethods = "GET_HEAD"
	CloudFrontAllowedCachedMethods_GET_HEAD_OPTIONS CloudFrontAllowedCachedMethods = "GET_HEAD_OPTIONS"
)

// An enum for the supported methods to a CloudFront distribution.
type CloudFrontAllowedMethods string

const (
	CloudFrontAllowedMethods_GET_HEAD CloudFrontAllowedMethods = "GET_HEAD"
	CloudFrontAllowedMethods_GET_HEAD_OPTIONS CloudFrontAllowedMethods = "GET_HEAD_OPTIONS"
	CloudFrontAllowedMethods_ALL CloudFrontAllowedMethods = "ALL"
)

// Amazon CloudFront is a global content delivery network (CDN) service that securely delivers data, videos, applications, and APIs to your viewers with low latency and high transfer speeds.
//
// CloudFront fronts user provided content and caches it at edge locations across the world.
//
// Here's how you can use this construct:
//
// ```ts
// const sourceBucket = new s3.Bucket(this, 'Bucket');
//
// const distribution = new cloudfront.CloudFrontWebDistribution(this, 'MyDistribution', {
//    originConfigs: [
//      {
//        s3OriginSource: {
//        s3BucketSource: sourceBucket,
//        },
//        behaviors : [ {isDefaultBehavior: true}],
//      },
//    ],
// });
// ```
//
// This will create a CloudFront distribution that uses your S3Bucket as it's origin.
//
// You can customize the distribution using additional properties from the CloudFrontWebDistributionProps interface.
//
// TODO: EXAMPLE
//
type CloudFrontWebDistribution interface {
	awscdk.Resource
	IDistribution
	DistributionDomainName() *string
	DistributionId() *string
	Env() *awscdk.ResourceEnvironment
	LoggingBucket() awss3.IBucket
	Node() constructs.Node
	PhysicalName() *string
	Stack() awscdk.Stack
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ToString() *string
}

// The jsii proxy struct for CloudFrontWebDistribution
type jsiiProxy_CloudFrontWebDistribution struct {
	internal.Type__awscdkResource
	jsiiProxy_IDistribution
}

func (j *jsiiProxy_CloudFrontWebDistribution) DistributionDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"distributionDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFrontWebDistribution) DistributionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"distributionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFrontWebDistribution) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFrontWebDistribution) LoggingBucket() awss3.IBucket {
	var returns awss3.IBucket
	_jsii_.Get(
		j,
		"loggingBucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFrontWebDistribution) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFrontWebDistribution) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFrontWebDistribution) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewCloudFrontWebDistribution(scope constructs.Construct, id *string, props *CloudFrontWebDistributionProps) CloudFrontWebDistribution {
	_init_.Initialize()

	j := jsiiProxy_CloudFrontWebDistribution{}

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront.CloudFrontWebDistribution",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCloudFrontWebDistribution_Override(c CloudFrontWebDistribution, scope constructs.Construct, id *string, props *CloudFrontWebDistributionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront.CloudFrontWebDistribution",
		[]interface{}{scope, id, props},
		c,
	)
}

// Creates a construct that represents an external (imported) distribution.
func CloudFrontWebDistribution_FromDistributionAttributes(scope constructs.Construct, id *string, attrs *CloudFrontWebDistributionAttributes) IDistribution {
	_init_.Initialize()

	var returns IDistribution

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.CloudFrontWebDistribution",
		"fromDistributionAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CloudFrontWebDistribution_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.CloudFrontWebDistribution",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func CloudFrontWebDistribution_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.CloudFrontWebDistribution",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (c *jsiiProxy_CloudFrontWebDistribution) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (c *jsiiProxy_CloudFrontWebDistribution) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
func (c *jsiiProxy_CloudFrontWebDistribution) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
func (c *jsiiProxy_CloudFrontWebDistribution) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (c *jsiiProxy_CloudFrontWebDistribution) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Attributes used to import a Distribution.
//
// TODO: EXAMPLE
//
type CloudFrontWebDistributionAttributes struct {
	// The distribution ID for this distribution.
	DistributionId *string `json:"distributionId" yaml:"distributionId"`
	// The generated domain name of the Distribution, such as d111111abcdef8.cloudfront.net.
	DomainName *string `json:"domainName" yaml:"domainName"`
}

// TODO: EXAMPLE
//
type CloudFrontWebDistributionProps struct {
	// The origin configurations for this distribution.
	//
	// Behaviors are a part of the origin.
	OriginConfigs *[]*SourceConfiguration `json:"originConfigs" yaml:"originConfigs"`
	// A comment for this distribution in the CloudFront console.
	Comment *string `json:"comment" yaml:"comment"`
	// The default object to serve.
	DefaultRootObject *string `json:"defaultRootObject" yaml:"defaultRootObject"`
	// Enable or disable the distribution.
	Enabled *bool `json:"enabled" yaml:"enabled"`
	// If your distribution should have IPv6 enabled.
	EnableIpV6 *bool `json:"enableIpV6" yaml:"enableIpV6"`
	// How CloudFront should handle requests that are not successful (eg PageNotFound).
	//
	// By default, CloudFront does not replace HTTP status codes in the 4xx and 5xx range
	// with custom error messages. CloudFront does not cache HTTP status codes.
	ErrorConfigurations *[]*CfnDistribution_CustomErrorResponseProperty `json:"errorConfigurations" yaml:"errorConfigurations"`
	// Controls the countries in which your content is distributed.
	GeoRestriction GeoRestriction `json:"geoRestriction" yaml:"geoRestriction"`
	// The max supported HTTP Versions.
	HttpVersion HttpVersion `json:"httpVersion" yaml:"httpVersion"`
	// Optional - if we should enable logging.
	//
	// You can pass an empty object ({}) to have us auto create a bucket for logging.
	// Omission of this property indicates no logging is to be enabled.
	LoggingConfig *LoggingConfiguration `json:"loggingConfig" yaml:"loggingConfig"`
	// The price class for the distribution (this impacts how many locations CloudFront uses for your distribution, and billing).
	PriceClass PriceClass `json:"priceClass" yaml:"priceClass"`
	// Specifies whether you want viewers to use HTTP or HTTPS to request your objects, whether you're using an alternate domain name with HTTPS, and if so, if you're using AWS Certificate Manager (ACM) or a third-party certificate authority.
	// See: https://aws.amazon.com/premiumsupport/knowledge-center/custom-ssl-certificate-cloudfront/
	//
	ViewerCertificate ViewerCertificate `json:"viewerCertificate" yaml:"viewerCertificate"`
	// The default viewer policy for incoming clients.
	ViewerProtocolPolicy ViewerProtocolPolicy `json:"viewerProtocolPolicy" yaml:"viewerProtocolPolicy"`
	// Unique identifier that specifies the AWS WAF web ACL to associate with this CloudFront distribution.
	//
	// To specify a web ACL created using the latest version of AWS WAF, use the ACL ARN, for example
	// `arn:aws:wafv2:us-east-1:123456789012:global/webacl/ExampleWebACL/473e64fd-f30b-4765-81a0-62ad96dd167a`.
	//
	// To specify a web ACL created using AWS WAF Classic, use the ACL ID, for example `473e64fd-f30b-4765-81a0-62ad96dd167a`.
	// See: https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_CreateDistribution.html#API_CreateDistribution_RequestParameters.
	//
	WebACLId *string `json:"webACLId" yaml:"webACLId"`
}

// A custom origin configuration.
//
// TODO: EXAMPLE
//
type CustomOriginConfig struct {
	// The domain name of the custom origin.
	//
	// Should not include the path - that should be in the parent SourceConfiguration
	DomainName *string `json:"domainName" yaml:"domainName"`
	// The SSL versions to use when interacting with the origin.
	AllowedOriginSSLVersions *[]OriginSslPolicy `json:"allowedOriginSSLVersions" yaml:"allowedOriginSSLVersions"`
	// The origin HTTP port.
	HttpPort *float64 `json:"httpPort" yaml:"httpPort"`
	// The origin HTTPS port.
	HttpsPort *float64 `json:"httpsPort" yaml:"httpsPort"`
	// Any additional headers to pass to the origin.
	OriginHeaders *map[string]*string `json:"originHeaders" yaml:"originHeaders"`
	// The keep alive timeout when making calls in seconds.
	OriginKeepaliveTimeout awscdk.Duration `json:"originKeepaliveTimeout" yaml:"originKeepaliveTimeout"`
	// The relative path to the origin root to use for sources.
	OriginPath *string `json:"originPath" yaml:"originPath"`
	// The protocol (http or https) policy to use when interacting with the origin.
	OriginProtocolPolicy OriginProtocolPolicy `json:"originProtocolPolicy" yaml:"originProtocolPolicy"`
	// The read timeout when calling the origin in seconds.
	OriginReadTimeout awscdk.Duration `json:"originReadTimeout" yaml:"originReadTimeout"`
	// When you enable Origin Shield in the AWS Region that has the lowest latency to your origin, you can get better network performance.
	OriginShieldRegion *string `json:"originShieldRegion" yaml:"originShieldRegion"`
}

// A CloudFront distribution with associated origin(s) and caching behavior(s).
//
// TODO: EXAMPLE
//
type Distribution interface {
	awscdk.Resource
	IDistribution
	DistributionDomainName() *string
	DistributionId() *string
	DomainName() *string
	Env() *awscdk.ResourceEnvironment
	Node() constructs.Node
	PhysicalName() *string
	Stack() awscdk.Stack
	AddBehavior(pathPattern *string, origin IOrigin, behaviorOptions *AddBehaviorOptions)
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ToString() *string
}

// The jsii proxy struct for Distribution
type jsiiProxy_Distribution struct {
	internal.Type__awscdkResource
	jsiiProxy_IDistribution
}

func (j *jsiiProxy_Distribution) DistributionDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"distributionDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Distribution) DistributionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"distributionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Distribution) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Distribution) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Distribution) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Distribution) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Distribution) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewDistribution(scope constructs.Construct, id *string, props *DistributionProps) Distribution {
	_init_.Initialize()

	j := jsiiProxy_Distribution{}

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront.Distribution",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewDistribution_Override(d Distribution, scope constructs.Construct, id *string, props *DistributionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront.Distribution",
		[]interface{}{scope, id, props},
		d,
	)
}

// Creates a Distribution construct that represents an external (imported) distribution.
func Distribution_FromDistributionAttributes(scope constructs.Construct, id *string, attrs *DistributionAttributes) IDistribution {
	_init_.Initialize()

	var returns IDistribution

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.Distribution",
		"fromDistributionAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func Distribution_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.Distribution",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func Distribution_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.Distribution",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Adds a new behavior to this distribution for the given pathPattern.
func (d *jsiiProxy_Distribution) AddBehavior(pathPattern *string, origin IOrigin, behaviorOptions *AddBehaviorOptions) {
	_jsii_.InvokeVoid(
		d,
		"addBehavior",
		[]interface{}{pathPattern, origin, behaviorOptions},
	)
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (d *jsiiProxy_Distribution) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		d,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (d *jsiiProxy_Distribution) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
func (d *jsiiProxy_Distribution) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
func (d *jsiiProxy_Distribution) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (d *jsiiProxy_Distribution) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Attributes used to import a Distribution.
//
// TODO: EXAMPLE
//
type DistributionAttributes struct {
	// The distribution ID for this distribution.
	DistributionId *string `json:"distributionId" yaml:"distributionId"`
	// The generated domain name of the Distribution, such as d111111abcdef8.cloudfront.net.
	DomainName *string `json:"domainName" yaml:"domainName"`
}

// Properties for a Distribution.
//
// TODO: EXAMPLE
//
type DistributionProps struct {
	// The default behavior for the distribution.
	DefaultBehavior *BehaviorOptions `json:"defaultBehavior" yaml:"defaultBehavior"`
	// Additional behaviors for the distribution, mapped by the pathPattern that specifies which requests to apply the behavior to.
	AdditionalBehaviors *map[string]*BehaviorOptions `json:"additionalBehaviors" yaml:"additionalBehaviors"`
	// A certificate to associate with the distribution.
	//
	// The certificate must be located in N. Virginia (us-east-1).
	Certificate awscertificatemanager.ICertificate `json:"certificate" yaml:"certificate"`
	// Any comments you want to include about the distribution.
	Comment *string `json:"comment" yaml:"comment"`
	// The object that you want CloudFront to request from your origin (for example, index.html) when a viewer requests the root URL for your distribution. If no default object is set, the request goes to the origin's root (e.g., example.com/).
	DefaultRootObject *string `json:"defaultRootObject" yaml:"defaultRootObject"`
	// Alternative domain names for this distribution.
	//
	// If you want to use your own domain name, such as www.example.com, instead of the cloudfront.net domain name,
	// you can add an alternate domain name to your distribution. If you attach a certificate to the distribution,
	// you must add (at least one of) the domain names of the certificate to this list.
	DomainNames *[]*string `json:"domainNames" yaml:"domainNames"`
	// Enable or disable the distribution.
	Enabled *bool `json:"enabled" yaml:"enabled"`
	// Whether CloudFront will respond to IPv6 DNS requests with an IPv6 address.
	//
	// If you specify false, CloudFront responds to IPv6 DNS requests with the DNS response code NOERROR and with no IP addresses.
	// This allows viewers to submit a second request, for an IPv4 address for your distribution.
	EnableIpv6 *bool `json:"enableIpv6" yaml:"enableIpv6"`
	// Enable access logging for the distribution.
	EnableLogging *bool `json:"enableLogging" yaml:"enableLogging"`
	// How CloudFront should handle requests that are not successful (e.g., PageNotFound).
	ErrorResponses *[]*ErrorResponse `json:"errorResponses" yaml:"errorResponses"`
	// Controls the countries in which your content is distributed.
	GeoRestriction GeoRestriction `json:"geoRestriction" yaml:"geoRestriction"`
	// Specify the maximum HTTP version that you want viewers to use to communicate with CloudFront.
	//
	// For viewers and CloudFront to use HTTP/2, viewers must support TLS 1.2 or later, and must support server name identification (SNI).
	HttpVersion HttpVersion `json:"httpVersion" yaml:"httpVersion"`
	// The Amazon S3 bucket to store the access logs in.
	LogBucket awss3.IBucket `json:"logBucket" yaml:"logBucket"`
	// An optional string that you want CloudFront to prefix to the access log filenames for this distribution.
	LogFilePrefix *string `json:"logFilePrefix" yaml:"logFilePrefix"`
	// Specifies whether you want CloudFront to include cookies in access logs.
	LogIncludesCookies *bool `json:"logIncludesCookies" yaml:"logIncludesCookies"`
	// The minimum version of the SSL protocol that you want CloudFront to use for HTTPS connections.
	//
	// CloudFront serves your objects only to browsers or devices that support at
	// least the SSL version that you specify.
	MinimumProtocolVersion SecurityPolicyProtocol `json:"minimumProtocolVersion" yaml:"minimumProtocolVersion"`
	// The price class that corresponds with the maximum price that you want to pay for CloudFront service.
	//
	// If you specify PriceClass_All, CloudFront responds to requests for your objects from all CloudFront edge locations.
	// If you specify a price class other than PriceClass_All, CloudFront serves your objects from the CloudFront edge location
	// that has the lowest latency among the edge locations in your price class.
	PriceClass PriceClass `json:"priceClass" yaml:"priceClass"`
	// Unique identifier that specifies the AWS WAF web ACL to associate with this CloudFront distribution.
	//
	// To specify a web ACL created using the latest version of AWS WAF, use the ACL ARN, for example
	// `arn:aws:wafv2:us-east-1:123456789012:global/webacl/ExampleWebACL/473e64fd-f30b-4765-81a0-62ad96dd167a`.
	// To specify a web ACL created using AWS WAF Classic, use the ACL ID, for example `473e64fd-f30b-4765-81a0-62ad96dd167a`.
	// See: https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_CreateDistribution.html#API_CreateDistribution_RequestParameters.
	//
	WebAclId *string `json:"webAclId" yaml:"webAclId"`
}

// Represents a Lambda function version and event type when using Lambda@Edge.
//
// The type of the {@link AddBehaviorOptions.edgeLambdas} property.
//
// TODO: EXAMPLE
//
type EdgeLambda struct {
	// The type of event in response to which should the function be invoked.
	EventType LambdaEdgeEventType `json:"eventType" yaml:"eventType"`
	// The version of the Lambda function that will be invoked.
	//
	// **Note**: it's not possible to use the '$LATEST' function version for Lambda@Edge!
	FunctionVersion awslambda.IVersion `json:"functionVersion" yaml:"functionVersion"`
	// Allows a Lambda function to have read access to the body content.
	//
	// Only valid for "request" event types (`ORIGIN_REQUEST` or `VIEWER_REQUEST`).
	// See https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/lambda-include-body-access.html
	IncludeBody *bool `json:"includeBody" yaml:"includeBody"`
}

// Options for configuring custom error responses.
//
// TODO: EXAMPLE
//
type ErrorResponse struct {
	// The HTTP status code for which you want to specify a custom error page and/or a caching duration.
	HttpStatus *float64 `json:"httpStatus" yaml:"httpStatus"`
	// The HTTP status code that you want CloudFront to return to the viewer along with the custom error page.
	//
	// If you specify a value for `responseHttpStatus`, you must also specify a value for `responsePagePath`.
	ResponseHttpStatus *float64 `json:"responseHttpStatus" yaml:"responseHttpStatus"`
	// The path to the custom error page that you want CloudFront to return to a viewer when your origin returns the `httpStatus`, for example, /4xx-errors/403-forbidden.html.
	ResponsePagePath *string `json:"responsePagePath" yaml:"responsePagePath"`
	// The minimum amount of time, in seconds, that you want CloudFront to cache the HTTP status code specified in ErrorCode.
	Ttl awscdk.Duration `json:"ttl" yaml:"ttl"`
}

// HTTP status code to failover to second origin.
//
// TODO: EXAMPLE
//
type FailoverStatusCode string

const (
	FailoverStatusCode_FORBIDDEN FailoverStatusCode = "FORBIDDEN"
	FailoverStatusCode_NOT_FOUND FailoverStatusCode = "NOT_FOUND"
	FailoverStatusCode_INTERNAL_SERVER_ERROR FailoverStatusCode = "INTERNAL_SERVER_ERROR"
	FailoverStatusCode_BAD_GATEWAY FailoverStatusCode = "BAD_GATEWAY"
	FailoverStatusCode_SERVICE_UNAVAILABLE FailoverStatusCode = "SERVICE_UNAVAILABLE"
	FailoverStatusCode_GATEWAY_TIMEOUT FailoverStatusCode = "GATEWAY_TIMEOUT"
)

// Options when reading the function's code from an external file.
//
// TODO: EXAMPLE
//
type FileCodeOptions struct {
	// The path of the file to read the code from.
	FilePath *string `json:"filePath" yaml:"filePath"`
}

// A CloudFront Function.
//
// TODO: EXAMPLE
//
type Function interface {
	awscdk.Resource
	IFunction
	Env() *awscdk.ResourceEnvironment
	FunctionArn() *string
	FunctionName() *string
	FunctionStage() *string
	Node() constructs.Node
	PhysicalName() *string
	Stack() awscdk.Stack
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ToString() *string
}

// The jsii proxy struct for Function
type jsiiProxy_Function struct {
	internal.Type__awscdkResource
	jsiiProxy_IFunction
}

func (j *jsiiProxy_Function) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Function) FunctionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Function) FunctionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Function) FunctionStage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionStage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Function) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Function) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Function) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewFunction(scope constructs.Construct, id *string, props *FunctionProps) Function {
	_init_.Initialize()

	j := jsiiProxy_Function{}

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront.Function",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewFunction_Override(f Function, scope constructs.Construct, id *string, props *FunctionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront.Function",
		[]interface{}{scope, id, props},
		f,
	)
}

// Imports a function by its name and ARN.
func Function_FromFunctionAttributes(scope constructs.Construct, id *string, attrs *FunctionAttributes) IFunction {
	_init_.Initialize()

	var returns IFunction

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.Function",
		"fromFunctionAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func Function_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.Function",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func Function_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.Function",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (f *jsiiProxy_Function) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		f,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (f *jsiiProxy_Function) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
func (f *jsiiProxy_Function) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
func (f *jsiiProxy_Function) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (f *jsiiProxy_Function) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Represents a CloudFront function and event type when using CF Functions.
//
// The type of the {@link AddBehaviorOptions.functionAssociations} property.
//
// TODO: EXAMPLE
//
type FunctionAssociation struct {
	// The type of event which should invoke the function.
	EventType FunctionEventType `json:"eventType" yaml:"eventType"`
	// The CloudFront function that will be invoked.
	Function IFunction `json:"function" yaml:"function"`
}

// Attributes of an existing CloudFront Function to import it.
//
// TODO: EXAMPLE
//
type FunctionAttributes struct {
	// The ARN of the function.
	FunctionArn *string `json:"functionArn" yaml:"functionArn"`
	// The name of the function.
	FunctionName *string `json:"functionName" yaml:"functionName"`
}

// Represents the function's source code.
//
// TODO: EXAMPLE
//
type FunctionCode interface {
	Render() *string
}

// The jsii proxy struct for FunctionCode
type jsiiProxy_FunctionCode struct {
	_ byte // padding
}

func NewFunctionCode_Override(f FunctionCode) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront.FunctionCode",
		nil, // no parameters
		f,
	)
}

// Code from external file for function.
//
// Returns: code object with contents from file.
func FunctionCode_FromFile(options *FileCodeOptions) FunctionCode {
	_init_.Initialize()

	var returns FunctionCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.FunctionCode",
		"fromFile",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Inline code for function.
//
// Returns: code object with inline code.
func FunctionCode_FromInline(code *string) FunctionCode {
	_init_.Initialize()

	var returns FunctionCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.FunctionCode",
		"fromInline",
		[]interface{}{code},
		&returns,
	)

	return returns
}

// renders the function code.
func (f *jsiiProxy_FunctionCode) Render() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"render",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The type of events that a CloudFront function can be invoked in response to.
//
// TODO: EXAMPLE
//
type FunctionEventType string

const (
	FunctionEventType_VIEWER_REQUEST FunctionEventType = "VIEWER_REQUEST"
	FunctionEventType_VIEWER_RESPONSE FunctionEventType = "VIEWER_RESPONSE"
)

// Properties for creating a CloudFront Function.
//
// TODO: EXAMPLE
//
type FunctionProps struct {
	// The source code of the function.
	Code FunctionCode `json:"code" yaml:"code"`
	// A comment to describe the function.
	Comment *string `json:"comment" yaml:"comment"`
	// A name to identify the function.
	FunctionName *string `json:"functionName" yaml:"functionName"`
}

// Controls the countries in which content is distributed.
//
// TODO: EXAMPLE
//
type GeoRestriction interface {
	Locations() *[]*string
	RestrictionType() *string
}

// The jsii proxy struct for GeoRestriction
type jsiiProxy_GeoRestriction struct {
	_ byte // padding
}

func (j *jsiiProxy_GeoRestriction) Locations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"locations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GeoRestriction) RestrictionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restrictionType",
		&returns,
	)
	return returns
}


// Allow specific countries which you want CloudFront to distribute your content.
func GeoRestriction_Allowlist(locations ...*string) GeoRestriction {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range locations {
		args = append(args, a)
	}

	var returns GeoRestriction

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.GeoRestriction",
		"allowlist",
		args,
		&returns,
	)

	return returns
}

// Deny specific countries which you don't want CloudFront to distribute your content.
func GeoRestriction_Denylist(locations ...*string) GeoRestriction {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range locations {
		args = append(args, a)
	}

	var returns GeoRestriction

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.GeoRestriction",
		"denylist",
		args,
		&returns,
	)

	return returns
}

// Enum representing possible values of the X-Frame-Options HTTP response header.
//
// TODO: EXAMPLE
//
type HeadersFrameOption string

const (
	HeadersFrameOption_DENY HeadersFrameOption = "DENY"
	HeadersFrameOption_SAMEORIGIN HeadersFrameOption = "SAMEORIGIN"
)

// Enum representing possible values of the Referrer-Policy HTTP response header.
//
// TODO: EXAMPLE
//
type HeadersReferrerPolicy string

const (
	HeadersReferrerPolicy_NO_REFERRER HeadersReferrerPolicy = "NO_REFERRER"
	HeadersReferrerPolicy_NO_REFERRER_WHEN_DOWNGRADE HeadersReferrerPolicy = "NO_REFERRER_WHEN_DOWNGRADE"
	HeadersReferrerPolicy_ORIGIN HeadersReferrerPolicy = "ORIGIN"
	HeadersReferrerPolicy_ORIGIN_WHEN_CROSS_ORIGIN HeadersReferrerPolicy = "ORIGIN_WHEN_CROSS_ORIGIN"
	HeadersReferrerPolicy_SAME_ORIGIN HeadersReferrerPolicy = "SAME_ORIGIN"
	HeadersReferrerPolicy_STRICT_ORIGIN HeadersReferrerPolicy = "STRICT_ORIGIN"
	HeadersReferrerPolicy_STRICT_ORIGIN_WHEN_CROSS_ORIGIN HeadersReferrerPolicy = "STRICT_ORIGIN_WHEN_CROSS_ORIGIN"
	HeadersReferrerPolicy_UNSAFE_URL HeadersReferrerPolicy = "UNSAFE_URL"
)

// Maximum HTTP version to support.
type HttpVersion string

const (
	HttpVersion_HTTP1_1 HttpVersion = "HTTP1_1"
	HttpVersion_HTTP2 HttpVersion = "HTTP2"
)

// Represents a Cache Policy.
type ICachePolicy interface {
	// The ID of the cache policy.
	CachePolicyId() *string
}

// The jsii proxy for ICachePolicy
type jsiiProxy_ICachePolicy struct {
	_ byte // padding
}

func (j *jsiiProxy_ICachePolicy) CachePolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cachePolicyId",
		&returns,
	)
	return returns
}

// Interface for CloudFront distributions.
type IDistribution interface {
	awscdk.IResource
	// The domain name of the Distribution, such as d111111abcdef8.cloudfront.net.
	DistributionDomainName() *string
	// The distribution ID for this distribution.
	DistributionId() *string
}

// The jsii proxy for IDistribution
type jsiiProxy_IDistribution struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IDistribution) DistributionDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"distributionDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDistribution) DistributionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"distributionId",
		&returns,
	)
	return returns
}

// Represents a CloudFront Function.
type IFunction interface {
	awscdk.IResource
	// The ARN of the function.
	FunctionArn() *string
	// The name of the function.
	FunctionName() *string
}

// The jsii proxy for IFunction
type jsiiProxy_IFunction struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IFunction) FunctionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFunction) FunctionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionName",
		&returns,
	)
	return returns
}

// Represents a Key Group.
type IKeyGroup interface {
	awscdk.IResource
	// The ID of the key group.
	KeyGroupId() *string
}

// The jsii proxy for IKeyGroup
type jsiiProxy_IKeyGroup struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IKeyGroup) KeyGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyGroupId",
		&returns,
	)
	return returns
}

// Represents the concept of a CloudFront Origin.
//
// You provide one or more origins when creating a Distribution.
type IOrigin interface {
	// The method called when a given Origin is added (for the first time) to a Distribution.
	Bind(scope constructs.Construct, options *OriginBindOptions) *OriginBindConfig
}

// The jsii proxy for IOrigin
type jsiiProxy_IOrigin struct {
	_ byte // padding
}

func (i *jsiiProxy_IOrigin) Bind(scope constructs.Construct, options *OriginBindOptions) *OriginBindConfig {
	var returns *OriginBindConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{scope, options},
		&returns,
	)

	return returns
}

// Interface for CloudFront OriginAccessIdentity.
type IOriginAccessIdentity interface {
	awsiam.IGrantable
	awscdk.IResource
	// The Origin Access Identity Name.
	OriginAccessIdentityName() *string
}

// The jsii proxy for IOriginAccessIdentity
type jsiiProxy_IOriginAccessIdentity struct {
	internal.Type__awsiamIGrantable
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IOriginAccessIdentity) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (j *jsiiProxy_IOriginAccessIdentity) OriginAccessIdentityName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originAccessIdentityName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOriginAccessIdentity) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOriginAccessIdentity) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOriginAccessIdentity) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOriginAccessIdentity) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

// Represents a Origin Request Policy.
type IOriginRequestPolicy interface {
	// The ID of the origin request policy.
	OriginRequestPolicyId() *string
}

// The jsii proxy for IOriginRequestPolicy
type jsiiProxy_IOriginRequestPolicy struct {
	_ byte // padding
}

func (j *jsiiProxy_IOriginRequestPolicy) OriginRequestPolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originRequestPolicyId",
		&returns,
	)
	return returns
}

// Represents a Public Key.
type IPublicKey interface {
	awscdk.IResource
	// The ID of the key group.
	PublicKeyId() *string
}

// The jsii proxy for IPublicKey
type jsiiProxy_IPublicKey struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IPublicKey) PublicKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicKeyId",
		&returns,
	)
	return returns
}

// Represents a response headers policy.
type IResponseHeadersPolicy interface {
	// The ID of the response headers policy.
	ResponseHeadersPolicyId() *string
}

// The jsii proxy for IResponseHeadersPolicy
type jsiiProxy_IResponseHeadersPolicy struct {
	_ byte // padding
}

func (j *jsiiProxy_IResponseHeadersPolicy) ResponseHeadersPolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseHeadersPolicyId",
		&returns,
	)
	return returns
}

// A Key Group configuration.
//
// TODO: EXAMPLE
//
type KeyGroup interface {
	awscdk.Resource
	IKeyGroup
	Env() *awscdk.ResourceEnvironment
	KeyGroupId() *string
	Node() constructs.Node
	PhysicalName() *string
	Stack() awscdk.Stack
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ToString() *string
}

// The jsii proxy struct for KeyGroup
type jsiiProxy_KeyGroup struct {
	internal.Type__awscdkResource
	jsiiProxy_IKeyGroup
}

func (j *jsiiProxy_KeyGroup) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyGroup) KeyGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyGroup) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyGroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewKeyGroup(scope constructs.Construct, id *string, props *KeyGroupProps) KeyGroup {
	_init_.Initialize()

	j := jsiiProxy_KeyGroup{}

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront.KeyGroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewKeyGroup_Override(k KeyGroup, scope constructs.Construct, id *string, props *KeyGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront.KeyGroup",
		[]interface{}{scope, id, props},
		k,
	)
}

// Imports a Key Group from its id.
func KeyGroup_FromKeyGroupId(scope constructs.Construct, id *string, keyGroupId *string) IKeyGroup {
	_init_.Initialize()

	var returns IKeyGroup

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.KeyGroup",
		"fromKeyGroupId",
		[]interface{}{scope, id, keyGroupId},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func KeyGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.KeyGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func KeyGroup_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.KeyGroup",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (k *jsiiProxy_KeyGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		k,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (k *jsiiProxy_KeyGroup) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
func (k *jsiiProxy_KeyGroup) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
func (k *jsiiProxy_KeyGroup) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (k *jsiiProxy_KeyGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for creating a Public Key.
//
// TODO: EXAMPLE
//
type KeyGroupProps struct {
	// A list of public keys to add to the key group.
	Items *[]IPublicKey `json:"items" yaml:"items"`
	// A comment to describe the key group.
	Comment *string `json:"comment" yaml:"comment"`
	// A name to identify the key group.
	KeyGroupName *string `json:"keyGroupName" yaml:"keyGroupName"`
}

// The type of events that a Lambda@Edge function can be invoked in response to.
//
// TODO: EXAMPLE
//
type LambdaEdgeEventType string

const (
	LambdaEdgeEventType_ORIGIN_REQUEST LambdaEdgeEventType = "ORIGIN_REQUEST"
	LambdaEdgeEventType_ORIGIN_RESPONSE LambdaEdgeEventType = "ORIGIN_RESPONSE"
	LambdaEdgeEventType_VIEWER_REQUEST LambdaEdgeEventType = "VIEWER_REQUEST"
	LambdaEdgeEventType_VIEWER_RESPONSE LambdaEdgeEventType = "VIEWER_RESPONSE"
)

// TODO: EXAMPLE
//
type LambdaFunctionAssociation struct {
	// The lambda event type defines at which event the lambda is called during the request lifecycle.
	EventType LambdaEdgeEventType `json:"eventType" yaml:"eventType"`
	// A version of the lambda to associate.
	LambdaFunction awslambda.IVersion `json:"lambdaFunction" yaml:"lambdaFunction"`
	// Allows a Lambda function to have read access to the body content.
	//
	// Only valid for "request" event types (`ORIGIN_REQUEST` or `VIEWER_REQUEST`).
	// See https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/lambda-include-body-access.html
	IncludeBody *bool `json:"includeBody" yaml:"includeBody"`
}

// Logging configuration for incoming requests.
//
// TODO: EXAMPLE
//
type LoggingConfiguration struct {
	// Bucket to log requests to.
	Bucket awss3.IBucket `json:"bucket" yaml:"bucket"`
	// Whether to include the cookies in the logs.
	IncludeCookies *bool `json:"includeCookies" yaml:"includeCookies"`
	// Where in the bucket to store logs.
	Prefix *string `json:"prefix" yaml:"prefix"`
}

// An origin access identity is a special CloudFront user that you can associate with Amazon S3 origins, so that you can secure all or just some of your Amazon S3 content.
//
// TODO: EXAMPLE
//
type OriginAccessIdentity interface {
	awscdk.Resource
	IOriginAccessIdentity
	CloudFrontOriginAccessIdentityS3CanonicalUserId() *string
	Env() *awscdk.ResourceEnvironment
	GrantPrincipal() awsiam.IPrincipal
	Node() constructs.Node
	OriginAccessIdentityName() *string
	PhysicalName() *string
	Stack() awscdk.Stack
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	Arn() *string
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ToString() *string
}

// The jsii proxy struct for OriginAccessIdentity
type jsiiProxy_OriginAccessIdentity struct {
	internal.Type__awscdkResource
	jsiiProxy_IOriginAccessIdentity
}

func (j *jsiiProxy_OriginAccessIdentity) CloudFrontOriginAccessIdentityS3CanonicalUserId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudFrontOriginAccessIdentityS3CanonicalUserId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OriginAccessIdentity) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OriginAccessIdentity) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OriginAccessIdentity) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OriginAccessIdentity) OriginAccessIdentityName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originAccessIdentityName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OriginAccessIdentity) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OriginAccessIdentity) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewOriginAccessIdentity(scope constructs.Construct, id *string, props *OriginAccessIdentityProps) OriginAccessIdentity {
	_init_.Initialize()

	j := jsiiProxy_OriginAccessIdentity{}

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront.OriginAccessIdentity",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewOriginAccessIdentity_Override(o OriginAccessIdentity, scope constructs.Construct, id *string, props *OriginAccessIdentityProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront.OriginAccessIdentity",
		[]interface{}{scope, id, props},
		o,
	)
}

// Creates a OriginAccessIdentity by providing the OriginAccessIdentityName.
func OriginAccessIdentity_FromOriginAccessIdentityName(scope constructs.Construct, id *string, originAccessIdentityName *string) IOriginAccessIdentity {
	_init_.Initialize()

	var returns IOriginAccessIdentity

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.OriginAccessIdentity",
		"fromOriginAccessIdentityName",
		[]interface{}{scope, id, originAccessIdentityName},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func OriginAccessIdentity_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.OriginAccessIdentity",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func OriginAccessIdentity_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.OriginAccessIdentity",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (o *jsiiProxy_OriginAccessIdentity) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		o,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// The ARN to include in S3 bucket policy to allow CloudFront access.
func (o *jsiiProxy_OriginAccessIdentity) Arn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"arn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OriginAccessIdentity) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
func (o *jsiiProxy_OriginAccessIdentity) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
func (o *jsiiProxy_OriginAccessIdentity) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (o *jsiiProxy_OriginAccessIdentity) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties of CloudFront OriginAccessIdentity.
//
// TODO: EXAMPLE
//
type OriginAccessIdentityProps struct {
	// Any comments you want to include about the origin access identity.
	Comment *string `json:"comment" yaml:"comment"`
}

// Represents a distribution origin, that describes the Amazon S3 bucket, HTTP server (for example, a web server), Amazon MediaStore, or other server from which CloudFront gets your files.
type OriginBase interface {
	IOrigin
	Bind(_scope constructs.Construct, options *OriginBindOptions) *OriginBindConfig
	RenderCustomOriginConfig() *CfnDistribution_CustomOriginConfigProperty
	RenderS3OriginConfig() *CfnDistribution_S3OriginConfigProperty
}

// The jsii proxy struct for OriginBase
type jsiiProxy_OriginBase struct {
	jsiiProxy_IOrigin
}

func NewOriginBase_Override(o OriginBase, domainName *string, props *OriginProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront.OriginBase",
		[]interface{}{domainName, props},
		o,
	)
}

// Binds the origin to the associated Distribution.
//
// Can be used to grant permissions, create dependent resources, etc.
func (o *jsiiProxy_OriginBase) Bind(_scope constructs.Construct, options *OriginBindOptions) *OriginBindConfig {
	var returns *OriginBindConfig

	_jsii_.Invoke(
		o,
		"bind",
		[]interface{}{_scope, options},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OriginBase) RenderCustomOriginConfig() *CfnDistribution_CustomOriginConfigProperty {
	var returns *CfnDistribution_CustomOriginConfigProperty

	_jsii_.Invoke(
		o,
		"renderCustomOriginConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OriginBase) RenderS3OriginConfig() *CfnDistribution_S3OriginConfigProperty {
	var returns *CfnDistribution_S3OriginConfigProperty

	_jsii_.Invoke(
		o,
		"renderS3OriginConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The struct returned from {@link IOrigin.bind}.
//
// TODO: EXAMPLE
//
type OriginBindConfig struct {
	// The failover configuration for this Origin.
	FailoverConfig *OriginFailoverConfig `json:"failoverConfig" yaml:"failoverConfig"`
	// The CloudFormation OriginProperty configuration for this Origin.
	OriginProperty *CfnDistribution_OriginProperty `json:"originProperty" yaml:"originProperty"`
}

// Options passed to Origin.bind().
//
// TODO: EXAMPLE
//
type OriginBindOptions struct {
	// The identifier of this Origin, as assigned by the Distribution this Origin has been used added to.
	OriginId *string `json:"originId" yaml:"originId"`
}

// The failover configuration used for Origin Groups, returned in {@link OriginBindConfig.failoverConfig}.
//
// TODO: EXAMPLE
//
type OriginFailoverConfig struct {
	// The origin to use as the fallback origin.
	FailoverOrigin IOrigin `json:"failoverOrigin" yaml:"failoverOrigin"`
	// The HTTP status codes of the response that trigger querying the failover Origin.
	StatusCodes *[]*float64 `json:"statusCodes" yaml:"statusCodes"`
}

// Properties to define an Origin.
//
// TODO: EXAMPLE
//
type OriginProps struct {
	// The number of times that CloudFront attempts to connect to the origin;
	//
	// valid values are 1, 2, or 3 attempts.
	ConnectionAttempts *float64 `json:"connectionAttempts" yaml:"connectionAttempts"`
	// The number of seconds that CloudFront waits when trying to establish a connection to the origin.
	//
	// Valid values are 1-10 seconds, inclusive.
	ConnectionTimeout awscdk.Duration `json:"connectionTimeout" yaml:"connectionTimeout"`
	// A list of HTTP header names and values that CloudFront adds to requests it sends to the origin.
	CustomHeaders *map[string]*string `json:"customHeaders" yaml:"customHeaders"`
	// An optional path that CloudFront appends to the origin domain name when CloudFront requests content from the origin.
	//
	// Must begin, but not end, with '/' (e.g., '/production/images').
	OriginPath *string `json:"originPath" yaml:"originPath"`
	// When you enable Origin Shield in the AWS Region that has the lowest latency to your origin, you can get better network performance.
	// See: https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/origin-shield.html
	//
	OriginShieldRegion *string `json:"originShieldRegion" yaml:"originShieldRegion"`
}

// Defines what protocols CloudFront will use to connect to an origin.
//
// TODO: EXAMPLE
//
type OriginProtocolPolicy string

const (
	OriginProtocolPolicy_HTTP_ONLY OriginProtocolPolicy = "HTTP_ONLY"
	OriginProtocolPolicy_MATCH_VIEWER OriginProtocolPolicy = "MATCH_VIEWER"
	OriginProtocolPolicy_HTTPS_ONLY OriginProtocolPolicy = "HTTPS_ONLY"
)

// Determines whether any cookies in viewer requests (and if so, which cookies) are included in requests that CloudFront sends to the origin.
//
// TODO: EXAMPLE
//
type OriginRequestCookieBehavior interface {
	Behavior() *string
	Cookies() *[]*string
}

// The jsii proxy struct for OriginRequestCookieBehavior
type jsiiProxy_OriginRequestCookieBehavior struct {
	_ byte // padding
}

func (j *jsiiProxy_OriginRequestCookieBehavior) Behavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"behavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OriginRequestCookieBehavior) Cookies() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cookies",
		&returns,
	)
	return returns
}


// All cookies in viewer requests are included in requests that CloudFront sends to the origin.
func OriginRequestCookieBehavior_All() OriginRequestCookieBehavior {
	_init_.Initialize()

	var returns OriginRequestCookieBehavior

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.OriginRequestCookieBehavior",
		"all",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Only the provided `cookies` are included in requests that CloudFront sends to the origin.
func OriginRequestCookieBehavior_AllowList(cookies ...*string) OriginRequestCookieBehavior {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range cookies {
		args = append(args, a)
	}

	var returns OriginRequestCookieBehavior

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.OriginRequestCookieBehavior",
		"allowList",
		args,
		&returns,
	)

	return returns
}

// Cookies in viewer requests are not included in requests that CloudFront sends to the origin.
//
// Any cookies that are listed in a CachePolicy are still included in origin requests.
func OriginRequestCookieBehavior_None() OriginRequestCookieBehavior {
	_init_.Initialize()

	var returns OriginRequestCookieBehavior

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.OriginRequestCookieBehavior",
		"none",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Determines whether any HTTP headers (and if so, which headers) are included in requests that CloudFront sends to the origin.
//
// TODO: EXAMPLE
//
type OriginRequestHeaderBehavior interface {
	Behavior() *string
	Headers() *[]*string
}

// The jsii proxy struct for OriginRequestHeaderBehavior
type jsiiProxy_OriginRequestHeaderBehavior struct {
	_ byte // padding
}

func (j *jsiiProxy_OriginRequestHeaderBehavior) Behavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"behavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OriginRequestHeaderBehavior) Headers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"headers",
		&returns,
	)
	return returns
}


// All HTTP headers in viewer requests are included in requests that CloudFront sends to the origin.
//
// Additionally, any additional CloudFront headers provided are included; the additional headers are added by CloudFront.
// See: https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/using-cloudfront-headers.html
//
func OriginRequestHeaderBehavior_All(cloudfrontHeaders ...*string) OriginRequestHeaderBehavior {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range cloudfrontHeaders {
		args = append(args, a)
	}

	var returns OriginRequestHeaderBehavior

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.OriginRequestHeaderBehavior",
		"all",
		args,
		&returns,
	)

	return returns
}

// Listed headers are included in requests that CloudFront sends to the origin.
func OriginRequestHeaderBehavior_AllowList(headers ...*string) OriginRequestHeaderBehavior {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range headers {
		args = append(args, a)
	}

	var returns OriginRequestHeaderBehavior

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.OriginRequestHeaderBehavior",
		"allowList",
		args,
		&returns,
	)

	return returns
}

// HTTP headers are not included in requests that CloudFront sends to the origin.
//
// Any headers that are listed in a CachePolicy are still included in origin requests.
func OriginRequestHeaderBehavior_None() OriginRequestHeaderBehavior {
	_init_.Initialize()

	var returns OriginRequestHeaderBehavior

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.OriginRequestHeaderBehavior",
		"none",
		nil, // no parameters
		&returns,
	)

	return returns
}

// A Origin Request Policy configuration.
//
// TODO: EXAMPLE
//
type OriginRequestPolicy interface {
	awscdk.Resource
	IOriginRequestPolicy
	Env() *awscdk.ResourceEnvironment
	Node() constructs.Node
	OriginRequestPolicyId() *string
	PhysicalName() *string
	Stack() awscdk.Stack
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ToString() *string
}

// The jsii proxy struct for OriginRequestPolicy
type jsiiProxy_OriginRequestPolicy struct {
	internal.Type__awscdkResource
	jsiiProxy_IOriginRequestPolicy
}

func (j *jsiiProxy_OriginRequestPolicy) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OriginRequestPolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OriginRequestPolicy) OriginRequestPolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originRequestPolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OriginRequestPolicy) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OriginRequestPolicy) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewOriginRequestPolicy(scope constructs.Construct, id *string, props *OriginRequestPolicyProps) OriginRequestPolicy {
	_init_.Initialize()

	j := jsiiProxy_OriginRequestPolicy{}

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront.OriginRequestPolicy",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewOriginRequestPolicy_Override(o OriginRequestPolicy, scope constructs.Construct, id *string, props *OriginRequestPolicyProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront.OriginRequestPolicy",
		[]interface{}{scope, id, props},
		o,
	)
}

// Imports a Origin Request Policy from its id.
func OriginRequestPolicy_FromOriginRequestPolicyId(scope constructs.Construct, id *string, originRequestPolicyId *string) IOriginRequestPolicy {
	_init_.Initialize()

	var returns IOriginRequestPolicy

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.OriginRequestPolicy",
		"fromOriginRequestPolicyId",
		[]interface{}{scope, id, originRequestPolicyId},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func OriginRequestPolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.OriginRequestPolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func OriginRequestPolicy_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.OriginRequestPolicy",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func OriginRequestPolicy_ALL_VIEWER() IOriginRequestPolicy {
	_init_.Initialize()
	var returns IOriginRequestPolicy
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cloudfront.OriginRequestPolicy",
		"ALL_VIEWER",
		&returns,
	)
	return returns
}

func OriginRequestPolicy_CORS_CUSTOM_ORIGIN() IOriginRequestPolicy {
	_init_.Initialize()
	var returns IOriginRequestPolicy
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cloudfront.OriginRequestPolicy",
		"CORS_CUSTOM_ORIGIN",
		&returns,
	)
	return returns
}

func OriginRequestPolicy_CORS_S3_ORIGIN() IOriginRequestPolicy {
	_init_.Initialize()
	var returns IOriginRequestPolicy
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cloudfront.OriginRequestPolicy",
		"CORS_S3_ORIGIN",
		&returns,
	)
	return returns
}

func OriginRequestPolicy_ELEMENTAL_MEDIA_TAILOR() IOriginRequestPolicy {
	_init_.Initialize()
	var returns IOriginRequestPolicy
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cloudfront.OriginRequestPolicy",
		"ELEMENTAL_MEDIA_TAILOR",
		&returns,
	)
	return returns
}

func OriginRequestPolicy_USER_AGENT_REFERER_HEADERS() IOriginRequestPolicy {
	_init_.Initialize()
	var returns IOriginRequestPolicy
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cloudfront.OriginRequestPolicy",
		"USER_AGENT_REFERER_HEADERS",
		&returns,
	)
	return returns
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (o *jsiiProxy_OriginRequestPolicy) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		o,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (o *jsiiProxy_OriginRequestPolicy) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
func (o *jsiiProxy_OriginRequestPolicy) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
func (o *jsiiProxy_OriginRequestPolicy) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (o *jsiiProxy_OriginRequestPolicy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for creating a Origin Request Policy.
//
// TODO: EXAMPLE
//
type OriginRequestPolicyProps struct {
	// A comment to describe the origin request policy.
	Comment *string `json:"comment" yaml:"comment"`
	// The cookies from viewer requests to include in origin requests.
	CookieBehavior OriginRequestCookieBehavior `json:"cookieBehavior" yaml:"cookieBehavior"`
	// The HTTP headers to include in origin requests.
	//
	// These can include headers from viewer requests and additional headers added by CloudFront.
	HeaderBehavior OriginRequestHeaderBehavior `json:"headerBehavior" yaml:"headerBehavior"`
	// A unique name to identify the origin request policy.
	//
	// The name must only include '-', '_', or alphanumeric characters.
	OriginRequestPolicyName *string `json:"originRequestPolicyName" yaml:"originRequestPolicyName"`
	// The URL query strings from viewer requests to include in origin requests.
	QueryStringBehavior OriginRequestQueryStringBehavior `json:"queryStringBehavior" yaml:"queryStringBehavior"`
}

// Determines whether any URL query strings in viewer requests (and if so, which query strings) are included in requests that CloudFront sends to the origin.
//
// TODO: EXAMPLE
//
type OriginRequestQueryStringBehavior interface {
	Behavior() *string
	QueryStrings() *[]*string
}

// The jsii proxy struct for OriginRequestQueryStringBehavior
type jsiiProxy_OriginRequestQueryStringBehavior struct {
	_ byte // padding
}

func (j *jsiiProxy_OriginRequestQueryStringBehavior) Behavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"behavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OriginRequestQueryStringBehavior) QueryStrings() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"queryStrings",
		&returns,
	)
	return returns
}


// All query strings in viewer requests are included in requests that CloudFront sends to the origin.
func OriginRequestQueryStringBehavior_All() OriginRequestQueryStringBehavior {
	_init_.Initialize()

	var returns OriginRequestQueryStringBehavior

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.OriginRequestQueryStringBehavior",
		"all",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Only the provided `queryStrings` are included in requests that CloudFront sends to the origin.
func OriginRequestQueryStringBehavior_AllowList(queryStrings ...*string) OriginRequestQueryStringBehavior {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range queryStrings {
		args = append(args, a)
	}

	var returns OriginRequestQueryStringBehavior

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.OriginRequestQueryStringBehavior",
		"allowList",
		args,
		&returns,
	)

	return returns
}

// Query strings in viewer requests are not included in requests that CloudFront sends to the origin.
//
// Any query strings that are listed in a CachePolicy are still included in origin requests.
func OriginRequestQueryStringBehavior_None() OriginRequestQueryStringBehavior {
	_init_.Initialize()

	var returns OriginRequestQueryStringBehavior

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.OriginRequestQueryStringBehavior",
		"none",
		nil, // no parameters
		&returns,
	)

	return returns
}

type OriginSslPolicy string

const (
	OriginSslPolicy_SSL_V3 OriginSslPolicy = "SSL_V3"
	OriginSslPolicy_TLS_V1 OriginSslPolicy = "TLS_V1"
	OriginSslPolicy_TLS_V1_1 OriginSslPolicy = "TLS_V1_1"
	OriginSslPolicy_TLS_V1_2 OriginSslPolicy = "TLS_V1_2"
)

// The price class determines how many edge locations CloudFront will use for your distribution.
//
// See https://aws.amazon.com/cloudfront/pricing/ for full list of supported regions.
type PriceClass string

const (
	PriceClass_PRICE_CLASS_100 PriceClass = "PRICE_CLASS_100"
	PriceClass_PRICE_CLASS_200 PriceClass = "PRICE_CLASS_200"
	PriceClass_PRICE_CLASS_ALL PriceClass = "PRICE_CLASS_ALL"
)

// A Public Key Configuration.
//
// TODO: EXAMPLE
//
type PublicKey interface {
	awscdk.Resource
	IPublicKey
	Env() *awscdk.ResourceEnvironment
	Node() constructs.Node
	PhysicalName() *string
	PublicKeyId() *string
	Stack() awscdk.Stack
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ToString() *string
}

// The jsii proxy struct for PublicKey
type jsiiProxy_PublicKey struct {
	internal.Type__awscdkResource
	jsiiProxy_IPublicKey
}

func (j *jsiiProxy_PublicKey) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublicKey) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublicKey) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublicKey) PublicKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublicKey) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewPublicKey(scope constructs.Construct, id *string, props *PublicKeyProps) PublicKey {
	_init_.Initialize()

	j := jsiiProxy_PublicKey{}

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront.PublicKey",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewPublicKey_Override(p PublicKey, scope constructs.Construct, id *string, props *PublicKeyProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront.PublicKey",
		[]interface{}{scope, id, props},
		p,
	)
}

// Imports a Public Key from its id.
func PublicKey_FromPublicKeyId(scope constructs.Construct, id *string, publicKeyId *string) IPublicKey {
	_init_.Initialize()

	var returns IPublicKey

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.PublicKey",
		"fromPublicKeyId",
		[]interface{}{scope, id, publicKeyId},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func PublicKey_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.PublicKey",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func PublicKey_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.PublicKey",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (p *jsiiProxy_PublicKey) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		p,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (p *jsiiProxy_PublicKey) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
func (p *jsiiProxy_PublicKey) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
func (p *jsiiProxy_PublicKey) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (p *jsiiProxy_PublicKey) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for creating a Public Key.
//
// TODO: EXAMPLE
//
type PublicKeyProps struct {
	// The public key that you can use with signed URLs and signed cookies, or with field-level encryption.
	//
	// The `encodedKey` parameter must include `-----BEGIN PUBLIC KEY-----` and `-----END PUBLIC KEY-----` lines.
	// See: https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/field-level-encryption.html
	//
	EncodedKey *string `json:"encodedKey" yaml:"encodedKey"`
	// A comment to describe the public key.
	Comment *string `json:"comment" yaml:"comment"`
	// A name to identify the public key.
	PublicKeyName *string `json:"publicKeyName" yaml:"publicKeyName"`
}

// An HTTP response header name and its value.
//
// CloudFront includes this header in HTTP responses that it sends for requests that match a cache behavior that’s associated with this response headers policy.
//
// TODO: EXAMPLE
//
type ResponseCustomHeader struct {
	// The HTTP response header name.
	Header *string `json:"header" yaml:"header"`
	// A Boolean that determines whether CloudFront overrides a response header with the same name received from the origin with the header specified here.
	Override *bool `json:"override" yaml:"override"`
	// The value for the HTTP response header.
	Value *string `json:"value" yaml:"value"`
}

// Configuration for a set of HTTP response headers that are sent for requests that match a cache behavior that’s associated with this response headers policy.
//
// TODO: EXAMPLE
//
type ResponseCustomHeadersBehavior struct {
	// The list of HTTP response headers and their values.
	CustomHeaders *[]*ResponseCustomHeader `json:"customHeaders" yaml:"customHeaders"`
}

// The policy directives and their values that CloudFront includes as values for the Content-Security-Policy HTTP response header.
//
// TODO: EXAMPLE
//
type ResponseHeadersContentSecurityPolicy struct {
	// The policy directives and their values that CloudFront includes as values for the Content-Security-Policy HTTP response header.
	ContentSecurityPolicy *string `json:"contentSecurityPolicy" yaml:"contentSecurityPolicy"`
	// A Boolean that determines whether CloudFront overrides the Content-Security-Policy HTTP response header received from the origin with the one specified in this response headers policy.
	Override *bool `json:"override" yaml:"override"`
}

// Determines whether CloudFront includes the X-Content-Type-Options HTTP response header with its value set to nosniff.
//
// TODO: EXAMPLE
//
type ResponseHeadersContentTypeOptions struct {
	// A Boolean that determines whether CloudFront overrides the X-Content-Type-Options HTTP response header received from the origin with the one specified in this response headers policy.
	Override *bool `json:"override" yaml:"override"`
}

// Configuration for a set of HTTP response headers that are used for cross-origin resource sharing (CORS).
//
// CloudFront adds these headers to HTTP responses that it sends for CORS requests that match a cache behavior
// associated with this response headers policy.
//
// TODO: EXAMPLE
//
type ResponseHeadersCorsBehavior struct {
	// A Boolean that CloudFront uses as the value for the Access-Control-Allow-Credentials HTTP response header.
	AccessControlAllowCredentials *bool `json:"accessControlAllowCredentials" yaml:"accessControlAllowCredentials"`
	// A list of HTTP header names that CloudFront includes as values for the Access-Control-Allow-Headers HTTP response header.
	//
	// You can specify `['*']` to allow all headers.
	AccessControlAllowHeaders *[]*string `json:"accessControlAllowHeaders" yaml:"accessControlAllowHeaders"`
	// A list of HTTP methods that CloudFront includes as values for the Access-Control-Allow-Methods HTTP response header.
	AccessControlAllowMethods *[]*string `json:"accessControlAllowMethods" yaml:"accessControlAllowMethods"`
	// A list of origins (domain names) that CloudFront can use as the value for the Access-Control-Allow-Origin HTTP response header.
	//
	// You can specify `['*']` to allow all origins.
	AccessControlAllowOrigins *[]*string `json:"accessControlAllowOrigins" yaml:"accessControlAllowOrigins"`
	// A Boolean that determines whether CloudFront overrides HTTP response headers received from the origin with the ones specified in this response headers policy.
	OriginOverride *bool `json:"originOverride" yaml:"originOverride"`
	// A list of HTTP headers that CloudFront includes as values for the Access-Control-Expose-Headers HTTP response header.
	//
	// You can specify `['*']` to expose all headers.
	AccessControlExposeHeaders *[]*string `json:"accessControlExposeHeaders" yaml:"accessControlExposeHeaders"`
	// A number that CloudFront uses as the value for the Access-Control-Max-Age HTTP response header.
	AccessControlMaxAge awscdk.Duration `json:"accessControlMaxAge" yaml:"accessControlMaxAge"`
}

// Determines whether CloudFront includes the X-Frame-Options HTTP response header and the header’s value.
//
// TODO: EXAMPLE
//
type ResponseHeadersFrameOptions struct {
	// The value of the X-Frame-Options HTTP response header.
	FrameOption HeadersFrameOption `json:"frameOption" yaml:"frameOption"`
	// A Boolean that determines whether CloudFront overrides the X-Frame-Options HTTP response header received from the origin with the one specified in this response headers policy.
	Override *bool `json:"override" yaml:"override"`
}

// A Response Headers Policy configuration.
//
// TODO: EXAMPLE
//
type ResponseHeadersPolicy interface {
	awscdk.Resource
	IResponseHeadersPolicy
	Env() *awscdk.ResourceEnvironment
	Node() constructs.Node
	PhysicalName() *string
	ResponseHeadersPolicyId() *string
	Stack() awscdk.Stack
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ToString() *string
}

// The jsii proxy struct for ResponseHeadersPolicy
type jsiiProxy_ResponseHeadersPolicy struct {
	internal.Type__awscdkResource
	jsiiProxy_IResponseHeadersPolicy
}

func (j *jsiiProxy_ResponseHeadersPolicy) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResponseHeadersPolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResponseHeadersPolicy) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResponseHeadersPolicy) ResponseHeadersPolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseHeadersPolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResponseHeadersPolicy) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewResponseHeadersPolicy(scope constructs.Construct, id *string, props *ResponseHeadersPolicyProps) ResponseHeadersPolicy {
	_init_.Initialize()

	j := jsiiProxy_ResponseHeadersPolicy{}

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront.ResponseHeadersPolicy",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewResponseHeadersPolicy_Override(r ResponseHeadersPolicy, scope constructs.Construct, id *string, props *ResponseHeadersPolicyProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront.ResponseHeadersPolicy",
		[]interface{}{scope, id, props},
		r,
	)
}

// Import an existing Response Headers Policy from its ID.
func ResponseHeadersPolicy_FromResponseHeadersPolicyId(scope constructs.Construct, id *string, responseHeadersPolicyId *string) IResponseHeadersPolicy {
	_init_.Initialize()

	var returns IResponseHeadersPolicy

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.ResponseHeadersPolicy",
		"fromResponseHeadersPolicyId",
		[]interface{}{scope, id, responseHeadersPolicyId},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ResponseHeadersPolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.ResponseHeadersPolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func ResponseHeadersPolicy_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.ResponseHeadersPolicy",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func ResponseHeadersPolicy_CORS_ALLOW_ALL_ORIGINS() IResponseHeadersPolicy {
	_init_.Initialize()
	var returns IResponseHeadersPolicy
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cloudfront.ResponseHeadersPolicy",
		"CORS_ALLOW_ALL_ORIGINS",
		&returns,
	)
	return returns
}

func ResponseHeadersPolicy_CORS_ALLOW_ALL_ORIGINS_AND_SECURITY_HEADERS() IResponseHeadersPolicy {
	_init_.Initialize()
	var returns IResponseHeadersPolicy
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cloudfront.ResponseHeadersPolicy",
		"CORS_ALLOW_ALL_ORIGINS_AND_SECURITY_HEADERS",
		&returns,
	)
	return returns
}

func ResponseHeadersPolicy_CORS_ALLOW_ALL_ORIGINS_WITH_PREFLIGHT() IResponseHeadersPolicy {
	_init_.Initialize()
	var returns IResponseHeadersPolicy
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cloudfront.ResponseHeadersPolicy",
		"CORS_ALLOW_ALL_ORIGINS_WITH_PREFLIGHT",
		&returns,
	)
	return returns
}

func ResponseHeadersPolicy_CORS_ALLOW_ALL_ORIGINS_WITH_PREFLIGHT_AND_SECURITY_HEADERS() IResponseHeadersPolicy {
	_init_.Initialize()
	var returns IResponseHeadersPolicy
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cloudfront.ResponseHeadersPolicy",
		"CORS_ALLOW_ALL_ORIGINS_WITH_PREFLIGHT_AND_SECURITY_HEADERS",
		&returns,
	)
	return returns
}

func ResponseHeadersPolicy_SECURITY_HEADERS() IResponseHeadersPolicy {
	_init_.Initialize()
	var returns IResponseHeadersPolicy
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cloudfront.ResponseHeadersPolicy",
		"SECURITY_HEADERS",
		&returns,
	)
	return returns
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (r *jsiiProxy_ResponseHeadersPolicy) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		r,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (r *jsiiProxy_ResponseHeadersPolicy) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
func (r *jsiiProxy_ResponseHeadersPolicy) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
func (r *jsiiProxy_ResponseHeadersPolicy) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (r *jsiiProxy_ResponseHeadersPolicy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for creating a Response Headers Policy.
//
// TODO: EXAMPLE
//
type ResponseHeadersPolicyProps struct {
	// A comment to describe the response headers policy.
	Comment *string `json:"comment" yaml:"comment"`
	// A configuration for a set of HTTP response headers that are used for cross-origin resource sharing (CORS).
	CorsBehavior *ResponseHeadersCorsBehavior `json:"corsBehavior" yaml:"corsBehavior"`
	// A configuration for a set of custom HTTP response headers.
	CustomHeadersBehavior *ResponseCustomHeadersBehavior `json:"customHeadersBehavior" yaml:"customHeadersBehavior"`
	// A unique name to identify the response headers policy.
	ResponseHeadersPolicyName *string `json:"responseHeadersPolicyName" yaml:"responseHeadersPolicyName"`
	// A configuration for a set of security-related HTTP response headers.
	SecurityHeadersBehavior *ResponseSecurityHeadersBehavior `json:"securityHeadersBehavior" yaml:"securityHeadersBehavior"`
}

// Determines whether CloudFront includes the Referrer-Policy HTTP response header and the header’s value.
//
// TODO: EXAMPLE
//
type ResponseHeadersReferrerPolicy struct {
	// A Boolean that determines whether CloudFront overrides the Referrer-Policy HTTP response header received from the origin with the one specified in this response headers policy.
	Override *bool `json:"override" yaml:"override"`
	// The value of the Referrer-Policy HTTP response header.
	ReferrerPolicy HeadersReferrerPolicy `json:"referrerPolicy" yaml:"referrerPolicy"`
}

// Determines whether CloudFront includes the Strict-Transport-Security HTTP response header and the header’s value.
//
// TODO: EXAMPLE
//
type ResponseHeadersStrictTransportSecurity struct {
	// A number that CloudFront uses as the value for the max-age directive in the Strict-Transport-Security HTTP response header.
	AccessControlMaxAge awscdk.Duration `json:"accessControlMaxAge" yaml:"accessControlMaxAge"`
	// A Boolean that determines whether CloudFront overrides the Strict-Transport-Security HTTP response header received from the origin with the one specified in this response headers policy.
	Override *bool `json:"override" yaml:"override"`
	// A Boolean that determines whether CloudFront includes the includeSubDomains directive in the Strict-Transport-Security HTTP response header.
	IncludeSubdomains *bool `json:"includeSubdomains" yaml:"includeSubdomains"`
	// A Boolean that determines whether CloudFront includes the preload directive in the Strict-Transport-Security HTTP response header.
	Preload *bool `json:"preload" yaml:"preload"`
}

// Determines whether CloudFront includes the X-XSS-Protection HTTP response header and the header’s value.
//
// TODO: EXAMPLE
//
type ResponseHeadersXSSProtection struct {
	// A Boolean that determines whether CloudFront overrides the X-XSS-Protection HTTP response header received from the origin with the one specified in this response headers policy.
	Override *bool `json:"override" yaml:"override"`
	// A Boolean that determines the value of the X-XSS-Protection HTTP response header.
	//
	// When this setting is true, the value of the X-XSS-Protection header is 1.
	// When this setting is false, the value of the X-XSS-Protection header is 0.
	Protection *bool `json:"protection" yaml:"protection"`
	// A Boolean that determines whether CloudFront includes the mode=block directive in the X-XSS-Protection header.
	ModeBlock *bool `json:"modeBlock" yaml:"modeBlock"`
	// A reporting URI, which CloudFront uses as the value of the report directive in the X-XSS-Protection header.
	//
	// You cannot specify a ReportUri when ModeBlock is true.
	ReportUri *string `json:"reportUri" yaml:"reportUri"`
}

// Configuration for a set of security-related HTTP response headers.
//
// CloudFront adds these headers to HTTP responses that it sends for requests that match a cache behavior
// associated with this response headers policy.
//
// TODO: EXAMPLE
//
type ResponseSecurityHeadersBehavior struct {
	// The policy directives and their values that CloudFront includes as values for the Content-Security-Policy HTTP response header.
	ContentSecurityPolicy *ResponseHeadersContentSecurityPolicy `json:"contentSecurityPolicy" yaml:"contentSecurityPolicy"`
	// Determines whether CloudFront includes the X-Content-Type-Options HTTP response header with its value set to nosniff.
	ContentTypeOptions *ResponseHeadersContentTypeOptions `json:"contentTypeOptions" yaml:"contentTypeOptions"`
	// Determines whether CloudFront includes the X-Frame-Options HTTP response header and the header’s value.
	FrameOptions *ResponseHeadersFrameOptions `json:"frameOptions" yaml:"frameOptions"`
	// Determines whether CloudFront includes the Referrer-Policy HTTP response header and the header’s value.
	ReferrerPolicy *ResponseHeadersReferrerPolicy `json:"referrerPolicy" yaml:"referrerPolicy"`
	// Determines whether CloudFront includes the Strict-Transport-Security HTTP response header and the header’s value.
	StrictTransportSecurity *ResponseHeadersStrictTransportSecurity `json:"strictTransportSecurity" yaml:"strictTransportSecurity"`
	// Determines whether CloudFront includes the X-XSS-Protection HTTP response header and the header’s value.
	XssProtection *ResponseHeadersXSSProtection `json:"xssProtection" yaml:"xssProtection"`
}

// S3 origin configuration for CloudFront.
//
// TODO: EXAMPLE
//
type S3OriginConfig struct {
	// The source bucket to serve content from.
	S3BucketSource awss3.IBucket `json:"s3BucketSource" yaml:"s3BucketSource"`
	// The optional Origin Access Identity of the origin identity cloudfront will use when calling your s3 bucket.
	OriginAccessIdentity IOriginAccessIdentity `json:"originAccessIdentity" yaml:"originAccessIdentity"`
	// Any additional headers to pass to the origin.
	OriginHeaders *map[string]*string `json:"originHeaders" yaml:"originHeaders"`
	// The relative path to the origin root to use for sources.
	OriginPath *string `json:"originPath" yaml:"originPath"`
	// When you enable Origin Shield in the AWS Region that has the lowest latency to your origin, you can get better network performance.
	OriginShieldRegion *string `json:"originShieldRegion" yaml:"originShieldRegion"`
}

// The SSL method CloudFront will use for your distribution.
//
// Server Name Indication (SNI) - is an extension to the TLS computer networking protocol by which a client indicates
//   which hostname it is attempting to connect to at the start of the handshaking process. This allows a server to present
//   multiple certificates on the same IP address and TCP port number and hence allows multiple secure (HTTPS) websites
// (or any other service over TLS) to be served by the same IP address without requiring all those sites to use the same certificate.
//
// CloudFront can use SNI to host multiple distributions on the same IP - which a large majority of clients will support.
//
// If your clients cannot support SNI however - CloudFront can use dedicated IPs for your distribution - but there is a prorated monthly charge for
// using this feature. By default, we use SNI - but you can optionally enable dedicated IPs (VIP).
//
// See the CloudFront SSL for more details about pricing : https://aws.amazon.com/cloudfront/custom-ssl-domains/
type SSLMethod string

const (
	SSLMethod_SNI SSLMethod = "SNI"
	SSLMethod_VIP SSLMethod = "VIP"
)

// The minimum version of the SSL protocol that you want CloudFront to use for HTTPS connections.
//
// CloudFront serves your objects only to browsers or devices that support at least the SSL version that you specify.
//
// TODO: EXAMPLE
//
type SecurityPolicyProtocol string

const (
	SecurityPolicyProtocol_SSL_V3 SecurityPolicyProtocol = "SSL_V3"
	SecurityPolicyProtocol_TLS_V1 SecurityPolicyProtocol = "TLS_V1"
	SecurityPolicyProtocol_TLS_V1_2016 SecurityPolicyProtocol = "TLS_V1_2016"
	SecurityPolicyProtocol_TLS_V1_1_2016 SecurityPolicyProtocol = "TLS_V1_1_2016"
	SecurityPolicyProtocol_TLS_V1_2_2018 SecurityPolicyProtocol = "TLS_V1_2_2018"
	SecurityPolicyProtocol_TLS_V1_2_2019 SecurityPolicyProtocol = "TLS_V1_2_2019"
	SecurityPolicyProtocol_TLS_V1_2_2021 SecurityPolicyProtocol = "TLS_V1_2_2021"
)

// A source configuration is a wrapper for CloudFront origins and behaviors.
//
// An origin is what CloudFront will "be in front of" - that is, CloudFront will pull it's assets from an origin.
//
// If you're using s3 as a source - pass the `s3Origin` property, otherwise, pass the `customOriginSource` property.
//
// One or the other must be passed, and it is invalid to pass both in the same SourceConfiguration.
//
// TODO: EXAMPLE
//
type SourceConfiguration struct {
	// The behaviors associated with this source.
	//
	// At least one (default) behavior must be included.
	Behaviors *[]*Behavior `json:"behaviors" yaml:"behaviors"`
	// The number of times that CloudFront attempts to connect to the origin.
	//
	// You can specify 1, 2, or 3 as the number of attempts.
	ConnectionAttempts *float64 `json:"connectionAttempts" yaml:"connectionAttempts"`
	// The number of seconds that CloudFront waits when trying to establish a connection to the origin.
	//
	// You can specify a number of seconds between 1 and 10 (inclusive).
	ConnectionTimeout awscdk.Duration `json:"connectionTimeout" yaml:"connectionTimeout"`
	// A custom origin source - for all non-s3 sources.
	CustomOriginSource *CustomOriginConfig `json:"customOriginSource" yaml:"customOriginSource"`
	// HTTP status code to failover to second origin.
	FailoverCriteriaStatusCodes *[]FailoverStatusCode `json:"failoverCriteriaStatusCodes" yaml:"failoverCriteriaStatusCodes"`
	// A custom origin source for failover in case the s3OriginSource returns invalid status code.
	FailoverCustomOriginSource *CustomOriginConfig `json:"failoverCustomOriginSource" yaml:"failoverCustomOriginSource"`
	// An s3 origin source for failover in case the s3OriginSource returns invalid status code.
	FailoverS3OriginSource *S3OriginConfig `json:"failoverS3OriginSource" yaml:"failoverS3OriginSource"`
	// When you enable Origin Shield in the AWS Region that has the lowest latency to your origin, you can get better network performance.
	// See: https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/origin-shield.html
	//
	OriginShieldRegion *string `json:"originShieldRegion" yaml:"originShieldRegion"`
	// An s3 origin source - if you're using s3 for your assets.
	S3OriginSource *S3OriginConfig `json:"s3OriginSource" yaml:"s3OriginSource"`
}

// Viewer certificate configuration class.
//
// TODO: EXAMPLE
//
type ViewerCertificate interface {
	Aliases() *[]*string
	Props() *CfnDistribution_ViewerCertificateProperty
}

// The jsii proxy struct for ViewerCertificate
type jsiiProxy_ViewerCertificate struct {
	_ byte // padding
}

func (j *jsiiProxy_ViewerCertificate) Aliases() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"aliases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ViewerCertificate) Props() *CfnDistribution_ViewerCertificateProperty {
	var returns *CfnDistribution_ViewerCertificateProperty
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}


// Generate an AWS Certificate Manager (ACM) viewer certificate configuration.
func ViewerCertificate_FromAcmCertificate(certificate awscertificatemanager.ICertificate, options *ViewerCertificateOptions) ViewerCertificate {
	_init_.Initialize()

	var returns ViewerCertificate

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.ViewerCertificate",
		"fromAcmCertificate",
		[]interface{}{certificate, options},
		&returns,
	)

	return returns
}

// Generate a viewer certifcate configuration using the CloudFront default certificate (e.g. d111111abcdef8.cloudfront.net) and a {@link SecurityPolicyProtocol.TLS_V1} security policy.
func ViewerCertificate_FromCloudFrontDefaultCertificate(aliases ...*string) ViewerCertificate {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range aliases {
		args = append(args, a)
	}

	var returns ViewerCertificate

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.ViewerCertificate",
		"fromCloudFrontDefaultCertificate",
		args,
		&returns,
	)

	return returns
}

// Generate an IAM viewer certificate configuration.
func ViewerCertificate_FromIamCertificate(iamCertificateId *string, options *ViewerCertificateOptions) ViewerCertificate {
	_init_.Initialize()

	var returns ViewerCertificate

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.ViewerCertificate",
		"fromIamCertificate",
		[]interface{}{iamCertificateId, options},
		&returns,
	)

	return returns
}

// TODO: EXAMPLE
//
type ViewerCertificateOptions struct {
	// Domain names on the certificate (both main domain name and Subject Alternative names).
	Aliases *[]*string `json:"aliases" yaml:"aliases"`
	// The minimum version of the SSL protocol that you want CloudFront to use for HTTPS connections.
	//
	// CloudFront serves your objects only to browsers or devices that support at
	// least the SSL version that you specify.
	SecurityPolicy SecurityPolicyProtocol `json:"securityPolicy" yaml:"securityPolicy"`
	// How CloudFront should serve HTTPS requests.
	//
	// See the notes on SSLMethod if you wish to use other SSL termination types.
	// See: https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_ViewerCertificate.html
	//
	SslMethod SSLMethod `json:"sslMethod" yaml:"sslMethod"`
}

// How HTTPs should be handled with your distribution.
//
// TODO: EXAMPLE
//
type ViewerProtocolPolicy string

const (
	ViewerProtocolPolicy_HTTPS_ONLY ViewerProtocolPolicy = "HTTPS_ONLY"
	ViewerProtocolPolicy_REDIRECT_TO_HTTPS ViewerProtocolPolicy = "REDIRECT_TO_HTTPS"
	ViewerProtocolPolicy_ALLOW_ALL ViewerProtocolPolicy = "ALLOW_ALL"
)

