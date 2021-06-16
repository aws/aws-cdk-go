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
// Experimental.
type AddBehaviorOptions struct {
	// HTTP methods to allow for this behavior.
	// Experimental.
	AllowedMethods AllowedMethods `json:"allowedMethods"`
	// HTTP methods to cache for this behavior.
	// Experimental.
	CachedMethods CachedMethods `json:"cachedMethods"`
	// The cache policy for this behavior.
	//
	// The cache policy determines what values are included in the cache key,
	// and the time-to-live (TTL) values for the cache.
	// See: https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-the-cache-key.html.
	//
	// Experimental.
	CachePolicy ICachePolicy `json:"cachePolicy"`
	// Whether you want CloudFront to automatically compress certain files for this cache behavior.
	//
	// See https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/ServingCompressedFiles.html#compressed-content-cloudfront-file-types
	// for file types CloudFront will compress.
	// Experimental.
	Compress *bool `json:"compress"`
	// The Lambda@Edge functions to invoke before serving the contents.
	// See: https://aws.amazon.com/lambda/edge
	//
	// Experimental.
	EdgeLambdas *[]*EdgeLambda `json:"edgeLambdas"`
	// The CloudFront functions to invoke before serving the contents.
	// Experimental.
	FunctionAssociations *[]*FunctionAssociation `json:"functionAssociations"`
	// The origin request policy for this behavior.
	//
	// The origin request policy determines which values (e.g., headers, cookies)
	// are included in requests that CloudFront sends to the origin.
	// Experimental.
	OriginRequestPolicy IOriginRequestPolicy `json:"originRequestPolicy"`
	// Set this to true to indicate you want to distribute media files in the Microsoft Smooth Streaming format using this behavior.
	// Experimental.
	SmoothStreaming *bool `json:"smoothStreaming"`
	// A list of Key Groups that CloudFront can use to validate signed URLs or signed cookies.
	// See: https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/PrivateContent.html
	//
	// Experimental.
	TrustedKeyGroups *[]IKeyGroup `json:"trustedKeyGroups"`
	// The protocol that viewers can use to access the files controlled by this behavior.
	// Experimental.
	ViewerProtocolPolicy ViewerProtocolPolicy `json:"viewerProtocolPolicy"`
}

// The HTTP methods that the Behavior will accept requests on.
// Experimental.
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
// Experimental.
type Behavior struct {
	// The method this CloudFront distribution responds do.
	// Experimental.
	AllowedMethods CloudFrontAllowedMethods `json:"allowedMethods"`
	// Which methods are cached by CloudFront by default.
	// Experimental.
	CachedMethods CloudFrontAllowedCachedMethods `json:"cachedMethods"`
	// If CloudFront should automatically compress some content types.
	// Experimental.
	Compress *bool `json:"compress"`
	// The default amount of time CloudFront will cache an object.
	//
	// This value applies only when your custom origin does not add HTTP headers,
	// such as Cache-Control max-age, Cache-Control s-maxage, and Expires to objects.
	// Experimental.
	DefaultTtl awscdk.Duration `json:"defaultTtl"`
	// The values CloudFront will forward to the origin when making a request.
	// Experimental.
	ForwardedValues *CfnDistribution_ForwardedValuesProperty `json:"forwardedValues"`
	// The CloudFront functions to invoke before serving the contents.
	// Experimental.
	FunctionAssociations *[]*FunctionAssociation `json:"functionAssociations"`
	// If this behavior is the default behavior for the distribution.
	//
	// You must specify exactly one default distribution per CloudFront distribution.
	// The default behavior is allowed to omit the "path" property.
	// Experimental.
	IsDefaultBehavior *bool `json:"isDefaultBehavior"`
	// Declares associated lambda@edge functions for this distribution behaviour.
	// Experimental.
	LambdaFunctionAssociations *[]*LambdaFunctionAssociation `json:"lambdaFunctionAssociations"`
	// The max amount of time you want objects to stay in the cache before CloudFront queries your origin.
	// Experimental.
	MaxTtl awscdk.Duration `json:"maxTtl"`
	// The minimum amount of time that you want objects to stay in the cache before CloudFront queries your origin.
	// Experimental.
	MinTtl awscdk.Duration `json:"minTtl"`
	// The path this behavior responds to.
	//
	// Required for all non-default behaviors. (The default behavior implicitly has "*" as the path pattern. )
	// Experimental.
	PathPattern *string `json:"pathPattern"`
	// A list of Key Groups that CloudFront can use to validate signed URLs or signed cookies.
	// See: https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/PrivateContent.html
	//
	// Experimental.
	TrustedKeyGroups *[]IKeyGroup `json:"trustedKeyGroups"`
}

// Options for creating a new behavior.
// Experimental.
type BehaviorOptions struct {
	// HTTP methods to allow for this behavior.
	// Experimental.
	AllowedMethods AllowedMethods `json:"allowedMethods"`
	// HTTP methods to cache for this behavior.
	// Experimental.
	CachedMethods CachedMethods `json:"cachedMethods"`
	// The cache policy for this behavior.
	//
	// The cache policy determines what values are included in the cache key,
	// and the time-to-live (TTL) values for the cache.
	// See: https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-the-cache-key.html.
	//
	// Experimental.
	CachePolicy ICachePolicy `json:"cachePolicy"`
	// Whether you want CloudFront to automatically compress certain files for this cache behavior.
	//
	// See https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/ServingCompressedFiles.html#compressed-content-cloudfront-file-types
	// for file types CloudFront will compress.
	// Experimental.
	Compress *bool `json:"compress"`
	// The Lambda@Edge functions to invoke before serving the contents.
	// See: https://aws.amazon.com/lambda/edge
	//
	// Experimental.
	EdgeLambdas *[]*EdgeLambda `json:"edgeLambdas"`
	// The CloudFront functions to invoke before serving the contents.
	// Experimental.
	FunctionAssociations *[]*FunctionAssociation `json:"functionAssociations"`
	// The origin request policy for this behavior.
	//
	// The origin request policy determines which values (e.g., headers, cookies)
	// are included in requests that CloudFront sends to the origin.
	// Experimental.
	OriginRequestPolicy IOriginRequestPolicy `json:"originRequestPolicy"`
	// Set this to true to indicate you want to distribute media files in the Microsoft Smooth Streaming format using this behavior.
	// Experimental.
	SmoothStreaming *bool `json:"smoothStreaming"`
	// A list of Key Groups that CloudFront can use to validate signed URLs or signed cookies.
	// See: https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/PrivateContent.html
	//
	// Experimental.
	TrustedKeyGroups *[]IKeyGroup `json:"trustedKeyGroups"`
	// The protocol that viewers can use to access the files controlled by this behavior.
	// Experimental.
	ViewerProtocolPolicy ViewerProtocolPolicy `json:"viewerProtocolPolicy"`
	// The origin that you want CloudFront to route requests to when they match this behavior.
	// Experimental.
	Origin IOrigin `json:"origin"`
}

// Determines whether any cookies in viewer requests are included in the cache key and automatically included in requests that CloudFront sends to the origin.
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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


// Experimental.
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

// Experimental.
func NewCachePolicy_Override(c CachePolicy, scope constructs.Construct, id *string, props *CachePolicyProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront.CachePolicy",
		[]interface{}{scope, id, props},
		c,
	)
}

// Imports a Cache Policy from its id.
// Experimental.
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
// Experimental.
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
// The resource can be deleted (`RemovalPolicy.DELETE`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
// Experimental.
func (c *jsiiProxy_CachePolicy) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
type CachePolicyProps struct {
	// A unique name to identify the cache policy.
	//
	// The name must only include '-', '_', or alphanumeric characters.
	// Experimental.
	CachePolicyName *string `json:"cachePolicyName"`
	// A comment to describe the cache policy.
	// Experimental.
	Comment *string `json:"comment"`
	// Determines whether any cookies in viewer requests are included in the cache key and automatically included in requests that CloudFront sends to the origin.
	// Experimental.
	CookieBehavior CacheCookieBehavior `json:"cookieBehavior"`
	// The default amount of time for objects to stay in the CloudFront cache.
	//
	// Only used when the origin does not send Cache-Control or Expires headers with the object.
	// Experimental.
	DefaultTtl awscdk.Duration `json:"defaultTtl"`
	// Whether to normalize and include the `Accept-Encoding` header in the cache key when the `Accept-Encoding` header is 'br'.
	// Experimental.
	EnableAcceptEncodingBrotli *bool `json:"enableAcceptEncodingBrotli"`
	// Whether to normalize and include the `Accept-Encoding` header in the cache key when the `Accept-Encoding` header is 'gzip'.
	// Experimental.
	EnableAcceptEncodingGzip *bool `json:"enableAcceptEncodingGzip"`
	// Determines whether any HTTP headers are included in the cache key and automatically included in requests that CloudFront sends to the origin.
	// Experimental.
	HeaderBehavior CacheHeaderBehavior `json:"headerBehavior"`
	// The maximum amount of time for objects to stay in the CloudFront cache.
	//
	// CloudFront uses this value only when the origin sends Cache-Control or Expires headers with the object.
	// Experimental.
	MaxTtl awscdk.Duration `json:"maxTtl"`
	// The minimum amount of time for objects to stay in the CloudFront cache.
	// Experimental.
	MinTtl awscdk.Duration `json:"minTtl"`
	// Determines whether any query strings are included in the cache key and automatically included in requests that CloudFront sends to the origin.
	// Experimental.
	QueryStringBehavior CacheQueryStringBehavior `json:"queryStringBehavior"`
}

// Determines whether any URL query strings in viewer requests are included in the cache key and automatically included in requests that CloudFront sends to the origin.
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
func (c *jsiiProxy_CfnCachePolicy) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
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
// Experimental.
func (c *jsiiProxy_CfnCachePolicy) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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

// Experimental.
func (c *jsiiProxy_CfnCachePolicy) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnCachePolicy_CachePolicyConfigProperty struct {
	// `CfnCachePolicy.CachePolicyConfigProperty.DefaultTTL`.
	DefaultTtl *float64 `json:"defaultTtl"`
	// `CfnCachePolicy.CachePolicyConfigProperty.MaxTTL`.
	MaxTtl *float64 `json:"maxTtl"`
	// `CfnCachePolicy.CachePolicyConfigProperty.MinTTL`.
	MinTtl *float64 `json:"minTtl"`
	// `CfnCachePolicy.CachePolicyConfigProperty.Name`.
	Name *string `json:"name"`
	// `CfnCachePolicy.CachePolicyConfigProperty.ParametersInCacheKeyAndForwardedToOrigin`.
	ParametersInCacheKeyAndForwardedToOrigin interface{} `json:"parametersInCacheKeyAndForwardedToOrigin"`
	// `CfnCachePolicy.CachePolicyConfigProperty.Comment`.
	Comment *string `json:"comment"`
}

type CfnCachePolicy_CookiesConfigProperty struct {
	// `CfnCachePolicy.CookiesConfigProperty.CookieBehavior`.
	CookieBehavior *string `json:"cookieBehavior"`
	// `CfnCachePolicy.CookiesConfigProperty.Cookies`.
	Cookies *[]*string `json:"cookies"`
}

type CfnCachePolicy_HeadersConfigProperty struct {
	// `CfnCachePolicy.HeadersConfigProperty.HeaderBehavior`.
	HeaderBehavior *string `json:"headerBehavior"`
	// `CfnCachePolicy.HeadersConfigProperty.Headers`.
	Headers *[]*string `json:"headers"`
}

type CfnCachePolicy_ParametersInCacheKeyAndForwardedToOriginProperty struct {
	// `CfnCachePolicy.ParametersInCacheKeyAndForwardedToOriginProperty.CookiesConfig`.
	CookiesConfig interface{} `json:"cookiesConfig"`
	// `CfnCachePolicy.ParametersInCacheKeyAndForwardedToOriginProperty.EnableAcceptEncodingGzip`.
	EnableAcceptEncodingGzip interface{} `json:"enableAcceptEncodingGzip"`
	// `CfnCachePolicy.ParametersInCacheKeyAndForwardedToOriginProperty.HeadersConfig`.
	HeadersConfig interface{} `json:"headersConfig"`
	// `CfnCachePolicy.ParametersInCacheKeyAndForwardedToOriginProperty.QueryStringsConfig`.
	QueryStringsConfig interface{} `json:"queryStringsConfig"`
	// `CfnCachePolicy.ParametersInCacheKeyAndForwardedToOriginProperty.EnableAcceptEncodingBrotli`.
	EnableAcceptEncodingBrotli interface{} `json:"enableAcceptEncodingBrotli"`
}

type CfnCachePolicy_QueryStringsConfigProperty struct {
	// `CfnCachePolicy.QueryStringsConfigProperty.QueryStringBehavior`.
	QueryStringBehavior *string `json:"queryStringBehavior"`
	// `CfnCachePolicy.QueryStringsConfigProperty.QueryStrings`.
	QueryStrings *[]*string `json:"queryStrings"`
}

// Properties for defining a `AWS::CloudFront::CachePolicy`.
type CfnCachePolicyProps struct {
	// `AWS::CloudFront::CachePolicy.CachePolicyConfig`.
	CachePolicyConfig interface{} `json:"cachePolicyConfig"`
}

// A CloudFormation `AWS::CloudFront::CloudFrontOriginAccessIdentity`.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
func (c *jsiiProxy_CfnCloudFrontOriginAccessIdentity) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
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
// Experimental.
func (c *jsiiProxy_CfnCloudFrontOriginAccessIdentity) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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

// Experimental.
func (c *jsiiProxy_CfnCloudFrontOriginAccessIdentity) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnCloudFrontOriginAccessIdentity_CloudFrontOriginAccessIdentityConfigProperty struct {
	// `CfnCloudFrontOriginAccessIdentity.CloudFrontOriginAccessIdentityConfigProperty.Comment`.
	Comment *string `json:"comment"`
}

// Properties for defining a `AWS::CloudFront::CloudFrontOriginAccessIdentity`.
type CfnCloudFrontOriginAccessIdentityProps struct {
	// `AWS::CloudFront::CloudFrontOriginAccessIdentity.CloudFrontOriginAccessIdentityConfig`.
	CloudFrontOriginAccessIdentityConfig interface{} `json:"cloudFrontOriginAccessIdentityConfig"`
}

// A CloudFormation `AWS::CloudFront::Distribution`.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
func (c *jsiiProxy_CfnDistribution) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
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
// Experimental.
func (c *jsiiProxy_CfnDistribution) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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

// Experimental.
func (c *jsiiProxy_CfnDistribution) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnDistribution_CacheBehaviorProperty struct {
	// `CfnDistribution.CacheBehaviorProperty.PathPattern`.
	PathPattern *string `json:"pathPattern"`
	// `CfnDistribution.CacheBehaviorProperty.TargetOriginId`.
	TargetOriginId *string `json:"targetOriginId"`
	// `CfnDistribution.CacheBehaviorProperty.ViewerProtocolPolicy`.
	ViewerProtocolPolicy *string `json:"viewerProtocolPolicy"`
	// `CfnDistribution.CacheBehaviorProperty.AllowedMethods`.
	AllowedMethods *[]*string `json:"allowedMethods"`
	// `CfnDistribution.CacheBehaviorProperty.CachedMethods`.
	CachedMethods *[]*string `json:"cachedMethods"`
	// `CfnDistribution.CacheBehaviorProperty.CachePolicyId`.
	CachePolicyId *string `json:"cachePolicyId"`
	// `CfnDistribution.CacheBehaviorProperty.Compress`.
	Compress interface{} `json:"compress"`
	// `CfnDistribution.CacheBehaviorProperty.DefaultTTL`.
	DefaultTtl *float64 `json:"defaultTtl"`
	// `CfnDistribution.CacheBehaviorProperty.FieldLevelEncryptionId`.
	FieldLevelEncryptionId *string `json:"fieldLevelEncryptionId"`
	// `CfnDistribution.CacheBehaviorProperty.ForwardedValues`.
	ForwardedValues interface{} `json:"forwardedValues"`
	// `CfnDistribution.CacheBehaviorProperty.FunctionAssociations`.
	FunctionAssociations interface{} `json:"functionAssociations"`
	// `CfnDistribution.CacheBehaviorProperty.LambdaFunctionAssociations`.
	LambdaFunctionAssociations interface{} `json:"lambdaFunctionAssociations"`
	// `CfnDistribution.CacheBehaviorProperty.MaxTTL`.
	MaxTtl *float64 `json:"maxTtl"`
	// `CfnDistribution.CacheBehaviorProperty.MinTTL`.
	MinTtl *float64 `json:"minTtl"`
	// `CfnDistribution.CacheBehaviorProperty.OriginRequestPolicyId`.
	OriginRequestPolicyId *string `json:"originRequestPolicyId"`
	// `CfnDistribution.CacheBehaviorProperty.RealtimeLogConfigArn`.
	RealtimeLogConfigArn *string `json:"realtimeLogConfigArn"`
	// `CfnDistribution.CacheBehaviorProperty.SmoothStreaming`.
	SmoothStreaming interface{} `json:"smoothStreaming"`
	// `CfnDistribution.CacheBehaviorProperty.TrustedKeyGroups`.
	TrustedKeyGroups *[]*string `json:"trustedKeyGroups"`
	// `CfnDistribution.CacheBehaviorProperty.TrustedSigners`.
	TrustedSigners *[]*string `json:"trustedSigners"`
}

type CfnDistribution_CookiesProperty struct {
	// `CfnDistribution.CookiesProperty.Forward`.
	Forward *string `json:"forward"`
	// `CfnDistribution.CookiesProperty.WhitelistedNames`.
	WhitelistedNames *[]*string `json:"whitelistedNames"`
}

type CfnDistribution_CustomErrorResponseProperty struct {
	// `CfnDistribution.CustomErrorResponseProperty.ErrorCode`.
	ErrorCode *float64 `json:"errorCode"`
	// `CfnDistribution.CustomErrorResponseProperty.ErrorCachingMinTTL`.
	ErrorCachingMinTtl *float64 `json:"errorCachingMinTtl"`
	// `CfnDistribution.CustomErrorResponseProperty.ResponseCode`.
	ResponseCode *float64 `json:"responseCode"`
	// `CfnDistribution.CustomErrorResponseProperty.ResponsePagePath`.
	ResponsePagePath *string `json:"responsePagePath"`
}

type CfnDistribution_CustomOriginConfigProperty struct {
	// `CfnDistribution.CustomOriginConfigProperty.OriginProtocolPolicy`.
	OriginProtocolPolicy *string `json:"originProtocolPolicy"`
	// `CfnDistribution.CustomOriginConfigProperty.HTTPPort`.
	HttpPort *float64 `json:"httpPort"`
	// `CfnDistribution.CustomOriginConfigProperty.HTTPSPort`.
	HttpsPort *float64 `json:"httpsPort"`
	// `CfnDistribution.CustomOriginConfigProperty.OriginKeepaliveTimeout`.
	OriginKeepaliveTimeout *float64 `json:"originKeepaliveTimeout"`
	// `CfnDistribution.CustomOriginConfigProperty.OriginReadTimeout`.
	OriginReadTimeout *float64 `json:"originReadTimeout"`
	// `CfnDistribution.CustomOriginConfigProperty.OriginSSLProtocols`.
	OriginSslProtocols *[]*string `json:"originSslProtocols"`
}

type CfnDistribution_DefaultCacheBehaviorProperty struct {
	// `CfnDistribution.DefaultCacheBehaviorProperty.TargetOriginId`.
	TargetOriginId *string `json:"targetOriginId"`
	// `CfnDistribution.DefaultCacheBehaviorProperty.ViewerProtocolPolicy`.
	ViewerProtocolPolicy *string `json:"viewerProtocolPolicy"`
	// `CfnDistribution.DefaultCacheBehaviorProperty.AllowedMethods`.
	AllowedMethods *[]*string `json:"allowedMethods"`
	// `CfnDistribution.DefaultCacheBehaviorProperty.CachedMethods`.
	CachedMethods *[]*string `json:"cachedMethods"`
	// `CfnDistribution.DefaultCacheBehaviorProperty.CachePolicyId`.
	CachePolicyId *string `json:"cachePolicyId"`
	// `CfnDistribution.DefaultCacheBehaviorProperty.Compress`.
	Compress interface{} `json:"compress"`
	// `CfnDistribution.DefaultCacheBehaviorProperty.DefaultTTL`.
	DefaultTtl *float64 `json:"defaultTtl"`
	// `CfnDistribution.DefaultCacheBehaviorProperty.FieldLevelEncryptionId`.
	FieldLevelEncryptionId *string `json:"fieldLevelEncryptionId"`
	// `CfnDistribution.DefaultCacheBehaviorProperty.ForwardedValues`.
	ForwardedValues interface{} `json:"forwardedValues"`
	// `CfnDistribution.DefaultCacheBehaviorProperty.FunctionAssociations`.
	FunctionAssociations interface{} `json:"functionAssociations"`
	// `CfnDistribution.DefaultCacheBehaviorProperty.LambdaFunctionAssociations`.
	LambdaFunctionAssociations interface{} `json:"lambdaFunctionAssociations"`
	// `CfnDistribution.DefaultCacheBehaviorProperty.MaxTTL`.
	MaxTtl *float64 `json:"maxTtl"`
	// `CfnDistribution.DefaultCacheBehaviorProperty.MinTTL`.
	MinTtl *float64 `json:"minTtl"`
	// `CfnDistribution.DefaultCacheBehaviorProperty.OriginRequestPolicyId`.
	OriginRequestPolicyId *string `json:"originRequestPolicyId"`
	// `CfnDistribution.DefaultCacheBehaviorProperty.RealtimeLogConfigArn`.
	RealtimeLogConfigArn *string `json:"realtimeLogConfigArn"`
	// `CfnDistribution.DefaultCacheBehaviorProperty.SmoothStreaming`.
	SmoothStreaming interface{} `json:"smoothStreaming"`
	// `CfnDistribution.DefaultCacheBehaviorProperty.TrustedKeyGroups`.
	TrustedKeyGroups *[]*string `json:"trustedKeyGroups"`
	// `CfnDistribution.DefaultCacheBehaviorProperty.TrustedSigners`.
	TrustedSigners *[]*string `json:"trustedSigners"`
}

type CfnDistribution_DistributionConfigProperty struct {
	// `CfnDistribution.DistributionConfigProperty.Enabled`.
	Enabled interface{} `json:"enabled"`
	// `CfnDistribution.DistributionConfigProperty.Aliases`.
	Aliases *[]*string `json:"aliases"`
	// `CfnDistribution.DistributionConfigProperty.CacheBehaviors`.
	CacheBehaviors interface{} `json:"cacheBehaviors"`
	// `CfnDistribution.DistributionConfigProperty.CNAMEs`.
	CnamEs *[]*string `json:"cnamEs"`
	// `CfnDistribution.DistributionConfigProperty.Comment`.
	Comment *string `json:"comment"`
	// `CfnDistribution.DistributionConfigProperty.CustomErrorResponses`.
	CustomErrorResponses interface{} `json:"customErrorResponses"`
	// `CfnDistribution.DistributionConfigProperty.CustomOrigin`.
	CustomOrigin interface{} `json:"customOrigin"`
	// `CfnDistribution.DistributionConfigProperty.DefaultCacheBehavior`.
	DefaultCacheBehavior interface{} `json:"defaultCacheBehavior"`
	// `CfnDistribution.DistributionConfigProperty.DefaultRootObject`.
	DefaultRootObject *string `json:"defaultRootObject"`
	// `CfnDistribution.DistributionConfigProperty.HttpVersion`.
	HttpVersion *string `json:"httpVersion"`
	// `CfnDistribution.DistributionConfigProperty.IPV6Enabled`.
	Ipv6Enabled interface{} `json:"ipv6Enabled"`
	// `CfnDistribution.DistributionConfigProperty.Logging`.
	Logging interface{} `json:"logging"`
	// `CfnDistribution.DistributionConfigProperty.OriginGroups`.
	OriginGroups interface{} `json:"originGroups"`
	// `CfnDistribution.DistributionConfigProperty.Origins`.
	Origins interface{} `json:"origins"`
	// `CfnDistribution.DistributionConfigProperty.PriceClass`.
	PriceClass *string `json:"priceClass"`
	// `CfnDistribution.DistributionConfigProperty.Restrictions`.
	Restrictions interface{} `json:"restrictions"`
	// `CfnDistribution.DistributionConfigProperty.S3Origin`.
	S3Origin interface{} `json:"s3Origin"`
	// `CfnDistribution.DistributionConfigProperty.ViewerCertificate`.
	ViewerCertificate interface{} `json:"viewerCertificate"`
	// `CfnDistribution.DistributionConfigProperty.WebACLId`.
	WebAclId *string `json:"webAclId"`
}

type CfnDistribution_ForwardedValuesProperty struct {
	// `CfnDistribution.ForwardedValuesProperty.QueryString`.
	QueryString interface{} `json:"queryString"`
	// `CfnDistribution.ForwardedValuesProperty.Cookies`.
	Cookies interface{} `json:"cookies"`
	// `CfnDistribution.ForwardedValuesProperty.Headers`.
	Headers *[]*string `json:"headers"`
	// `CfnDistribution.ForwardedValuesProperty.QueryStringCacheKeys`.
	QueryStringCacheKeys *[]*string `json:"queryStringCacheKeys"`
}

type CfnDistribution_FunctionAssociationProperty struct {
	// `CfnDistribution.FunctionAssociationProperty.EventType`.
	EventType *string `json:"eventType"`
	// `CfnDistribution.FunctionAssociationProperty.FunctionARN`.
	FunctionArn *string `json:"functionArn"`
}

type CfnDistribution_GeoRestrictionProperty struct {
	// `CfnDistribution.GeoRestrictionProperty.RestrictionType`.
	RestrictionType *string `json:"restrictionType"`
	// `CfnDistribution.GeoRestrictionProperty.Locations`.
	Locations *[]*string `json:"locations"`
}

type CfnDistribution_LambdaFunctionAssociationProperty struct {
	// `CfnDistribution.LambdaFunctionAssociationProperty.EventType`.
	EventType *string `json:"eventType"`
	// `CfnDistribution.LambdaFunctionAssociationProperty.IncludeBody`.
	IncludeBody interface{} `json:"includeBody"`
	// `CfnDistribution.LambdaFunctionAssociationProperty.LambdaFunctionARN`.
	LambdaFunctionArn *string `json:"lambdaFunctionArn"`
}

type CfnDistribution_LegacyCustomOriginProperty struct {
	// `CfnDistribution.LegacyCustomOriginProperty.DNSName`.
	DnsName *string `json:"dnsName"`
	// `CfnDistribution.LegacyCustomOriginProperty.OriginProtocolPolicy`.
	OriginProtocolPolicy *string `json:"originProtocolPolicy"`
	// `CfnDistribution.LegacyCustomOriginProperty.OriginSSLProtocols`.
	OriginSslProtocols *[]*string `json:"originSslProtocols"`
	// `CfnDistribution.LegacyCustomOriginProperty.HTTPPort`.
	HttpPort *float64 `json:"httpPort"`
	// `CfnDistribution.LegacyCustomOriginProperty.HTTPSPort`.
	HttpsPort *float64 `json:"httpsPort"`
}

type CfnDistribution_LegacyS3OriginProperty struct {
	// `CfnDistribution.LegacyS3OriginProperty.DNSName`.
	DnsName *string `json:"dnsName"`
	// `CfnDistribution.LegacyS3OriginProperty.OriginAccessIdentity`.
	OriginAccessIdentity *string `json:"originAccessIdentity"`
}

type CfnDistribution_LoggingProperty struct {
	// `CfnDistribution.LoggingProperty.Bucket`.
	Bucket *string `json:"bucket"`
	// `CfnDistribution.LoggingProperty.IncludeCookies`.
	IncludeCookies interface{} `json:"includeCookies"`
	// `CfnDistribution.LoggingProperty.Prefix`.
	Prefix *string `json:"prefix"`
}

type CfnDistribution_OriginCustomHeaderProperty struct {
	// `CfnDistribution.OriginCustomHeaderProperty.HeaderName`.
	HeaderName *string `json:"headerName"`
	// `CfnDistribution.OriginCustomHeaderProperty.HeaderValue`.
	HeaderValue *string `json:"headerValue"`
}

type CfnDistribution_OriginGroupFailoverCriteriaProperty struct {
	// `CfnDistribution.OriginGroupFailoverCriteriaProperty.StatusCodes`.
	StatusCodes interface{} `json:"statusCodes"`
}

type CfnDistribution_OriginGroupMemberProperty struct {
	// `CfnDistribution.OriginGroupMemberProperty.OriginId`.
	OriginId *string `json:"originId"`
}

type CfnDistribution_OriginGroupMembersProperty struct {
	// `CfnDistribution.OriginGroupMembersProperty.Items`.
	Items interface{} `json:"items"`
	// `CfnDistribution.OriginGroupMembersProperty.Quantity`.
	Quantity *float64 `json:"quantity"`
}

type CfnDistribution_OriginGroupProperty struct {
	// `CfnDistribution.OriginGroupProperty.FailoverCriteria`.
	FailoverCriteria interface{} `json:"failoverCriteria"`
	// `CfnDistribution.OriginGroupProperty.Id`.
	Id *string `json:"id"`
	// `CfnDistribution.OriginGroupProperty.Members`.
	Members interface{} `json:"members"`
}

type CfnDistribution_OriginGroupsProperty struct {
	// `CfnDistribution.OriginGroupsProperty.Quantity`.
	Quantity *float64 `json:"quantity"`
	// `CfnDistribution.OriginGroupsProperty.Items`.
	Items interface{} `json:"items"`
}

type CfnDistribution_OriginProperty struct {
	// `CfnDistribution.OriginProperty.DomainName`.
	DomainName *string `json:"domainName"`
	// `CfnDistribution.OriginProperty.Id`.
	Id *string `json:"id"`
	// `CfnDistribution.OriginProperty.ConnectionAttempts`.
	ConnectionAttempts *float64 `json:"connectionAttempts"`
	// `CfnDistribution.OriginProperty.ConnectionTimeout`.
	ConnectionTimeout *float64 `json:"connectionTimeout"`
	// `CfnDistribution.OriginProperty.CustomOriginConfig`.
	CustomOriginConfig interface{} `json:"customOriginConfig"`
	// `CfnDistribution.OriginProperty.OriginCustomHeaders`.
	OriginCustomHeaders interface{} `json:"originCustomHeaders"`
	// `CfnDistribution.OriginProperty.OriginPath`.
	OriginPath *string `json:"originPath"`
	// `CfnDistribution.OriginProperty.OriginShield`.
	OriginShield interface{} `json:"originShield"`
	// `CfnDistribution.OriginProperty.S3OriginConfig`.
	S3OriginConfig interface{} `json:"s3OriginConfig"`
}

type CfnDistribution_OriginShieldProperty struct {
	// `CfnDistribution.OriginShieldProperty.Enabled`.
	Enabled interface{} `json:"enabled"`
	// `CfnDistribution.OriginShieldProperty.OriginShieldRegion`.
	OriginShieldRegion *string `json:"originShieldRegion"`
}

type CfnDistribution_RestrictionsProperty struct {
	// `CfnDistribution.RestrictionsProperty.GeoRestriction`.
	GeoRestriction interface{} `json:"geoRestriction"`
}

type CfnDistribution_S3OriginConfigProperty struct {
	// `CfnDistribution.S3OriginConfigProperty.OriginAccessIdentity`.
	OriginAccessIdentity *string `json:"originAccessIdentity"`
}

type CfnDistribution_StatusCodesProperty struct {
	// `CfnDistribution.StatusCodesProperty.Items`.
	Items interface{} `json:"items"`
	// `CfnDistribution.StatusCodesProperty.Quantity`.
	Quantity *float64 `json:"quantity"`
}

type CfnDistribution_ViewerCertificateProperty struct {
	// `CfnDistribution.ViewerCertificateProperty.AcmCertificateArn`.
	AcmCertificateArn *string `json:"acmCertificateArn"`
	// `CfnDistribution.ViewerCertificateProperty.CloudFrontDefaultCertificate`.
	CloudFrontDefaultCertificate interface{} `json:"cloudFrontDefaultCertificate"`
	// `CfnDistribution.ViewerCertificateProperty.IamCertificateId`.
	IamCertificateId *string `json:"iamCertificateId"`
	// `CfnDistribution.ViewerCertificateProperty.MinimumProtocolVersion`.
	MinimumProtocolVersion *string `json:"minimumProtocolVersion"`
	// `CfnDistribution.ViewerCertificateProperty.SslSupportMethod`.
	SslSupportMethod *string `json:"sslSupportMethod"`
}

// Properties for defining a `AWS::CloudFront::Distribution`.
type CfnDistributionProps struct {
	// `AWS::CloudFront::Distribution.DistributionConfig`.
	DistributionConfig interface{} `json:"distributionConfig"`
	// `AWS::CloudFront::Distribution.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags"`
}

// A CloudFormation `AWS::CloudFront::Function`.
type CfnFunction interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrFunctionArn() *string
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
	FunctionMetadata() interface{}
	SetFunctionMetadata(val interface{})
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

func (j *jsiiProxy_CfnFunction) FunctionMetadata() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"functionMetadata",
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

func (j *jsiiProxy_CfnFunction) SetFunctionMetadata(val interface{}) {
	_jsii_.Set(
		j,
		"functionMetadata",
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
func (c *jsiiProxy_CfnFunction) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
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
// Experimental.
func (c *jsiiProxy_CfnFunction) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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

// Experimental.
func (c *jsiiProxy_CfnFunction) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnFunction_FunctionConfigProperty struct {
	// `CfnFunction.FunctionConfigProperty.Comment`.
	Comment *string `json:"comment"`
	// `CfnFunction.FunctionConfigProperty.Runtime`.
	Runtime *string `json:"runtime"`
}

type CfnFunction_FunctionMetadataProperty struct {
	// `CfnFunction.FunctionMetadataProperty.FunctionARN`.
	FunctionArn *string `json:"functionArn"`
}

// Properties for defining a `AWS::CloudFront::Function`.
type CfnFunctionProps struct {
	// `AWS::CloudFront::Function.Name`.
	Name *string `json:"name"`
	// `AWS::CloudFront::Function.AutoPublish`.
	AutoPublish interface{} `json:"autoPublish"`
	// `AWS::CloudFront::Function.FunctionCode`.
	FunctionCode *string `json:"functionCode"`
	// `AWS::CloudFront::Function.FunctionConfig`.
	FunctionConfig interface{} `json:"functionConfig"`
	// `AWS::CloudFront::Function.FunctionMetadata`.
	FunctionMetadata interface{} `json:"functionMetadata"`
}

// A CloudFormation `AWS::CloudFront::KeyGroup`.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
func (c *jsiiProxy_CfnKeyGroup) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
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
// Experimental.
func (c *jsiiProxy_CfnKeyGroup) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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

// Experimental.
func (c *jsiiProxy_CfnKeyGroup) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnKeyGroup_KeyGroupConfigProperty struct {
	// `CfnKeyGroup.KeyGroupConfigProperty.Items`.
	Items *[]*string `json:"items"`
	// `CfnKeyGroup.KeyGroupConfigProperty.Name`.
	Name *string `json:"name"`
	// `CfnKeyGroup.KeyGroupConfigProperty.Comment`.
	Comment *string `json:"comment"`
}

// Properties for defining a `AWS::CloudFront::KeyGroup`.
type CfnKeyGroupProps struct {
	// `AWS::CloudFront::KeyGroup.KeyGroupConfig`.
	KeyGroupConfig interface{} `json:"keyGroupConfig"`
}

// A CloudFormation `AWS::CloudFront::OriginRequestPolicy`.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
func (c *jsiiProxy_CfnOriginRequestPolicy) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
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
// Experimental.
func (c *jsiiProxy_CfnOriginRequestPolicy) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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

// Experimental.
func (c *jsiiProxy_CfnOriginRequestPolicy) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnOriginRequestPolicy_CookiesConfigProperty struct {
	// `CfnOriginRequestPolicy.CookiesConfigProperty.CookieBehavior`.
	CookieBehavior *string `json:"cookieBehavior"`
	// `CfnOriginRequestPolicy.CookiesConfigProperty.Cookies`.
	Cookies *[]*string `json:"cookies"`
}

type CfnOriginRequestPolicy_HeadersConfigProperty struct {
	// `CfnOriginRequestPolicy.HeadersConfigProperty.HeaderBehavior`.
	HeaderBehavior *string `json:"headerBehavior"`
	// `CfnOriginRequestPolicy.HeadersConfigProperty.Headers`.
	Headers *[]*string `json:"headers"`
}

type CfnOriginRequestPolicy_OriginRequestPolicyConfigProperty struct {
	// `CfnOriginRequestPolicy.OriginRequestPolicyConfigProperty.CookiesConfig`.
	CookiesConfig interface{} `json:"cookiesConfig"`
	// `CfnOriginRequestPolicy.OriginRequestPolicyConfigProperty.HeadersConfig`.
	HeadersConfig interface{} `json:"headersConfig"`
	// `CfnOriginRequestPolicy.OriginRequestPolicyConfigProperty.Name`.
	Name *string `json:"name"`
	// `CfnOriginRequestPolicy.OriginRequestPolicyConfigProperty.QueryStringsConfig`.
	QueryStringsConfig interface{} `json:"queryStringsConfig"`
	// `CfnOriginRequestPolicy.OriginRequestPolicyConfigProperty.Comment`.
	Comment *string `json:"comment"`
}

type CfnOriginRequestPolicy_QueryStringsConfigProperty struct {
	// `CfnOriginRequestPolicy.QueryStringsConfigProperty.QueryStringBehavior`.
	QueryStringBehavior *string `json:"queryStringBehavior"`
	// `CfnOriginRequestPolicy.QueryStringsConfigProperty.QueryStrings`.
	QueryStrings *[]*string `json:"queryStrings"`
}

// Properties for defining a `AWS::CloudFront::OriginRequestPolicy`.
type CfnOriginRequestPolicyProps struct {
	// `AWS::CloudFront::OriginRequestPolicy.OriginRequestPolicyConfig`.
	OriginRequestPolicyConfig interface{} `json:"originRequestPolicyConfig"`
}

// A CloudFormation `AWS::CloudFront::PublicKey`.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
func (c *jsiiProxy_CfnPublicKey) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
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
// Experimental.
func (c *jsiiProxy_CfnPublicKey) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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

// Experimental.
func (c *jsiiProxy_CfnPublicKey) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnPublicKey_PublicKeyConfigProperty struct {
	// `CfnPublicKey.PublicKeyConfigProperty.CallerReference`.
	CallerReference *string `json:"callerReference"`
	// `CfnPublicKey.PublicKeyConfigProperty.EncodedKey`.
	EncodedKey *string `json:"encodedKey"`
	// `CfnPublicKey.PublicKeyConfigProperty.Name`.
	Name *string `json:"name"`
	// `CfnPublicKey.PublicKeyConfigProperty.Comment`.
	Comment *string `json:"comment"`
}

// Properties for defining a `AWS::CloudFront::PublicKey`.
type CfnPublicKeyProps struct {
	// `AWS::CloudFront::PublicKey.PublicKeyConfig`.
	PublicKeyConfig interface{} `json:"publicKeyConfig"`
}

// A CloudFormation `AWS::CloudFront::RealtimeLogConfig`.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
func (c *jsiiProxy_CfnRealtimeLogConfig) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
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
// Experimental.
func (c *jsiiProxy_CfnRealtimeLogConfig) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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

// Experimental.
func (c *jsiiProxy_CfnRealtimeLogConfig) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnRealtimeLogConfig_EndPointProperty struct {
	// `CfnRealtimeLogConfig.EndPointProperty.KinesisStreamConfig`.
	KinesisStreamConfig interface{} `json:"kinesisStreamConfig"`
	// `CfnRealtimeLogConfig.EndPointProperty.StreamType`.
	StreamType *string `json:"streamType"`
}

type CfnRealtimeLogConfig_KinesisStreamConfigProperty struct {
	// `CfnRealtimeLogConfig.KinesisStreamConfigProperty.RoleArn`.
	RoleArn *string `json:"roleArn"`
	// `CfnRealtimeLogConfig.KinesisStreamConfigProperty.StreamArn`.
	StreamArn *string `json:"streamArn"`
}

// Properties for defining a `AWS::CloudFront::RealtimeLogConfig`.
type CfnRealtimeLogConfigProps struct {
	// `AWS::CloudFront::RealtimeLogConfig.EndPoints`.
	EndPoints interface{} `json:"endPoints"`
	// `AWS::CloudFront::RealtimeLogConfig.Fields`.
	Fields *[]*string `json:"fields"`
	// `AWS::CloudFront::RealtimeLogConfig.Name`.
	Name *string `json:"name"`
	// `AWS::CloudFront::RealtimeLogConfig.SamplingRate`.
	SamplingRate *float64 `json:"samplingRate"`
}

// A CloudFormation `AWS::CloudFront::StreamingDistribution`.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
func (c *jsiiProxy_CfnStreamingDistribution) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
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
// Experimental.
func (c *jsiiProxy_CfnStreamingDistribution) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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

// Experimental.
func (c *jsiiProxy_CfnStreamingDistribution) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnStreamingDistribution_LoggingProperty struct {
	// `CfnStreamingDistribution.LoggingProperty.Bucket`.
	Bucket *string `json:"bucket"`
	// `CfnStreamingDistribution.LoggingProperty.Enabled`.
	Enabled interface{} `json:"enabled"`
	// `CfnStreamingDistribution.LoggingProperty.Prefix`.
	Prefix *string `json:"prefix"`
}

type CfnStreamingDistribution_S3OriginProperty struct {
	// `CfnStreamingDistribution.S3OriginProperty.DomainName`.
	DomainName *string `json:"domainName"`
	// `CfnStreamingDistribution.S3OriginProperty.OriginAccessIdentity`.
	OriginAccessIdentity *string `json:"originAccessIdentity"`
}

type CfnStreamingDistribution_StreamingDistributionConfigProperty struct {
	// `CfnStreamingDistribution.StreamingDistributionConfigProperty.Comment`.
	Comment *string `json:"comment"`
	// `CfnStreamingDistribution.StreamingDistributionConfigProperty.Enabled`.
	Enabled interface{} `json:"enabled"`
	// `CfnStreamingDistribution.StreamingDistributionConfigProperty.S3Origin`.
	S3Origin interface{} `json:"s3Origin"`
	// `CfnStreamingDistribution.StreamingDistributionConfigProperty.TrustedSigners`.
	TrustedSigners interface{} `json:"trustedSigners"`
	// `CfnStreamingDistribution.StreamingDistributionConfigProperty.Aliases`.
	Aliases *[]*string `json:"aliases"`
	// `CfnStreamingDistribution.StreamingDistributionConfigProperty.Logging`.
	Logging interface{} `json:"logging"`
	// `CfnStreamingDistribution.StreamingDistributionConfigProperty.PriceClass`.
	PriceClass *string `json:"priceClass"`
}

type CfnStreamingDistribution_TrustedSignersProperty struct {
	// `CfnStreamingDistribution.TrustedSignersProperty.Enabled`.
	Enabled interface{} `json:"enabled"`
	// `CfnStreamingDistribution.TrustedSignersProperty.AwsAccountNumbers`.
	AwsAccountNumbers *[]*string `json:"awsAccountNumbers"`
}

// Properties for defining a `AWS::CloudFront::StreamingDistribution`.
type CfnStreamingDistributionProps struct {
	// `AWS::CloudFront::StreamingDistribution.StreamingDistributionConfig`.
	StreamingDistributionConfig interface{} `json:"streamingDistributionConfig"`
	// `AWS::CloudFront::StreamingDistribution.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags"`
}

// Enums for the methods CloudFront can cache.
// Experimental.
type CloudFrontAllowedCachedMethods string

const (
	CloudFrontAllowedCachedMethods_GET_HEAD CloudFrontAllowedCachedMethods = "GET_HEAD"
	CloudFrontAllowedCachedMethods_GET_HEAD_OPTIONS CloudFrontAllowedCachedMethods = "GET_HEAD_OPTIONS"
)

// An enum for the supported methods to a CloudFront distribution.
// Experimental.
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
// import { CloudFrontWebDistribution } from '@aws-cdk/aws-cloudfront'
//
// const sourceBucket = new Bucket(this, 'Bucket');
//
// const distribution = new CloudFrontWebDistribution(this, 'MyDistribution', {
//   originConfigs: [
//     {
//       s3OriginSource: {
//       s3BucketSource: sourceBucket
//       },
//       behaviors : [ {isDefaultBehavior: true}]
//     }
//   ]
// });
// ```
//
// This will create a CloudFront distribution that uses your S3Bucket as it's origin.
//
// You can customize the distribution using additional properties from the CloudFrontWebDistributionProps interface.
// Experimental.
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


// Experimental.
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

// Experimental.
func NewCloudFrontWebDistribution_Override(c CloudFrontWebDistribution, scope constructs.Construct, id *string, props *CloudFrontWebDistributionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront.CloudFrontWebDistribution",
		[]interface{}{scope, id, props},
		c,
	)
}

// Creates a construct that represents an external (imported) distribution.
// Experimental.
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
// Experimental.
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
// The resource can be deleted (`RemovalPolicy.DELETE`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
// Experimental.
func (c *jsiiProxy_CloudFrontWebDistribution) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
type CloudFrontWebDistributionAttributes struct {
	// The distribution ID for this distribution.
	// Experimental.
	DistributionId *string `json:"distributionId"`
	// The generated domain name of the Distribution, such as d111111abcdef8.cloudfront.net.
	// Experimental.
	DomainName *string `json:"domainName"`
}

// Experimental.
type CloudFrontWebDistributionProps struct {
	// The origin configurations for this distribution.
	//
	// Behaviors are a part of the origin.
	// Experimental.
	OriginConfigs *[]*SourceConfiguration `json:"originConfigs"`
	// A comment for this distribution in the CloudFront console.
	// Experimental.
	Comment *string `json:"comment"`
	// The default object to serve.
	// Experimental.
	DefaultRootObject *string `json:"defaultRootObject"`
	// If your distribution should have IPv6 enabled.
	// Experimental.
	EnableIpV6 *bool `json:"enableIpV6"`
	// How CloudFront should handle requests that are not successful (eg PageNotFound).
	//
	// By default, CloudFront does not replace HTTP status codes in the 4xx and 5xx range
	// with custom error messages. CloudFront does not cache HTTP status codes.
	// Experimental.
	ErrorConfigurations *[]*CfnDistribution_CustomErrorResponseProperty `json:"errorConfigurations"`
	// Controls the countries in which your content is distributed.
	// Experimental.
	GeoRestriction GeoRestriction `json:"geoRestriction"`
	// The max supported HTTP Versions.
	// Experimental.
	HttpVersion HttpVersion `json:"httpVersion"`
	// Optional - if we should enable logging.
	//
	// You can pass an empty object ({}) to have us auto create a bucket for logging.
	// Omission of this property indicates no logging is to be enabled.
	// Experimental.
	LoggingConfig *LoggingConfiguration `json:"loggingConfig"`
	// The price class for the distribution (this impacts how many locations CloudFront uses for your distribution, and billing).
	// Experimental.
	PriceClass PriceClass `json:"priceClass"`
	// Specifies whether you want viewers to use HTTP or HTTPS to request your objects, whether you're using an alternate domain name with HTTPS, and if so, if you're using AWS Certificate Manager (ACM) or a third-party certificate authority.
	// See: https://aws.amazon.com/premiumsupport/knowledge-center/custom-ssl-certificate-cloudfront/
	//
	// Experimental.
	ViewerCertificate ViewerCertificate `json:"viewerCertificate"`
	// The default viewer policy for incoming clients.
	// Experimental.
	ViewerProtocolPolicy ViewerProtocolPolicy `json:"viewerProtocolPolicy"`
	// Unique identifier that specifies the AWS WAF web ACL to associate with this CloudFront distribution.
	//
	// To specify a web ACL created using the latest version of AWS WAF, use the ACL ARN, for example
	// `arn:aws:wafv2:us-east-1:123456789012:global/webacl/ExampleWebACL/473e64fd-f30b-4765-81a0-62ad96dd167a`.
	//
	// To specify a web ACL created using AWS WAF Classic, use the ACL ID, for example `473e64fd-f30b-4765-81a0-62ad96dd167a`.
	// See: https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_CreateDistribution.html#API_CreateDistribution_RequestParameters.
	//
	// Experimental.
	WebACLId *string `json:"webACLId"`
}

// A custom origin configuration.
// Experimental.
type CustomOriginConfig struct {
	// The domain name of the custom origin.
	//
	// Should not include the path - that should be in the parent SourceConfiguration
	// Experimental.
	DomainName *string `json:"domainName"`
	// The SSL versions to use when interacting with the origin.
	// Experimental.
	AllowedOriginSSLVersions *[]OriginSslPolicy `json:"allowedOriginSSLVersions"`
	// The origin HTTP port.
	// Experimental.
	HttpPort *float64 `json:"httpPort"`
	// The origin HTTPS port.
	// Experimental.
	HttpsPort *float64 `json:"httpsPort"`
	// Any additional headers to pass to the origin.
	// Experimental.
	OriginHeaders *map[string]*string `json:"originHeaders"`
	// The keep alive timeout when making calls in seconds.
	// Experimental.
	OriginKeepaliveTimeout awscdk.Duration `json:"originKeepaliveTimeout"`
	// The relative path to the origin root to use for sources.
	// Experimental.
	OriginPath *string `json:"originPath"`
	// The protocol (http or https) policy to use when interacting with the origin.
	// Experimental.
	OriginProtocolPolicy OriginProtocolPolicy `json:"originProtocolPolicy"`
	// The read timeout when calling the origin in seconds.
	// Experimental.
	OriginReadTimeout awscdk.Duration `json:"originReadTimeout"`
}

// A CloudFront distribution with associated origin(s) and caching behavior(s).
// Experimental.
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


// Experimental.
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

// Experimental.
func NewDistribution_Override(d Distribution, scope constructs.Construct, id *string, props *DistributionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront.Distribution",
		[]interface{}{scope, id, props},
		d,
	)
}

// Creates a Distribution construct that represents an external (imported) distribution.
// Experimental.
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
// Experimental.
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
// Experimental.
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
// The resource can be deleted (`RemovalPolicy.DELETE`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
// Experimental.
func (d *jsiiProxy_Distribution) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		d,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
type DistributionAttributes struct {
	// The distribution ID for this distribution.
	// Experimental.
	DistributionId *string `json:"distributionId"`
	// The generated domain name of the Distribution, such as d111111abcdef8.cloudfront.net.
	// Experimental.
	DomainName *string `json:"domainName"`
}

// Properties for a Distribution.
// Experimental.
type DistributionProps struct {
	// The default behavior for the distribution.
	// Experimental.
	DefaultBehavior *BehaviorOptions `json:"defaultBehavior"`
	// Additional behaviors for the distribution, mapped by the pathPattern that specifies which requests to apply the behavior to.
	// Experimental.
	AdditionalBehaviors *map[string]*BehaviorOptions `json:"additionalBehaviors"`
	// A certificate to associate with the distribution.
	//
	// The certificate must be located in N. Virginia (us-east-1).
	// Experimental.
	Certificate awscertificatemanager.ICertificate `json:"certificate"`
	// Any comments you want to include about the distribution.
	// Experimental.
	Comment *string `json:"comment"`
	// The object that you want CloudFront to request from your origin (for example, index.html) when a viewer requests the root URL for your distribution. If no default object is set, the request goes to the origin's root (e.g., example.com/).
	// Experimental.
	DefaultRootObject *string `json:"defaultRootObject"`
	// Alternative domain names for this distribution.
	//
	// If you want to use your own domain name, such as www.example.com, instead of the cloudfront.net domain name,
	// you can add an alternate domain name to your distribution. If you attach a certificate to the distribution,
	// you must add (at least one of) the domain names of the certificate to this list.
	// Experimental.
	DomainNames *[]*string `json:"domainNames"`
	// Enable or disable the distribution.
	// Experimental.
	Enabled *bool `json:"enabled"`
	// Whether CloudFront will respond to IPv6 DNS requests with an IPv6 address.
	//
	// If you specify false, CloudFront responds to IPv6 DNS requests with the DNS response code NOERROR and with no IP addresses.
	// This allows viewers to submit a second request, for an IPv4 address for your distribution.
	// Experimental.
	EnableIpv6 *bool `json:"enableIpv6"`
	// Enable access logging for the distribution.
	// Experimental.
	EnableLogging *bool `json:"enableLogging"`
	// How CloudFront should handle requests that are not successful (e.g., PageNotFound).
	// Experimental.
	ErrorResponses *[]*ErrorResponse `json:"errorResponses"`
	// Controls the countries in which your content is distributed.
	// Experimental.
	GeoRestriction GeoRestriction `json:"geoRestriction"`
	// Specify the maximum HTTP version that you want viewers to use to communicate with CloudFront.
	//
	// For viewers and CloudFront to use HTTP/2, viewers must support TLS 1.2 or later, and must support server name identification (SNI).
	// Experimental.
	HttpVersion HttpVersion `json:"httpVersion"`
	// The Amazon S3 bucket to store the access logs in.
	// Experimental.
	LogBucket awss3.IBucket `json:"logBucket"`
	// An optional string that you want CloudFront to prefix to the access log filenames for this distribution.
	// Experimental.
	LogFilePrefix *string `json:"logFilePrefix"`
	// Specifies whether you want CloudFront to include cookies in access logs.
	// Experimental.
	LogIncludesCookies *bool `json:"logIncludesCookies"`
	// The minimum version of the SSL protocol that you want CloudFront to use for HTTPS connections.
	//
	// CloudFront serves your objects only to browsers or devices that support at
	// least the SSL version that you specify.
	// Experimental.
	MinimumProtocolVersion SecurityPolicyProtocol `json:"minimumProtocolVersion"`
	// The price class that corresponds with the maximum price that you want to pay for CloudFront service.
	//
	// If you specify PriceClass_All, CloudFront responds to requests for your objects from all CloudFront edge locations.
	// If you specify a price class other than PriceClass_All, CloudFront serves your objects from the CloudFront edge location
	// that has the lowest latency among the edge locations in your price class.
	// Experimental.
	PriceClass PriceClass `json:"priceClass"`
	// Unique identifier that specifies the AWS WAF web ACL to associate with this CloudFront distribution.
	//
	// To specify a web ACL created using the latest version of AWS WAF, use the ACL ARN, for example
	// `arn:aws:wafv2:us-east-1:123456789012:global/webacl/ExampleWebACL/473e64fd-f30b-4765-81a0-62ad96dd167a`.
	// To specify a web ACL created using AWS WAF Classic, use the ACL ID, for example `473e64fd-f30b-4765-81a0-62ad96dd167a`.
	// See: https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_CreateDistribution.html#API_CreateDistribution_RequestParameters.
	//
	// Experimental.
	WebAclId *string `json:"webAclId"`
}

// Represents a Lambda function version and event type when using Lambda@Edge.
//
// The type of the {@link AddBehaviorOptions.edgeLambdas} property.
// Experimental.
type EdgeLambda struct {
	// The type of event in response to which should the function be invoked.
	// Experimental.
	EventType LambdaEdgeEventType `json:"eventType"`
	// The version of the Lambda function that will be invoked.
	//
	// **Note**: it's not possible to use the '$LATEST' function version for Lambda@Edge!
	// Experimental.
	FunctionVersion awslambda.IVersion `json:"functionVersion"`
	// Allows a Lambda function to have read access to the body content.
	//
	// Only valid for "request" event types (`ORIGIN_REQUEST` or `VIEWER_REQUEST`).
	// See https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/lambda-include-body-access.html
	// Experimental.
	IncludeBody *bool `json:"includeBody"`
}

// Options for configuring custom error responses.
// Experimental.
type ErrorResponse struct {
	// The HTTP status code for which you want to specify a custom error page and/or a caching duration.
	// Experimental.
	HttpStatus *float64 `json:"httpStatus"`
	// The HTTP status code that you want CloudFront to return to the viewer along with the custom error page.
	//
	// If you specify a value for `responseHttpStatus`, you must also specify a value for `responsePagePath`.
	// Experimental.
	ResponseHttpStatus *float64 `json:"responseHttpStatus"`
	// The path to the custom error page that you want CloudFront to return to a viewer when your origin returns the `httpStatus`, for example, /4xx-errors/403-forbidden.html.
	// Experimental.
	ResponsePagePath *string `json:"responsePagePath"`
	// The minimum amount of time, in seconds, that you want CloudFront to cache the HTTP status code specified in ErrorCode.
	// Experimental.
	Ttl awscdk.Duration `json:"ttl"`
}

// HTTP status code to failover to second origin.
// Experimental.
type FailoverStatusCode string

const (
	FailoverStatusCode_FORBIDDEN FailoverStatusCode = "FORBIDDEN"
	FailoverStatusCode_NOT_FOUND FailoverStatusCode = "NOT_FOUND"
	FailoverStatusCode_INTERNAL_SERVER_ERROR FailoverStatusCode = "INTERNAL_SERVER_ERROR"
	FailoverStatusCode_BAD_GATEWAY FailoverStatusCode = "BAD_GATEWAY"
	FailoverStatusCode_SERVICE_UNAVAILABLE FailoverStatusCode = "SERVICE_UNAVAILABLE"
	FailoverStatusCode_GATEWAY_TIMEOUT FailoverStatusCode = "GATEWAY_TIMEOUT"
)

// A CloudFront Function.
// Experimental.
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


// Experimental.
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

// Experimental.
func NewFunction_Override(f Function, scope constructs.Construct, id *string, props *FunctionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront.Function",
		[]interface{}{scope, id, props},
		f,
	)
}

// Imports a function by its name and ARN.
// Experimental.
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
// Experimental.
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
// The resource can be deleted (`RemovalPolicy.DELETE`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
// Experimental.
func (f *jsiiProxy_Function) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		f,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
type FunctionAssociation struct {
	// The type of event which should invoke the function.
	// Experimental.
	EventType FunctionEventType `json:"eventType"`
	// The CloudFront function that will be invoked.
	// Experimental.
	Function IFunction `json:"function"`
}

// Attributes of an existing CloudFront Function to import it.
// Experimental.
type FunctionAttributes struct {
	// The ARN of the function.
	// Experimental.
	FunctionArn *string `json:"functionArn"`
	// The name of the function.
	// Experimental.
	FunctionName *string `json:"functionName"`
}

// Represents the function's source code.
// Experimental.
type FunctionCode interface {
	Render() *string
}

// The jsii proxy struct for FunctionCode
type jsiiProxy_FunctionCode struct {
	_ byte // padding
}

// Experimental.
func NewFunctionCode_Override(f FunctionCode) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront.FunctionCode",
		nil, // no parameters
		f,
	)
}

// Inline code for function.
//
// Returns: `InlineCode` with inline code.
// Experimental.
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
// Experimental.
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
// Experimental.
type FunctionEventType string

const (
	FunctionEventType_VIEWER_REQUEST FunctionEventType = "VIEWER_REQUEST"
	FunctionEventType_VIEWER_RESPONSE FunctionEventType = "VIEWER_RESPONSE"
)

// Properties for creating a CloudFront Function.
// Experimental.
type FunctionProps struct {
	// The source code of the function.
	// Experimental.
	Code FunctionCode `json:"code"`
	// A comment to describe the function.
	// Experimental.
	Comment *string `json:"comment"`
	// A name to identify the function.
	// Experimental.
	FunctionName *string `json:"functionName"`
}

// Controls the countries in which content is distributed.
// Experimental.
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
// Experimental.
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
// Experimental.
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

// Maximum HTTP version to support.
// Experimental.
type HttpVersion string

const (
	HttpVersion_HTTP1_1 HttpVersion = "HTTP1_1"
	HttpVersion_HTTP2 HttpVersion = "HTTP2"
)

// Represents a Cache Policy.
// Experimental.
type ICachePolicy interface {
	// The ID of the cache policy.
	// Experimental.
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
// Experimental.
type IDistribution interface {
	awscdk.IResource
	// The domain name of the Distribution, such as d111111abcdef8.cloudfront.net.
	// Experimental.
	DistributionDomainName() *string
	// The distribution ID for this distribution.
	// Experimental.
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
// Experimental.
type IFunction interface {
	awscdk.IResource
	// The ARN of the function.
	// Experimental.
	FunctionArn() *string
	// The name of the function.
	// Experimental.
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
// Experimental.
type IKeyGroup interface {
	awscdk.IResource
	// The ID of the key group.
	// Experimental.
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
// Experimental.
type IOrigin interface {
	// The method called when a given Origin is added (for the first time) to a Distribution.
	// Experimental.
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
// Experimental.
type IOriginAccessIdentity interface {
	awsiam.IGrantable
	awscdk.IResource
	// The Origin Access Identity Name.
	// Experimental.
	OriginAccessIdentityName() *string
}

// The jsii proxy for IOriginAccessIdentity
type jsiiProxy_IOriginAccessIdentity struct {
	internal.Type__awsiamIGrantable
	internal.Type__awscdkIResource
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
// Experimental.
type IOriginRequestPolicy interface {
	// The ID of the origin request policy.
	// Experimental.
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
// Experimental.
type IPublicKey interface {
	awscdk.IResource
	// The ID of the key group.
	// Experimental.
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

// A Key Group configuration.
// Experimental.
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


// Experimental.
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

// Experimental.
func NewKeyGroup_Override(k KeyGroup, scope constructs.Construct, id *string, props *KeyGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront.KeyGroup",
		[]interface{}{scope, id, props},
		k,
	)
}

// Imports a Key Group from its id.
// Experimental.
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
// Experimental.
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
// The resource can be deleted (`RemovalPolicy.DELETE`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
// Experimental.
func (k *jsiiProxy_KeyGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		k,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
type KeyGroupProps struct {
	// A list of public keys to add to the key group.
	// Experimental.
	Items *[]IPublicKey `json:"items"`
	// A comment to describe the key group.
	// Experimental.
	Comment *string `json:"comment"`
	// A name to identify the key group.
	// Experimental.
	KeyGroupName *string `json:"keyGroupName"`
}

// The type of events that a Lambda@Edge function can be invoked in response to.
// Experimental.
type LambdaEdgeEventType string

const (
	LambdaEdgeEventType_ORIGIN_REQUEST LambdaEdgeEventType = "ORIGIN_REQUEST"
	LambdaEdgeEventType_ORIGIN_RESPONSE LambdaEdgeEventType = "ORIGIN_RESPONSE"
	LambdaEdgeEventType_VIEWER_REQUEST LambdaEdgeEventType = "VIEWER_REQUEST"
	LambdaEdgeEventType_VIEWER_RESPONSE LambdaEdgeEventType = "VIEWER_RESPONSE"
)

// Experimental.
type LambdaFunctionAssociation struct {
	// The lambda event type defines at which event the lambda is called during the request lifecycle.
	// Experimental.
	EventType LambdaEdgeEventType `json:"eventType"`
	// A version of the lambda to associate.
	// Experimental.
	LambdaFunction awslambda.IVersion `json:"lambdaFunction"`
	// Allows a Lambda function to have read access to the body content.
	//
	// Only valid for "request" event types (`ORIGIN_REQUEST` or `VIEWER_REQUEST`).
	// See https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/lambda-include-body-access.html
	// Experimental.
	IncludeBody *bool `json:"includeBody"`
}

// Logging configuration for incoming requests.
// Experimental.
type LoggingConfiguration struct {
	// Bucket to log requests to.
	// Experimental.
	Bucket awss3.IBucket `json:"bucket"`
	// Whether to include the cookies in the logs.
	// Experimental.
	IncludeCookies *bool `json:"includeCookies"`
	// Where in the bucket to store logs.
	// Experimental.
	Prefix *string `json:"prefix"`
}

// An origin access identity is a special CloudFront user that you can associate with Amazon S3 origins, so that you can secure all or just some of your Amazon S3 content.
// Experimental.
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


// Experimental.
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

// Experimental.
func NewOriginAccessIdentity_Override(o OriginAccessIdentity, scope constructs.Construct, id *string, props *OriginAccessIdentityProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront.OriginAccessIdentity",
		[]interface{}{scope, id, props},
		o,
	)
}

// Creates a OriginAccessIdentity by providing the OriginAccessIdentityName.
// Experimental.
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
// Experimental.
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
// The resource can be deleted (`RemovalPolicy.DELETE`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
// Experimental.
func (o *jsiiProxy_OriginAccessIdentity) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		o,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// The ARN to include in S3 bucket policy to allow CloudFront access.
// Experimental.
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

// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
type OriginAccessIdentityProps struct {
	// Any comments you want to include about the origin access identity.
	// Experimental.
	Comment *string `json:"comment"`
}

// Represents a distribution origin, that describes the Amazon S3 bucket, HTTP server (for example, a web server), Amazon MediaStore, or other server from which CloudFront gets your files.
// Experimental.
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

// Experimental.
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
// Experimental.
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

// Experimental.
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

// Experimental.
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
// Experimental.
type OriginBindConfig struct {
	// The failover configuration for this Origin.
	// Experimental.
	FailoverConfig *OriginFailoverConfig `json:"failoverConfig"`
	// The CloudFormation OriginProperty configuration for this Origin.
	// Experimental.
	OriginProperty *CfnDistribution_OriginProperty `json:"originProperty"`
}

// Options passed to Origin.bind().
// Experimental.
type OriginBindOptions struct {
	// The identifier of this Origin, as assigned by the Distribution this Origin has been used added to.
	// Experimental.
	OriginId *string `json:"originId"`
}

// The failover configuration used for Origin Groups, returned in {@link OriginBindConfig.failoverConfig}.
// Experimental.
type OriginFailoverConfig struct {
	// The origin to use as the fallback origin.
	// Experimental.
	FailoverOrigin IOrigin `json:"failoverOrigin"`
	// The HTTP status codes of the response that trigger querying the failover Origin.
	// Experimental.
	StatusCodes *[]*float64 `json:"statusCodes"`
}

// Properties to define an Origin.
// Experimental.
type OriginProps struct {
	// The number of times that CloudFront attempts to connect to the origin;
	//
	// valid values are 1, 2, or 3 attempts.
	// Experimental.
	ConnectionAttempts *float64 `json:"connectionAttempts"`
	// The number of seconds that CloudFront waits when trying to establish a connection to the origin.
	//
	// Valid values are 1-10 seconds, inclusive.
	// Experimental.
	ConnectionTimeout awscdk.Duration `json:"connectionTimeout"`
	// A list of HTTP header names and values that CloudFront adds to requests it sends to the origin.
	// Experimental.
	CustomHeaders *map[string]*string `json:"customHeaders"`
	// An optional path that CloudFront appends to the origin domain name when CloudFront requests content from the origin.
	//
	// Must begin, but not end, with '/' (e.g., '/production/images').
	// Experimental.
	OriginPath *string `json:"originPath"`
}

// Defines what protocols CloudFront will use to connect to an origin.
// Experimental.
type OriginProtocolPolicy string

const (
	OriginProtocolPolicy_HTTP_ONLY OriginProtocolPolicy = "HTTP_ONLY"
	OriginProtocolPolicy_MATCH_VIEWER OriginProtocolPolicy = "MATCH_VIEWER"
	OriginProtocolPolicy_HTTPS_ONLY OriginProtocolPolicy = "HTTPS_ONLY"
)

// Determines whether any cookies in viewer requests (and if so, which cookies) are included in requests that CloudFront sends to the origin.
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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


// Experimental.
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

// Experimental.
func NewOriginRequestPolicy_Override(o OriginRequestPolicy, scope constructs.Construct, id *string, props *OriginRequestPolicyProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront.OriginRequestPolicy",
		[]interface{}{scope, id, props},
		o,
	)
}

// Imports a Origin Request Policy from its id.
// Experimental.
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
// Experimental.
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
// The resource can be deleted (`RemovalPolicy.DELETE`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
// Experimental.
func (o *jsiiProxy_OriginRequestPolicy) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		o,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
type OriginRequestPolicyProps struct {
	// A comment to describe the origin request policy.
	// Experimental.
	Comment *string `json:"comment"`
	// The cookies from viewer requests to include in origin requests.
	// Experimental.
	CookieBehavior OriginRequestCookieBehavior `json:"cookieBehavior"`
	// The HTTP headers to include in origin requests.
	//
	// These can include headers from viewer requests and additional headers added by CloudFront.
	// Experimental.
	HeaderBehavior OriginRequestHeaderBehavior `json:"headerBehavior"`
	// A unique name to identify the origin request policy.
	//
	// The name must only include '-', '_', or alphanumeric characters.
	// Experimental.
	OriginRequestPolicyName *string `json:"originRequestPolicyName"`
	// The URL query strings from viewer requests to include in origin requests.
	// Experimental.
	QueryStringBehavior OriginRequestQueryStringBehavior `json:"queryStringBehavior"`
}

// Determines whether any URL query strings in viewer requests (and if so, which query strings) are included in requests that CloudFront sends to the origin.
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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

// Experimental.
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
// Experimental.
type PriceClass string

const (
	PriceClass_PRICE_CLASS_100 PriceClass = "PRICE_CLASS_100"
	PriceClass_PRICE_CLASS_200 PriceClass = "PRICE_CLASS_200"
	PriceClass_PRICE_CLASS_ALL PriceClass = "PRICE_CLASS_ALL"
)

// A Public Key Configuration.
// Experimental.
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


// Experimental.
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

// Experimental.
func NewPublicKey_Override(p PublicKey, scope constructs.Construct, id *string, props *PublicKeyProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront.PublicKey",
		[]interface{}{scope, id, props},
		p,
	)
}

// Imports a Public Key from its id.
// Experimental.
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
// Experimental.
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
// The resource can be deleted (`RemovalPolicy.DELETE`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
// Experimental.
func (p *jsiiProxy_PublicKey) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		p,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
type PublicKeyProps struct {
	// The public key that you can use with signed URLs and signed cookies, or with field-level encryption.
	//
	// The `encodedKey` parameter must include `-----BEGIN PUBLIC KEY-----` and `-----END PUBLIC KEY-----` lines.
	// See: https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/field-level-encryption.html
	//
	// Experimental.
	EncodedKey *string `json:"encodedKey"`
	// A comment to describe the public key.
	// Experimental.
	Comment *string `json:"comment"`
	// A name to identify the public key.
	// Experimental.
	PublicKeyName *string `json:"publicKeyName"`
}

// S3 origin configuration for CloudFront.
// Experimental.
type S3OriginConfig struct {
	// The source bucket to serve content from.
	// Experimental.
	S3BucketSource awss3.IBucket `json:"s3BucketSource"`
	// The optional Origin Access Identity of the origin identity cloudfront will use when calling your s3 bucket.
	// Experimental.
	OriginAccessIdentity IOriginAccessIdentity `json:"originAccessIdentity"`
	// Any additional headers to pass to the origin.
	// Experimental.
	OriginHeaders *map[string]*string `json:"originHeaders"`
	// The relative path to the origin root to use for sources.
	// Experimental.
	OriginPath *string `json:"originPath"`
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
// Experimental.
type SSLMethod string

const (
	SSLMethod_SNI SSLMethod = "SNI"
	SSLMethod_VIP SSLMethod = "VIP"
)

// The minimum version of the SSL protocol that you want CloudFront to use for HTTPS connections.
//
// CloudFront serves your objects only to browsers or devices that support at least the SSL version that you specify.
// Experimental.
type SecurityPolicyProtocol string

const (
	SecurityPolicyProtocol_SSL_V3 SecurityPolicyProtocol = "SSL_V3"
	SecurityPolicyProtocol_TLS_V1 SecurityPolicyProtocol = "TLS_V1"
	SecurityPolicyProtocol_TLS_V1_2016 SecurityPolicyProtocol = "TLS_V1_2016"
	SecurityPolicyProtocol_TLS_V1_1_2016 SecurityPolicyProtocol = "TLS_V1_1_2016"
	SecurityPolicyProtocol_TLS_V1_2_2018 SecurityPolicyProtocol = "TLS_V1_2_2018"
	SecurityPolicyProtocol_TLS_V1_2_2019 SecurityPolicyProtocol = "TLS_V1_2_2019"
)

// A source configuration is a wrapper for CloudFront origins and behaviors.
//
// An origin is what CloudFront will "be in front of" - that is, CloudFront will pull it's assets from an origin.
//
// If you're using s3 as a source - pass the `s3Origin` property, otherwise, pass the `customOriginSource` property.
//
// One or the other must be passed, and it is invalid to pass both in the same SourceConfiguration.
// Experimental.
type SourceConfiguration struct {
	// The behaviors associated with this source.
	//
	// At least one (default) behavior must be included.
	// Experimental.
	Behaviors *[]*Behavior `json:"behaviors"`
	// The number of times that CloudFront attempts to connect to the origin.
	//
	// You can specify 1, 2, or 3 as the number of attempts.
	// Experimental.
	ConnectionAttempts *float64 `json:"connectionAttempts"`
	// The number of seconds that CloudFront waits when trying to establish a connection to the origin.
	//
	// You can specify a number of seconds between 1 and 10 (inclusive).
	// Experimental.
	ConnectionTimeout awscdk.Duration `json:"connectionTimeout"`
	// A custom origin source - for all non-s3 sources.
	// Experimental.
	CustomOriginSource *CustomOriginConfig `json:"customOriginSource"`
	// HTTP status code to failover to second origin.
	// Experimental.
	FailoverCriteriaStatusCodes *[]FailoverStatusCode `json:"failoverCriteriaStatusCodes"`
	// A custom origin source for failover in case the s3OriginSource returns invalid status code.
	// Experimental.
	FailoverCustomOriginSource *CustomOriginConfig `json:"failoverCustomOriginSource"`
	// An s3 origin source for failover in case the s3OriginSource returns invalid status code.
	// Experimental.
	FailoverS3OriginSource *S3OriginConfig `json:"failoverS3OriginSource"`
	// An s3 origin source - if you're using s3 for your assets.
	// Experimental.
	S3OriginSource *S3OriginConfig `json:"s3OriginSource"`
}

// Viewer certificate configuration class.
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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

// Experimental.
type ViewerCertificateOptions struct {
	// Domain names on the certificate (both main domain name and Subject Alternative names).
	// Experimental.
	Aliases *[]*string `json:"aliases"`
	// The minimum version of the SSL protocol that you want CloudFront to use for HTTPS connections.
	//
	// CloudFront serves your objects only to browsers or devices that support at
	// least the SSL version that you specify.
	// Experimental.
	SecurityPolicy SecurityPolicyProtocol `json:"securityPolicy"`
	// How CloudFront should serve HTTPS requests.
	//
	// See the notes on SSLMethod if you wish to use other SSL termination types.
	// See: https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_ViewerCertificate.html
	//
	// Experimental.
	SslMethod SSLMethod `json:"sslMethod"`
}

// How HTTPs should be handled with your distribution.
// Experimental.
type ViewerProtocolPolicy string

const (
	ViewerProtocolPolicy_HTTPS_ONLY ViewerProtocolPolicy = "HTTPS_ONLY"
	ViewerProtocolPolicy_REDIRECT_TO_HTTPS ViewerProtocolPolicy = "REDIRECT_TO_HTTPS"
	ViewerProtocolPolicy_ALLOW_ALL ViewerProtocolPolicy = "ALLOW_ALL"
)

