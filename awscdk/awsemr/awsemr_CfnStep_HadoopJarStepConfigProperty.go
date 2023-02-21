package awsemr


// A job flow step consisting of a JAR file whose main function will be executed.
//
// The main function submits a job for Hadoop to execute and waits for the job to finish or fail.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hadoopJarStepConfigProperty := &hadoopJarStepConfigProperty{
//   	jar: jsii.String("jar"),
//
//   	// the properties below are optional
//   	args: []*string{
//   		jsii.String("args"),
//   	},
//   	mainClass: jsii.String("mainClass"),
//   	stepProperties: []interface{}{
//   		&keyValueProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnStep_HadoopJarStepConfigProperty struct {
	// A path to a JAR file run during the step.
	Jar *string `field:"required" json:"jar" yaml:"jar"`
	// A list of command line arguments passed to the JAR file's main function when executed.
	Args *[]*string `field:"optional" json:"args" yaml:"args"`
	// The name of the main class in the specified Java file.
	//
	// If not specified, the JAR file should specify a Main-Class in its manifest file.
	MainClass *string `field:"optional" json:"mainClass" yaml:"mainClass"`
	// A list of Java properties that are set when the step runs.
	//
	// You can use these properties to pass key value pairs to your main function.
	StepProperties interface{} `field:"optional" json:"stepProperties" yaml:"stepProperties"`
}

