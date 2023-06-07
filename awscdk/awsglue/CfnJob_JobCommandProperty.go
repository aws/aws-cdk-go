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
	// For an Apache Spark ETL job, this must be `glueetl` . For a Python shell job, it must be `pythonshell` . For an Apache Spark streaming ETL job, this must be `gluestreaming` .
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The Python version being used to execute a Python shell job.
	//
	// Allowed values are 3 or 3.9. Version 2 is deprecated.
	PythonVersion *string `field:"optional" json:"pythonVersion" yaml:"pythonVersion"`
	// `CfnJob.JobCommandProperty.Runtime`.
	Runtime *string `field:"optional" json:"runtime" yaml:"runtime"`
	// Specifies the Amazon Simple Storage Service (Amazon S3) path to a script that executes a job (required).
	ScriptLocation *string `field:"optional" json:"scriptLocation" yaml:"scriptLocation"`
}

