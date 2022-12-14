package awspipes


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ecsEnvironmentVariableProperty := &ecsEnvironmentVariableProperty{
//   	name: jsii.String("name"),
//   	value: jsii.String("value"),
//   }
//
type CfnPipe_EcsEnvironmentVariableProperty struct {
	// `CfnPipe.EcsEnvironmentVariableProperty.Name`.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// `CfnPipe.EcsEnvironmentVariableProperty.Value`.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

