package awscdkgluealpha


// Props for creating a Python Ray job executable.
//
// Example:
//   glue.NewJob(this, jsii.String("RayJob"), &JobProps{
//   	Executable: glue.JobExecutable_PythonRay(&PythonRayExecutableProps{
//   		GlueVersion: glue.GlueVersion_V4_0(),
//   		PythonVersion: glue.PythonVersion_THREE_NINE,
//   		Runtime: glue.Runtime_RAY_TWO_FOUR(),
//   		Script: glue.Code_FromAsset(path.join(__dirname, jsii.String("job-script"), jsii.String("hello_world.py"))),
//   	}),
//   	WorkerType: glue.WorkerType_Z_2X(),
//   	WorkerCount: jsii.Number(2),
//   	Description: jsii.String("an example Ray job"),
//   })
//
// Experimental.
type PythonRayExecutableProps struct {
	// Glue version.
	// See: https://docs.aws.amazon.com/glue/latest/dg/release-notes.html
	//
	// Experimental.
	GlueVersion GlueVersion `field:"required" json:"glueVersion" yaml:"glueVersion"`
	// The Python version to use.
	// Experimental.
	PythonVersion PythonVersion `field:"required" json:"pythonVersion" yaml:"pythonVersion"`
	// The script that executes a job.
	// Experimental.
	Script Code `field:"required" json:"script" yaml:"script"`
	// Additional files, such as configuration files that AWS Glue copies to the working directory of your script before executing it.
	//
	// Only individual files are supported, directories are not supported.
	// Equivalent to a job parameter `--extra-files`.
	// See: https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	// Default: [] - no extra files are copied to the working directory.
	//
	// Experimental.
	ExtraFiles *[]Code `field:"optional" json:"extraFiles" yaml:"extraFiles"`
	// Runtime.
	//
	// It is required for Ray jobs.
	// Experimental.
	Runtime Runtime `field:"optional" json:"runtime" yaml:"runtime"`
	// Additional Python modules that AWS Glue adds to the Python path before executing your script.
	//
	// Equivalent to a job parameter `--s3-py-modules`.
	// See: https://docs.aws.amazon.com/glue/latest/dg/author-job-ray-job-parameters.html
	//
	// Default: - no extra python files and argument is not set.
	//
	// Experimental.
	S3PythonModules *[]Code `field:"optional" json:"s3PythonModules" yaml:"s3PythonModules"`
}

