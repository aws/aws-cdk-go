package awsstepfunctions

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Interface for Result Writer configuration properties.
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
type ResultWriterProps struct {
	// S3 Bucket in which to save Map Run results.
	Bucket awss3.IBucket `field:"required" json:"bucket" yaml:"bucket"`
	// S3 prefix in which to save Map Run results.
	// Default: - No prefix.
	//
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

