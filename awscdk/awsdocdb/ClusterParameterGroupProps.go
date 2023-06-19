package awsdocdb


// Properties for a cluster parameter group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clusterParameterGroupProps := &ClusterParameterGroupProps{
//   	Family: jsii.String("family"),
//   	Parameters: map[string]*string{
//   		"parametersKey": jsii.String("parameters"),
//   	},
//
//   	// the properties below are optional
//   	DbClusterParameterGroupName: jsii.String("dbClusterParameterGroupName"),
//   	Description: jsii.String("description"),
//   }
//
// Experimental.
type ClusterParameterGroupProps struct {
	// Database family of this parameter group.
	// Experimental.
	Family *string `field:"required" json:"family" yaml:"family"`
	// The parameters in this parameter group.
	// Experimental.
	Parameters *map[string]*string `field:"required" json:"parameters" yaml:"parameters"`
	// The name of the cluster parameter group.
	// Experimental.
	DbClusterParameterGroupName *string `field:"optional" json:"dbClusterParameterGroupName" yaml:"dbClusterParameterGroupName"`
	// Description for this parameter group.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

