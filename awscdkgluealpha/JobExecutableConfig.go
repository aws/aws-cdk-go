package awscdkgluealpha


// Result of binding a `JobExecutable` into a `Job`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkgluealpha"
//
//   var code code
//   var glueVersion glueVersion
//   var jobType jobType
//   var runtime runtime
//
//   jobExecutableConfig := &JobExecutableConfig{
//   	GlueVersion: glueVersion,
//   	Language: glue_alpha.JobLanguage_SCALA,
//   	Script: code,
//   	Type: jobType,
//
//   	// the properties below are optional
//   	ClassName: jsii.String("className"),
//   	ExtraFiles: []*code{
//   		code,
//   	},
//   	ExtraJars: []*code{
//   		code,
//   	},
//   	ExtraJarsFirst: jsii.Boolean(false),
//   	ExtraPythonFiles: []*code{
//   		code,
//   	},
//   	PythonVersion: glue_alpha.PythonVersion_TWO,
//   	Runtime: runtime,
//   	S3PythonModules: []*code{
//   		code,
//   	},
//   }
//
// Experimental.
type JobExecutableConfig struct {
	// Glue version.
	// See: https://docs.aws.amazon.com/glue/latest/dg/release-notes.html
	//
	// Experimental.
	GlueVersion GlueVersion `field:"required" json:"glueVersion" yaml:"glueVersion"`
	// The language of the job (Scala or Python).
	//
	// Equivalent to a job parameter `--job-language`.
	// See: https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	// Experimental.
	Language JobLanguage `field:"required" json:"language" yaml:"language"`
	// The script that is executed by a job.
	// Experimental.
	Script Code `field:"required" json:"script" yaml:"script"`
	// Specify the type of the job whether it's an Apache Spark ETL or streaming one or if it's a Python shell job.
	// Experimental.
	Type JobType `field:"required" json:"type" yaml:"type"`
	// The Scala class that serves as the entry point for the job.
	//
	// This applies only if your the job langauage is Scala.
	// Equivalent to a job parameter `--class`.
	// See: https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	// Default: - no scala className specified.
	//
	// Experimental.
	ClassName *string `field:"optional" json:"className" yaml:"className"`
	// Additional files, such as configuration files that AWS Glue copies to the working directory of your script before executing it.
	//
	// Equivalent to a job parameter `--extra-files`.
	// See: https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	// Default: - no extra files specified.
	//
	// Experimental.
	ExtraFiles *[]Code `field:"optional" json:"extraFiles" yaml:"extraFiles"`
	// Additional Java .jar files that AWS Glue adds to the Java classpath before executing your script. Equivalent to a job parameter `--extra-jars`.
	// See: https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	// Default: - no extra jars specified.
	//
	// Experimental.
	ExtraJars *[]Code `field:"optional" json:"extraJars" yaml:"extraJars"`
	// Setting this value to true prioritizes the customer's extra JAR files in the classpath.
	//
	// Equivalent to a job parameter `--user-jars-first`.
	// See: https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	// Default: - extra jars are not prioritized.
	//
	// Experimental.
	ExtraJarsFirst *bool `field:"optional" json:"extraJarsFirst" yaml:"extraJarsFirst"`
	// Additional Python files that AWS Glue adds to the Python path before executing your script.
	//
	// Equivalent to a job parameter `--extra-py-files`.
	// See: https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	// Default: - no extra python files specified.
	//
	// Experimental.
	ExtraPythonFiles *[]Code `field:"optional" json:"extraPythonFiles" yaml:"extraPythonFiles"`
	// The Python version to use.
	// Default: - no python version specified.
	//
	// Experimental.
	PythonVersion PythonVersion `field:"optional" json:"pythonVersion" yaml:"pythonVersion"`
	// The Runtime to use.
	// Default: - no runtime specified.
	//
	// Experimental.
	Runtime Runtime `field:"optional" json:"runtime" yaml:"runtime"`
	// Additional Python modules that AWS Glue adds to the Python path before executing your script.
	//
	// Equivalent to a job parameter `--s3-py-modules`.
	// See: https://docs.aws.amazon.com/glue/latest/dg/author-job-ray-job-parameters.html
	//
	// Default: - no extra python files specified.
	//
	// Experimental.
	S3PythonModules *[]Code `field:"optional" json:"s3PythonModules" yaml:"s3PythonModules"`
}

