package awsssm


// Properties for defining a `CfnResourceDataSync`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnResourceDataSyncProps := &CfnResourceDataSyncProps{
//   	SyncName: jsii.String("syncName"),
//
//   	// the properties below are optional
//   	BucketName: jsii.String("bucketName"),
//   	BucketPrefix: jsii.String("bucketPrefix"),
//   	BucketRegion: jsii.String("bucketRegion"),
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   	S3Destination: &S3DestinationProperty{
//   		BucketName: jsii.String("bucketName"),
//   		BucketRegion: jsii.String("bucketRegion"),
//   		SyncFormat: jsii.String("syncFormat"),
//
//   		// the properties below are optional
//   		BucketPrefix: jsii.String("bucketPrefix"),
//   		KmsKeyArn: jsii.String("kmsKeyArn"),
//   	},
//   	SyncFormat: jsii.String("syncFormat"),
//   	SyncSource: &SyncSourceProperty{
//   		SourceRegions: []*string{
//   			jsii.String("sourceRegions"),
//   		},
//   		SourceType: jsii.String("sourceType"),
//
//   		// the properties below are optional
//   		AwsOrganizationsSource: &AwsOrganizationsSourceProperty{
//   			OrganizationSourceType: jsii.String("organizationSourceType"),
//
//   			// the properties below are optional
//   			OrganizationalUnits: []*string{
//   				jsii.String("organizationalUnits"),
//   			},
//   		},
//   		IncludeFutureRegions: jsii.Boolean(false),
//   	},
//   	SyncType: jsii.String("syncType"),
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

