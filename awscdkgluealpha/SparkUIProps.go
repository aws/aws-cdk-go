package awscdkgluealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Properties for enabling Spark UI monitoring feature for Spark-based Glue jobs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import glue_alpha "github.com/aws/aws-cdk-go/awscdkgluealpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bucket bucket
//
//   sparkUIProps := &SparkUIProps{
//   	Bucket: bucket,
//   	JobRunQueuingEnabled: jsii.Boolean(false),
//   	Prefix: jsii.String("prefix"),
//   }
//
// See: https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
//
// Experimental.
type SparkUIProps struct {
	// The bucket where the Glue job stores the logs.
	// Default: a new bucket will be created.
	//
	// Experimental.
	Bucket awss3.IBucket `field:"optional" json:"bucket" yaml:"bucket"`
	// Specifies whether job run queuing is enabled for the job runs for this job.
	//
	// A value of true means job run queuing is enabled for the job runs.
	// If false or not populated, the job runs will not be considered for queueing.
	// If this field does not match the value set in the job run, then the value from
	// the job run field will be used. This property must be set to false for flex jobs.
	// If this property is enabled, maxRetries must be set to zero.
	// Default: - no job run queuing.
	//
	// Experimental.
	JobRunQueuingEnabled *bool `field:"optional" json:"jobRunQueuingEnabled" yaml:"jobRunQueuingEnabled"`
	// The path inside the bucket (objects prefix) where the Glue job stores the logs.
	//
	// Use format `'/foo/bar'`.
	// Default: - the logs will be written at the root of the bucket.
	//
	// Experimental.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

