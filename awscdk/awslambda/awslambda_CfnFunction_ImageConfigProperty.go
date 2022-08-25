package awslambda


// Configuration values that override the container image Dockerfile settings.
//
// See [Container settings](https://docs.aws.amazon.com/lambda/latest/dg/images-create.html#images-parms) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   imageConfigProperty := &imageConfigProperty{
//   	command: []*string{
//   		jsii.String("command"),
//   	},
//   	entryPoint: []*string{
//   		jsii.String("entryPoint"),
//   	},
//   	workingDirectory: jsii.String("workingDirectory"),
//   }
//
type CfnFunction_ImageConfigProperty struct {
	// Specifies parameters that you want to pass in with ENTRYPOINT.
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// Specifies the entry point to their application, which is typically the location of the runtime executable.
	EntryPoint *[]*string `field:"optional" json:"entryPoint" yaml:"entryPoint"`
	// Specifies the working directory.
	WorkingDirectory *string `field:"optional" json:"workingDirectory" yaml:"workingDirectory"`
}

