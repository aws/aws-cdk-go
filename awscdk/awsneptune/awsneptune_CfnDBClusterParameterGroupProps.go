package awsneptune

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
	// Provides the customer-specified description for this DB cluster parameter group.
	Description *string `field:"required" json:"description" yaml:"description"`
	// Must be `neptune1` .
	Family *string `field:"required" json:"family" yaml:"family"`
	// The parameters to set for this DB cluster parameter group.
	//
	// The parameters are expressed as a JSON object consisting of key-value pairs.
	//
	// If you update the parameters, some interruption may occur depending on which parameters you update.
	Parameters interface{} `field:"required" json:"parameters" yaml:"parameters"`
	// Provides the name of the DB cluster parameter group.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The tags that you want to attach to this parameter group.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

