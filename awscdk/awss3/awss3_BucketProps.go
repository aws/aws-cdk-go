package awss3

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
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
type BucketProps struct {
	// Specifies a canned ACL that grants predefined permissions to the bucket.
	AccessControl BucketAccessControl `field:"optional" json:"accessControl" yaml:"accessControl"`
	// Whether all objects should be automatically deleted when the bucket is removed from the stack or when the stack is deleted.
	//
	// Requires the `removalPolicy` to be set to `RemovalPolicy.DESTROY`.
	//
	// **Warning** if you have deployed a bucket with `autoDeleteObjects: true`,
	// switching this to `false` in a CDK version *before* `1.126.0` will lead to
	// all objects in the bucket being deleted. Be sure to update your bucket resources
	// by deploying with CDK version `1.126.0` or later **before** switching this value to `false`.
	AutoDeleteObjects *bool `field:"optional" json:"autoDeleteObjects" yaml:"autoDeleteObjects"`
	// The block public access configuration of this bucket.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html
	//
	BlockPublicAccess BlockPublicAccess `field:"optional" json:"blockPublicAccess" yaml:"blockPublicAccess"`
	// Whether Amazon S3 should use its own intermediary key to generate data keys.
	//
	// Only relevant when using KMS for encryption.
	//
	// - If not enabled, every object GET and PUT will cause an API call to KMS (with the
	//    attendant cost implications of that).
	// - If enabled, S3 will use its own time-limited key instead.
	//
	// Only relevant, when Encryption is set to `BucketEncryption.KMS` or `BucketEncryption.KMS_MANAGED`.
	BucketKeyEnabled *bool `field:"optional" json:"bucketKeyEnabled" yaml:"bucketKeyEnabled"`
	// Physical name of this bucket.
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// The CORS configuration of this bucket.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-cors.html
	//
	Cors *[]*CorsRule `field:"optional" json:"cors" yaml:"cors"`
	// The kind of server-side encryption to apply to this bucket.
	//
	// If you choose KMS, you can specify a KMS key via `encryptionKey`. If
	// encryption key is not specified, a key will automatically be created.
	Encryption BucketEncryption `field:"optional" json:"encryption" yaml:"encryption"`
	// External KMS key to use for bucket encryption.
	//
	// The 'encryption' property must be either not specified or set to "Kms".
	// An error will be emitted if encryption is set to "Unencrypted" or
	// "Managed".
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// Enforces SSL for requests.
	//
	// S3.5 of the AWS Foundational Security Best Practices Regarding S3.
	// See: https://docs.aws.amazon.com/config/latest/developerguide/s3-bucket-ssl-requests-only.html
	//
	EnforceSSL *bool `field:"optional" json:"enforceSSL" yaml:"enforceSSL"`
	// Whether this bucket should send notifications to Amazon EventBridge or not.
	EventBridgeEnabled *bool `field:"optional" json:"eventBridgeEnabled" yaml:"eventBridgeEnabled"`
	// Inteligent Tiering Configurations.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/userguide/intelligent-tiering.html
	//
	IntelligentTieringConfigurations *[]*IntelligentTieringConfiguration `field:"optional" json:"intelligentTieringConfigurations" yaml:"intelligentTieringConfigurations"`
	// The inventory configuration of the bucket.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/storage-inventory.html
	//
	Inventories *[]*Inventory `field:"optional" json:"inventories" yaml:"inventories"`
	// Rules that define how Amazon S3 manages objects during their lifetime.
	LifecycleRules *[]*LifecycleRule `field:"optional" json:"lifecycleRules" yaml:"lifecycleRules"`
	// The metrics configuration of this bucket.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-metricsconfiguration.html
	//
	Metrics *[]*BucketMetrics `field:"optional" json:"metrics" yaml:"metrics"`
	// The role to be used by the notifications handler.
	NotificationsHandlerRole awsiam.IRole `field:"optional" json:"notificationsHandlerRole" yaml:"notificationsHandlerRole"`
	// The objectOwnership of the bucket.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/about-object-ownership.html
	//
	ObjectOwnership ObjectOwnership `field:"optional" json:"objectOwnership" yaml:"objectOwnership"`
	// Grants public read access to all objects in the bucket.
	//
	// Similar to calling `bucket.grantPublicAccess()`
	PublicReadAccess *bool `field:"optional" json:"publicReadAccess" yaml:"publicReadAccess"`
	// Policy to apply when the bucket is removed from this stack.
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// Destination bucket for the server access logs.
	ServerAccessLogsBucket IBucket `field:"optional" json:"serverAccessLogsBucket" yaml:"serverAccessLogsBucket"`
	// Optional log file prefix to use for the bucket's access logs.
	//
	// If defined without "serverAccessLogsBucket", enables access logs to current bucket with this prefix.
	ServerAccessLogsPrefix *string `field:"optional" json:"serverAccessLogsPrefix" yaml:"serverAccessLogsPrefix"`
	// Whether this bucket should have transfer acceleration turned on or not.
	TransferAcceleration *bool `field:"optional" json:"transferAcceleration" yaml:"transferAcceleration"`
	// Whether this bucket should have versioning turned on or not.
	Versioned *bool `field:"optional" json:"versioned" yaml:"versioned"`
	// The name of the error document (e.g. "404.html") for the website. `websiteIndexDocument` must also be set if this is set.
	WebsiteErrorDocument *string `field:"optional" json:"websiteErrorDocument" yaml:"websiteErrorDocument"`
	// The name of the index document (e.g. "index.html") for the website. Enables static website hosting for this bucket.
	WebsiteIndexDocument *string `field:"optional" json:"websiteIndexDocument" yaml:"websiteIndexDocument"`
	// Specifies the redirect behavior of all requests to a website endpoint of a bucket.
	//
	// If you specify this property, you can't specify "websiteIndexDocument", "websiteErrorDocument" nor , "websiteRoutingRules".
	WebsiteRedirect *RedirectTarget `field:"optional" json:"websiteRedirect" yaml:"websiteRedirect"`
	// Rules that define when a redirect is applied and the redirect behavior.
	WebsiteRoutingRules *[]*RoutingRule `field:"optional" json:"websiteRoutingRules" yaml:"websiteRoutingRules"`
}

