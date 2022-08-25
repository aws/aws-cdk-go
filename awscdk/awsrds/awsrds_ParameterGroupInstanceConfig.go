package awsrds


// The type returned from {@link IParameterGroup.bindToInstance}.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   parameterGroupInstanceConfig := &parameterGroupInstanceConfig{
//   	parameterGroupName: jsii.String("parameterGroupName"),
//   }
//
type ParameterGroupInstanceConfig struct {
	// The name of this parameter group.
	ParameterGroupName *string `field:"required" json:"parameterGroupName" yaml:"parameterGroupName"`
}

