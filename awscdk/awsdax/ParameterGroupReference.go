package awsdax


// A reference to a ParameterGroup resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   parameterGroupReference := &ParameterGroupReference{
//   	ParameterGroupId: jsii.String("parameterGroupId"),
//   }
//
type ParameterGroupReference struct {
	// The Id of the ParameterGroup resource.
	ParameterGroupId *string `field:"required" json:"parameterGroupId" yaml:"parameterGroupId"`
}

