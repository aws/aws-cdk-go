package awsappstream


// The details of the script.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scriptDetailsProperty := &ScriptDetailsProperty{
//   	ExecutablePath: jsii.String("executablePath"),
//   	ScriptS3Location: &S3LocationProperty{
//   		S3Bucket: jsii.String("s3Bucket"),
//
//   		// the properties below are optional
//   		S3Key: jsii.String("s3Key"),
//   	},
//   	TimeoutInSeconds: jsii.Number(123),
//
//   	// the properties below are optional
//   	ExecutableParameters: jsii.String("executableParameters"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-appblock-scriptdetails.html
//
type CfnAppBlock_ScriptDetailsProperty struct {
	// The run path for the script.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-appblock-scriptdetails.html#cfn-appstream-appblock-scriptdetails-executablepath
	//
	ExecutablePath *string `field:"required" json:"executablePath" yaml:"executablePath"`
	// The S3 object location of the script.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-appblock-scriptdetails.html#cfn-appstream-appblock-scriptdetails-scripts3location
	//
	ScriptS3Location interface{} `field:"required" json:"scriptS3Location" yaml:"scriptS3Location"`
	// The run timeout, in seconds, for the script.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-appblock-scriptdetails.html#cfn-appstream-appblock-scriptdetails-timeoutinseconds
	//
	TimeoutInSeconds *float64 `field:"required" json:"timeoutInSeconds" yaml:"timeoutInSeconds"`
	// The parameters used in the run path for the script.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-appblock-scriptdetails.html#cfn-appstream-appblock-scriptdetails-executableparameters
	//
	ExecutableParameters *string `field:"optional" json:"executableParameters" yaml:"executableParameters"`
}

