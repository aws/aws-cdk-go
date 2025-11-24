package mixinsawsssm


// Properties for CfnResourceDataSyncPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnResourceDataSyncMixinProps := &CfnResourceDataSyncMixinProps{
//   	BucketName: jsii.String("bucketName"),
//   	BucketPrefix: jsii.String("bucketPrefix"),
//   	BucketRegion: jsii.String("bucketRegion"),
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   	S3Destination: &S3DestinationProperty{
//   		BucketName: jsii.String("bucketName"),
//   		BucketPrefix: jsii.String("bucketPrefix"),
//   		BucketRegion: jsii.String("bucketRegion"),
//   		KmsKeyArn: jsii.String("kmsKeyArn"),
//   		SyncFormat: jsii.String("syncFormat"),
//   	},
//   	SyncFormat: jsii.String("syncFormat"),
//   	SyncName: jsii.String("syncName"),
//   	SyncSource: &SyncSourceProperty{
//   		AwsOrganizationsSource: &AwsOrganizationsSourceProperty{
//   			OrganizationalUnits: []*string{
//   				jsii.String("organizationalUnits"),
//   			},
//   			OrganizationSourceType: jsii.String("organizationSourceType"),
//   		},
//   		IncludeFutureRegions: jsii.Boolean(false),
//   		SourceRegions: []*string{
//   			jsii.String("sourceRegions"),
//   		},
//   		SourceType: jsii.String("sourceType"),
//   	},
//   	SyncType: jsii.String("syncType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-resourcedatasync.html
//
type CfnResourceDataSyncMixinProps struct {
	// The name of the S3 bucket where the aggregated data is stored.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-resourcedatasync.html#cfn-ssm-resourcedatasync-bucketname
	//
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// An Amazon S3 prefix for the bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-resourcedatasync.html#cfn-ssm-resourcedatasync-bucketprefix
	//
	BucketPrefix *string `field:"optional" json:"bucketPrefix" yaml:"bucketPrefix"`
	// The AWS Region with the S3 bucket targeted by the resource data sync.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-resourcedatasync.html#cfn-ssm-resourcedatasync-bucketregion
	//
	BucketRegion *string `field:"optional" json:"bucketRegion" yaml:"bucketRegion"`
	// The Amazon Resource Name (ARN) of an encryption key for a destination in Amazon S3 .
	//
	// You can use a KMS key to encrypt inventory data in Amazon S3 . You must specify a key that exist in the same AWS Region as the destination Amazon S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-resourcedatasync.html#cfn-ssm-resourcedatasync-kmskeyarn
	//
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// Configuration information for the target S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-resourcedatasync.html#cfn-ssm-resourcedatasync-s3destination
	//
	S3Destination interface{} `field:"optional" json:"s3Destination" yaml:"s3Destination"`
	// A supported sync format.
	//
	// The following format is currently supported: JsonSerDe.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-resourcedatasync.html#cfn-ssm-resourcedatasync-syncformat
	//
	SyncFormat *string `field:"optional" json:"syncFormat" yaml:"syncFormat"`
	// A name for the resource data sync.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-resourcedatasync.html#cfn-ssm-resourcedatasync-syncname
	//
	SyncName *string `field:"optional" json:"syncName" yaml:"syncName"`
	// Information about the source where the data was synchronized.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-resourcedatasync.html#cfn-ssm-resourcedatasync-syncsource
	//
	SyncSource interface{} `field:"optional" json:"syncSource" yaml:"syncSource"`
	// The type of resource data sync.
	//
	// If `SyncType` is `SyncToDestination` , then the resource data sync synchronizes data to an S3 bucket. If the `SyncType` is `SyncFromSource` then the resource data sync synchronizes data from AWS Organizations or from multiple AWS Regions .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-resourcedatasync.html#cfn-ssm-resourcedatasync-synctype
	//
	SyncType *string `field:"optional" json:"syncType" yaml:"syncType"`
}

