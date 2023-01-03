// The CDK Construct Library for AWS::Glue
package awscdkgluealpha


// Result of binding a `JobExecutable` into a `Job`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import glue_alpha "github.com/aws/aws-cdk-go/awscdkgluealpha"
//
//   var code code
//   var glueVersion glueVersion
//   var jobType jobType
//
//   jobExecutableConfig := &jobExecutableConfig{
//   	glueVersion: glueVersion,
//   	language: glue_alpha.jobLanguage_SCALA,
//   	script: code,
//   	type: jobType,
//
//   	// the properties below are optional
//   	className: jsii.String("className"),
//   	extraFiles: []*code{
//   		code,
//   	},
//   	extraJars: []*code{
//   		code,
//   	},
//   	extraJarsFirst: jsii.Boolean(false),
//   	extraPythonFiles: []*code{
//   		code,
//   	},
//   	pythonVersion: glue_alpha.pythonVersion_TWO,
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
	// See: `--job-language` in https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
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
	// See: `--class` in https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	// Experimental.
	ClassName *string `field:"optional" json:"className" yaml:"className"`
	// Additional files, such as configuration files that AWS Glue copies to the working directory of your script before executing it.
	// See: `--extra-files` in https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	// Experimental.
	ExtraFiles *[]Code `field:"optional" json:"extraFiles" yaml:"extraFiles"`
	// Additional Java .jar files that AWS Glue adds to the Java classpath before executing your script.
	// See: `--extra-jars` in https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	// Experimental.
	ExtraJars *[]Code `field:"optional" json:"extraJars" yaml:"extraJars"`
	// Setting this value to true prioritizes the customer's extra JAR files in the classpath.
	// See: `--user-jars-first` in https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	// Experimental.
	ExtraJarsFirst *bool `field:"optional" json:"extraJarsFirst" yaml:"extraJarsFirst"`
	// Additional Python files that AWS Glue adds to the Python path before executing your script.
	// See: `--extra-py-files` in https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	// Experimental.
	ExtraPythonFiles *[]Code `field:"optional" json:"extraPythonFiles" yaml:"extraPythonFiles"`
	// The Python version to use.
	// Experimental.
	PythonVersion PythonVersion `field:"optional" json:"pythonVersion" yaml:"pythonVersion"`
}

