// The CDK Construct Library for AWS::Neptune
package awscdkneptunealpha


// Marker class for cluster parameter group.
//
// Example:
//   clusterParams := neptune.NewClusterParameterGroup(this, jsii.String("ClusterParams"), &clusterParameterGroupProps{
//   	description: jsii.String("Cluster parameter group"),
//   	parameters: map[string]*string{
//   		"neptune_enable_audit_log": jsii.String("1"),
//   	},
//   })
//
//   dbParams := neptune.NewParameterGroup(this, jsii.String("DbParams"), &parameterGroupProps{
//   	description: jsii.String("Db parameter group"),
//   	parameters: map[string]*string{
//   		"neptune_query_timeout": jsii.String("120000"),
//   	},
//   })
//
//   cluster := neptune.NewDatabaseCluster(this, jsii.String("Database"), &databaseClusterProps{
//   	vpc: vpc,
//   	instanceType: neptune.instanceType_R5_LARGE(),
//   	clusterParameterGroup: clusterParams,
//   	parameterGroup: dbParams,
//   })
//
// Experimental.
type ParameterGroupProps struct {
	// The parameters in this parameter group.
	// Experimental.
	Parameters *map[string]*string `field:"required" json:"parameters" yaml:"parameters"`
	// Description for this parameter group.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Parameter group family.
	// Experimental.
	Family ParameterGroupFamily `field:"optional" json:"family" yaml:"family"`
	// The name of the parameter group.
	// Experimental.
	ParameterGroupName *string `field:"optional" json:"parameterGroupName" yaml:"parameterGroupName"`
}

