package awsstepfunctions

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Properties for configuring an Item Reader that iterates over objects in an S3 bucket.
//
// Example:
//   import s3 "github.com/aws/aws-cdk-go/awscdk"
//
//
//   /**
//    * Tree view of bucket:
//    *  my-bucket
//    *  |
//    *  +--item1
//    *  |
//    *  +--otherItem
//    *  |
//    *  +--item2
//    *  |
//    *  ...
//    */
//   bucket := s3.NewBucket(this, jsii.String("Bucket"), &BucketProps{
//   	BucketName: jsii.String("my-bucket"),
//   })
//
//   distributedMap := sfn.NewDistributedMap(this, jsii.String("DistributedMap"), &DistributedMapProps{
//   	ItemReader: sfn.NewS3ObjectsItemReader(&S3ObjectsItemReaderProps{
//   		Bucket: *Bucket,
//   		Prefix: jsii.String("item"),
//   	}),
//   })
//   distributedMap.ItemProcessor(sfn.NewPass(this, jsii.String("Pass")))
//
type S3ObjectsItemReaderProps struct {
	// S3 Bucket containing objects to iterate over or a file with a list to iterate over.
	// See: bucketNamePath.
	//
	// Default: - S3 bucket will be determined from.
	//
	Bucket awss3.IBucket `field:"optional" json:"bucket" yaml:"bucket"`
	// S3 bucket name containing objects to iterate over or a file with a list to iterate over, as JsonPath.
	// See: bucket.
	//
	// Default: - S3 bucket will be determined from.
	//
	BucketNamePath *string `field:"optional" json:"bucketNamePath" yaml:"bucketNamePath"`
	// Limits the number of items passed to the Distributed Map state.
	// Default: - Distributed Map state will iterate over all items provided by the ItemReader.
	//
	MaxItems *float64 `field:"optional" json:"maxItems" yaml:"maxItems"`
	// S3 prefix used to limit objects to iterate over.
	// Default: - No prefix.
	//
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

