package awsglue


// Specifies code executed when a job is run.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   jobCommandProperty := &JobCommandProperty{
//   	Name: jsii.String("name"),
//   	PythonVersion: jsii.String("pythonVersion"),
//   	Runtime: jsii.String("runtime"),
//   	ScriptLocation: jsii.String("scriptLocation"),
//   }
//
type CfnJob_JobCommandProperty struct {
	// The name of the job command.
	//
	// For an Apache Spark ETL job, this must be `glueetl` . For a Python shell job, it must be `pythonshell` . For an Apache Spark streaming ETL job, this must be `gluestreaming` . For a Ray job, this must be `glueray` .
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The Python version being used to execute a Python shell job.
	//
	// Allowed values are 3 or 3.9. Version 2 is deprecated.
	PythonVersion *string `field:"optional" json:"pythonVersion" yaml:"pythonVersion"`
	// In Ray jobs, Runtime is used to specify the versions of Ray, Python and additional libraries available in your environment.
	//
	// This field is not used in other job types. For supported runtime environment values, see [Working with Ray jobs](https://docs.aws.amazon.com/glue/latest/dg/ray-jobs-section.html) in the AWS Glue Developer Guide.
	Runtime *string `field:"optional" json:"runtime" yaml:"runtime"`
	// Specifies the Amazon Simple Storage Service (Amazon S3) path to a script that executes a job (required).
	ScriptLocation *string `field:"optional" json:"scriptLocation" yaml:"scriptLocation"`
}

