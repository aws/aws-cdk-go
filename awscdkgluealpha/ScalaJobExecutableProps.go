package awscdkgluealpha


// Props for creating a Scala Spark (ETL or Streaming) job executable.
//
// Example:
//   var bucket bucket
//
//   glue.NewJob(this, jsii.String("ScalaSparkEtlJob"), &JobProps{
//   	Executable: glue.JobExecutable_ScalaEtl(&ScalaJobExecutableProps{
//   		GlueVersion: glue.GlueVersion_V4_0(),
//   		Script: glue.Code_FromBucket(bucket, jsii.String("src/com/example/HelloWorld.scala")),
//   		ClassName: jsii.String("com.example.HelloWorld"),
//   		ExtraJars: []code{
//   			glue.*code_*FromBucket(bucket, jsii.String("jars/HelloWorld.jar")),
//   		},
//   	}),
//   	WorkerType: glue.WorkerType_G_8X(),
//   	Description: jsii.String("an example Scala ETL job"),
//   })
//
// Experimental.
type ScalaJobExecutableProps struct {
	// The fully qualified Scala class name that serves as the entry point for the job.
	// See:  `--class` in https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	// Experimental.
	ClassName *string `field:"required" json:"className" yaml:"className"`
	// Glue version.
	// See: https://docs.aws.amazon.com/glue/latest/dg/release-notes.html
	//
	// Experimental.
	GlueVersion GlueVersion `field:"required" json:"glueVersion" yaml:"glueVersion"`
	// The script that executes a job.
	// Experimental.
	Script Code `field:"required" json:"script" yaml:"script"`
	// Additional files, such as configuration files that AWS Glue copies to the working directory of your script before executing it.
	//
	// Only individual files are supported, directories are not supported.
	// See:  `--extra-files` in https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	// Experimental.
	ExtraFiles *[]Code `field:"optional" json:"extraFiles" yaml:"extraFiles"`
	// Additional Java .jar files that AWS Glue adds to the Java classpath before executing your script. Only individual files are supported, directories are not supported.
	// See:  `--extra-jars` in https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	// Experimental.
	ExtraJars *[]Code `field:"optional" json:"extraJars" yaml:"extraJars"`
	// Setting this value to true prioritizes the customer's extra JAR files in the classpath.
	// See:  `--user-jars-first` in https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	// Experimental.
	ExtraJarsFirst *bool `field:"optional" json:"extraJarsFirst" yaml:"extraJarsFirst"`
}

