package awscloudfrontorigins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfrontorigins/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
)

// An Origin for an HTTP server or S3 bucket configured for website hosting.
//
// Example:
//   myBucket := s3.NewBucket(this, jsii.String("myBucket"))
//   cloudfront.NewDistribution(this, jsii.String("myDist"), &distributionProps{
//   	defaultBehavior: &behaviorOptions{
//   		origin: origins.NewOriginGroup(&originGroupProps{
//   			primaryOrigin: origins.NewS3Origin(myBucket),
//   			fallbackOrigin: origins.NewHttpOrigin(jsii.String("www.example.com")),
//   			// optional, defaults to: 500, 502, 503 and 504
//   			fallbackStatusCodes: []*f64{
//   				jsii.Number(404),
//   			},
//   		}),
//   	},
//   })
//
type HttpOrigin interface {
	awscloudfront.OriginBase
	// Binds the origin to the associated Distribution.
	//
	// Can be used to grant permissions, create dependent resources, etc.
	Bind(_scope constructs.Construct, options *awscloudfront.OriginBindOptions) *awscloudfront.OriginBindConfig
	RenderCustomOriginConfig() *awscloudfront.CfnDistribution_CustomOriginConfigProperty
	RenderS3OriginConfig() *awscloudfront.CfnDistribution_S3OriginConfigProperty
}

// The jsii proxy struct for HttpOrigin
type jsiiProxy_HttpOrigin struct {
	internal.Type__awscloudfrontOriginBase
}

func NewHttpOrigin(domainName *string, props *HttpOriginProps) HttpOrigin {
	_init_.Initialize()

	j := jsiiProxy_HttpOrigin{}

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront_origins.HttpOrigin",
		[]interface{}{domainName, props},
		&j,
	)

	return &j
}

func NewHttpOrigin_Override(h HttpOrigin, domainName *string, props *HttpOriginProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront_origins.HttpOrigin",
		[]interface{}{domainName, props},
		h,
	)
}

func (h *jsiiProxy_HttpOrigin) Bind(_scope constructs.Construct, options *awscloudfront.OriginBindOptions) *awscloudfront.OriginBindConfig {
	var returns *awscloudfront.OriginBindConfig

	_jsii_.Invoke(
		h,
		"bind",
		[]interface{}{_scope, options},
		&returns,
	)

	return returns
}

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
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   httpOriginProps := &httpOriginProps{
//   	connectionAttempts: jsii.Number(123),
//   	connectionTimeout: cdk.duration.minutes(jsii.Number(30)),
//   	customHeaders: map[string]*string{
//   		"customHeadersKey": jsii.String("customHeaders"),
//   	},
//   	httpPort: jsii.Number(123),
//   	httpsPort: jsii.Number(123),
//   	keepaliveTimeout: cdk.*duration.minutes(jsii.Number(30)),
//   	originPath: jsii.String("originPath"),
//   	originShieldRegion: jsii.String("originShieldRegion"),
//   	originSslProtocols: []originSslPolicy{
//   		awscdk.Aws_cloudfront.*originSslPolicy_SSL_V3,
//   	},
//   	protocolPolicy: awscdk.*Aws_cloudfront.originProtocolPolicy_HTTP_ONLY,
//   	readTimeout: cdk.*duration.minutes(jsii.Number(30)),
//   }
//
type HttpOriginProps struct {
	// The number of times that CloudFront attempts to connect to the origin;
	//
	// valid values are 1, 2, or 3 attempts.
	ConnectionAttempts *float64 `field:"optional" json:"connectionAttempts" yaml:"connectionAttempts"`
	// The number of seconds that CloudFront waits when trying to establish a connection to the origin.
	//
	// Valid values are 1-10 seconds, inclusive.
	ConnectionTimeout awscdk.Duration `field:"optional" json:"connectionTimeout" yaml:"connectionTimeout"`
	// A list of HTTP header names and values that CloudFront adds to requests it sends to the origin.
	CustomHeaders *map[string]*string `field:"optional" json:"customHeaders" yaml:"customHeaders"`
	// An optional path that CloudFront appends to the origin domain name when CloudFront requests content from the origin.
	//
	// Must begin, but not end, with '/' (e.g., '/production/images').
	OriginPath *string `field:"optional" json:"originPath" yaml:"originPath"`
	// When you enable Origin Shield in the AWS Region that has the lowest latency to your origin, you can get better network performance.
	// See: https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/origin-shield.html
	//
	OriginShieldRegion *string `field:"optional" json:"originShieldRegion" yaml:"originShieldRegion"`
	// The HTTP port that CloudFront uses to connect to the origin.
	HttpPort *float64 `field:"optional" json:"httpPort" yaml:"httpPort"`
	// The HTTPS port that CloudFront uses to connect to the origin.
	HttpsPort *float64 `field:"optional" json:"httpsPort" yaml:"httpsPort"`
	// Specifies how long, in seconds, CloudFront persists its connection to the origin.
	//
	// The valid range is from 1 to 180 seconds, inclusive.
	//
	// Note that values over 60 seconds are possible only after a limit increase request for the origin response timeout quota
	// has been approved in the target account; otherwise, values over 60 seconds will produce an error at deploy time.
	KeepaliveTimeout awscdk.Duration `field:"optional" json:"keepaliveTimeout" yaml:"keepaliveTimeout"`
	// The SSL versions to use when interacting with the origin.
	OriginSslProtocols *[]awscloudfront.OriginSslPolicy `field:"optional" json:"originSslProtocols" yaml:"originSslProtocols"`
	// Specifies the protocol (HTTP or HTTPS) that CloudFront uses to connect to the origin.
	ProtocolPolicy awscloudfront.OriginProtocolPolicy `field:"optional" json:"protocolPolicy" yaml:"protocolPolicy"`
	// Specifies how long, in seconds, CloudFront waits for a response from the origin, also known as the origin response timeout.
	//
	// The valid range is from 1 to 180 seconds, inclusive.
	//
	// Note that values over 60 seconds are possible only after a limit increase request for the origin response timeout quota
	// has been approved in the target account; otherwise, values over 60 seconds will produce an error at deploy time.
	ReadTimeout awscdk.Duration `field:"optional" json:"readTimeout" yaml:"readTimeout"`
}

// An Origin for a v2 load balancer.
//
// Example:
//   import ec2 "github.com/aws/aws-cdk-go/awscdk"
//   import elbv2 "github.com/aws/aws-cdk-go/awscdk"
//
//   var vpc vpc
//
//   // Create an application load balancer in a VPC. 'internetFacing' must be 'true'
//   // for CloudFront to access the load balancer and use it as an origin.
//   lb := elbv2.NewApplicationLoadBalancer(this, jsii.String("LB"), &applicationLoadBalancerProps{
//   	vpc: vpc,
//   	internetFacing: jsii.Boolean(true),
//   })
//   cloudfront.NewDistribution(this, jsii.String("myDist"), &distributionProps{
//   	defaultBehavior: &behaviorOptions{
//   		origin: origins.NewLoadBalancerV2Origin(lb),
//   	},
//   })
//
type LoadBalancerV2Origin interface {
	HttpOrigin
	// Binds the origin to the associated Distribution.
	//
	// Can be used to grant permissions, create dependent resources, etc.
	Bind(_scope constructs.Construct, options *awscloudfront.OriginBindOptions) *awscloudfront.OriginBindConfig
	RenderCustomOriginConfig() *awscloudfront.CfnDistribution_CustomOriginConfigProperty
	RenderS3OriginConfig() *awscloudfront.CfnDistribution_S3OriginConfigProperty
}

// The jsii proxy struct for LoadBalancerV2Origin
type jsiiProxy_LoadBalancerV2Origin struct {
	jsiiProxy_HttpOrigin
}

func NewLoadBalancerV2Origin(loadBalancer awselasticloadbalancingv2.ILoadBalancerV2, props *LoadBalancerV2OriginProps) LoadBalancerV2Origin {
	_init_.Initialize()

	j := jsiiProxy_LoadBalancerV2Origin{}

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront_origins.LoadBalancerV2Origin",
		[]interface{}{loadBalancer, props},
		&j,
	)

	return &j
}

func NewLoadBalancerV2Origin_Override(l LoadBalancerV2Origin, loadBalancer awselasticloadbalancingv2.ILoadBalancerV2, props *LoadBalancerV2OriginProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront_origins.LoadBalancerV2Origin",
		[]interface{}{loadBalancer, props},
		l,
	)
}

func (l *jsiiProxy_LoadBalancerV2Origin) Bind(_scope constructs.Construct, options *awscloudfront.OriginBindOptions) *awscloudfront.OriginBindConfig {
	var returns *awscloudfront.OriginBindConfig

	_jsii_.Invoke(
		l,
		"bind",
		[]interface{}{_scope, options},
		&returns,
	)

	return returns
}

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
//
// Example:
//   import elbv2 "github.com/aws/aws-cdk-go/awscdk"
//
//   var loadBalancer applicationLoadBalancer
//
//   origin := origins.NewLoadBalancerV2Origin(loadBalancer, &loadBalancerV2OriginProps{
//   	connectionAttempts: jsii.Number(3),
//   	connectionTimeout: awscdk.Duration.seconds(jsii.Number(5)),
//   	readTimeout: awscdk.Duration.seconds(jsii.Number(45)),
//   	keepaliveTimeout: awscdk.Duration.seconds(jsii.Number(45)),
//   	protocolPolicy: cloudfront.originProtocolPolicy_MATCH_VIEWER,
//   })
//
type LoadBalancerV2OriginProps struct {
	// The number of times that CloudFront attempts to connect to the origin;
	//
	// valid values are 1, 2, or 3 attempts.
	ConnectionAttempts *float64 `field:"optional" json:"connectionAttempts" yaml:"connectionAttempts"`
	// The number of seconds that CloudFront waits when trying to establish a connection to the origin.
	//
	// Valid values are 1-10 seconds, inclusive.
	ConnectionTimeout awscdk.Duration `field:"optional" json:"connectionTimeout" yaml:"connectionTimeout"`
	// A list of HTTP header names and values that CloudFront adds to requests it sends to the origin.
	CustomHeaders *map[string]*string `field:"optional" json:"customHeaders" yaml:"customHeaders"`
	// An optional path that CloudFront appends to the origin domain name when CloudFront requests content from the origin.
	//
	// Must begin, but not end, with '/' (e.g., '/production/images').
	OriginPath *string `field:"optional" json:"originPath" yaml:"originPath"`
	// When you enable Origin Shield in the AWS Region that has the lowest latency to your origin, you can get better network performance.
	// See: https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/origin-shield.html
	//
	OriginShieldRegion *string `field:"optional" json:"originShieldRegion" yaml:"originShieldRegion"`
	// The HTTP port that CloudFront uses to connect to the origin.
	HttpPort *float64 `field:"optional" json:"httpPort" yaml:"httpPort"`
	// The HTTPS port that CloudFront uses to connect to the origin.
	HttpsPort *float64 `field:"optional" json:"httpsPort" yaml:"httpsPort"`
	// Specifies how long, in seconds, CloudFront persists its connection to the origin.
	//
	// The valid range is from 1 to 180 seconds, inclusive.
	//
	// Note that values over 60 seconds are possible only after a limit increase request for the origin response timeout quota
	// has been approved in the target account; otherwise, values over 60 seconds will produce an error at deploy time.
	KeepaliveTimeout awscdk.Duration `field:"optional" json:"keepaliveTimeout" yaml:"keepaliveTimeout"`
	// The SSL versions to use when interacting with the origin.
	OriginSslProtocols *[]awscloudfront.OriginSslPolicy `field:"optional" json:"originSslProtocols" yaml:"originSslProtocols"`
	// Specifies the protocol (HTTP or HTTPS) that CloudFront uses to connect to the origin.
	ProtocolPolicy awscloudfront.OriginProtocolPolicy `field:"optional" json:"protocolPolicy" yaml:"protocolPolicy"`
	// Specifies how long, in seconds, CloudFront waits for a response from the origin, also known as the origin response timeout.
	//
	// The valid range is from 1 to 180 seconds, inclusive.
	//
	// Note that values over 60 seconds are possible only after a limit increase request for the origin response timeout quota
	// has been approved in the target account; otherwise, values over 60 seconds will produce an error at deploy time.
	ReadTimeout awscdk.Duration `field:"optional" json:"readTimeout" yaml:"readTimeout"`
}

// An Origin that represents a group.
//
// Consists of a primary Origin,
// and a fallback Origin called when the primary returns one of the provided HTTP status codes.
//
// Example:
//   myBucket := s3.NewBucket(this, jsii.String("myBucket"))
//   cloudfront.NewDistribution(this, jsii.String("myDist"), &distributionProps{
//   	defaultBehavior: &behaviorOptions{
//   		origin: origins.NewOriginGroup(&originGroupProps{
//   			primaryOrigin: origins.NewS3Origin(myBucket),
//   			fallbackOrigin: origins.NewHttpOrigin(jsii.String("www.example.com")),
//   			// optional, defaults to: 500, 502, 503 and 504
//   			fallbackStatusCodes: []*f64{
//   				jsii.Number(404),
//   			},
//   		}),
//   	},
//   })
//
type OriginGroup interface {
	awscloudfront.IOrigin
	// The method called when a given Origin is added (for the first time) to a Distribution.
	Bind(scope constructs.Construct, options *awscloudfront.OriginBindOptions) *awscloudfront.OriginBindConfig
}

// The jsii proxy struct for OriginGroup
type jsiiProxy_OriginGroup struct {
	internal.Type__awscloudfrontIOrigin
}

func NewOriginGroup(props *OriginGroupProps) OriginGroup {
	_init_.Initialize()

	j := jsiiProxy_OriginGroup{}

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront_origins.OriginGroup",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewOriginGroup_Override(o OriginGroup, props *OriginGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront_origins.OriginGroup",
		[]interface{}{props},
		o,
	)
}

func (o *jsiiProxy_OriginGroup) Bind(scope constructs.Construct, options *awscloudfront.OriginBindOptions) *awscloudfront.OriginBindConfig {
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
//
// Example:
//   myBucket := s3.NewBucket(this, jsii.String("myBucket"))
//   cloudfront.NewDistribution(this, jsii.String("myDist"), &distributionProps{
//   	defaultBehavior: &behaviorOptions{
//   		origin: origins.NewOriginGroup(&originGroupProps{
//   			primaryOrigin: origins.NewS3Origin(myBucket),
//   			fallbackOrigin: origins.NewHttpOrigin(jsii.String("www.example.com")),
//   			// optional, defaults to: 500, 502, 503 and 504
//   			fallbackStatusCodes: []*f64{
//   				jsii.Number(404),
//   			},
//   		}),
//   	},
//   })
//
type OriginGroupProps struct {
	// The fallback origin that should serve requests when the primary fails.
	FallbackOrigin awscloudfront.IOrigin `field:"required" json:"fallbackOrigin" yaml:"fallbackOrigin"`
	// The primary origin that should serve requests for this group.
	PrimaryOrigin awscloudfront.IOrigin `field:"required" json:"primaryOrigin" yaml:"primaryOrigin"`
	// The list of HTTP status codes that, when returned from the primary origin, would cause querying the fallback origin.
	FallbackStatusCodes *[]*float64 `field:"optional" json:"fallbackStatusCodes" yaml:"fallbackStatusCodes"`
}

// An Origin that is backed by an S3 bucket.
//
// If the bucket is configured for website hosting, this origin will be configured to use the bucket as an
// HTTP server origin and will use the bucket's configured website redirects and error handling. Otherwise,
// the origin is created as a bucket origin and will use CloudFront's redirect and error handling.
//
// Example:
//   // Adding an existing Lambda@Edge function created in a different stack
//   // to a CloudFront distribution.
//   var s3Bucket bucket
//
//   functionVersion := lambda.version.fromVersionArn(this, jsii.String("Version"), jsii.String("arn:aws:lambda:us-east-1:123456789012:function:functionName:1"))
//
//   cloudfront.NewDistribution(this, jsii.String("distro"), &distributionProps{
//   	defaultBehavior: &behaviorOptions{
//   		origin: origins.NewS3Origin(s3Bucket),
//   		edgeLambdas: []edgeLambda{
//   			&edgeLambda{
//   				functionVersion: functionVersion,
//   				eventType: cloudfront.lambdaEdgeEventType_VIEWER_REQUEST,
//   			},
//   		},
//   	},
//   })
//
type S3Origin interface {
	awscloudfront.IOrigin
	// The method called when a given Origin is added (for the first time) to a Distribution.
	Bind(scope constructs.Construct, options *awscloudfront.OriginBindOptions) *awscloudfront.OriginBindConfig
}

// The jsii proxy struct for S3Origin
type jsiiProxy_S3Origin struct {
	internal.Type__awscloudfrontIOrigin
}

func NewS3Origin(bucket awss3.IBucket, props *S3OriginProps) S3Origin {
	_init_.Initialize()

	j := jsiiProxy_S3Origin{}

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront_origins.S3Origin",
		[]interface{}{bucket, props},
		&j,
	)

	return &j
}

func NewS3Origin_Override(s S3Origin, bucket awss3.IBucket, props *S3OriginProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront_origins.S3Origin",
		[]interface{}{bucket, props},
		s,
	)
}

func (s *jsiiProxy_S3Origin) Bind(scope constructs.Construct, options *awscloudfront.OriginBindOptions) *awscloudfront.OriginBindConfig {
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
//
// Example:
//   myBucket := s3.NewBucket(this, jsii.String("myBucket"))
//   cloudfront.NewDistribution(this, jsii.String("myDist"), &distributionProps{
//   	defaultBehavior: &behaviorOptions{
//   		origin: origins.NewS3Origin(myBucket, &s3OriginProps{
//   			customHeaders: map[string]*string{
//   				"Foo": jsii.String("bar"),
//   			},
//   		}),
//   	},
//   })
//
type S3OriginProps struct {
	// The number of times that CloudFront attempts to connect to the origin;
	//
	// valid values are 1, 2, or 3 attempts.
	ConnectionAttempts *float64 `field:"optional" json:"connectionAttempts" yaml:"connectionAttempts"`
	// The number of seconds that CloudFront waits when trying to establish a connection to the origin.
	//
	// Valid values are 1-10 seconds, inclusive.
	ConnectionTimeout awscdk.Duration `field:"optional" json:"connectionTimeout" yaml:"connectionTimeout"`
	// A list of HTTP header names and values that CloudFront adds to requests it sends to the origin.
	CustomHeaders *map[string]*string `field:"optional" json:"customHeaders" yaml:"customHeaders"`
	// An optional path that CloudFront appends to the origin domain name when CloudFront requests content from the origin.
	//
	// Must begin, but not end, with '/' (e.g., '/production/images').
	OriginPath *string `field:"optional" json:"originPath" yaml:"originPath"`
	// When you enable Origin Shield in the AWS Region that has the lowest latency to your origin, you can get better network performance.
	// See: https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/origin-shield.html
	//
	OriginShieldRegion *string `field:"optional" json:"originShieldRegion" yaml:"originShieldRegion"`
	// An optional Origin Access Identity of the origin identity cloudfront will use when calling your s3 bucket.
	OriginAccessIdentity awscloudfront.IOriginAccessIdentity `field:"optional" json:"originAccessIdentity" yaml:"originAccessIdentity"`
}

