package awscdkgluealpha


// Code props for different {@link Code} assets used by different types of Spark jobs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import glue_alpha "github.com/aws/aws-cdk-go/awscdkgluealpha"
//
//   var code code
//
//   sparkExtraCodeProps := &SparkExtraCodeProps{
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
//   }
//
// Experimental.
type SparkExtraCodeProps struct {
	// Additional files, such as configuration files that AWS Glue copies to the working directory of your script before executing it.
	// See: https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	// Default: - no extra files specified.
	//
	// Experimental.
	ExtraFiles *[]Code `field:"optional" json:"extraFiles" yaml:"extraFiles"`
	// Extra Jars S3 URL (optional) S3 URL where additional jar dependencies are located.
	// Default: - no extra jar files.
	//
	// Experimental.
	ExtraJars *[]Code `field:"optional" json:"extraJars" yaml:"extraJars"`
	// Setting this value to true prioritizes the customer's extra JAR files in the classpath.
	// See:  `--user-jars-first` in https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	// Default: false - priority is not given to user-provided jars.
	//
	// Experimental.
	ExtraJarsFirst *bool `field:"optional" json:"extraJarsFirst" yaml:"extraJarsFirst"`
	// Extra Python Files S3 URL (optional) S3 URL where additional python dependencies are located.
	// Default: - no extra files.
	//
	// Experimental.
	ExtraPythonFiles *[]Code `field:"optional" json:"extraPythonFiles" yaml:"extraPythonFiles"`
}

