package awss3

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Example:
//   sourceBucket := s3.NewBucket(this, jsii.String("MyBucket"), &BucketProps{
//   	Versioned: jsii.Boolean(true),
//   })
//
//   pipeline := codepipeline.NewPipeline(this, jsii.String("MyPipeline"))
//   sourceOutput := codepipeline.NewArtifact()
//   sourceAction := codepipeline_actions.NewS3SourceAction(&S3SourceActionProps{
//   	ActionName: jsii.String("S3Source"),
//   	Bucket: sourceBucket,
//   	BucketKey: jsii.String("path/to/file.zip"),
//   	Output: sourceOutput,
//   })
//   pipeline.AddStage(&StageOptions{
//   	StageName: jsii.String("Source"),
//   	Actions: []iAction{
//   		sourceAction,
//   	},
//   })
//
type BucketProps struct {
	// Specifies a canned ACL that grants predefined permissions to the bucket.
	// Default: BucketAccessControl.PRIVATE
	//
	AccessControl BucketAccessControl `field:"optional" json:"accessControl" yaml:"accessControl"`
	// Whether all objects should be automatically deleted when the bucket is removed from the stack or when the stack is deleted.
	//
	// Requires the `removalPolicy` to be set to `RemovalPolicy.DESTROY`.
	//
	// **Warning** if you have deployed a bucket with `autoDeleteObjects: true`,
	// switching this to `false` in a CDK version *before* `1.126.0` will lead to
	// all objects in the bucket being deleted. Be sure to update your bucket resources
	// by deploying with CDK version `1.126.0` or later **before** switching this value to `false`.
	// Default: false.
	//
	AutoDeleteObjects *bool `field:"optional" json:"autoDeleteObjects" yaml:"autoDeleteObjects"`
	// The block public access configuration of this bucket.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html
	//
	// Default: - CloudFormation defaults will apply. New buckets and objects don't allow public access, but users can modify bucket policies or object permissions to allow public access
	//
	BlockPublicAccess BlockPublicAccess `field:"optional" json:"blockPublicAccess" yaml:"blockPublicAccess"`
	// Whether Amazon S3 should use its own intermediary key to generate data keys.
	//
	// Only relevant when using KMS for encryption.
	//
	// - If not enabled, every object GET and PUT will cause an API call to KMS (with the
	//   attendant cost implications of that).
	// - If enabled, S3 will use its own time-limited key instead.
	//
	// Only relevant, when Encryption is set to `BucketEncryption.KMS` or `BucketEncryption.KMS_MANAGED`.
	// Default: - false.
	//
	BucketKeyEnabled *bool `field:"optional" json:"bucketKeyEnabled" yaml:"bucketKeyEnabled"`
	// Physical name of this bucket.
	// Default: - Assigned by CloudFormation (recommended).
	//
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// The CORS configuration of this bucket.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-cors.html
	//
	// Default: - No CORS configuration.
	//
	Cors *[]*CorsRule `field:"optional" json:"cors" yaml:"cors"`
	// The kind of server-side encryption to apply to this bucket.
	//
	// If you choose KMS, you can specify a KMS key via `encryptionKey`. If
	// encryption key is not specified, a key will automatically be created.
	// Default: - `KMS` if `encryptionKey` is specified, or `UNENCRYPTED` otherwise.
	// But if `UNENCRYPTED` is specified, the bucket will be encrypted as `S3_MANAGED` automatically.
	//
	Encryption BucketEncryption `field:"optional" json:"encryption" yaml:"encryption"`
	// External KMS key to use for bucket encryption.
	//
	// The `encryption` property must be either not specified or set to `KMS` or `DSSE`.
	// An error will be emitted if `encryption` is set to `UNENCRYPTED` or `S3_MANAGED`.
	// Default: - If `encryption` is set to `KMS` and this property is undefined,
	// a new KMS key will be created and associated with this bucket.
	//
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// Enforces SSL for requests.
	//
	// S3.5 of the AWS Foundational Security Best Practices Regarding S3.
	// See: https://docs.aws.amazon.com/config/latest/developerguide/s3-bucket-ssl-requests-only.html
	//
	// Default: false.
	//
	EnforceSSL *bool `field:"optional" json:"enforceSSL" yaml:"enforceSSL"`
	// Whether this bucket should send notifications to Amazon EventBridge or not.
	// Default: false.
	//
	EventBridgeEnabled *bool `field:"optional" json:"eventBridgeEnabled" yaml:"eventBridgeEnabled"`
	// Inteligent Tiering Configurations.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/userguide/intelligent-tiering.html
	//
	// Default: No Intelligent Tiiering Configurations.
	//
	IntelligentTieringConfigurations *[]*IntelligentTieringConfiguration `field:"optional" json:"intelligentTieringConfigurations" yaml:"intelligentTieringConfigurations"`
	// The inventory configuration of the bucket.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/storage-inventory.html
	//
	// Default: - No inventory configuration.
	//
	Inventories *[]*Inventory `field:"optional" json:"inventories" yaml:"inventories"`
	// Rules that define how Amazon S3 manages objects during their lifetime.
	// Default: - No lifecycle rules.
	//
	LifecycleRules *[]*LifecycleRule `field:"optional" json:"lifecycleRules" yaml:"lifecycleRules"`
	// The metrics configuration of this bucket.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-metricsconfiguration.html
	//
	// Default: - No metrics configuration.
	//
	Metrics *[]*BucketMetrics `field:"optional" json:"metrics" yaml:"metrics"`
	// Enforces minimum TLS version for requests.
	//
	// Requires `enforceSSL` to be enabled.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/userguide/amazon-s3-policy-keys.html#example-object-tls-version
	//
	// Default: No minimum TLS version is enforced.
	//
	MinimumTLSVersion *float64 `field:"optional" json:"minimumTLSVersion" yaml:"minimumTLSVersion"`
	// The role to be used by the notifications handler.
	// Default: - a new role will be created.
	//
	NotificationsHandlerRole awsiam.IRole `field:"optional" json:"notificationsHandlerRole" yaml:"notificationsHandlerRole"`
	// The default retention mode and rules for S3 Object Lock.
	//
	// Default retention can be configured after a bucket is created if the bucket already
	// has object lock enabled. Enabling object lock for existing buckets is not supported.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-lock-overview.html#object-lock-bucket-config-enable
	//
	// Default: no default retention period.
	//
	ObjectLockDefaultRetention ObjectLockRetention `field:"optional" json:"objectLockDefaultRetention" yaml:"objectLockDefaultRetention"`
	// Enable object lock on the bucket.
	//
	// Enabling object lock for existing buckets is not supported. Object lock must be
	// enabled when the bucket is created.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-lock-overview.html#object-lock-bucket-config-enable
	//
	// Default: false, unless objectLockDefaultRetention is set (then, true).
	//
	ObjectLockEnabled *bool `field:"optional" json:"objectLockEnabled" yaml:"objectLockEnabled"`
	// The objectOwnership of the bucket.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/about-object-ownership.html
	//
	// Default: - No ObjectOwnership configuration, uploading account will own the object.
	//
	ObjectOwnership ObjectOwnership `field:"optional" json:"objectOwnership" yaml:"objectOwnership"`
	// Grants public read access to all objects in the bucket.
	//
	// Similar to calling `bucket.grantPublicAccess()`
	// Default: false.
	//
	PublicReadAccess *bool `field:"optional" json:"publicReadAccess" yaml:"publicReadAccess"`
	// Policy to apply when the bucket is removed from this stack.
	// Default: - The bucket will be orphaned.
	//
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// Destination bucket for the server access logs.
	// Default: - If "serverAccessLogsPrefix" undefined - access logs disabled, otherwise - log to current bucket.
	//
	ServerAccessLogsBucket IBucket `field:"optional" json:"serverAccessLogsBucket" yaml:"serverAccessLogsBucket"`
	// Optional log file prefix to use for the bucket's access logs.
	//
	// If defined without "serverAccessLogsBucket", enables access logs to current bucket with this prefix.
	// Default: - No log file prefix.
	//
	ServerAccessLogsPrefix *string `field:"optional" json:"serverAccessLogsPrefix" yaml:"serverAccessLogsPrefix"`
	// Optional key format for log objects.
	// Default: - the default key format is: [DestinationPrefix][YYYY]-[MM]-[DD]-[hh]-[mm]-[ss]-[UniqueString].
	//
	TargetObjectKeyFormat TargetObjectKeyFormat `field:"optional" json:"targetObjectKeyFormat" yaml:"targetObjectKeyFormat"`
	// Whether this bucket should have transfer acceleration turned on or not.
	// Default: false.
	//
	TransferAcceleration *bool `field:"optional" json:"transferAcceleration" yaml:"transferAcceleration"`
	// Whether this bucket should have versioning turned on or not.
	// Default: false (unless object lock is enabled, then true).
	//
	Versioned *bool `field:"optional" json:"versioned" yaml:"versioned"`
	// The name of the error document (e.g. "404.html") for the website. `websiteIndexDocument` must also be set if this is set.
	// Default: - No error document.
	//
	WebsiteErrorDocument *string `field:"optional" json:"websiteErrorDocument" yaml:"websiteErrorDocument"`
	// The name of the index document (e.g. "index.html") for the website. Enables static website hosting for this bucket.
	// Default: - No index document.
	//
	WebsiteIndexDocument *string `field:"optional" json:"websiteIndexDocument" yaml:"websiteIndexDocument"`
	// Specifies the redirect behavior of all requests to a website endpoint of a bucket.
	//
	// If you specify this property, you can't specify "websiteIndexDocument", "websiteErrorDocument" nor , "websiteRoutingRules".
	// Default: - No redirection.
	//
	WebsiteRedirect *RedirectTarget `field:"optional" json:"websiteRedirect" yaml:"websiteRedirect"`
	// Rules that define when a redirect is applied and the redirect behavior.
	// Default: - No redirection rules.
	//
	WebsiteRoutingRules *[]*RoutingRule `field:"optional" json:"websiteRoutingRules" yaml:"websiteRoutingRules"`
}

