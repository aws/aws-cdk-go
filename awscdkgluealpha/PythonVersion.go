package awscdkgluealpha


// Python version.
//
// Example:
//   glue.NewJob(this, jsii.String("EnableRunQueuing"), &JobProps{
//   	JobName: jsii.String("EtlJobWithRunQueuing"),
//   	Executable: glue.JobExecutable_PythonEtl(&PythonSparkJobExecutableProps{
//   		GlueVersion: glue.GlueVersion_V4_0(),
//   		PythonVersion: glue.PythonVersion_THREE,
//   		Script: glue.Code_FromAsset(path.join(__dirname, jsii.String("job-script"), jsii.String("hello_world.py"))),
//   	}),
//   	JobRunQueuingEnabled: jsii.Boolean(true),
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

