package awslambda


// Configuration values that override the container image Dockerfile settings.
//
// For more information, see [Container image settings](https://docs.aws.amazon.com/lambda/latest/dg/images-create.html#images-parms) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   imageConfigProperty := &ImageConfigProperty{
//   	Command: []*string{
//   		jsii.String("command"),
//   	},
//   	EntryPoint: []*string{
//   		jsii.String("entryPoint"),
//   	},
//   	WorkingDirectory: jsii.String("workingDirectory"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-imageconfig.html
//
type CfnFunction_ImageConfigProperty struct {
	// Specifies parameters that you want to pass in with ENTRYPOINT.
	//
	// You can specify a maximum of 1,500 parameters in the list.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-imageconfig.html#cfn-lambda-function-imageconfig-command
	//
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// Specifies the entry point to their application, which is typically the location of the runtime executable.
	//
	// You can specify a maximum of 1,500 string entries in the list.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-imageconfig.html#cfn-lambda-function-imageconfig-entrypoint
	//
	EntryPoint *[]*string `field:"optional" json:"entryPoint" yaml:"entryPoint"`
	// Specifies the working directory.
	//
	// The length of the directory string cannot exceed 1,000 characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-imageconfig.html#cfn-lambda-function-imageconfig-workingdirectory
	//
	WorkingDirectory *string `field:"optional" json:"workingDirectory" yaml:"workingDirectory"`
}

