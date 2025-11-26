package previewawsemrmixins


// A job flow step consisting of a JAR file whose main function will be executed.
//
// The main function submits a job for Hadoop to execute and waits for the job to finish or fail.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   hadoopJarStepConfigProperty := &HadoopJarStepConfigProperty{
//   	Args: []*string{
//   		jsii.String("args"),
//   	},
//   	Jar: jsii.String("jar"),
//   	MainClass: jsii.String("mainClass"),
//   	StepProperties: []interface{}{
//   		&KeyValueProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-step-hadoopjarstepconfig.html
//
type CfnStepPropsMixin_HadoopJarStepConfigProperty struct {
	// A list of command line arguments passed to the JAR file's main function when executed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-step-hadoopjarstepconfig.html#cfn-emr-step-hadoopjarstepconfig-args
	//
	Args *[]*string `field:"optional" json:"args" yaml:"args"`
	// A path to a JAR file run during the step.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-step-hadoopjarstepconfig.html#cfn-emr-step-hadoopjarstepconfig-jar
	//
	Jar *string `field:"optional" json:"jar" yaml:"jar"`
	// The name of the main class in the specified Java file.
	//
	// If not specified, the JAR file should specify a Main-Class in its manifest file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-step-hadoopjarstepconfig.html#cfn-emr-step-hadoopjarstepconfig-mainclass
	//
	MainClass *string `field:"optional" json:"mainClass" yaml:"mainClass"`
	// A list of Java properties that are set when the step runs.
	//
	// You can use these properties to pass key value pairs to your main function.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-step-hadoopjarstepconfig.html#cfn-emr-step-hadoopjarstepconfig-stepproperties
	//
	StepProperties interface{} `field:"optional" json:"stepProperties" yaml:"stepProperties"`
}

