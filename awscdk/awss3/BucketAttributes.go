package awss3

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// A reference to a bucket outside this stack.
//
// Example:
//   var myLambda function
//
//   bucket := s3.Bucket_FromBucketAttributes(this, jsii.String("ImportedBucket"), &BucketAttributes{
//   	BucketArn: jsii.String("arn:aws:s3:::my-bucket"),
//   })
//
//   // now you can just call methods on the bucket
//   bucket.AddEventNotification(s3.EventType_OBJECT_CREATED, s3n.NewLambdaDestination(myLambda), &NotificationKeyFilter{
//   	Prefix: jsii.String("home/myusername/*"),
//   })
//
type BucketAttributes struct {
	// The account this existing bucket belongs to.
	// Default: - it's assumed the bucket belongs to the same account as the scope it's being imported into.
	//
	Account *string `field:"optional" json:"account" yaml:"account"`
	// The ARN of the bucket.
	//
	// At least one of bucketArn or bucketName must be
	// defined in order to initialize a bucket ref.
	BucketArn *string `field:"optional" json:"bucketArn" yaml:"bucketArn"`
	// The domain name of the bucket.
	// Default: - Inferred from bucket name.
	//
	BucketDomainName *string `field:"optional" json:"bucketDomainName" yaml:"bucketDomainName"`
	// The IPv6 DNS name of the specified bucket.
	BucketDualStackDomainName *string `field:"optional" json:"bucketDualStackDomainName" yaml:"bucketDualStackDomainName"`
	// The name of the bucket.
	//
	// If the underlying value of ARN is a string, the
	// name will be parsed from the ARN. Otherwise, the name is optional, but
	// some features that require the bucket name such as auto-creating a bucket
	// policy, won't work.
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// The regional domain name of the specified bucket.
	BucketRegionalDomainName *string `field:"optional" json:"bucketRegionalDomainName" yaml:"bucketRegionalDomainName"`
	// Force the format of the website URL of the bucket.
	//
	// This should be true for
	// regions launched since 2014.
	// Default: - inferred from available region information, `false` otherwise.
	//
	// Deprecated: The correct website url format can be inferred automatically from the bucket `region`.
	// Always provide the bucket region if the `bucketWebsiteUrl` will be used.
	// Alternatively provide the full `bucketWebsiteUrl` manually.
	BucketWebsiteNewUrlFormat *bool `field:"optional" json:"bucketWebsiteNewUrlFormat" yaml:"bucketWebsiteNewUrlFormat"`
	// The website URL of the bucket (if static web hosting is enabled).
	// Default: - Inferred from bucket name and region.
	//
	BucketWebsiteUrl *string `field:"optional" json:"bucketWebsiteUrl" yaml:"bucketWebsiteUrl"`
	// KMS encryption key associated with this bucket.
	// Default: - no encryption key.
	//
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// If this bucket has been configured for static website hosting.
	// Default: false.
	//
	IsWebsite *bool `field:"optional" json:"isWebsite" yaml:"isWebsite"`
	// The role to be used by the notifications handler.
	// Default: - a new role will be created.
	//
	NotificationsHandlerRole awsiam.IRole `field:"optional" json:"notificationsHandlerRole" yaml:"notificationsHandlerRole"`
	// The region this existing bucket is in.
	//
	// Features that require the region (e.g. `bucketWebsiteUrl`) won't fully work
	// if the region cannot be correctly inferred.
	// Default: - it's assumed the bucket is in the same region as the scope it's being imported into.
	//
	Region *string `field:"optional" json:"region" yaml:"region"`
}

