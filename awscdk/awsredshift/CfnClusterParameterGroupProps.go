package awsredshift

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnClusterParameterGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnClusterParameterGroupProps := &CfnClusterParameterGroupProps{
//   	Description: jsii.String("description"),
//   	ParameterGroupFamily: jsii.String("parameterGroupFamily"),
//
//   	// the properties below are optional
//   	Parameters: []interface{}{
//   		&ParameterProperty{
//   			ParameterName: jsii.String("parameterName"),
//   			ParameterValue: jsii.String("parameterValue"),
//   		},
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnClusterParameterGroupProps struct {
	// The description of the parameter group.
	Description *string `field:"required" json:"description" yaml:"description"`
	// The name of the cluster parameter group family that this cluster parameter group is compatible with.
	ParameterGroupFamily *string `field:"required" json:"parameterGroupFamily" yaml:"parameterGroupFamily"`
	// An array of parameters to be modified. A maximum of 20 parameters can be modified in a single request.
	//
	// For each parameter to be modified, you must supply at least the parameter name and parameter value; other name-value pairs of the parameter are optional.
	//
	// For the workload management (WLM) configuration, you must supply all the name-value pairs in the wlm_json_configuration parameter.
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// The list of tags for the cluster parameter group.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

