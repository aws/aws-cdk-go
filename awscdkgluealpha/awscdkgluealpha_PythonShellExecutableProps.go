// The CDK Construct Library for AWS::Glue
package awscdkgluealpha


// Props for creating a Python shell job executable.
//
// Example:
//   var bucket bucket
//
//   glue.NewJob(this, jsii.String("PythonShellJob"), &jobProps{
//   	executable: glue.jobExecutable.pythonShell(&pythonShellExecutableProps{
//   		glueVersion: glue.glueVersion_V1_0(),
//   		pythonVersion: glue.pythonVersion_THREE,
//   		script: glue.code.fromBucket(bucket, jsii.String("script.py")),
//   	}),
//   	description: jsii.String("an example Python Shell job"),
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
	// See: `--extra-files` in https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	// Experimental.
	ExtraFiles *[]Code `field:"optional" json:"extraFiles" yaml:"extraFiles"`
	// Additional Python files that AWS Glue adds to the Python path before executing your script.
	//
	// Only individual files are supported, directories are not supported.
	// See: `--extra-py-files` in https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	// Experimental.
	ExtraPythonFiles *[]Code `field:"optional" json:"extraPythonFiles" yaml:"extraPythonFiles"`
}

