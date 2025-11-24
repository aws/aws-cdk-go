package mixinsawsglue


// Specifies code executed when a job is run.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   jobCommandProperty := &JobCommandProperty{
//   	Name: jsii.String("name"),
//   	PythonVersion: jsii.String("pythonVersion"),
//   	Runtime: jsii.String("runtime"),
//   	ScriptLocation: jsii.String("scriptLocation"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-job-jobcommand.html
//
type CfnJobPropsMixin_JobCommandProperty struct {
	// The name of the job command.
	//
	// For an Apache Spark ETL job, this must be `glueetl` . For a Python shell job, it must be `pythonshell` . For an Apache Spark streaming ETL job, this must be `gluestreaming` . For a Ray job, this must be `glueray` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-job-jobcommand.html#cfn-glue-job-jobcommand-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The Python version being used to execute a Python shell job.
	//
	// Allowed values are 3 or 3.9. Version 2 is deprecated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-job-jobcommand.html#cfn-glue-job-jobcommand-pythonversion
	//
	PythonVersion *string `field:"optional" json:"pythonVersion" yaml:"pythonVersion"`
	// In Ray jobs, Runtime is used to specify the versions of Ray, Python and additional libraries available in your environment.
	//
	// This field is not used in other job types. For supported runtime environment values, see [Working with Ray jobs](https://docs.aws.amazon.com/glue/latest/dg/ray-jobs-section.html) in the AWS Glue Developer Guide.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-job-jobcommand.html#cfn-glue-job-jobcommand-runtime
	//
	Runtime *string `field:"optional" json:"runtime" yaml:"runtime"`
	// Specifies the Amazon Simple Storage Service (Amazon S3) path to a script that executes a job (required).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-job-jobcommand.html#cfn-glue-job-jobcommand-scriptlocation
	//
	ScriptLocation *string `field:"optional" json:"scriptLocation" yaml:"scriptLocation"`
}

