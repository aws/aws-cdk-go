package awscloudfrontorigins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscloudfront"
	"github.com/aws/aws-cdk-go/awscdk/awscloudfrontorigins/internal"
	"github.com/aws/aws-cdk-go/awscdk/awselasticloadbalancingv2"
	"github.com/aws/aws-cdk-go/awscdk/awss3"
)

// An Origin for an HTTP server or S3 bucket configured for website hosting.
// Experimental.
type HttpOrigin interface {
	awscloudfront.OriginBase
	Bind(_scope awscdk.Construct, options *awscloudfront.OriginBindOptions) *awscloudfront.OriginBindConfig
	RenderCustomOriginConfig() *awscloudfront.CfnDistribution_CustomOriginConfigProperty
	RenderS3OriginConfig() *awscloudfront.CfnDistribution_S3OriginConfigProperty
}

// The jsii proxy struct for HttpOrigin
type jsiiProxy_HttpOrigin struct {
	internal.Type__awscloudfrontOriginBase
}

// Experimental.
func NewHttpOrigin(domainName *string, props *HttpOriginProps) HttpOrigin {
	_init_.Initialize()

	j := jsiiProxy_HttpOrigin{}

	_jsii_.Create(
		"monocdk.aws_cloudfront_origins.HttpOrigin",
		[]interface{}{domainName, props},
		&j,
	)

	return &j
}

// Experimental.
func NewHttpOrigin_Override(h HttpOrigin, domainName *string, props *HttpOriginProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_cloudfront_origins.HttpOrigin",
		[]interface{}{domainName, props},
		h,
	)
}

// Binds the origin to the associated Distribution.
//
// Can be used to grant permissions, create dependent resources, etc.
// Experimental.
func (h *jsiiProxy_HttpOrigin) Bind(_scope awscdk.Construct, options *awscloudfront.OriginBindOptions) *awscloudfront.OriginBindConfig {
	var returns *awscloudfront.OriginBindConfig

	_jsii_.Invoke(
		h,
		"bind",
		[]interface{}{_scope, options},
		&returns,
	)

	return returns
}

// Experimental.
func (h *jsiiProxy_HttpOrigin) RenderCustomOriginConfig() *awscloudfront.CfnDistribution_CustomOriginConfigProperty {
	var returns *awscloudfront.CfnDistribution_CustomOriginConfigProperty

	_jsii_.Invoke(
		h,
		"renderCustomOriginConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (h *jsiiProxy_HttpOrigin) RenderS3OriginConfig() *awscloudfront.CfnDistribution_S3OriginConfigProperty {
	var returns *awscloudfront.CfnDistribution_S3OriginConfigProperty

	_jsii_.Invoke(
		h,
		"renderS3OriginConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for an Origin backed by an S3 website-configured bucket, load balancer, or custom HTTP server.
// Experimental.
type HttpOriginProps struct {
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
	// The HTTP port that CloudFront uses to connect to the origin.
	// Experimental.
	HttpPort *float64 `json:"httpPort"`
	// The HTTPS port that CloudFront uses to connect to the origin.
	// Experimental.
	HttpsPort *float64 `json:"httpsPort"`
	// Specifies how long, in seconds, CloudFront persists its connection to the origin.
	//
	// The valid range is from 1 to 60 seconds, inclusive.
	// Experimental.
	KeepaliveTimeout awscdk.Duration `json:"keepaliveTimeout"`
	// The SSL versions to use when interacting with the origin.
	// Experimental.
	OriginSslProtocols *[]awscloudfront.OriginSslPolicy `json:"originSslProtocols"`
	// Specifies the protocol (HTTP or HTTPS) that CloudFront uses to connect to the origin.
	// Experimental.
	ProtocolPolicy awscloudfront.OriginProtocolPolicy `json:"protocolPolicy"`
	// Specifies how long, in seconds, CloudFront waits for a response from the origin, also known as the origin response timeout.
	//
	// The valid range is from 1 to 60 seconds, inclusive.
	// Experimental.
	ReadTimeout awscdk.Duration `json:"readTimeout"`
}

// An Origin for a v2 load balancer.
// Experimental.
type LoadBalancerV2Origin interface {
	HttpOrigin
	Bind(_scope awscdk.Construct, options *awscloudfront.OriginBindOptions) *awscloudfront.OriginBindConfig
	RenderCustomOriginConfig() *awscloudfront.CfnDistribution_CustomOriginConfigProperty
	RenderS3OriginConfig() *awscloudfront.CfnDistribution_S3OriginConfigProperty
}

// The jsii proxy struct for LoadBalancerV2Origin
type jsiiProxy_LoadBalancerV2Origin struct {
	jsiiProxy_HttpOrigin
}

// Experimental.
func NewLoadBalancerV2Origin(loadBalancer awselasticloadbalancingv2.ILoadBalancerV2, props *LoadBalancerV2OriginProps) LoadBalancerV2Origin {
	_init_.Initialize()

	j := jsiiProxy_LoadBalancerV2Origin{}

	_jsii_.Create(
		"monocdk.aws_cloudfront_origins.LoadBalancerV2Origin",
		[]interface{}{loadBalancer, props},
		&j,
	)

	return &j
}

// Experimental.
func NewLoadBalancerV2Origin_Override(l LoadBalancerV2Origin, loadBalancer awselasticloadbalancingv2.ILoadBalancerV2, props *LoadBalancerV2OriginProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_cloudfront_origins.LoadBalancerV2Origin",
		[]interface{}{loadBalancer, props},
		l,
	)
}

// Binds the origin to the associated Distribution.
//
// Can be used to grant permissions, create dependent resources, etc.
// Experimental.
func (l *jsiiProxy_LoadBalancerV2Origin) Bind(_scope awscdk.Construct, options *awscloudfront.OriginBindOptions) *awscloudfront.OriginBindConfig {
	var returns *awscloudfront.OriginBindConfig

	_jsii_.Invoke(
		l,
		"bind",
		[]interface{}{_scope, options},
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LoadBalancerV2Origin) RenderCustomOriginConfig() *awscloudfront.CfnDistribution_CustomOriginConfigProperty {
	var returns *awscloudfront.CfnDistribution_CustomOriginConfigProperty

	_jsii_.Invoke(
		l,
		"renderCustomOriginConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (l *jsiiProxy_LoadBalancerV2Origin) RenderS3OriginConfig() *awscloudfront.CfnDistribution_S3OriginConfigProperty {
	var returns *awscloudfront.CfnDistribution_S3OriginConfigProperty

	_jsii_.Invoke(
		l,
		"renderS3OriginConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for an Origin backed by a v2 load balancer.
// Experimental.
type LoadBalancerV2OriginProps struct {
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
	// The HTTP port that CloudFront uses to connect to the origin.
	// Experimental.
	HttpPort *float64 `json:"httpPort"`
	// The HTTPS port that CloudFront uses to connect to the origin.
	// Experimental.
	HttpsPort *float64 `json:"httpsPort"`
	// Specifies how long, in seconds, CloudFront persists its connection to the origin.
	//
	// The valid range is from 1 to 60 seconds, inclusive.
	// Experimental.
	KeepaliveTimeout awscdk.Duration `json:"keepaliveTimeout"`
	// The SSL versions to use when interacting with the origin.
	// Experimental.
	OriginSslProtocols *[]awscloudfront.OriginSslPolicy `json:"originSslProtocols"`
	// Specifies the protocol (HTTP or HTTPS) that CloudFront uses to connect to the origin.
	// Experimental.
	ProtocolPolicy awscloudfront.OriginProtocolPolicy `json:"protocolPolicy"`
	// Specifies how long, in seconds, CloudFront waits for a response from the origin, also known as the origin response timeout.
	//
	// The valid range is from 1 to 60 seconds, inclusive.
	// Experimental.
	ReadTimeout awscdk.Duration `json:"readTimeout"`
}

// An Origin that represents a group.
//
// Consists of a primary Origin,
// and a fallback Origin called when the primary returns one of the provided HTTP status codes.
// Experimental.
type OriginGroup interface {
	awscloudfront.IOrigin
	Bind(scope awscdk.Construct, options *awscloudfront.OriginBindOptions) *awscloudfront.OriginBindConfig
}

// The jsii proxy struct for OriginGroup
type jsiiProxy_OriginGroup struct {
	internal.Type__awscloudfrontIOrigin
}

// Experimental.
func NewOriginGroup(props *OriginGroupProps) OriginGroup {
	_init_.Initialize()

	j := jsiiProxy_OriginGroup{}

	_jsii_.Create(
		"monocdk.aws_cloudfront_origins.OriginGroup",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewOriginGroup_Override(o OriginGroup, props *OriginGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_cloudfront_origins.OriginGroup",
		[]interface{}{props},
		o,
	)
}

// The method called when a given Origin is added (for the first time) to a Distribution.
// Experimental.
func (o *jsiiProxy_OriginGroup) Bind(scope awscdk.Construct, options *awscloudfront.OriginBindOptions) *awscloudfront.OriginBindConfig {
	var returns *awscloudfront.OriginBindConfig

	_jsii_.Invoke(
		o,
		"bind",
		[]interface{}{scope, options},
		&returns,
	)

	return returns
}

// Construction properties for {@link OriginGroup}.
// Experimental.
type OriginGroupProps struct {
	// The fallback origin that should serve requests when the primary fails.
	// Experimental.
	FallbackOrigin awscloudfront.IOrigin `json:"fallbackOrigin"`
	// The primary origin that should serve requests for this group.
	// Experimental.
	PrimaryOrigin awscloudfront.IOrigin `json:"primaryOrigin"`
	// The list of HTTP status codes that, when returned from the primary origin, would cause querying the fallback origin.
	// Experimental.
	FallbackStatusCodes *[]*float64 `json:"fallbackStatusCodes"`
}

// An Origin that is backed by an S3 bucket.
//
// If the bucket is configured for website hosting, this origin will be configured to use the bucket as an
// HTTP server origin and will use the bucket's configured website redirects and error handling. Otherwise,
// the origin is created as a bucket origin and will use CloudFront's redirect and error handling.
// Experimental.
type S3Origin interface {
	awscloudfront.IOrigin
	Bind(scope awscdk.Construct, options *awscloudfront.OriginBindOptions) *awscloudfront.OriginBindConfig
}

// The jsii proxy struct for S3Origin
type jsiiProxy_S3Origin struct {
	internal.Type__awscloudfrontIOrigin
}

// Experimental.
func NewS3Origin(bucket awss3.IBucket, props *S3OriginProps) S3Origin {
	_init_.Initialize()

	j := jsiiProxy_S3Origin{}

	_jsii_.Create(
		"monocdk.aws_cloudfront_origins.S3Origin",
		[]interface{}{bucket, props},
		&j,
	)

	return &j
}

// Experimental.
func NewS3Origin_Override(s S3Origin, bucket awss3.IBucket, props *S3OriginProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_cloudfront_origins.S3Origin",
		[]interface{}{bucket, props},
		s,
	)
}

// The method called when a given Origin is added (for the first time) to a Distribution.
// Experimental.
func (s *jsiiProxy_S3Origin) Bind(scope awscdk.Construct, options *awscloudfront.OriginBindOptions) *awscloudfront.OriginBindConfig {
	var returns *awscloudfront.OriginBindConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{scope, options},
		&returns,
	)

	return returns
}

// Properties to use to customize an S3 Origin.
// Experimental.
type S3OriginProps struct {
	// An optional Origin Access Identity of the origin identity cloudfront will use when calling your s3 bucket.
	// Experimental.
	OriginAccessIdentity awscloudfront.IOriginAccessIdentity `json:"originAccessIdentity"`
	// An optional path that CloudFront appends to the origin domain name when CloudFront requests content from the origin.
	//
	// Must begin, but not end, with '/' (e.g., '/production/images').
	// Experimental.
	OriginPath *string `json:"originPath"`
}

