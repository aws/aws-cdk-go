package awsneptune

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnDBParameterGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var parameters interface{}
//
//   cfnDBParameterGroupProps := &CfnDBParameterGroupProps{
//   	Description: jsii.String("description"),
//   	Family: jsii.String("family"),
//   	Parameters: parameters,
//
//   	// the properties below are optional
//   	Name: jsii.String("name"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnDBParameterGroupProps struct {
	// Provides the customer-specified description for this DB parameter group.
	Description *string `field:"required" json:"description" yaml:"description"`
	// Must be `neptune1` for engine versions prior to [1.2.0.0](https://docs.aws.amazon.com/neptune/latest/userguide/engine-releases-1.2.0.0.html) , or `neptune1.2` for engine version `1.2.0.0` and higher.
	Family *string `field:"required" json:"family" yaml:"family"`
	// The parameters to set for this DB parameter group.
	//
	// The parameters are expressed as a JSON object consisting of key-value pairs.
	//
	// Changes to dynamic parameters are applied immediately. During an update, if you have static parameters (whether they were changed or not), it triggers AWS CloudFormation to reboot the associated DB instance without failover.
	Parameters interface{} `field:"required" json:"parameters" yaml:"parameters"`
	// Provides the name of the DB parameter group.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The tags that you want to attach to this parameter group.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

