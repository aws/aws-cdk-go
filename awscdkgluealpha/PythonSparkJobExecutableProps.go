package awscdkgluealpha


// Props for creating a Python Spark (ETL or Streaming) job executable.
//
// Example:
//   glue.NewJob(this, jsii.String("EnableRunQueuing"), &JobProps{
//   	JobName: jsii.String("EtlJobWithRunQueuing"),
//   	Executable: glue.JobExecutable_PythonEtl(&PythonSparkJobExecutableProps{
//   		GlueVersion: glue.GlueVersion_V5_0(),
//   		PythonVersion: glue.PythonVersion_THREE,
//   		Script: glue.Code_FromAsset(path.join(__dirname, jsii.String("job-script"), jsii.String("hello_world.py"))),
//   	}),
//   	JobRunQueuingEnabled: jsii.Boolean(true),
//   })
//
// Experimental.
type PythonSparkJobExecutableProps struct {
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
	// Additional Java .jar files that AWS Glue adds to the Java classpath before executing your script. Only individual files are supported, directories are not supported. Equivalent to a job parameter `--extra-jars`.
	// See: https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	// Default: [] - no extra jars are added to the classpath.
	//
	// Experimental.
	ExtraJars *[]Code `field:"optional" json:"extraJars" yaml:"extraJars"`
	// Setting this value to true prioritizes the customer's extra JAR files in the classpath.
	//
	// Equivalent to a job parameter `--user-jars-first`.
	// See: https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	// Default: false - priority is not given to user-provided jars.
	//
	// Experimental.
	ExtraJarsFirst *bool `field:"optional" json:"extraJarsFirst" yaml:"extraJarsFirst"`
	// Additional Python files that AWS Glue adds to the Python path before executing your script.
	//
	// Only individual files are supported, directories are not supported.
	// Equivalent to a job parameter `--extra-py-files`.
	// See: https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
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

