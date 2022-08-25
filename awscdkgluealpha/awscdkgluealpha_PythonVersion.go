// The CDK Construct Library for AWS::Glue
package awscdkgluealpha


// Python version.
//
// Example:
//   glue.NewJob(this, jsii.String("PythonSparkStreamingJob"), &jobProps{
//   	executable: glue.jobExecutable.pythonStreaming(&pythonSparkJobExecutableProps{
//   		glueVersion: glue.glueVersion_V2_0(),
//   		pythonVersion: glue.pythonVersion_THREE,
//   		script: glue.code.fromAsset(path.join(__dirname, jsii.String("job-script/hello_world.py"))),
//   	}),
//   	description: jsii.String("an example Python Streaming job"),
//   })
//
// Experimental.
type PythonVersion string

const (
	// Python 2 (the exact version depends on GlueVersion and JobCommand used).
	// Experimental.
	PythonVersion_TWO PythonVersion = "TWO"
	// Python 3 (the exact version depends on GlueVersion and JobCommand used).
	// Experimental.
	PythonVersion_THREE PythonVersion = "THREE"
)

