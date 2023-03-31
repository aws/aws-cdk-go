package awssam


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
	// `CfnFunction.ImageConfigProperty.Command`.
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// `CfnFunction.ImageConfigProperty.EntryPoint`.
	EntryPoint *[]*string `field:"optional" json:"entryPoint" yaml:"entryPoint"`
	// `CfnFunction.ImageConfigProperty.WorkingDirectory`.
	WorkingDirectory *string `field:"optional" json:"workingDirectory" yaml:"workingDirectory"`
}

