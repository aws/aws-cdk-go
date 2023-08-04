package awscdkgluealpha


// Props for creating a Python shell job executable.
//
// Example:
//   var bucket bucket
//
//   glue.NewJob(this, jsii.String("PythonShellJob"), &JobProps{
//   	Executable: glue.JobExecutable_PythonShell(&PythonShellExecutableProps{
//   		GlueVersion: glue.GlueVersion_V1_0(),
//   		PythonVersion: glue.PythonVersion_THREE,
//   		Script: glue.Code_FromBucket(bucket, jsii.String("script.py")),
//   	}),
//   	Description: jsii.String("an example Python Shell job"),
//   })
//
// Experimental.
type PythonShellExecutableProps struct {
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
	// See:  `--extra-files` in https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	// Default: [] - no extra files are copied to the working directory.
	//
	// Experimental.
	ExtraFiles *[]Code `field:"optional" json:"extraFiles" yaml:"extraFiles"`
	// Additional Python files that AWS Glue adds to the Python path before executing your script.
	//
	// Only individual files are supported, directories are not supported.
	// See:  `--extra-py-files` in https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	// Default: - no extra python files and argument is not set.
	//
	// Experimental.
	ExtraPythonFiles *[]Code `field:"optional" json:"extraPythonFiles" yaml:"extraPythonFiles"`
	// Runtime.
	//
	// It is required for Ray jobs.
	// Experimental.
	Runtime Runtime `field:"optional" json:"runtime" yaml:"runtime"`
}

