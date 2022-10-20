package awsredshift

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnClusterParameterGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnClusterParameterGroupProps := &cfnClusterParameterGroupProps{
//   	description: jsii.String("description"),
//   	parameterGroupFamily: jsii.String("parameterGroupFamily"),
//
//   	// the properties below are optional
//   	parameters: []interface{}{
//   		&parameterProperty{
//   			parameterName: jsii.String("parameterName"),
//   			parameterValue: jsii.String("parameterValue"),
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
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

