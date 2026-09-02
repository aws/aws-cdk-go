package awsglue


// The SessionCommand that runs the job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sessionCommandProperty := &SessionCommandProperty{
//   	Name: jsii.String("name"),
//   	PythonVersion: jsii.String("pythonVersion"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-session-sessioncommand.html
//
type CfnSession_SessionCommandProperty struct {
	// Specifies the name of the SessionCommand.
	//
	// Can be 'glueetl' or 'gluestreaming'.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-session-sessioncommand.html#cfn-glue-session-sessioncommand-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Specifies the Python version.
	//
	// The Python version indicates the version supported for jobs of type Spark.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-session-sessioncommand.html#cfn-glue-session-sessioncommand-pythonversion
	//
	PythonVersion *string `field:"optional" json:"pythonVersion" yaml:"pythonVersion"`
}

