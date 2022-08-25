package awsappstream


// The details of the script.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scriptDetailsProperty := &scriptDetailsProperty{
//   	executablePath: jsii.String("executablePath"),
//   	scriptS3Location: &s3LocationProperty{
//   		s3Bucket: jsii.String("s3Bucket"),
//   		s3Key: jsii.String("s3Key"),
//   	},
//   	timeoutInSeconds: jsii.Number(123),
//
//   	// the properties below are optional
//   	executableParameters: jsii.String("executableParameters"),
//   }
//
type CfnAppBlock_ScriptDetailsProperty struct {
	// The run path for the script.
	ExecutablePath *string `field:"required" json:"executablePath" yaml:"executablePath"`
	// The S3 object location of the script.
	ScriptS3Location interface{} `field:"required" json:"scriptS3Location" yaml:"scriptS3Location"`
	// The run timeout, in seconds, for the script.
	TimeoutInSeconds *float64 `field:"required" json:"timeoutInSeconds" yaml:"timeoutInSeconds"`
	// The parameters used in the run path for the script.
	ExecutableParameters *string `field:"optional" json:"executableParameters" yaml:"executableParameters"`
}

