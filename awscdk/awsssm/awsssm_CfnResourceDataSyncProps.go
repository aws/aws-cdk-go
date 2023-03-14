package awsssm


// Properties for defining a `CfnResourceDataSync`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnResourceDataSyncProps := &cfnResourceDataSyncProps{
//   	syncName: jsii.String("syncName"),
//
//   	// the properties below are optional
//   	bucketName: jsii.String("bucketName"),
//   	bucketPrefix: jsii.String("bucketPrefix"),
//   	bucketRegion: jsii.String("bucketRegion"),
//   	kmsKeyArn: jsii.String("kmsKeyArn"),
//   	s3Destination: &s3DestinationProperty{
//   		bucketName: jsii.String("bucketName"),
//   		bucketRegion: jsii.String("bucketRegion"),
//   		syncFormat: jsii.String("syncFormat"),
//
//   		// the properties below are optional
//   		bucketPrefix: jsii.String("bucketPrefix"),
//   		kmsKeyArn: jsii.String("kmsKeyArn"),
//   	},
//   	syncFormat: jsii.String("syncFormat"),
//   	syncSource: &syncSourceProperty{
//   		sourceRegions: []*string{
//   			jsii.String("sourceRegions"),
//   		},
//   		sourceType: jsii.String("sourceType"),
//
//   		// the properties below are optional
//   		awsOrganizationsSource: &awsOrganizationsSourceProperty{
//   			organizationSourceType: jsii.String("organizationSourceType"),
//
//   			// the properties below are optional
//   			organizationalUnits: []*string{
//   				jsii.String("organizationalUnits"),
//   			},
//   		},
//   		includeFutureRegions: jsii.Boolean(false),
//   	},
//   	syncType: jsii.String("syncType"),
//   }
//
type CfnResourceDataSyncProps struct {
	// A name for the resource data sync.
	SyncName *string `field:"required" json:"syncName" yaml:"syncName"`
	// The name of the S3 bucket where the aggregated data is stored.
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// An Amazon S3 prefix for the bucket.
	BucketPrefix *string `field:"optional" json:"bucketPrefix" yaml:"bucketPrefix"`
	// The AWS Region with the S3 bucket targeted by the resource data sync.
	BucketRegion *string `field:"optional" json:"bucketRegion" yaml:"bucketRegion"`
	// The ARN of an encryption key for a destination in Amazon S3 .
	//
	// You can use a KMS key to encrypt inventory data in Amazon S3 . You must specify a key that exist in the same region as the destination Amazon S3 bucket.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// Configuration information for the target S3 bucket.
	S3Destination interface{} `field:"optional" json:"s3Destination" yaml:"s3Destination"`
	// A supported sync format.
	//
	// The following format is currently supported: JsonSerDe.
	SyncFormat *string `field:"optional" json:"syncFormat" yaml:"syncFormat"`
	// Information about the source where the data was synchronized.
	SyncSource interface{} `field:"optional" json:"syncSource" yaml:"syncSource"`
	// The type of resource data sync.
	//
	// If `SyncType` is `SyncToDestination` , then the resource data sync synchronizes data to an S3 bucket. If the `SyncType` is `SyncFromSource` then the resource data sync synchronizes data from AWS Organizations or from multiple AWS Regions .
	SyncType *string `field:"optional" json:"syncType" yaml:"syncType"`
}

