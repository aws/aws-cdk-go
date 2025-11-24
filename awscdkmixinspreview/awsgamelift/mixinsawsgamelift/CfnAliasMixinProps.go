package mixinsawsgamelift

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnAliasPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAliasMixinProps := &CfnAliasMixinProps{
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	RoutingStrategy: &RoutingStrategyProperty{
//   		FleetId: jsii.String("fleetId"),
//   		Message: jsii.String("message"),
//   		Type: jsii.String("type"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-alias.html
//
type CfnAliasMixinProps struct {
	// A human-readable description of the alias.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-alias.html#cfn-gamelift-alias-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A descriptive label that is associated with an alias.
	//
	// Alias names do not need to be unique.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-alias.html#cfn-gamelift-alias-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The routing configuration, including routing type and fleet target, for the alias.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-alias.html#cfn-gamelift-alias-routingstrategy
	//
	RoutingStrategy interface{} `field:"optional" json:"routingStrategy" yaml:"routingStrategy"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-alias.html#cfn-gamelift-alias-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

