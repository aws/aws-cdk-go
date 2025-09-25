package awsredshift


// A reference to a ClusterParameterGroup resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clusterParameterGroupReference := &ClusterParameterGroupReference{
//   	ParameterGroupName: jsii.String("parameterGroupName"),
//   }
//
type ClusterParameterGroupReference struct {
	// The ParameterGroupName of the ClusterParameterGroup resource.
	ParameterGroupName *string `field:"required" json:"parameterGroupName" yaml:"parameterGroupName"`
}

