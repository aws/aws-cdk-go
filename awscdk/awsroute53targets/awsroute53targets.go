package awsroute53targets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancing"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsglobalaccelerator"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53targets/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
)

// Defines an API Gateway REST API as the alias target. Requires that the domain name will be defined through `RestApiProps.domainName`.
//
// You can direct the alias to any `apigateway.DomainName` resource through the
// `ApiGatewayDomain` class.
//
// TODO: EXAMPLE
//
type ApiGateway interface {
	ApiGatewayDomain
	Bind(_record awsroute53.IRecordSet, _zone awsroute53.IHostedZone) *awsroute53.AliasRecordTargetConfig
}

// The jsii proxy struct for ApiGateway
type jsiiProxy_ApiGateway struct {
	jsiiProxy_ApiGatewayDomain
}

func NewApiGateway(api awsapigateway.RestApiBase) ApiGateway {
	_init_.Initialize()

	j := jsiiProxy_ApiGateway{}

	_jsii_.Create(
		"aws-cdk-lib.aws_route53_targets.ApiGateway",
		[]interface{}{api},
		&j,
	)

	return &j
}

func NewApiGateway_Override(a ApiGateway, api awsapigateway.RestApiBase) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_route53_targets.ApiGateway",
		[]interface{}{api},
		a,
	)
}

// Return hosted zone ID and DNS name, usable for Route53 alias targets.
func (a *jsiiProxy_ApiGateway) Bind(_record awsroute53.IRecordSet, _zone awsroute53.IHostedZone) *awsroute53.AliasRecordTargetConfig {
	var returns *awsroute53.AliasRecordTargetConfig

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{_record, _zone},
		&returns,
	)

	return returns
}

// Defines an API Gateway domain name as the alias target.
//
// Use the `ApiGateway` class if you wish to map the alias to an REST API with a
// domain name defined through the `RestApiProps.domainName` prop.
//
// TODO: EXAMPLE
//
type ApiGatewayDomain interface {
	awsroute53.IAliasRecordTarget
	Bind(_record awsroute53.IRecordSet, _zone awsroute53.IHostedZone) *awsroute53.AliasRecordTargetConfig
}

// The jsii proxy struct for ApiGatewayDomain
type jsiiProxy_ApiGatewayDomain struct {
	internal.Type__awsroute53IAliasRecordTarget
}

func NewApiGatewayDomain(domainName awsapigateway.IDomainName) ApiGatewayDomain {
	_init_.Initialize()

	j := jsiiProxy_ApiGatewayDomain{}

	_jsii_.Create(
		"aws-cdk-lib.aws_route53_targets.ApiGatewayDomain",
		[]interface{}{domainName},
		&j,
	)

	return &j
}

func NewApiGatewayDomain_Override(a ApiGatewayDomain, domainName awsapigateway.IDomainName) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_route53_targets.ApiGatewayDomain",
		[]interface{}{domainName},
		a,
	)
}

// Return hosted zone ID and DNS name, usable for Route53 alias targets.
func (a *jsiiProxy_ApiGatewayDomain) Bind(_record awsroute53.IRecordSet, _zone awsroute53.IHostedZone) *awsroute53.AliasRecordTargetConfig {
	var returns *awsroute53.AliasRecordTargetConfig

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{_record, _zone},
		&returns,
	)

	return returns
}

// Defines an API Gateway V2 domain name as the alias target.
//
// TODO: EXAMPLE
//
type ApiGatewayv2DomainProperties interface {
	awsroute53.IAliasRecordTarget
	Bind(_record awsroute53.IRecordSet, _zone awsroute53.IHostedZone) *awsroute53.AliasRecordTargetConfig
}

// The jsii proxy struct for ApiGatewayv2DomainProperties
type jsiiProxy_ApiGatewayv2DomainProperties struct {
	internal.Type__awsroute53IAliasRecordTarget
}

func NewApiGatewayv2DomainProperties(regionalDomainName *string, regionalHostedZoneId *string) ApiGatewayv2DomainProperties {
	_init_.Initialize()

	j := jsiiProxy_ApiGatewayv2DomainProperties{}

	_jsii_.Create(
		"aws-cdk-lib.aws_route53_targets.ApiGatewayv2DomainProperties",
		[]interface{}{regionalDomainName, regionalHostedZoneId},
		&j,
	)

	return &j
}

func NewApiGatewayv2DomainProperties_Override(a ApiGatewayv2DomainProperties, regionalDomainName *string, regionalHostedZoneId *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_route53_targets.ApiGatewayv2DomainProperties",
		[]interface{}{regionalDomainName, regionalHostedZoneId},
		a,
	)
}

// Return hosted zone ID and DNS name, usable for Route53 alias targets.
func (a *jsiiProxy_ApiGatewayv2DomainProperties) Bind(_record awsroute53.IRecordSet, _zone awsroute53.IHostedZone) *awsroute53.AliasRecordTargetConfig {
	var returns *awsroute53.AliasRecordTargetConfig

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{_record, _zone},
		&returns,
	)

	return returns
}

// Use a S3 as an alias record target.
//
// TODO: EXAMPLE
//
type BucketWebsiteTarget interface {
	awsroute53.IAliasRecordTarget
	Bind(_record awsroute53.IRecordSet, _zone awsroute53.IHostedZone) *awsroute53.AliasRecordTargetConfig
}

// The jsii proxy struct for BucketWebsiteTarget
type jsiiProxy_BucketWebsiteTarget struct {
	internal.Type__awsroute53IAliasRecordTarget
}

func NewBucketWebsiteTarget(bucket awss3.IBucket) BucketWebsiteTarget {
	_init_.Initialize()

	j := jsiiProxy_BucketWebsiteTarget{}

	_jsii_.Create(
		"aws-cdk-lib.aws_route53_targets.BucketWebsiteTarget",
		[]interface{}{bucket},
		&j,
	)

	return &j
}

func NewBucketWebsiteTarget_Override(b BucketWebsiteTarget, bucket awss3.IBucket) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_route53_targets.BucketWebsiteTarget",
		[]interface{}{bucket},
		b,
	)
}

// Return hosted zone ID and DNS name, usable for Route53 alias targets.
func (b *jsiiProxy_BucketWebsiteTarget) Bind(_record awsroute53.IRecordSet, _zone awsroute53.IHostedZone) *awsroute53.AliasRecordTargetConfig {
	var returns *awsroute53.AliasRecordTargetConfig

	_jsii_.Invoke(
		b,
		"bind",
		[]interface{}{_record, _zone},
		&returns,
	)

	return returns
}

// Use a classic ELB as an alias record target.
//
// TODO: EXAMPLE
//
type ClassicLoadBalancerTarget interface {
	awsroute53.IAliasRecordTarget
	Bind(_record awsroute53.IRecordSet, _zone awsroute53.IHostedZone) *awsroute53.AliasRecordTargetConfig
}

// The jsii proxy struct for ClassicLoadBalancerTarget
type jsiiProxy_ClassicLoadBalancerTarget struct {
	internal.Type__awsroute53IAliasRecordTarget
}

func NewClassicLoadBalancerTarget(loadBalancer awselasticloadbalancing.LoadBalancer) ClassicLoadBalancerTarget {
	_init_.Initialize()

	j := jsiiProxy_ClassicLoadBalancerTarget{}

	_jsii_.Create(
		"aws-cdk-lib.aws_route53_targets.ClassicLoadBalancerTarget",
		[]interface{}{loadBalancer},
		&j,
	)

	return &j
}

func NewClassicLoadBalancerTarget_Override(c ClassicLoadBalancerTarget, loadBalancer awselasticloadbalancing.LoadBalancer) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_route53_targets.ClassicLoadBalancerTarget",
		[]interface{}{loadBalancer},
		c,
	)
}

// Return hosted zone ID and DNS name, usable for Route53 alias targets.
func (c *jsiiProxy_ClassicLoadBalancerTarget) Bind(_record awsroute53.IRecordSet, _zone awsroute53.IHostedZone) *awsroute53.AliasRecordTargetConfig {
	var returns *awsroute53.AliasRecordTargetConfig

	_jsii_.Invoke(
		c,
		"bind",
		[]interface{}{_record, _zone},
		&returns,
	)

	return returns
}

// Use a CloudFront Distribution as an alias record target.
//
// TODO: EXAMPLE
//
type CloudFrontTarget interface {
	awsroute53.IAliasRecordTarget
	Bind(_record awsroute53.IRecordSet, _zone awsroute53.IHostedZone) *awsroute53.AliasRecordTargetConfig
}

// The jsii proxy struct for CloudFrontTarget
type jsiiProxy_CloudFrontTarget struct {
	internal.Type__awsroute53IAliasRecordTarget
}

func NewCloudFrontTarget(distribution awscloudfront.IDistribution) CloudFrontTarget {
	_init_.Initialize()

	j := jsiiProxy_CloudFrontTarget{}

	_jsii_.Create(
		"aws-cdk-lib.aws_route53_targets.CloudFrontTarget",
		[]interface{}{distribution},
		&j,
	)

	return &j
}

func NewCloudFrontTarget_Override(c CloudFrontTarget, distribution awscloudfront.IDistribution) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_route53_targets.CloudFrontTarget",
		[]interface{}{distribution},
		c,
	)
}

// Get the hosted zone id for the current scope.
func CloudFrontTarget_GetHostedZoneId(scope constructs.IConstruct) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53_targets.CloudFrontTarget",
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
		"aws-cdk-lib.aws_route53_targets.CloudFrontTarget",
		"CLOUDFRONT_ZONE_ID",
		&returns,
	)
	return returns
}

// Return hosted zone ID and DNS name, usable for Route53 alias targets.
func (c *jsiiProxy_CloudFrontTarget) Bind(_record awsroute53.IRecordSet, _zone awsroute53.IHostedZone) *awsroute53.AliasRecordTargetConfig {
	var returns *awsroute53.AliasRecordTargetConfig

	_jsii_.Invoke(
		c,
		"bind",
		[]interface{}{_record, _zone},
		&returns,
	)

	return returns
}

// Use an Elastic Beanstalk environment URL as an alias record target. E.g. mysampleenvironment.xyz.us-east-1.elasticbeanstalk.com or mycustomcnameprefix.us-east-1.elasticbeanstalk.com.
//
// Only supports Elastic Beanstalk environments created after 2016 that have a regional endpoint.
//
// TODO: EXAMPLE
//
type ElasticBeanstalkEnvironmentEndpointTarget interface {
	awsroute53.IAliasRecordTarget
	Bind(_record awsroute53.IRecordSet, _zone awsroute53.IHostedZone) *awsroute53.AliasRecordTargetConfig
}

// The jsii proxy struct for ElasticBeanstalkEnvironmentEndpointTarget
type jsiiProxy_ElasticBeanstalkEnvironmentEndpointTarget struct {
	internal.Type__awsroute53IAliasRecordTarget
}

func NewElasticBeanstalkEnvironmentEndpointTarget(environmentEndpoint *string) ElasticBeanstalkEnvironmentEndpointTarget {
	_init_.Initialize()

	j := jsiiProxy_ElasticBeanstalkEnvironmentEndpointTarget{}

	_jsii_.Create(
		"aws-cdk-lib.aws_route53_targets.ElasticBeanstalkEnvironmentEndpointTarget",
		[]interface{}{environmentEndpoint},
		&j,
	)

	return &j
}

func NewElasticBeanstalkEnvironmentEndpointTarget_Override(e ElasticBeanstalkEnvironmentEndpointTarget, environmentEndpoint *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_route53_targets.ElasticBeanstalkEnvironmentEndpointTarget",
		[]interface{}{environmentEndpoint},
		e,
	)
}

// Return hosted zone ID and DNS name, usable for Route53 alias targets.
func (e *jsiiProxy_ElasticBeanstalkEnvironmentEndpointTarget) Bind(_record awsroute53.IRecordSet, _zone awsroute53.IHostedZone) *awsroute53.AliasRecordTargetConfig {
	var returns *awsroute53.AliasRecordTargetConfig

	_jsii_.Invoke(
		e,
		"bind",
		[]interface{}{_record, _zone},
		&returns,
	)

	return returns
}

// Use a Global Accelerator domain name as an alias record target.
//
// TODO: EXAMPLE
//
type GlobalAcceleratorDomainTarget interface {
	awsroute53.IAliasRecordTarget
	Bind(_record awsroute53.IRecordSet, _zone awsroute53.IHostedZone) *awsroute53.AliasRecordTargetConfig
}

// The jsii proxy struct for GlobalAcceleratorDomainTarget
type jsiiProxy_GlobalAcceleratorDomainTarget struct {
	internal.Type__awsroute53IAliasRecordTarget
}

// Create an Alias Target for a Global Accelerator domain name.
func NewGlobalAcceleratorDomainTarget(acceleratorDomainName *string) GlobalAcceleratorDomainTarget {
	_init_.Initialize()

	j := jsiiProxy_GlobalAcceleratorDomainTarget{}

	_jsii_.Create(
		"aws-cdk-lib.aws_route53_targets.GlobalAcceleratorDomainTarget",
		[]interface{}{acceleratorDomainName},
		&j,
	)

	return &j
}

// Create an Alias Target for a Global Accelerator domain name.
func NewGlobalAcceleratorDomainTarget_Override(g GlobalAcceleratorDomainTarget, acceleratorDomainName *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_route53_targets.GlobalAcceleratorDomainTarget",
		[]interface{}{acceleratorDomainName},
		g,
	)
}

func GlobalAcceleratorDomainTarget_GLOBAL_ACCELERATOR_ZONE_ID() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_route53_targets.GlobalAcceleratorDomainTarget",
		"GLOBAL_ACCELERATOR_ZONE_ID",
		&returns,
	)
	return returns
}

// Return hosted zone ID and DNS name, usable for Route53 alias targets.
func (g *jsiiProxy_GlobalAcceleratorDomainTarget) Bind(_record awsroute53.IRecordSet, _zone awsroute53.IHostedZone) *awsroute53.AliasRecordTargetConfig {
	var returns *awsroute53.AliasRecordTargetConfig

	_jsii_.Invoke(
		g,
		"bind",
		[]interface{}{_record, _zone},
		&returns,
	)

	return returns
}

// Use a Global Accelerator instance domain name as an alias record target.
//
// TODO: EXAMPLE
//
type GlobalAcceleratorTarget interface {
	GlobalAcceleratorDomainTarget
	Bind(_record awsroute53.IRecordSet, _zone awsroute53.IHostedZone) *awsroute53.AliasRecordTargetConfig
}

// The jsii proxy struct for GlobalAcceleratorTarget
type jsiiProxy_GlobalAcceleratorTarget struct {
	jsiiProxy_GlobalAcceleratorDomainTarget
}

// Create an Alias Target for a Global Accelerator instance.
func NewGlobalAcceleratorTarget(accelerator awsglobalaccelerator.IAccelerator) GlobalAcceleratorTarget {
	_init_.Initialize()

	j := jsiiProxy_GlobalAcceleratorTarget{}

	_jsii_.Create(
		"aws-cdk-lib.aws_route53_targets.GlobalAcceleratorTarget",
		[]interface{}{accelerator},
		&j,
	)

	return &j
}

// Create an Alias Target for a Global Accelerator instance.
func NewGlobalAcceleratorTarget_Override(g GlobalAcceleratorTarget, accelerator awsglobalaccelerator.IAccelerator) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_route53_targets.GlobalAcceleratorTarget",
		[]interface{}{accelerator},
		g,
	)
}

func GlobalAcceleratorTarget_GLOBAL_ACCELERATOR_ZONE_ID() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_route53_targets.GlobalAcceleratorTarget",
		"GLOBAL_ACCELERATOR_ZONE_ID",
		&returns,
	)
	return returns
}

// Return hosted zone ID and DNS name, usable for Route53 alias targets.
func (g *jsiiProxy_GlobalAcceleratorTarget) Bind(_record awsroute53.IRecordSet, _zone awsroute53.IHostedZone) *awsroute53.AliasRecordTargetConfig {
	var returns *awsroute53.AliasRecordTargetConfig

	_jsii_.Invoke(
		g,
		"bind",
		[]interface{}{_record, _zone},
		&returns,
	)

	return returns
}

// Set an InterfaceVpcEndpoint as a target for an ARecord.
//
// TODO: EXAMPLE
//
type InterfaceVpcEndpointTarget interface {
	awsroute53.IAliasRecordTarget
	Bind(_record awsroute53.IRecordSet, _zone awsroute53.IHostedZone) *awsroute53.AliasRecordTargetConfig
}

// The jsii proxy struct for InterfaceVpcEndpointTarget
type jsiiProxy_InterfaceVpcEndpointTarget struct {
	internal.Type__awsroute53IAliasRecordTarget
}

func NewInterfaceVpcEndpointTarget(vpcEndpoint awsec2.IInterfaceVpcEndpoint) InterfaceVpcEndpointTarget {
	_init_.Initialize()

	j := jsiiProxy_InterfaceVpcEndpointTarget{}

	_jsii_.Create(
		"aws-cdk-lib.aws_route53_targets.InterfaceVpcEndpointTarget",
		[]interface{}{vpcEndpoint},
		&j,
	)

	return &j
}

func NewInterfaceVpcEndpointTarget_Override(i InterfaceVpcEndpointTarget, vpcEndpoint awsec2.IInterfaceVpcEndpoint) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_route53_targets.InterfaceVpcEndpointTarget",
		[]interface{}{vpcEndpoint},
		i,
	)
}

// Return hosted zone ID and DNS name, usable for Route53 alias targets.
func (i *jsiiProxy_InterfaceVpcEndpointTarget) Bind(_record awsroute53.IRecordSet, _zone awsroute53.IHostedZone) *awsroute53.AliasRecordTargetConfig {
	var returns *awsroute53.AliasRecordTargetConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{_record, _zone},
		&returns,
	)

	return returns
}

// Use an ELBv2 as an alias record target.
//
// TODO: EXAMPLE
//
type LoadBalancerTarget interface {
	awsroute53.IAliasRecordTarget
	Bind(_record awsroute53.IRecordSet, _zone awsroute53.IHostedZone) *awsroute53.AliasRecordTargetConfig
}

// The jsii proxy struct for LoadBalancerTarget
type jsiiProxy_LoadBalancerTarget struct {
	internal.Type__awsroute53IAliasRecordTarget
}

func NewLoadBalancerTarget(loadBalancer awselasticloadbalancingv2.ILoadBalancerV2) LoadBalancerTarget {
	_init_.Initialize()

	j := jsiiProxy_LoadBalancerTarget{}

	_jsii_.Create(
		"aws-cdk-lib.aws_route53_targets.LoadBalancerTarget",
		[]interface{}{loadBalancer},
		&j,
	)

	return &j
}

func NewLoadBalancerTarget_Override(l LoadBalancerTarget, loadBalancer awselasticloadbalancingv2.ILoadBalancerV2) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_route53_targets.LoadBalancerTarget",
		[]interface{}{loadBalancer},
		l,
	)
}

// Return hosted zone ID and DNS name, usable for Route53 alias targets.
func (l *jsiiProxy_LoadBalancerTarget) Bind(_record awsroute53.IRecordSet, _zone awsroute53.IHostedZone) *awsroute53.AliasRecordTargetConfig {
	var returns *awsroute53.AliasRecordTargetConfig

	_jsii_.Invoke(
		l,
		"bind",
		[]interface{}{_record, _zone},
		&returns,
	)

	return returns
}

// Use another Route 53 record as an alias record target.
//
// TODO: EXAMPLE
//
type Route53RecordTarget interface {
	awsroute53.IAliasRecordTarget
	Bind(_record awsroute53.IRecordSet, zone awsroute53.IHostedZone) *awsroute53.AliasRecordTargetConfig
}

// The jsii proxy struct for Route53RecordTarget
type jsiiProxy_Route53RecordTarget struct {
	internal.Type__awsroute53IAliasRecordTarget
}

func NewRoute53RecordTarget(record awsroute53.IRecordSet) Route53RecordTarget {
	_init_.Initialize()

	j := jsiiProxy_Route53RecordTarget{}

	_jsii_.Create(
		"aws-cdk-lib.aws_route53_targets.Route53RecordTarget",
		[]interface{}{record},
		&j,
	)

	return &j
}

func NewRoute53RecordTarget_Override(r Route53RecordTarget, record awsroute53.IRecordSet) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_route53_targets.Route53RecordTarget",
		[]interface{}{record},
		r,
	)
}

// Return hosted zone ID and DNS name, usable for Route53 alias targets.
func (r *jsiiProxy_Route53RecordTarget) Bind(_record awsroute53.IRecordSet, zone awsroute53.IHostedZone) *awsroute53.AliasRecordTargetConfig {
	var returns *awsroute53.AliasRecordTargetConfig

	_jsii_.Invoke(
		r,
		"bind",
		[]interface{}{_record, zone},
		&returns,
	)

	return returns
}

// Use a user pool domain as an alias record target.
//
// TODO: EXAMPLE
//
type UserPoolDomainTarget interface {
	awsroute53.IAliasRecordTarget
	Bind(_record awsroute53.IRecordSet, _zone awsroute53.IHostedZone) *awsroute53.AliasRecordTargetConfig
}

// The jsii proxy struct for UserPoolDomainTarget
type jsiiProxy_UserPoolDomainTarget struct {
	internal.Type__awsroute53IAliasRecordTarget
}

func NewUserPoolDomainTarget(domain awscognito.UserPoolDomain) UserPoolDomainTarget {
	_init_.Initialize()

	j := jsiiProxy_UserPoolDomainTarget{}

	_jsii_.Create(
		"aws-cdk-lib.aws_route53_targets.UserPoolDomainTarget",
		[]interface{}{domain},
		&j,
	)

	return &j
}

func NewUserPoolDomainTarget_Override(u UserPoolDomainTarget, domain awscognito.UserPoolDomain) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_route53_targets.UserPoolDomainTarget",
		[]interface{}{domain},
		u,
	)
}

// Return hosted zone ID and DNS name, usable for Route53 alias targets.
func (u *jsiiProxy_UserPoolDomainTarget) Bind(_record awsroute53.IRecordSet, _zone awsroute53.IHostedZone) *awsroute53.AliasRecordTargetConfig {
	var returns *awsroute53.AliasRecordTargetConfig

	_jsii_.Invoke(
		u,
		"bind",
		[]interface{}{_record, _zone},
		&returns,
	)

	return returns
}

