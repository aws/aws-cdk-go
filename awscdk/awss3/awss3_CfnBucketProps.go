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
//   cfnBucketProps := &CfnBucketProps{
//   	AccelerateConfiguration: &AccelerateConfigurationProperty{
//   		AccelerationStatus: jsii.String("accelerationStatus"),
//   	},
//   	AccessControl: jsii.String("accessControl"),
//   	AnalyticsConfigurations: []interface{}{
//   		&AnalyticsConfigurationProperty{
//   			Id: jsii.String("id"),
//   			StorageClassAnalysis: &StorageClassAnalysisProperty{
//   				DataExport: &DataExportProperty{
//   					Destination: &DestinationProperty{
//   						BucketArn: jsii.String("bucketArn"),
//   						Format: jsii.String("format"),
//
//   						// the properties below are optional
//   						BucketAccountId: jsii.String("bucketAccountId"),
//   						Prefix: jsii.String("prefix"),
//   					},
//   					OutputSchemaVersion: jsii.String("outputSchemaVersion"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			Prefix: jsii.String("prefix"),
//   			TagFilters: []interface{}{
//   				&TagFilterProperty{
//   					Key: jsii.String("key"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   	BucketEncryption: &BucketEncryptionProperty{
//   		ServerSideEncryptionConfiguration: []interface{}{
//   			&ServerSideEncryptionRuleProperty{
//   				BucketKeyEnabled: jsii.Boolean(false),
//   				ServerSideEncryptionByDefault: &ServerSideEncryptionByDefaultProperty{
//   					SseAlgorithm: jsii.String("sseAlgorithm"),
//
//   					// the properties below are optional
//   					KmsMasterKeyId: jsii.String("kmsMasterKeyId"),
//   				},
//   			},
//   		},
//   	},
//   	BucketName: jsii.String("bucketName"),
//   	CorsConfiguration: &CorsConfigurationProperty{
//   		CorsRules: []interface{}{
//   			&CorsRuleProperty{
//   				AllowedMethods: []*string{
//   					jsii.String("allowedMethods"),
//   				},
//   				AllowedOrigins: []*string{
//   					jsii.String("allowedOrigins"),
//   				},
//
//   				// the properties below are optional
//   				AllowedHeaders: []*string{
//   					jsii.String("allowedHeaders"),
//   				},
//   				ExposedHeaders: []*string{
//   					jsii.String("exposedHeaders"),
//   				},
//   				Id: jsii.String("id"),
//   				MaxAge: jsii.Number(123),
//   			},
//   		},
//   	},
//   	IntelligentTieringConfigurations: []interface{}{
//   		&IntelligentTieringConfigurationProperty{
//   			Id: jsii.String("id"),
//   			Status: jsii.String("status"),
//   			Tierings: []interface{}{
//   				&TieringProperty{
//   					AccessTier: jsii.String("accessTier"),
//   					Days: jsii.Number(123),
//   				},
//   			},
//
//   			// the properties below are optional
//   			Prefix: jsii.String("prefix"),
//   			TagFilters: []interface{}{
//   				&TagFilterProperty{
//   					Key: jsii.String("key"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   	InventoryConfigurations: []interface{}{
//   		&InventoryConfigurationProperty{
//   			Destination: &DestinationProperty{
//   				BucketArn: jsii.String("bucketArn"),
//   				Format: jsii.String("format"),
//
//   				// the properties below are optional
//   				BucketAccountId: jsii.String("bucketAccountId"),
//   				Prefix: jsii.String("prefix"),
//   			},
//   			Enabled: jsii.Boolean(false),
//   			Id: jsii.String("id"),
//   			IncludedObjectVersions: jsii.String("includedObjectVersions"),
//   			ScheduleFrequency: jsii.String("scheduleFrequency"),
//
//   			// the properties below are optional
//   			OptionalFields: []*string{
//   				jsii.String("optionalFields"),
//   			},
//   			Prefix: jsii.String("prefix"),
//   		},
//   	},
//   	LifecycleConfiguration: &LifecycleConfigurationProperty{
//   		Rules: []interface{}{
//   			&RuleProperty{
//   				Status: jsii.String("status"),
//
//   				// the properties below are optional
//   				AbortIncompleteMultipartUpload: &AbortIncompleteMultipartUploadProperty{
//   					DaysAfterInitiation: jsii.Number(123),
//   				},
//   				ExpirationDate: NewDate(),
//   				ExpirationInDays: jsii.Number(123),
//   				ExpiredObjectDeleteMarker: jsii.Boolean(false),
//   				Id: jsii.String("id"),
//   				NoncurrentVersionExpiration: &NoncurrentVersionExpirationProperty{
//   					NoncurrentDays: jsii.Number(123),
//
//   					// the properties below are optional
//   					NewerNoncurrentVersions: jsii.Number(123),
//   				},
//   				NoncurrentVersionExpirationInDays: jsii.Number(123),
//   				NoncurrentVersionTransition: &NoncurrentVersionTransitionProperty{
//   					StorageClass: jsii.String("storageClass"),
//   					TransitionInDays: jsii.Number(123),
//
//   					// the properties below are optional
//   					NewerNoncurrentVersions: jsii.Number(123),
//   				},
//   				NoncurrentVersionTransitions: []interface{}{
//   					&NoncurrentVersionTransitionProperty{
//   						StorageClass: jsii.String("storageClass"),
//   						TransitionInDays: jsii.Number(123),
//
//   						// the properties below are optional
//   						NewerNoncurrentVersions: jsii.Number(123),
//   					},
//   				},
//   				ObjectSizeGreaterThan: jsii.Number(123),
//   				ObjectSizeLessThan: jsii.Number(123),
//   				Prefix: jsii.String("prefix"),
//   				TagFilters: []interface{}{
//   					&TagFilterProperty{
//   						Key: jsii.String("key"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   				Transition: &TransitionProperty{
//   					StorageClass: jsii.String("storageClass"),
//
//   					// the properties below are optional
//   					TransitionDate: NewDate(),
//   					TransitionInDays: jsii.Number(123),
//   				},
//   				Transitions: []interface{}{
//   					&TransitionProperty{
//   						StorageClass: jsii.String("storageClass"),
//
//   						// the properties below are optional
//   						TransitionDate: NewDate(),
//   						TransitionInDays: jsii.Number(123),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	LoggingConfiguration: &LoggingConfigurationProperty{
//   		DestinationBucketName: jsii.String("destinationBucketName"),
//   		LogFilePrefix: jsii.String("logFilePrefix"),
//   	},
//   	MetricsConfigurations: []interface{}{
//   		&MetricsConfigurationProperty{
//   			Id: jsii.String("id"),
//
//   			// the properties below are optional
//   			AccessPointArn: jsii.String("accessPointArn"),
//   			Prefix: jsii.String("prefix"),
//   			TagFilters: []interface{}{
//   				&TagFilterProperty{
//   					Key: jsii.String("key"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   	NotificationConfiguration: &NotificationConfigurationProperty{
//   		EventBridgeConfiguration: &EventBridgeConfigurationProperty{
//   			EventBridgeEnabled: jsii.Boolean(false),
//   		},
//   		LambdaConfigurations: []interface{}{
//   			&LambdaConfigurationProperty{
//   				Event: jsii.String("event"),
//   				Function: jsii.String("function"),
//
//   				// the properties below are optional
//   				Filter: &NotificationFilterProperty{
//   					S3Key: &S3KeyFilterProperty{
//   						Rules: []interface{}{
//   							&FilterRuleProperty{
//   								Name: jsii.String("name"),
//   								Value: jsii.String("value"),
//   							},
//   						},
//   					},
//   				},
//   			},
//   		},
//   		QueueConfigurations: []interface{}{
//   			&QueueConfigurationProperty{
//   				Event: jsii.String("event"),
//   				Queue: jsii.String("queue"),
//
//   				// the properties below are optional
//   				Filter: &NotificationFilterProperty{
//   					S3Key: &S3KeyFilterProperty{
//   						Rules: []interface{}{
//   							&FilterRuleProperty{
//   								Name: jsii.String("name"),
//   								Value: jsii.String("value"),
//   							},
//   						},
//   					},
//   				},
//   			},
//   		},
//   		TopicConfigurations: []interface{}{
//   			&TopicConfigurationProperty{
//   				Event: jsii.String("event"),
//   				Topic: jsii.String("topic"),
//
//   				// the properties below are optional
//   				Filter: &NotificationFilterProperty{
//   					S3Key: &S3KeyFilterProperty{
//   						Rules: []interface{}{
//   							&FilterRuleProperty{
//   								Name: jsii.String("name"),
//   								Value: jsii.String("value"),
//   							},
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	ObjectLockConfiguration: &ObjectLockConfigurationProperty{
//   		ObjectLockEnabled: jsii.String("objectLockEnabled"),
//   		Rule: &ObjectLockRuleProperty{
//   			DefaultRetention: &DefaultRetentionProperty{
//   				Days: jsii.Number(123),
//   				Mode: jsii.String("mode"),
//   				Years: jsii.Number(123),
//   			},
//   		},
//   	},
//   	ObjectLockEnabled: jsii.Boolean(false),
//   	OwnershipControls: &OwnershipControlsProperty{
//   		Rules: []interface{}{
//   			&OwnershipControlsRuleProperty{
//   				ObjectOwnership: jsii.String("objectOwnership"),
//   			},
//   		},
//   	},
//   	PublicAccessBlockConfiguration: &PublicAccessBlockConfigurationProperty{
//   		BlockPublicAcls: jsii.Boolean(false),
//   		BlockPublicPolicy: jsii.Boolean(false),
//   		IgnorePublicAcls: jsii.Boolean(false),
//   		RestrictPublicBuckets: jsii.Boolean(false),
//   	},
//   	ReplicationConfiguration: &ReplicationConfigurationProperty{
//   		Role: jsii.String("role"),
//   		Rules: []interface{}{
//   			&ReplicationRuleProperty{
//   				Destination: &ReplicationDestinationProperty{
//   					Bucket: jsii.String("bucket"),
//
//   					// the properties below are optional
//   					AccessControlTranslation: &AccessControlTranslationProperty{
//   						Owner: jsii.String("owner"),
//   					},
//   					Account: jsii.String("account"),
//   					EncryptionConfiguration: &EncryptionConfigurationProperty{
//   						ReplicaKmsKeyId: jsii.String("replicaKmsKeyId"),
//   					},
//   					Metrics: &MetricsProperty{
//   						Status: jsii.String("status"),
//
//   						// the properties below are optional
//   						EventThreshold: &ReplicationTimeValueProperty{
//   							Minutes: jsii.Number(123),
//   						},
//   					},
//   					ReplicationTime: &ReplicationTimeProperty{
//   						Status: jsii.String("status"),
//   						Time: &ReplicationTimeValueProperty{
//   							Minutes: jsii.Number(123),
//   						},
//   					},
//   					StorageClass: jsii.String("storageClass"),
//   				},
//   				Status: jsii.String("status"),
//
//   				// the properties below are optional
//   				DeleteMarkerReplication: &DeleteMarkerReplicationProperty{
//   					Status: jsii.String("status"),
//   				},
//   				Filter: &ReplicationRuleFilterProperty{
//   					And: &ReplicationRuleAndOperatorProperty{
//   						Prefix: jsii.String("prefix"),
//   						TagFilters: []interface{}{
//   							&TagFilterProperty{
//   								Key: jsii.String("key"),
//   								Value: jsii.String("value"),
//   							},
//   						},
//   					},
//   					Prefix: jsii.String("prefix"),
//   					TagFilter: &TagFilterProperty{
//   						Key: jsii.String("key"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   				Id: jsii.String("id"),
//   				Prefix: jsii.String("prefix"),
//   				Priority: jsii.Number(123),
//   				SourceSelectionCriteria: &SourceSelectionCriteriaProperty{
//   					ReplicaModifications: &ReplicaModificationsProperty{
//   						Status: jsii.String("status"),
//   					},
//   					SseKmsEncryptedObjects: &SseKmsEncryptedObjectsProperty{
//   						Status: jsii.String("status"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VersioningConfiguration: &VersioningConfigurationProperty{
//   		Status: jsii.String("status"),
//   	},
//   	WebsiteConfiguration: &WebsiteConfigurationProperty{
//   		ErrorDocument: jsii.String("errorDocument"),
//   		IndexDocument: jsii.String("indexDocument"),
//   		RedirectAllRequestsTo: &RedirectAllRequestsToProperty{
//   			HostName: jsii.String("hostName"),
//
//   			// the properties below are optional
//   			Protocol: jsii.String("protocol"),
//   		},
//   		RoutingRules: []interface{}{
//   			&RoutingRuleProperty{
//   				RedirectRule: &RedirectRuleProperty{
//   					HostName: jsii.String("hostName"),
//   					HttpRedirectCode: jsii.String("httpRedirectCode"),
//   					Protocol: jsii.String("protocol"),
//   					ReplaceKeyPrefixWith: jsii.String("replaceKeyPrefixWith"),
//   					ReplaceKeyWith: jsii.String("replaceKeyWith"),
//   				},
//
//   				// the properties below are optional
//   				RoutingRuleCondition: &RoutingRuleConditionProperty{
//   					HttpErrorCodeReturnedEquals: jsii.String("httpErrorCodeReturnedEquals"),
//   					KeyPrefixEquals: jsii.String("keyPrefixEquals"),
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

