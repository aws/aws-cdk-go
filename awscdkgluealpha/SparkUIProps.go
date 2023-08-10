package awscdkgluealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Properties for enabling Spark UI monitoring feature for Spark-based Glue jobs.
//
// Example:
//   glue.NewJob(this, jsii.String("EnableSparkUI"), &JobProps{
//   	JobName: jsii.String("EtlJobWithSparkUIPrefix"),
//   	SparkUI: &SparkUIProps{
//   		Enabled: jsii.Boolean(true),
//   	},
//   	Executable: glue.JobExecutable_PythonEtl(&PythonSparkJobExecutableProps{
//   		GlueVersion: glue.GlueVersion_V3_0(),
//   		PythonVersion: glue.PythonVersion_THREE,
//   		Script: glue.Code_FromAsset(path.join(__dirname, jsii.String("job-script/hello_world.py"))),
//   	}),
//   })
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
	//
	// Use format `'/foo/bar'`.
	// Experimental.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

