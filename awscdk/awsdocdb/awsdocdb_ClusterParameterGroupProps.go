package awsdocdb


// Properties for a cluster parameter group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clusterParameterGroupProps := &clusterParameterGroupProps{
//   	family: jsii.String("family"),
//   	parameters: map[string]*string{
//   		"parametersKey": jsii.String("parameters"),
//   	},
//
//   	// the properties below are optional
//   	dbClusterParameterGroupName: jsii.String("dbClusterParameterGroupName"),
//   	description: jsii.String("description"),
//   }
//
type ClusterParameterGroupProps struct {
	// Database family of this parameter group.
	Family *string `field:"required" json:"family" yaml:"family"`
	// The parameters in this parameter group.
	Parameters *map[string]*string `field:"required" json:"parameters" yaml:"parameters"`
	// The name of the cluster parameter group.
	DbClusterParameterGroupName *string `field:"optional" json:"dbClusterParameterGroupName" yaml:"dbClusterParameterGroupName"`
	// Description for this parameter group.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

