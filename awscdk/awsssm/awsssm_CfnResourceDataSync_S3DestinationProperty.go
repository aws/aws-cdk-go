package awsssm


// Information about the target S3 bucket for the resource data sync.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3DestinationProperty := &s3DestinationProperty{
//   	bucketName: jsii.String("bucketName"),
//   	bucketRegion: jsii.String("bucketRegion"),
//   	syncFormat: jsii.String("syncFormat"),
//
//   	// the properties below are optional
//   	bucketPrefix: jsii.String("bucketPrefix"),
//   	kmsKeyArn: jsii.String("kmsKeyArn"),
//   }
//
type CfnResourceDataSync_S3DestinationProperty struct {
	// The name of the S3 bucket where the aggregated data is stored.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// The AWS Region with the S3 bucket targeted by the resource data sync.
	BucketRegion *string `field:"required" json:"bucketRegion" yaml:"bucketRegion"`
	// A supported sync format.
	//
	// The following format is currently supported: JsonSerDe.
	SyncFormat *string `field:"required" json:"syncFormat" yaml:"syncFormat"`
	// An Amazon S3 prefix for the bucket.
	BucketPrefix *string `field:"optional" json:"bucketPrefix" yaml:"bucketPrefix"`
	// The ARN of an encryption key for a destination in Amazon S3.
	//
	// Must belong to the same Region as the destination S3 bucket.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
}

