package awsroute53targets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53targets/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Use a S3 as an alias record target.
//
// Example:
//   import s3 "github.com/aws/aws-cdk-go/awscdk"
//
//
//   recordName := "www"
//   domainName := "example.com"
//
//   bucketWebsite := s3.NewBucket(this, jsii.String("BucketWebsite"), &BucketProps{
//   	BucketName: []*string{
//   		recordName,
//   		domainName,
//   	}.join(jsii.String(".")),
//   	 // www.example.com
//   	PublicReadAccess: jsii.Boolean(true),
//   	WebsiteIndexDocument: jsii.String("index.html"),
//   })
//
//   zone := route53.HostedZone_FromLookup(this, jsii.String("Zone"), &HostedZoneProviderProps{
//   	DomainName: jsii.String(DomainName),
//   }) // example.com
//
//    // example.com
//   route53.NewARecord(this, jsii.String("AliasRecord"), &ARecordProps{
//   	Zone: Zone,
//   	RecordName: jsii.String(RecordName),
//   	 // www
//   	Target: route53.RecordTarget_FromAlias(
//   	targets.NewBucketWebsiteTarget(bucketWebsite, map[string]*bool{
//   		"evaluateTargetHealth": jsii.Boolean(true),
//   	})),
//   })
//
type BucketWebsiteTarget interface {
	awsroute53.IAliasRecordTarget
	// Return hosted zone ID and DNS name, usable for Route53 alias targets.
	Bind(record awsroute53.IRecordSet, _zone awsroute53.IHostedZone) *awsroute53.AliasRecordTargetConfig
}

// The jsii proxy struct for BucketWebsiteTarget
type jsiiProxy_BucketWebsiteTarget struct {
	internal.Type__awsroute53IAliasRecordTarget
}

func NewBucketWebsiteTarget(bucket awss3.IBucket, props IAliasRecordTargetProps) BucketWebsiteTarget {
	_init_.Initialize()

	if err := validateNewBucketWebsiteTargetParameters(bucket); err != nil {
		panic(err)
	}
	j := jsiiProxy_BucketWebsiteTarget{}

	_jsii_.Create(
		"aws-cdk-lib.aws_route53_targets.BucketWebsiteTarget",
		[]interface{}{bucket, props},
		&j,
	)

	return &j
}

func NewBucketWebsiteTarget_Override(b BucketWebsiteTarget, bucket awss3.IBucket, props IAliasRecordTargetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_route53_targets.BucketWebsiteTarget",
		[]interface{}{bucket, props},
		b,
	)
}

func (b *jsiiProxy_BucketWebsiteTarget) Bind(record awsroute53.IRecordSet, _zone awsroute53.IHostedZone) *awsroute53.AliasRecordTargetConfig {
	if err := b.validateBindParameters(record); err != nil {
		panic(err)
	}
	var returns *awsroute53.AliasRecordTargetConfig

	_jsii_.Invoke(
		b,
		"bind",
		[]interface{}{record, _zone},
		&returns,
	)

	return returns
}

