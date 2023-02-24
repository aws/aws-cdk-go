// The CDK Construct Library for AWS::Glue
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
//   	Enabled: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	Bucket: bucket,
//   	Prefix: jsii.String("prefix"),
//   }
//
// See: https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
//
// Experimental.
type SparkUIProps struct {
	// Enable Spark UI.
	// Experimental.
	Enabled *bool `field:"required" json:"enabled" yaml:"enabled"`
	// The bucket where the Glue job stores the logs.
	// Experimental.
	Bucket awss3.IBucket `field:"optional" json:"bucket" yaml:"bucket"`
	// The path inside the bucket (objects prefix) where the Glue job stores the logs.
	// Experimental.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

