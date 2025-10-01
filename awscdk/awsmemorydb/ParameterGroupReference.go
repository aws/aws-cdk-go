package awsmemorydb


// A reference to a ParameterGroup resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   parameterGroupReference := &ParameterGroupReference{
//   	ParameterGroupName: jsii.String("parameterGroupName"),
//   }
//
type ParameterGroupReference struct {
	// The ParameterGroupName of the ParameterGroup resource.
	ParameterGroupName *string `field:"required" json:"parameterGroupName" yaml:"parameterGroupName"`
}

