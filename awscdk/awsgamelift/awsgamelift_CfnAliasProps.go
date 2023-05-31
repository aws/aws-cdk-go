package awsgamelift


// Properties for defining a `CfnAlias`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAliasProps := &CfnAliasProps{
//   	Name: jsii.String("name"),
//   	RoutingStrategy: &RoutingStrategyProperty{
//   		Type: jsii.String("type"),
//
//   		// the properties below are optional
//   		FleetId: jsii.String("fleetId"),
//   		Message: jsii.String("message"),
//   	},
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   }
//
type CfnAliasProps struct {
	// A descriptive label that is associated with an alias.
	//
	// Alias names do not need to be unique.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The routing configuration, including routing type and fleet target, for the alias.
	RoutingStrategy interface{} `field:"required" json:"routingStrategy" yaml:"routingStrategy"`
	// A human-readable description of the alias.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

