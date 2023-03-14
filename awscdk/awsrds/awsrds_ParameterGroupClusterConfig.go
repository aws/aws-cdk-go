package awsrds


// The type returned from {@link IParameterGroup.bindToCluster}.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   parameterGroupClusterConfig := &parameterGroupClusterConfig{
//   	parameterGroupName: jsii.String("parameterGroupName"),
//   }
//
// Experimental.
type ParameterGroupClusterConfig struct {
	// The name of this parameter group.
	// Experimental.
	ParameterGroupName *string `field:"required" json:"parameterGroupName" yaml:"parameterGroupName"`
}

