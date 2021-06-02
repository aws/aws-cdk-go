package awsroute53targets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/awscloudfront"
	"github.com/aws/aws-cdk-go/awscdk/awscognito"
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awselasticloadbalancing"
	"github.com/aws/aws-cdk-go/awscdk/awselasticloadbalancingv2"
	"github.com/aws/aws-cdk-go/awscdk/awsglobalaccelerator"
	"github.com/aws/aws-cdk-go/awscdk/awsroute53"
	"github.com/aws/aws-cdk-go/awscdk/awsroute53targets/internal"
	"github.com/aws/aws-cdk-go/awscdk/awss3"
)

// Defines an API Gateway REST API as the alias target. Requires that the domain name will be defined through `RestApiProps.domainName`.
//
// You can direct the alias to any `apigateway.DomainName` resource through the
// `ApiGatewayDomain` class.
// Experimental.
type ApiGateway interface {
	ApiGatewayDomain
	Bind(_record awsroute53.IRecordSet) *awsroute53.AliasRecordTargetConfig
}

// The jsii proxy struct for ApiGateway
type jsiiProxy_ApiGateway struct {
	jsiiProxy_ApiGatewayDomain
}

// Experimental.
func NewApiGateway(api awsapigateway.RestApi) ApiGateway {
	_init_.Initialize()

	j := jsiiProxy_ApiGateway{}

	_jsii_.Create(
		"monocdk.aws_route53_targets.ApiGateway",
		[]interface{}{api},
		&j,
	)

	return &j
}

// Experimental.
func NewApiGateway_Override(a ApiGateway, api awsapigateway.RestApi) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_route53_targets.ApiGateway",
		[]interface{}{api},
		a,
	)
}

// Return hosted zone ID and DNS name, usable for Route53 alias targets.
// Experimental.
func (a *jsiiProxy_ApiGateway) Bind(_record awsroute53.IRecordSet) *awsroute53.AliasRecordTargetConfig {
	var returns *awsroute53.AliasRecordTargetConfig

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{_record},
		&returns,
	)

	return returns
}

// Defines an API Gateway domain name as the alias target.
//
// Use the `ApiGateway` class if you wish to map the alias to an REST API with a
// domain name defined through the `RestApiProps.domainName` prop.
// Experimental.
type ApiGatewayDomain interface {
	awsroute53.IAliasRecordTarget
	Bind(_record awsroute53.IRecordSet) *awsroute53.AliasRecordTargetConfig
}

// The jsii proxy struct for ApiGatewayDomain
type jsiiProxy_ApiGatewayDomain struct {
	internal.Type__awsroute53IAliasRecordTarget
}

// Experimental.
func NewApiGatewayDomain(domainName awsapigateway.IDomainName) ApiGatewayDomain {
	_init_.Initialize()

	j := jsiiProxy_ApiGatewayDomain{}

	_jsii_.Create(
		"monocdk.aws_route53_targets.ApiGatewayDomain",
		[]interface{}{domainName},
		&j,
	)

	return &j
}

// Experimental.
func NewApiGatewayDomain_Override(a ApiGatewayDomain, domainName awsapigateway.IDomainName) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_route53_targets.ApiGatewayDomain",
		[]interface{}{domainName},
		a,
	)
}

// Return hosted zone ID and DNS name, usable for Route53 alias targets.
// Experimental.
func (a *jsiiProxy_ApiGatewayDomain) Bind(_record awsroute53.IRecordSet) *awsroute53.AliasRecordTargetConfig {
	var returns *awsroute53.AliasRecordTargetConfig

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{_record},
		&returns,
	)

	return returns
}

// Defines an API Gateway V2 domain name as the alias target.
// Experimental.
type ApiGatewayv2DomainProperties interface {
	awsroute53.IAliasRecordTarget
	Bind(_record awsroute53.IRecordSet) *awsroute53.AliasRecordTargetConfig
}

// The jsii proxy struct for ApiGatewayv2DomainProperties
type jsiiProxy_ApiGatewayv2DomainProperties struct {
	internal.Type__awsroute53IAliasRecordTarget
}

// Experimental.
func NewApiGatewayv2DomainProperties(regionalDomainName *string, regionalHostedZoneId *string) ApiGatewayv2DomainProperties {
	_init_.Initialize()

	j := jsiiProxy_ApiGatewayv2DomainProperties{}

	_jsii_.Create(
		"monocdk.aws_route53_targets.ApiGatewayv2DomainProperties",
		[]interface{}{regionalDomainName, regionalHostedZoneId},
		&j,
	)

	return &j
}

// Experimental.
func NewApiGatewayv2DomainProperties_Override(a ApiGatewayv2DomainProperties, regionalDomainName *string, regionalHostedZoneId *string) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_route53_targets.ApiGatewayv2DomainProperties",
		[]interface{}{regionalDomainName, regionalHostedZoneId},
		a,
	)
}

// Return hosted zone ID and DNS name, usable for Route53 alias targets.
// Experimental.
func (a *jsiiProxy_ApiGatewayv2DomainProperties) Bind(_record awsroute53.IRecordSet) *awsroute53.AliasRecordTargetConfig {
	var returns *awsroute53.AliasRecordTargetConfig

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{_record},
		&returns,
	)

	return returns
}

// Use a S3 as an alias record target.
// Experimental.
type BucketWebsiteTarget interface {
	awsroute53.IAliasRecordTarget
	Bind(_record awsroute53.IRecordSet) *awsroute53.AliasRecordTargetConfig
}

// The jsii proxy struct for BucketWebsiteTarget
type jsiiProxy_BucketWebsiteTarget struct {
	internal.Type__awsroute53IAliasRecordTarget
}

// Experimental.
func NewBucketWebsiteTarget(bucket awss3.IBucket) BucketWebsiteTarget {
	_init_.Initialize()

	j := jsiiProxy_BucketWebsiteTarget{}

	_jsii_.Create(
		"monocdk.aws_route53_targets.BucketWebsiteTarget",
		[]interface{}{bucket},
		&j,
	)

	return &j
}

// Experimental.
func NewBucketWebsiteTarget_Override(b BucketWebsiteTarget, bucket awss3.IBucket) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_route53_targets.BucketWebsiteTarget",
		[]interface{}{bucket},
		b,
	)
}

// Return hosted zone ID and DNS name, usable for Route53 alias targets.
// Experimental.
func (b *jsiiProxy_BucketWebsiteTarget) Bind(_record awsroute53.IRecordSet) *awsroute53.AliasRecordTargetConfig {
	var returns *awsroute53.AliasRecordTargetConfig

	_jsii_.Invoke(
		b,
		"bind",
		[]interface{}{_record},
		&returns,
	)

	return returns
}

// Use a classic ELB as an alias record target.
// Experimental.
type ClassicLoadBalancerTarget interface {
	awsroute53.IAliasRecordTarget
	Bind(_record awsroute53.IRecordSet) *awsroute53.AliasRecordTargetConfig
}

// The jsii proxy struct for ClassicLoadBalancerTarget
type jsiiProxy_ClassicLoadBalancerTarget struct {
	internal.Type__awsroute53IAliasRecordTarget
}

// Experimental.
func NewClassicLoadBalancerTarget(loadBalancer awselasticloadbalancing.LoadBalancer) ClassicLoadBalancerTarget {
	_init_.Initialize()

	j := jsiiProxy_ClassicLoadBalancerTarget{}

	_jsii_.Create(
		"monocdk.aws_route53_targets.ClassicLoadBalancerTarget",
		[]interface{}{loadBalancer},
		&j,
	)

	return &j
}

// Experimental.
func NewClassicLoadBalancerTarget_Override(c ClassicLoadBalancerTarget, loadBalancer awselasticloadbalancing.LoadBalancer) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_route53_targets.ClassicLoadBalancerTarget",
		[]interface{}{loadBalancer},
		c,
	)
}

// Return hosted zone ID and DNS name, usable for Route53 alias targets.
// Experimental.
func (c *jsiiProxy_ClassicLoadBalancerTarget) Bind(_record awsroute53.IRecordSet) *awsroute53.AliasRecordTargetConfig {
	var returns *awsroute53.AliasRecordTargetConfig

	_jsii_.Invoke(
		c,
		"bind",
		[]interface{}{_record},
		&returns,
	)

	return returns
}

// Use a CloudFront Distribution as an alias record target.
// Experimental.
type CloudFrontTarget interface {
	awsroute53.IAliasRecordTarget
	Bind(_record awsroute53.IRecordSet) *awsroute53.AliasRecordTargetConfig
}

// The jsii proxy struct for CloudFrontTarget
type jsiiProxy_CloudFrontTarget struct {
	internal.Type__awsroute53IAliasRecordTarget
}

// Experimental.
func NewCloudFrontTarget(distribution awscloudfront.IDistribution) CloudFrontTarget {
	_init_.Initialize()

	j := jsiiProxy_CloudFrontTarget{}

	_jsii_.Create(
		"monocdk.aws_route53_targets.CloudFrontTarget",
		[]interface{}{distribution},
		&j,
	)

	return &j
}

// Experimental.
func NewCloudFrontTarget_Override(c CloudFrontTarget, distribution awscloudfront.IDistribution) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_route53_targets.CloudFrontTarget",
		[]interface{}{distribution},
		c,
	)
}

// Get the hosted zone id for the current scope.
// Experimental.
func CloudFrontTarget_GetHostedZoneId(scope awscdk.IConstruct) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"monocdk.aws_route53_targets.CloudFrontTarget",
		"getHostedZoneId",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

func CloudFrontTarget_CLOUDFRONT_ZONE_ID() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_route53_targets.CloudFrontTarget",
		"CLOUDFRONT_ZONE_ID",
		&returns,
	)
	return returns
}

// Return hosted zone ID and DNS name, usable for Route53 alias targets.
// Experimental.
func (c *jsiiProxy_CloudFrontTarget) Bind(_record awsroute53.IRecordSet) *awsroute53.AliasRecordTargetConfig {
	var returns *awsroute53.AliasRecordTargetConfig

	_jsii_.Invoke(
		c,
		"bind",
		[]interface{}{_record},
		&returns,
	)

	return returns
}

// Use a Global Accelerator domain name as an alias record target.
// Experimental.
type GlobalAcceleratorDomainTarget interface {
	awsroute53.IAliasRecordTarget
	Bind(_record awsroute53.IRecordSet) *awsroute53.AliasRecordTargetConfig
}

// The jsii proxy struct for GlobalAcceleratorDomainTarget
type jsiiProxy_GlobalAcceleratorDomainTarget struct {
	internal.Type__awsroute53IAliasRecordTarget
}

// Create an Alias Target for a Global Accelerator domain name.
// Experimental.
func NewGlobalAcceleratorDomainTarget(acceleratorDomainName *string) GlobalAcceleratorDomainTarget {
	_init_.Initialize()

	j := jsiiProxy_GlobalAcceleratorDomainTarget{}

	_jsii_.Create(
		"monocdk.aws_route53_targets.GlobalAcceleratorDomainTarget",
		[]interface{}{acceleratorDomainName},
		&j,
	)

	return &j
}

// Create an Alias Target for a Global Accelerator domain name.
// Experimental.
func NewGlobalAcceleratorDomainTarget_Override(g GlobalAcceleratorDomainTarget, acceleratorDomainName *string) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_route53_targets.GlobalAcceleratorDomainTarget",
		[]interface{}{acceleratorDomainName},
		g,
	)
}

func GlobalAcceleratorDomainTarget_GLOBAL_ACCELERATOR_ZONE_ID() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_route53_targets.GlobalAcceleratorDomainTarget",
		"GLOBAL_ACCELERATOR_ZONE_ID",
		&returns,
	)
	return returns
}

// Return hosted zone ID and DNS name, usable for Route53 alias targets.
// Experimental.
func (g *jsiiProxy_GlobalAcceleratorDomainTarget) Bind(_record awsroute53.IRecordSet) *awsroute53.AliasRecordTargetConfig {
	var returns *awsroute53.AliasRecordTargetConfig

	_jsii_.Invoke(
		g,
		"bind",
		[]interface{}{_record},
		&returns,
	)

	return returns
}

// Use a Global Accelerator instance domain name as an alias record target.
// Experimental.
type GlobalAcceleratorTarget interface {
	GlobalAcceleratorDomainTarget
	Bind(_record awsroute53.IRecordSet) *awsroute53.AliasRecordTargetConfig
}

// The jsii proxy struct for GlobalAcceleratorTarget
type jsiiProxy_GlobalAcceleratorTarget struct {
	jsiiProxy_GlobalAcceleratorDomainTarget
}

// Create an Alias Target for a Global Accelerator instance.
// Experimental.
func NewGlobalAcceleratorTarget(accelerator awsglobalaccelerator.IAccelerator) GlobalAcceleratorTarget {
	_init_.Initialize()

	j := jsiiProxy_GlobalAcceleratorTarget{}

	_jsii_.Create(
		"monocdk.aws_route53_targets.GlobalAcceleratorTarget",
		[]interface{}{accelerator},
		&j,
	)

	return &j
}

// Create an Alias Target for a Global Accelerator instance.
// Experimental.
func NewGlobalAcceleratorTarget_Override(g GlobalAcceleratorTarget, accelerator awsglobalaccelerator.IAccelerator) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_route53_targets.GlobalAcceleratorTarget",
		[]interface{}{accelerator},
		g,
	)
}

func GlobalAcceleratorTarget_GLOBAL_ACCELERATOR_ZONE_ID() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_route53_targets.GlobalAcceleratorTarget",
		"GLOBAL_ACCELERATOR_ZONE_ID",
		&returns,
	)
	return returns
}

// Return hosted zone ID and DNS name, usable for Route53 alias targets.
// Experimental.
func (g *jsiiProxy_GlobalAcceleratorTarget) Bind(_record awsroute53.IRecordSet) *awsroute53.AliasRecordTargetConfig {
	var returns *awsroute53.AliasRecordTargetConfig

	_jsii_.Invoke(
		g,
		"bind",
		[]interface{}{_record},
		&returns,
	)

	return returns
}

// Set an InterfaceVpcEndpoint as a target for an ARecord.
// Experimental.
type InterfaceVpcEndpointTarget interface {
	awsroute53.IAliasRecordTarget
	Bind(_record awsroute53.IRecordSet) *awsroute53.AliasRecordTargetConfig
}

// The jsii proxy struct for InterfaceVpcEndpointTarget
type jsiiProxy_InterfaceVpcEndpointTarget struct {
	internal.Type__awsroute53IAliasRecordTarget
}

// Experimental.
func NewInterfaceVpcEndpointTarget(vpcEndpoint awsec2.IInterfaceVpcEndpoint) InterfaceVpcEndpointTarget {
	_init_.Initialize()

	j := jsiiProxy_InterfaceVpcEndpointTarget{}

	_jsii_.Create(
		"monocdk.aws_route53_targets.InterfaceVpcEndpointTarget",
		[]interface{}{vpcEndpoint},
		&j,
	)

	return &j
}

// Experimental.
func NewInterfaceVpcEndpointTarget_Override(i InterfaceVpcEndpointTarget, vpcEndpoint awsec2.IInterfaceVpcEndpoint) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_route53_targets.InterfaceVpcEndpointTarget",
		[]interface{}{vpcEndpoint},
		i,
	)
}

// Return hosted zone ID and DNS name, usable for Route53 alias targets.
// Experimental.
func (i *jsiiProxy_InterfaceVpcEndpointTarget) Bind(_record awsroute53.IRecordSet) *awsroute53.AliasRecordTargetConfig {
	var returns *awsroute53.AliasRecordTargetConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{_record},
		&returns,
	)

	return returns
}

// Use an ELBv2 as an alias record target.
// Experimental.
type LoadBalancerTarget interface {
	awsroute53.IAliasRecordTarget
	Bind(_record awsroute53.IRecordSet) *awsroute53.AliasRecordTargetConfig
}

// The jsii proxy struct for LoadBalancerTarget
type jsiiProxy_LoadBalancerTarget struct {
	internal.Type__awsroute53IAliasRecordTarget
}

// Experimental.
func NewLoadBalancerTarget(loadBalancer awselasticloadbalancingv2.ILoadBalancerV2) LoadBalancerTarget {
	_init_.Initialize()

	j := jsiiProxy_LoadBalancerTarget{}

	_jsii_.Create(
		"monocdk.aws_route53_targets.LoadBalancerTarget",
		[]interface{}{loadBalancer},
		&j,
	)

	return &j
}

// Experimental.
func NewLoadBalancerTarget_Override(l LoadBalancerTarget, loadBalancer awselasticloadbalancingv2.ILoadBalancerV2) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_route53_targets.LoadBalancerTarget",
		[]interface{}{loadBalancer},
		l,
	)
}

// Return hosted zone ID and DNS name, usable for Route53 alias targets.
// Experimental.
func (l *jsiiProxy_LoadBalancerTarget) Bind(_record awsroute53.IRecordSet) *awsroute53.AliasRecordTargetConfig {
	var returns *awsroute53.AliasRecordTargetConfig

	_jsii_.Invoke(
		l,
		"bind",
		[]interface{}{_record},
		&returns,
	)

	return returns
}

// Use a user pool domain as an alias record target.
// Experimental.
type UserPoolDomainTarget interface {
	awsroute53.IAliasRecordTarget
	Bind(_record awsroute53.IRecordSet) *awsroute53.AliasRecordTargetConfig
}

// The jsii proxy struct for UserPoolDomainTarget
type jsiiProxy_UserPoolDomainTarget struct {
	internal.Type__awsroute53IAliasRecordTarget
}

// Experimental.
func NewUserPoolDomainTarget(domain awscognito.UserPoolDomain) UserPoolDomainTarget {
	_init_.Initialize()

	j := jsiiProxy_UserPoolDomainTarget{}

	_jsii_.Create(
		"monocdk.aws_route53_targets.UserPoolDomainTarget",
		[]interface{}{domain},
		&j,
	)

	return &j
}

// Experimental.
func NewUserPoolDomainTarget_Override(u UserPoolDomainTarget, domain awscognito.UserPoolDomain) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_route53_targets.UserPoolDomainTarget",
		[]interface{}{domain},
		u,
	)
}

// Return hosted zone ID and DNS name, usable for Route53 alias targets.
// Experimental.
func (u *jsiiProxy_UserPoolDomainTarget) Bind(_record awsroute53.IRecordSet) *awsroute53.AliasRecordTargetConfig {
	var returns *awsroute53.AliasRecordTargetConfig

	_jsii_.Invoke(
		u,
		"bind",
		[]interface{}{_record},
		&returns,
	)

	return returns
}

