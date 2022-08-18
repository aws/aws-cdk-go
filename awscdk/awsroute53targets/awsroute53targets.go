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
// Example:
//   import route53 "github.com/aws/aws-cdk-go/awscdk"
//   import targets "github.com/aws/aws-cdk-go/awscdk"
//
//   var api restApi
//   var hostedZoneForExampleCom interface{}
//
//
//   route53.NewARecord(this, jsii.String("CustomDomainAliasRecord"), &aRecordProps{
//   	zone: hostedZoneForExampleCom,
//   	target: route53.recordTarget.fromAlias(targets.NewApiGateway(api)),
//   })
//
type ApiGateway interface {
	ApiGatewayDomain
	// Return hosted zone ID and DNS name, usable for Route53 alias targets.
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
// Example:
//   var hostedZoneForExampleCom interface{}
//   var domainName domainName
//
//   import route53 "github.com/aws/aws-cdk-go/awscdk"
//   import targets "github.com/aws/aws-cdk-go/awscdk"
//
//
//   route53.NewARecord(this, jsii.String("CustomDomainAliasRecord"), &aRecordProps{
//   	zone: hostedZoneForExampleCom,
//   	target: route53.recordTarget.fromAlias(targets.NewApiGatewayDomain(domainName)),
//   })
//
type ApiGatewayDomain interface {
	awsroute53.IAliasRecordTarget
	// Return hosted zone ID and DNS name, usable for Route53 alias targets.
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
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   import apigwv2 "github.com/aws/aws-cdk-go/awscdk"
//
//   var zone hostedZone
//   var domainName apigwv2.DomainName
//
//
//   route53.NewARecord(this, jsii.String("AliasRecord"), &aRecordProps{
//   	zone: zone,
//   	target: route53.recordTarget.fromAlias(targets.NewApiGatewayv2DomainProperties(domainName.regionalDomainName, domainName.regionalHostedZoneId)),
//   })
//
type ApiGatewayv2DomainProperties interface {
	awsroute53.IAliasRecordTarget
	// Return hosted zone ID and DNS name, usable for Route53 alias targets.
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
// Example:
//   import s3 "github.com/aws/aws-cdk-go/awscdk"
//
//
//   recordName := "www"
//   domainName := "example.com"
//
//   bucketWebsite := s3.NewBucket(this, jsii.String("BucketWebsite"), &bucketProps{
//   	bucketName: []*string{
//   		recordName,
//   		domainName,
//   	}.join(jsii.String(".")),
//   	 // www.example.com
//   	publicReadAccess: jsii.Boolean(true),
//   	websiteIndexDocument: jsii.String("index.html"),
//   })
//
//   zone := route53.hostedZone.fromLookup(this, jsii.String("Zone"), &hostedZoneProviderProps{
//   	domainName: jsii.String(domainName),
//   }) // example.com
//
//    // example.com
//   route53.NewARecord(this, jsii.String("AliasRecord"), &aRecordProps{
//   	zone: zone,
//   	recordName: jsii.String(recordName),
//   	 // www
//   	target: route53.recordTarget.fromAlias(targets.NewBucketWebsiteTarget(bucketWebsite)),
//   })
//
type BucketWebsiteTarget interface {
	awsroute53.IAliasRecordTarget
	// Return hosted zone ID and DNS name, usable for Route53 alias targets.
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
// Example:
//   import elb "github.com/aws/aws-cdk-go/awscdk"
//
//   var zone hostedZone
//   var lb loadBalancer
//
//
//   route53.NewARecord(this, jsii.String("AliasRecord"), &aRecordProps{
//   	zone: zone,
//   	target: route53.recordTarget.fromAlias(targets.NewClassicLoadBalancerTarget(lb)),
//   })
//
type ClassicLoadBalancerTarget interface {
	awsroute53.IAliasRecordTarget
	// Return hosted zone ID and DNS name, usable for Route53 alias targets.
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
// Example:
//   import cloudfront "github.com/aws/aws-cdk-go/awscdk"
//
//   var myZone hostedZone
//   var distribution cloudFrontWebDistribution
//
//   route53.NewAaaaRecord(this, jsii.String("Alias"), &aaaaRecordProps{
//   	zone: myZone,
//   	target: route53.recordTarget.fromAlias(targets.NewCloudFrontTarget(distribution)),
//   })
//
type CloudFrontTarget interface {
	awsroute53.IAliasRecordTarget
	// Return hosted zone ID and DNS name, usable for Route53 alias targets.
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
// Example:
//   var zone hostedZone
//   var ebsEnvironmentUrl string
//
//
//   route53.NewARecord(this, jsii.String("AliasRecord"), &aRecordProps{
//   	zone: zone,
//   	target: route53.recordTarget.fromAlias(targets.NewElasticBeanstalkEnvironmentEndpointTarget(ebsEnvironmentUrl)),
//   })
//
type ElasticBeanstalkEnvironmentEndpointTarget interface {
	awsroute53.IAliasRecordTarget
	// Return hosted zone ID and DNS name, usable for Route53 alias targets.
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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   globalAcceleratorDomainTarget := awscdk.Aws_route53_targets.NewGlobalAcceleratorDomainTarget(jsii.String("acceleratorDomainName"))
//
type GlobalAcceleratorDomainTarget interface {
	awsroute53.IAliasRecordTarget
	// Return hosted zone ID and DNS name, usable for Route53 alias targets.
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
// Example:
//   import globalaccelerator "github.com/aws/aws-cdk-go/awscdk"
//
//   var zone hostedZone
//   var accelerator accelerator
//
//
//   route53.NewARecord(this, jsii.String("AliasRecord"), &aRecordProps{
//   	zone: zone,
//   	target: route53.recordTarget.fromAlias(targets.NewGlobalAcceleratorTarget(accelerator)),
//   })
//
type GlobalAcceleratorTarget interface {
	GlobalAcceleratorDomainTarget
	// Return hosted zone ID and DNS name, usable for Route53 alias targets.
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
// Example:
//   import ec2 "github.com/aws/aws-cdk-go/awscdk"
//
//   var zone hostedZone
//   var interfaceVpcEndpoint interfaceVpcEndpoint
//
//
//   route53.NewARecord(this, jsii.String("AliasRecord"), &aRecordProps{
//   	zone: zone,
//   	target: route53.recordTarget.fromAlias(targets.NewInterfaceVpcEndpointTarget(interfaceVpcEndpoint)),
//   })
//
type InterfaceVpcEndpointTarget interface {
	awsroute53.IAliasRecordTarget
	// Return hosted zone ID and DNS name, usable for Route53 alias targets.
	Bind(_record awsroute53.IRecordSet, _zone awsroute53.IHostedZone) *awsroute53.AliasRecordTargetConfig
}

// The jsii proxy struct for InterfaceVpcEndpointTarget
type jsiiProxy_InterfaceVpcEndpointTarget struct {
	internal.Type__awsroute53IAliasRecordTarget
}

func NewInterfaceVpcEndpointTarget(vpcEndpoint awsec2.InterfaceVpcEndpoint) InterfaceVpcEndpointTarget {
	_init_.Initialize()

	j := jsiiProxy_InterfaceVpcEndpointTarget{}

	_jsii_.Create(
		"aws-cdk-lib.aws_route53_targets.InterfaceVpcEndpointTarget",
		[]interface{}{vpcEndpoint},
		&j,
	)

	return &j
}

func NewInterfaceVpcEndpointTarget_Override(i InterfaceVpcEndpointTarget, vpcEndpoint awsec2.InterfaceVpcEndpoint) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_route53_targets.InterfaceVpcEndpointTarget",
		[]interface{}{vpcEndpoint},
		i,
	)
}

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
// Example:
//   import elbv2 "github.com/aws/aws-cdk-go/awscdk"
//
//   var zone hostedZone
//   var lb applicationLoadBalancer
//
//
//   route53.NewARecord(this, jsii.String("AliasRecord"), &aRecordProps{
//   	zone: zone,
//   	target: route53.recordTarget.fromAlias(targets.NewLoadBalancerTarget(lb)),
//   })
//
type LoadBalancerTarget interface {
	awsroute53.IAliasRecordTarget
	// Return hosted zone ID and DNS name, usable for Route53 alias targets.
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
// Example:
//   var zone hostedZone
//   var record aRecord
//
//   route53.NewARecord(this, jsii.String("AliasRecord"), &aRecordProps{
//   	zone: zone,
//   	target: route53.recordTarget.fromAlias(targets.NewRoute53RecordTarget(record)),
//   })
//
type Route53RecordTarget interface {
	awsroute53.IAliasRecordTarget
	// Return hosted zone ID and DNS name, usable for Route53 alias targets.
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
// Example:
//   import cognito "github.com/aws/aws-cdk-go/awscdk"
//
//   var zone hostedZone
//   var domain userPoolDomain
//
//   route53.NewARecord(this, jsii.String("AliasRecord"), &aRecordProps{
//   	zone: zone,
//   	target: route53.recordTarget.fromAlias(targets.NewUserPoolDomainTarget(domain)),
//   })
//
type UserPoolDomainTarget interface {
	awsroute53.IAliasRecordTarget
	// Return hosted zone ID and DNS name, usable for Route53 alias targets.
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

