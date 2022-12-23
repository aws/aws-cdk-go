package awsbatch


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eksContainerEnvironmentVariableProperty := &eksContainerEnvironmentVariableProperty{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	value: jsii.String("value"),
//   }
//
type CfnJobDefinition_EksContainerEnvironmentVariableProperty struct {
	// `CfnJobDefinition.EksContainerEnvironmentVariableProperty.Name`.
	Name *string `field:"required" json:"name" yaml:"name"`
	// `CfnJobDefinition.EksContainerEnvironmentVariableProperty.Value`.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

