package awss3


// Specifies the inventory configuration for an Amazon S3 bucket.
//
// For more information, see [GET Bucket inventory](https://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketGETInventoryConfig.html) in the *Amazon S3 API Reference* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inventoryConfigurationProperty := &inventoryConfigurationProperty{
//   	destination: &destinationProperty{
//   		bucketArn: jsii.String("bucketArn"),
//   		format: jsii.String("format"),
//
//   		// the properties below are optional
//   		bucketAccountId: jsii.String("bucketAccountId"),
//   		prefix: jsii.String("prefix"),
//   	},
//   	enabled: jsii.Boolean(false),
//   	id: jsii.String("id"),
//   	includedObjectVersions: jsii.String("includedObjectVersions"),
//   	scheduleFrequency: jsii.String("scheduleFrequency"),
//
//   	// the properties below are optional
//   	optionalFields: []*string{
//   		jsii.String("optionalFields"),
//   	},
//   	prefix: jsii.String("prefix"),
//   }
//
type CfnBucket_InventoryConfigurationProperty struct {
	// Contains information about where to publish the inventory results.
	Destination interface{} `field:"required" json:"destination" yaml:"destination"`
	// Specifies whether the inventory is enabled or disabled.
	//
	// If set to `True` , an inventory list is generated. If set to `False` , no inventory list is generated.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// The ID used to identify the inventory configuration.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Object versions to include in the inventory list.
	//
	// If set to `All` , the list includes all the object versions, which adds the version-related fields `VersionId` , `IsLatest` , and `DeleteMarker` to the list. If set to `Current` , the list does not contain these version-related fields.
	IncludedObjectVersions *string `field:"required" json:"includedObjectVersions" yaml:"includedObjectVersions"`
	// Specifies the schedule for generating inventory results.
	//
	// *Allowed values* : `Daily` | `Weekly`.
	ScheduleFrequency *string `field:"required" json:"scheduleFrequency" yaml:"scheduleFrequency"`
	// Contains the optional fields that are included in the inventory results.
	//
	// *Valid values* : `Size | LastModifiedDate | StorageClass | ETag | IsMultipartUploaded | ReplicationStatus | EncryptionStatus | ObjectLockRetainUntilDate | ObjectLockMode | ObjectLockLegalHoldStatus | IntelligentTieringAccessTier | BucketKeyStatus`.
	OptionalFields *[]*string `field:"optional" json:"optionalFields" yaml:"optionalFields"`
	// Specifies the inventory filter prefix.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

