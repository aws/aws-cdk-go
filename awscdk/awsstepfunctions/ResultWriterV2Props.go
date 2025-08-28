package awsstepfunctions

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Interface for Result Writer configuration props.
//
// Example:
//   import s3 "github.com/aws/aws-cdk-go/awscdk"
//
//
//   // create a bucket
//   bucket := s3.NewBucket(this, jsii.String("Bucket"))
//
//   // create a WriterConfig
//
//   distributedMap := sfn.NewDistributedMap(this, jsii.String("Distributed Map State"), &DistributedMapProps{
//   	ResultWriterV2: sfn.NewResultWriterV2(&ResultWriterV2Props{
//   		Bucket: bucket,
//   		Prefix: jsii.String("my-prefix"),
//   		WriterConfig: map[string]interface{}{
//   			"outputType": sfn.OutputType_JSONL,
//   			"transformation": sfn.Transformation_NONE,
//   		},
//   	}),
//   })
//   distributedMap.ItemProcessor(sfn.NewPass(this, jsii.String("Pass State")))
//
type ResultWriterV2Props struct {
	// S3 Bucket in which to save Map Run results.
	// Default: - specify a bucket.
	//
	Bucket awss3.IBucket `field:"optional" json:"bucket" yaml:"bucket"`
	// S3 bucket name in which to save Map Run results, as JsonPath.
	// Default: - no bucket path.
	//
	BucketNamePath *string `field:"optional" json:"bucketNamePath" yaml:"bucketNamePath"`
	// S3 prefix in which to save Map Run results.
	// Default: - No prefix.
	//
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// Configuration to format the output of the Child Workflow executions.
	// Default: - Specify both Transformation and OutputType.
	//
	WriterConfig WriterConfig `field:"optional" json:"writerConfig" yaml:"writerConfig"`
}

