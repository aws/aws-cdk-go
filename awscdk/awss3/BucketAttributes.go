package awss3

import (
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awskms"
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
// Experimental.
type BucketAttributes struct {
	// The account this existing bucket belongs to.
	// Experimental.
	Account *string `field:"optional" json:"account" yaml:"account"`
	// The ARN of the bucket.
	//
	// At least one of bucketArn or bucketName must be
	// defined in order to initialize a bucket ref.
	// Experimental.
	BucketArn *string `field:"optional" json:"bucketArn" yaml:"bucketArn"`
	// The domain name of the bucket.
	// Experimental.
	BucketDomainName *string `field:"optional" json:"bucketDomainName" yaml:"bucketDomainName"`
	// The IPv6 DNS name of the specified bucket.
	// Experimental.
	BucketDualStackDomainName *string `field:"optional" json:"bucketDualStackDomainName" yaml:"bucketDualStackDomainName"`
	// The name of the bucket.
	//
	// If the underlying value of ARN is a string, the
	// name will be parsed from the ARN. Otherwise, the name is optional, but
	// some features that require the bucket name such as auto-creating a bucket
	// policy, won't work.
	// Experimental.
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// The regional domain name of the specified bucket.
	// Experimental.
	BucketRegionalDomainName *string `field:"optional" json:"bucketRegionalDomainName" yaml:"bucketRegionalDomainName"`
	// The format of the website URL of the bucket.
	//
	// This should be true for
	// regions launched since 2014.
	// Experimental.
	BucketWebsiteNewUrlFormat *bool `field:"optional" json:"bucketWebsiteNewUrlFormat" yaml:"bucketWebsiteNewUrlFormat"`
	// The website URL of the bucket (if static web hosting is enabled).
	// Experimental.
	BucketWebsiteUrl *string `field:"optional" json:"bucketWebsiteUrl" yaml:"bucketWebsiteUrl"`
	// Experimental.
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// If this bucket has been configured for static website hosting.
	// Experimental.
	IsWebsite *bool `field:"optional" json:"isWebsite" yaml:"isWebsite"`
	// The role to be used by the notifications handler.
	// Experimental.
	NotificationsHandlerRole awsiam.IRole `field:"optional" json:"notificationsHandlerRole" yaml:"notificationsHandlerRole"`
	// The region this existing bucket is in.
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
}

