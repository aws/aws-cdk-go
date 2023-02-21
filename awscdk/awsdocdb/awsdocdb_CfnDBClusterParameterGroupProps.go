package awsdocdb

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnDBClusterParameterGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var parameters interface{}
//
//   cfnDBClusterParameterGroupProps := &cfnDBClusterParameterGroupProps{
//   	description: jsii.String("description"),
//   	family: jsii.String("family"),
//   	parameters: parameters,
//
//   	// the properties below are optional
//   	name: jsii.String("name"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnDBClusterParameterGroupProps struct {
	// The description for the cluster parameter group.
	Description *string `field:"required" json:"description" yaml:"description"`
	// The cluster parameter group family name.
	Family *string `field:"required" json:"family" yaml:"family"`
	// Provides a list of parameters for the cluster parameter group.
	Parameters interface{} `field:"required" json:"parameters" yaml:"parameters"`
	// The name of the DB cluster parameter group.
	//
	// Constraints:
	//
	// - Must not match the name of an existing `DBClusterParameterGroup` .
	//
	// > This value is stored as a lowercase string.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The tags to be assigned to the cluster parameter group.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

