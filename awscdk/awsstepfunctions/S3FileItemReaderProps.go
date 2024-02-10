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
//   // create a bucket
//   bucket := s3.NewBucket(this, jsii.String("Bucket"))
//
//   distributedMap := sfn.NewDistributedMap(this, jsii.String("Distributed Map State"), &DistributedMapProps{
//   	ItemReader: sfn.NewS3JsonItemReader(&S3FileItemReaderProps{
//   		Bucket: bucket,
//   		Key: jsii.String("my-key.json"),
//   	}),
//   	ResultWriter: sfn.NewResultWriter(&ResultWriterProps{
//   		Bucket: bucket,
//   		Prefix: jsii.String("my-prefix"),
//   	}),
//   })
//   distributedMap.ItemProcessor(sfn.NewPass(this, jsii.String("Pass State")))
//
type S3FileItemReaderProps struct {
	// S3 Bucket containing objects to iterate over or a file with a list to iterate over.
	Bucket awss3.IBucket `field:"required" json:"bucket" yaml:"bucket"`
	// Limits the number of items passed to the Distributed Map state.
	// Default: - Distributed Map state will iterate over all items provided by the ItemReader.
	//
	MaxItems *float64 `field:"optional" json:"maxItems" yaml:"maxItems"`
	// Key of file stored in S3 bucket containing an array to iterate over.
	Key *string `field:"required" json:"key" yaml:"key"`
}

