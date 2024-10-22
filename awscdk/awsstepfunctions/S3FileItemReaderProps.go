package awsstepfunctions

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Base interface for Item Reader configuration properties the iterate over entries in a S3 file.
//
// Example:
//   import s3 "github.com/aws/aws-cdk-go/awscdk"
//
//
//   /**
//    * Tree view of bucket:
//    *  my-bucket
//    *  |
//    *  +--input.json
//    *  |
//    *  ...
//    *
//    * File content of input.json:
//    *  [
//    *    "item1",
//    *    "item2"
//    *  ]
//    */
//   bucket := s3.NewBucket(this, jsii.String("Bucket"), &BucketProps{
//   	BucketName: jsii.String("my-bucket"),
//   })
//
//   distributedMap := sfn.NewDistributedMap(this, jsii.String("DistributedMap"), &DistributedMapProps{
//   	ItemReader: sfn.NewS3JsonItemReader(&S3FileItemReaderProps{
//   		Bucket: *Bucket,
//   		Key: jsii.String("input.json"),
//   	}),
//   })
//   distributedMap.ItemProcessor(sfn.NewPass(this, jsii.String("Pass")))
//
type S3FileItemReaderProps struct {
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
	// Key of file stored in S3 bucket containing an array to iterate over.
	Key *string `field:"required" json:"key" yaml:"key"`
}

