package awsproton


// A reference to a EnvironmentAccountConnection resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   environmentAccountConnectionReference := &EnvironmentAccountConnectionReference{
//   	EnvironmentAccountConnectionArn: jsii.String("environmentAccountConnectionArn"),
//   }
//
type EnvironmentAccountConnectionReference struct {
	// The Arn of the EnvironmentAccountConnection resource.
	EnvironmentAccountConnectionArn *string `field:"required" json:"environmentAccountConnectionArn" yaml:"environmentAccountConnectionArn"`
}

