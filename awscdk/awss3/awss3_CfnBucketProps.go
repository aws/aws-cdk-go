package awss3

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnBucket`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnBucketProps := &cfnBucketProps{
//   	accelerateConfiguration: &accelerateConfigurationProperty{
//   		accelerationStatus: jsii.String("accelerationStatus"),
//   	},
//   	accessControl: jsii.String("accessControl"),
//   	analyticsConfigurations: []interface{}{
//   		&analyticsConfigurationProperty{
//   			id: jsii.String("id"),
//   			storageClassAnalysis: &storageClassAnalysisProperty{
//   				dataExport: &dataExportProperty{
//   					destination: &destinationProperty{
//   						bucketArn: jsii.String("bucketArn"),
//   						format: jsii.String("format"),
//
//   						// the properties below are optional
//   						bucketAccountId: jsii.String("bucketAccountId"),
//   						prefix: jsii.String("prefix"),
//   					},
//   					outputSchemaVersion: jsii.String("outputSchemaVersion"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			prefix: jsii.String("prefix"),
//   			tagFilters: []interface{}{
//   				&tagFilterProperty{
//   					key: jsii.String("key"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   	bucketEncryption: &bucketEncryptionProperty{
//   		serverSideEncryptionConfiguration: []interface{}{
//   			&serverSideEncryptionRuleProperty{
//   				bucketKeyEnabled: jsii.Boolean(false),
//   				serverSideEncryptionByDefault: &serverSideEncryptionByDefaultProperty{
//   					sseAlgorithm: jsii.String("sseAlgorithm"),
//
//   					// the properties below are optional
//   					kmsMasterKeyId: jsii.String("kmsMasterKeyId"),
//   				},
//   			},
//   		},
//   	},
//   	bucketName: jsii.String("bucketName"),
//   	corsConfiguration: &corsConfigurationProperty{
//   		corsRules: []interface{}{
//   			&corsRuleProperty{
//   				allowedMethods: []*string{
//   					jsii.String("allowedMethods"),
//   				},
//   				allowedOrigins: []*string{
//   					jsii.String("allowedOrigins"),
//   				},
//
//   				// the properties below are optional
//   				allowedHeaders: []*string{
//   					jsii.String("allowedHeaders"),
//   				},
//   				exposedHeaders: []*string{
//   					jsii.String("exposedHeaders"),
//   				},
//   				id: jsii.String("id"),
//   				maxAge: jsii.Number(123),
//   			},
//   		},
//   	},
//   	intelligentTieringConfigurations: []interface{}{
//   		&intelligentTieringConfigurationProperty{
//   			id: jsii.String("id"),
//   			status: jsii.String("status"),
//   			tierings: []interface{}{
//   				&tieringProperty{
//   					accessTier: jsii.String("accessTier"),
//   					days: jsii.Number(123),
//   				},
//   			},
//
//   			// the properties below are optional
//   			prefix: jsii.String("prefix"),
//   			tagFilters: []interface{}{
//   				&tagFilterProperty{
//   					key: jsii.String("key"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   	inventoryConfigurations: []interface{}{
//   		&inventoryConfigurationProperty{
//   			destination: &destinationProperty{
//   				bucketArn: jsii.String("bucketArn"),
//   				format: jsii.String("format"),
//
//   				// the properties below are optional
//   				bucketAccountId: jsii.String("bucketAccountId"),
//   				prefix: jsii.String("prefix"),
//   			},
//   			enabled: jsii.Boolean(false),
//   			id: jsii.String("id"),
//   			includedObjectVersions: jsii.String("includedObjectVersions"),
//   			scheduleFrequency: jsii.String("scheduleFrequency"),
//
//   			// the properties below are optional
//   			optionalFields: []*string{
//   				jsii.String("optionalFields"),
//   			},
//   			prefix: jsii.String("prefix"),
//   		},
//   	},
//   	lifecycleConfiguration: &lifecycleConfigurationProperty{
//   		rules: []interface{}{
//   			&ruleProperty{
//   				status: jsii.String("status"),
//
//   				// the properties below are optional
//   				abortIncompleteMultipartUpload: &abortIncompleteMultipartUploadProperty{
//   					daysAfterInitiation: jsii.Number(123),
//   				},
//   				expirationDate: NewDate(),
//   				expirationInDays: jsii.Number(123),
//   				expiredObjectDeleteMarker: jsii.Boolean(false),
//   				id: jsii.String("id"),
//   				noncurrentVersionExpiration: &noncurrentVersionExpirationProperty{
//   					noncurrentDays: jsii.Number(123),
//
//   					// the properties below are optional
//   					newerNoncurrentVersions: jsii.Number(123),
//   				},
//   				noncurrentVersionExpirationInDays: jsii.Number(123),
//   				noncurrentVersionTransition: &noncurrentVersionTransitionProperty{
//   					storageClass: jsii.String("storageClass"),
//   					transitionInDays: jsii.Number(123),
//
//   					// the properties below are optional
//   					newerNoncurrentVersions: jsii.Number(123),
//   				},
//   				noncurrentVersionTransitions: []interface{}{
//   					&noncurrentVersionTransitionProperty{
//   						storageClass: jsii.String("storageClass"),
//   						transitionInDays: jsii.Number(123),
//
//   						// the properties below are optional
//   						newerNoncurrentVersions: jsii.Number(123),
//   					},
//   				},
//   				objectSizeGreaterThan: jsii.Number(123),
//   				objectSizeLessThan: jsii.Number(123),
//   				prefix: jsii.String("prefix"),
//   				tagFilters: []interface{}{
//   					&tagFilterProperty{
//   						key: jsii.String("key"),
//   						value: jsii.String("value"),
//   					},
//   				},
//   				transition: &transitionProperty{
//   					storageClass: jsii.String("storageClass"),
//
//   					// the properties below are optional
//   					transitionDate: NewDate(),
//   					transitionInDays: jsii.Number(123),
//   				},
//   				transitions: []interface{}{
//   					&transitionProperty{
//   						storageClass: jsii.String("storageClass"),
//
//   						// the properties below are optional
//   						transitionDate: NewDate(),
//   						transitionInDays: jsii.Number(123),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	loggingConfiguration: &loggingConfigurationProperty{
//   		destinationBucketName: jsii.String("destinationBucketName"),
//   		logFilePrefix: jsii.String("logFilePrefix"),
//   	},
//   	metricsConfigurations: []interface{}{
//   		&metricsConfigurationProperty{
//   			id: jsii.String("id"),
//
//   			// the properties below are optional
//   			accessPointArn: jsii.String("accessPointArn"),
//   			prefix: jsii.String("prefix"),
//   			tagFilters: []interface{}{
//   				&tagFilterProperty{
//   					key: jsii.String("key"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   	notificationConfiguration: &notificationConfigurationProperty{
//   		eventBridgeConfiguration: &eventBridgeConfigurationProperty{
//   			eventBridgeEnabled: jsii.Boolean(false),
//   		},
//   		lambdaConfigurations: []interface{}{
//   			&lambdaConfigurationProperty{
//   				event: jsii.String("event"),
//   				function: jsii.String("function"),
//
//   				// the properties below are optional
//   				filter: &notificationFilterProperty{
//   					s3Key: &s3KeyFilterProperty{
//   						rules: []interface{}{
//   							&filterRuleProperty{
//   								name: jsii.String("name"),
//   								value: jsii.String("value"),
//   							},
//   						},
//   					},
//   				},
//   			},
//   		},
//   		queueConfigurations: []interface{}{
//   			&queueConfigurationProperty{
//   				event: jsii.String("event"),
//   				queue: jsii.String("queue"),
//
//   				// the properties below are optional
//   				filter: &notificationFilterProperty{
//   					s3Key: &s3KeyFilterProperty{
//   						rules: []interface{}{
//   							&filterRuleProperty{
//   								name: jsii.String("name"),
//   								value: jsii.String("value"),
//   							},
//   						},
//   					},
//   				},
//   			},
//   		},
//   		topicConfigurations: []interface{}{
//   			&topicConfigurationProperty{
//   				event: jsii.String("event"),
//   				topic: jsii.String("topic"),
//
//   				// the properties below are optional
//   				filter: &notificationFilterProperty{
//   					s3Key: &s3KeyFilterProperty{
//   						rules: []interface{}{
//   							&filterRuleProperty{
//   								name: jsii.String("name"),
//   								value: jsii.String("value"),
//   							},
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	objectLockConfiguration: &objectLockConfigurationProperty{
//   		objectLockEnabled: jsii.String("objectLockEnabled"),
//   		rule: &objectLockRuleProperty{
//   			defaultRetention: &defaultRetentionProperty{
//   				days: jsii.Number(123),
//   				mode: jsii.String("mode"),
//   				years: jsii.Number(123),
//   			},
//   		},
//   	},
//   	objectLockEnabled: jsii.Boolean(false),
//   	ownershipControls: &ownershipControlsProperty{
//   		rules: []interface{}{
//   			&ownershipControlsRuleProperty{
//   				objectOwnership: jsii.String("objectOwnership"),
//   			},
//   		},
//   	},
//   	publicAccessBlockConfiguration: &publicAccessBlockConfigurationProperty{
//   		blockPublicAcls: jsii.Boolean(false),
//   		blockPublicPolicy: jsii.Boolean(false),
//   		ignorePublicAcls: jsii.Boolean(false),
//   		restrictPublicBuckets: jsii.Boolean(false),
//   	},
//   	replicationConfiguration: &replicationConfigurationProperty{
//   		role: jsii.String("role"),
//   		rules: []interface{}{
//   			&replicationRuleProperty{
//   				destination: &replicationDestinationProperty{
//   					bucket: jsii.String("bucket"),
//
//   					// the properties below are optional
//   					accessControlTranslation: &accessControlTranslationProperty{
//   						owner: jsii.String("owner"),
//   					},
//   					account: jsii.String("account"),
//   					encryptionConfiguration: &encryptionConfigurationProperty{
//   						replicaKmsKeyId: jsii.String("replicaKmsKeyId"),
//   					},
//   					metrics: &metricsProperty{
//   						status: jsii.String("status"),
//
//   						// the properties below are optional
//   						eventThreshold: &replicationTimeValueProperty{
//   							minutes: jsii.Number(123),
//   						},
//   					},
//   					replicationTime: &replicationTimeProperty{
//   						status: jsii.String("status"),
//   						time: &replicationTimeValueProperty{
//   							minutes: jsii.Number(123),
//   						},
//   					},
//   					storageClass: jsii.String("storageClass"),
//   				},
//   				status: jsii.String("status"),
//
//   				// the properties below are optional
//   				deleteMarkerReplication: &deleteMarkerReplicationProperty{
//   					status: jsii.String("status"),
//   				},
//   				filter: &replicationRuleFilterProperty{
//   					and: &replicationRuleAndOperatorProperty{
//   						prefix: jsii.String("prefix"),
//   						tagFilters: []interface{}{
//   							&tagFilterProperty{
//   								key: jsii.String("key"),
//   								value: jsii.String("value"),
//   							},
//   						},
//   					},
//   					prefix: jsii.String("prefix"),
//   					tagFilter: &tagFilterProperty{
//   						key: jsii.String("key"),
//   						value: jsii.String("value"),
//   					},
//   				},
//   				id: jsii.String("id"),
//   				prefix: jsii.String("prefix"),
//   				priority: jsii.Number(123),
//   				sourceSelectionCriteria: &sourceSelectionCriteriaProperty{
//   					replicaModifications: &replicaModificationsProperty{
//   						status: jsii.String("status"),
//   					},
//   					sseKmsEncryptedObjects: &sseKmsEncryptedObjectsProperty{
//   						status: jsii.String("status"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	versioningConfiguration: &versioningConfigurationProperty{
//   		status: jsii.String("status"),
//   	},
//   	websiteConfiguration: &websiteConfigurationProperty{
//   		errorDocument: jsii.String("errorDocument"),
//   		indexDocument: jsii.String("indexDocument"),
//   		redirectAllRequestsTo: &redirectAllRequestsToProperty{
//   			hostName: jsii.String("hostName"),
//
//   			// the properties below are optional
//   			protocol: jsii.String("protocol"),
//   		},
//   		routingRules: []interface{}{
//   			&routingRuleProperty{
//   				redirectRule: &redirectRuleProperty{
//   					hostName: jsii.String("hostName"),
//   					httpRedirectCode: jsii.String("httpRedirectCode"),
//   					protocol: jsii.String("protocol"),
//   					replaceKeyPrefixWith: jsii.String("replaceKeyPrefixWith"),
//   					replaceKeyWith: jsii.String("replaceKeyWith"),
//   				},
//
//   				// the properties below are optional
//   				routingRuleCondition: &routingRuleConditionProperty{
//   					httpErrorCodeReturnedEquals: jsii.String("httpErrorCodeReturnedEquals"),
//   					keyPrefixEquals: jsii.String("keyPrefixEquals"),
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnBucketProps struct {
	// Configures the transfer acceleration state for an Amazon S3 bucket.
	//
	// For more information, see [Amazon S3 Transfer Acceleration](https://docs.aws.amazon.com/AmazonS3/latest/dev/transfer-acceleration.html) in the *Amazon S3 User Guide* .
	AccelerateConfiguration interface{} `field:"optional" json:"accelerateConfiguration" yaml:"accelerateConfiguration"`
	// A canned access control list (ACL) that grants predefined permissions to the bucket.
	//
	// For more information about canned ACLs, see [Canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl) in the *Amazon S3 User Guide* .
	//
	// Be aware that the syntax for this property differs from the information provided in the *Amazon S3 User Guide* . The AccessControl property is case-sensitive and must be one of the following values: Private, PublicRead, PublicReadWrite, AuthenticatedRead, LogDeliveryWrite, BucketOwnerRead, BucketOwnerFullControl, or AwsExecRead.
	AccessControl *string `field:"optional" json:"accessControl" yaml:"accessControl"`
	// Specifies the configuration and any analyses for the analytics filter of an Amazon S3 bucket.
	AnalyticsConfigurations interface{} `field:"optional" json:"analyticsConfigurations" yaml:"analyticsConfigurations"`
	// Specifies default encryption for a bucket using server-side encryption with Amazon S3-managed keys (SSE-S3) or AWS KMS-managed keys (SSE-KMS) bucket.
	//
	// For information about the Amazon S3 default encryption feature, see [Amazon S3 Default Encryption for S3 Buckets](https://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-encryption.html) in the *Amazon S3 User Guide* .
	BucketEncryption interface{} `field:"optional" json:"bucketEncryption" yaml:"bucketEncryption"`
	// A name for the bucket.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique ID and uses that ID for the bucket name. The bucket name must contain only lowercase letters, numbers, periods (.), and dashes (-) and must follow [Amazon S3 bucket restrictions and limitations](https://docs.aws.amazon.com/AmazonS3/latest/dev/BucketRestrictions.html) . For more information, see [Rules for naming Amazon S3 buckets](https://docs.aws.amazon.com/AmazonS3/latest/dev/BucketRestrictions.html#bucketnamingrules) in the *Amazon S3 User Guide* .
	//
	// > If you specify a name, you can't perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you need to replace the resource, specify a new name.
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// Describes the cross-origin access configuration for objects in an Amazon S3 bucket.
	//
	// For more information, see [Enabling Cross-Origin Resource Sharing](https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html) in the *Amazon S3 User Guide* .
	CorsConfiguration interface{} `field:"optional" json:"corsConfiguration" yaml:"corsConfiguration"`
	// Defines how Amazon S3 handles Intelligent-Tiering storage.
	IntelligentTieringConfigurations interface{} `field:"optional" json:"intelligentTieringConfigurations" yaml:"intelligentTieringConfigurations"`
	// Specifies the inventory configuration for an Amazon S3 bucket.
	//
	// For more information, see [GET Bucket inventory](https://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketGETInventoryConfig.html) in the *Amazon S3 API Reference* .
	InventoryConfigurations interface{} `field:"optional" json:"inventoryConfigurations" yaml:"inventoryConfigurations"`
	// Specifies the lifecycle configuration for objects in an Amazon S3 bucket.
	//
	// For more information, see [Object Lifecycle Management](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html) in the *Amazon S3 User Guide* .
	LifecycleConfiguration interface{} `field:"optional" json:"lifecycleConfiguration" yaml:"lifecycleConfiguration"`
	// Settings that define where logs are stored.
	LoggingConfiguration interface{} `field:"optional" json:"loggingConfiguration" yaml:"loggingConfiguration"`
	// Specifies a metrics configuration for the CloudWatch request metrics (specified by the metrics configuration ID) from an Amazon S3 bucket.
	//
	// If you're updating an existing metrics configuration, note that this is a full replacement of the existing metrics configuration. If you don't include the elements you want to keep, they are erased. For more information, see [PutBucketMetricsConfiguration](https://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketPUTMetricConfiguration.html) .
	MetricsConfigurations interface{} `field:"optional" json:"metricsConfigurations" yaml:"metricsConfigurations"`
	// Configuration that defines how Amazon S3 handles bucket notifications.
	NotificationConfiguration interface{} `field:"optional" json:"notificationConfiguration" yaml:"notificationConfiguration"`
	// Places an Object Lock configuration on the specified bucket.
	//
	// The rule specified in the Object Lock configuration will be applied by default to every new object placed in the specified bucket. For more information, see [Locking Objects](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock.html) .
	//
	// > - The `DefaultRetention` settings require both a mode and a period.
	// > - The `DefaultRetention` period can be either `Days` or `Years` but you must select one. You cannot specify `Days` and `Years` at the same time.
	// > - You can only enable Object Lock for new buckets. If you want to turn on Object Lock for an existing bucket, contact AWS Support.
	ObjectLockConfiguration interface{} `field:"optional" json:"objectLockConfiguration" yaml:"objectLockConfiguration"`
	// Indicates whether this bucket has an Object Lock configuration enabled.
	//
	// Enable `ObjectLockEnabled` when you apply `ObjectLockConfiguration` to a bucket.
	ObjectLockEnabled interface{} `field:"optional" json:"objectLockEnabled" yaml:"objectLockEnabled"`
	// Configuration that defines how Amazon S3 handles Object Ownership rules.
	OwnershipControls interface{} `field:"optional" json:"ownershipControls" yaml:"ownershipControls"`
	// Configuration that defines how Amazon S3 handles public access.
	PublicAccessBlockConfiguration interface{} `field:"optional" json:"publicAccessBlockConfiguration" yaml:"publicAccessBlockConfiguration"`
	// Configuration for replicating objects in an S3 bucket.
	//
	// To enable replication, you must also enable versioning by using the `VersioningConfiguration` property.
	//
	// Amazon S3 can store replicated objects in a single destination bucket or multiple destination buckets. The destination bucket or buckets must already exist.
	ReplicationConfiguration interface{} `field:"optional" json:"replicationConfiguration" yaml:"replicationConfiguration"`
	// An arbitrary set of tags (key-value pairs) for this S3 bucket.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Enables multiple versions of all objects in this bucket.
	//
	// You might enable versioning to prevent objects from being deleted or overwritten by mistake or to archive objects so that you can retrieve previous versions of them.
	VersioningConfiguration interface{} `field:"optional" json:"versioningConfiguration" yaml:"versioningConfiguration"`
	// Information used to configure the bucket as a static website.
	//
	// For more information, see [Hosting Websites on Amazon S3](https://docs.aws.amazon.com/AmazonS3/latest/dev/WebsiteHosting.html) .
	WebsiteConfiguration interface{} `field:"optional" json:"websiteConfiguration" yaml:"websiteConfiguration"`
}

