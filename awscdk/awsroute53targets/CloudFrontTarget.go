package awsroute53targets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscloudfront"
	"github.com/aws/aws-cdk-go/awscdk/awsroute53"
	"github.com/aws/aws-cdk-go/awscdk/awsroute53targets/internal"
)

// Use a CloudFront Distribution as an alias record target.
//
// Example:
//   import cloudfront "github.com/aws/aws-cdk-go/awscdk"
//
//   var myZone hostedZone
//   var distribution cloudFrontWebDistribution
//
//   route53.NewAaaaRecord(this, jsii.String("Alias"), &AaaaRecordProps{
//   	Zone: myZone,
//   	Target: route53.RecordTarget_FromAlias(targets.NewCloudFrontTarget(distribution)),
//   })
//
// Experimental.
type CloudFrontTarget interface {
	awsroute53.IAliasRecordTarget
	// Return hosted zone ID and DNS name, usable for Route53 alias targets.
	// Experimental.
	Bind(_record awsroute53.IRecordSet, _zone awsroute53.IHostedZone) *awsroute53.AliasRecordTargetConfig
}

// The jsii proxy struct for CloudFrontTarget
type jsiiProxy_CloudFrontTarget struct {
	internal.Type__awsroute53IAliasRecordTarget
}

// Experimental.
func NewCloudFrontTarget(distribution awscloudfront.IDistribution) CloudFrontTarget {
	_init_.Initialize()

	if err := validateNewCloudFrontTargetParameters(distribution); err != nil {
		panic(err)
	}
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

	if err := validateCloudFrontTarget_GetHostedZoneIdParameters(scope); err != nil {
		panic(err)
	}
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

func (c *jsiiProxy_CloudFrontTarget) Bind(_record awsroute53.IRecordSet, _zone awsroute53.IHostedZone) *awsroute53.AliasRecordTargetConfig {
	if err := c.validateBindParameters(_record); err != nil {
		panic(err)
	}
	var returns *awsroute53.AliasRecordTargetConfig

	_jsii_.Invoke(
		c,
		"bind",
		[]interface{}{_record, _zone},
		&returns,
	)

	return returns
}

