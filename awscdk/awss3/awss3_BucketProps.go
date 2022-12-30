package awss3

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awskms"
)

// Example:
//   sourceBucket := s3.NewBucket(this, jsii.String("MyBucket"), &bucketProps{
//   	versioned: jsii.Boolean(true),
//   })
//
//   pipeline := codepipeline.NewPipeline(this, jsii.String("MyPipeline"))
//   sourceOutput := codepipeline.NewArtifact()
//   sourceAction := codepipeline_actions.NewS3SourceAction(&s3SourceActionProps{
//   	actionName: jsii.String("S3Source"),
//   	bucket: sourceBucket,
//   	bucketKey: jsii.String("path/to/file.zip"),
//   	output: sourceOutput,
//   })
//   pipeline.addStage(&stageOptions{
//   	stageName: jsii.String("Source"),
//   	actions: []iAction{
//   		sourceAction,
//   	},
//   })
//
// Experimental.
type BucketProps struct {
	// Specifies a canned ACL that grants predefined permissions to the bucket.
	// Experimental.
	AccessControl BucketAccessControl `field:"optional" json:"accessControl" yaml:"accessControl"`
	// Whether all objects should be automatically deleted when the bucket is removed from the stack or when the stack is deleted.
	//
	// Requires the `removalPolicy` to be set to `RemovalPolicy.DESTROY`.
	//
	// **Warning** if you have deployed a bucket with `autoDeleteObjects: true`,
	// switching this to `false` in a CDK version *before* `1.126.0` will lead to
	// all objects in the bucket being deleted. Be sure to update your bucket resources
	// by deploying with CDK version `1.126.0` or later **before** switching this value to `false`.
	// Experimental.
	AutoDeleteObjects *bool `field:"optional" json:"autoDeleteObjects" yaml:"autoDeleteObjects"`
	// The block public access configuration of this bucket.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html
	//
	// Experimental.
	BlockPublicAccess BlockPublicAccess `field:"optional" json:"blockPublicAccess" yaml:"blockPublicAccess"`
	// Specifies whether Amazon S3 should use an S3 Bucket Key with server-side encryption using KMS (SSE-KMS) for new objects in the bucket.
	//
	// Only relevant, when Encryption is set to {@link BucketEncryption.KMS}
	// Experimental.
	BucketKeyEnabled *bool `field:"optional" json:"bucketKeyEnabled" yaml:"bucketKeyEnabled"`
	// Physical name of this bucket.
	// Experimental.
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// The CORS configuration of this bucket.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-cors.html
	//
	// Experimental.
	Cors *[]*CorsRule `field:"optional" json:"cors" yaml:"cors"`
	// The kind of server-side encryption to apply to this bucket.
	//
	// If you choose KMS, you can specify a KMS key via `encryptionKey`. If
	// encryption key is not specified, a key will automatically be created.
	// Experimental.
	Encryption BucketEncryption `field:"optional" json:"encryption" yaml:"encryption"`
	// External KMS key to use for bucket encryption.
	//
	// The 'encryption' property must be either not specified or set to "Kms".
	// An error will be emitted if encryption is set to "Unencrypted" or
	// "Managed".
	// Experimental.
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// Enforces SSL for requests.
	//
	// S3.5 of the AWS Foundational Security Best Practices Regarding S3.
	// See: https://docs.aws.amazon.com/config/latest/developerguide/s3-bucket-ssl-requests-only.html
	//
	// Experimental.
	EnforceSSL *bool `field:"optional" json:"enforceSSL" yaml:"enforceSSL"`
	// Whether this bucket should send notifications to Amazon EventBridge or not.
	// Experimental.
	EventBridgeEnabled *bool `field:"optional" json:"eventBridgeEnabled" yaml:"eventBridgeEnabled"`
	// Inteligent Tiering Configurations.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/userguide/intelligent-tiering.html
	//
	// Experimental.
	IntelligentTieringConfigurations *[]*IntelligentTieringConfiguration `field:"optional" json:"intelligentTieringConfigurations" yaml:"intelligentTieringConfigurations"`
	// The inventory configuration of the bucket.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/storage-inventory.html
	//
	// Experimental.
	Inventories *[]*Inventory `field:"optional" json:"inventories" yaml:"inventories"`
	// Rules that define how Amazon S3 manages objects during their lifetime.
	// Experimental.
	LifecycleRules *[]*LifecycleRule `field:"optional" json:"lifecycleRules" yaml:"lifecycleRules"`
	// The metrics configuration of this bucket.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-metricsconfiguration.html
	//
	// Experimental.
	Metrics *[]*BucketMetrics `field:"optional" json:"metrics" yaml:"metrics"`
	// The role to be used by the notifications handler.
	// Experimental.
	NotificationsHandlerRole awsiam.IRole `field:"optional" json:"notificationsHandlerRole" yaml:"notificationsHandlerRole"`
	// The objectOwnership of the bucket.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/about-object-ownership.html
	//
	// Experimental.
	ObjectOwnership ObjectOwnership `field:"optional" json:"objectOwnership" yaml:"objectOwnership"`
	// Grants public read access to all objects in the bucket.
	//
	// Similar to calling `bucket.grantPublicAccess()`
	// Experimental.
	PublicReadAccess *bool `field:"optional" json:"publicReadAccess" yaml:"publicReadAccess"`
	// Policy to apply when the bucket is removed from this stack.
	// Experimental.
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// Destination bucket for the server access logs.
	// Experimental.
	ServerAccessLogsBucket IBucket `field:"optional" json:"serverAccessLogsBucket" yaml:"serverAccessLogsBucket"`
	// Optional log file prefix to use for the bucket's access logs.
	//
	// If defined without "serverAccessLogsBucket", enables access logs to current bucket with this prefix.
	// Experimental.
	ServerAccessLogsPrefix *string `field:"optional" json:"serverAccessLogsPrefix" yaml:"serverAccessLogsPrefix"`
	// Whether this bucket should have transfer acceleration turned on or not.
	// Experimental.
	TransferAcceleration *bool `field:"optional" json:"transferAcceleration" yaml:"transferAcceleration"`
	// Whether this bucket should have versioning turned on or not.
	// Experimental.
	Versioned *bool `field:"optional" json:"versioned" yaml:"versioned"`
	// The name of the error document (e.g. "404.html") for the website. `websiteIndexDocument` must also be set if this is set.
	// Experimental.
	WebsiteErrorDocument *string `field:"optional" json:"websiteErrorDocument" yaml:"websiteErrorDocument"`
	// The name of the index document (e.g. "index.html") for the website. Enables static website hosting for this bucket.
	// Experimental.
	WebsiteIndexDocument *string `field:"optional" json:"websiteIndexDocument" yaml:"websiteIndexDocument"`
	// Specifies the redirect behavior of all requests to a website endpoint of a bucket.
	//
	// If you specify this property, you can't specify "websiteIndexDocument", "websiteErrorDocument" nor , "websiteRoutingRules".
	// Experimental.
	WebsiteRedirect *RedirectTarget `field:"optional" json:"websiteRedirect" yaml:"websiteRedirect"`
	// Rules that define when a redirect is applied and the redirect behavior.
	// Experimental.
	WebsiteRoutingRules *[]*RoutingRule `field:"optional" json:"websiteRoutingRules" yaml:"websiteRoutingRules"`
}

