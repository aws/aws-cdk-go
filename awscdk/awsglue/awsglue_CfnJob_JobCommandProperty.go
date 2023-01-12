package awsglue


// Specifies code executed when a job is run.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   jobCommandProperty := &jobCommandProperty{
//   	name: jsii.String("name"),
//   	pythonVersion: jsii.String("pythonVersion"),
//   	scriptLocation: jsii.String("scriptLocation"),
//   }
//
type CfnJob_JobCommandProperty struct {
	// The name of the job command.
	//
	// For an Apache Spark ETL job, this must be `glueetl` . For a Python shell job, it must be `pythonshell` .
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The Python version being used to execute a Python shell job.
	//
	// Allowed values are 2 or 3.
	PythonVersion *string `field:"optional" json:"pythonVersion" yaml:"pythonVersion"`
	// Specifies the Amazon Simple Storage Service (Amazon S3) path to a script that executes a job (required).
	ScriptLocation *string `field:"optional" json:"scriptLocation" yaml:"scriptLocation"`
}

