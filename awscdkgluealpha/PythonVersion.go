package awscdkgluealpha


// Python version.
//
// Example:
//   glue.NewJob(this, jsii.String("PythonSparkStreamingJob"), &JobProps{
//   	Executable: glue.JobExecutable_PythonStreaming(&PythonSparkJobExecutableProps{
//   		GlueVersion: glue.GlueVersion_V4_0(),
//   		PythonVersion: glue.PythonVersion_THREE,
//   		Script: glue.Code_FromAsset(path.join(__dirname, jsii.String("job-script/hello_world.py"))),
//   	}),
//   	Description: jsii.String("an example Python Streaming job"),
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
	// Python 3.9 (the exact version depends on GlueVersion and JobCommand used).
	// Experimental.
	PythonVersion_THREE_NINE PythonVersion = "THREE_NINE"
)

