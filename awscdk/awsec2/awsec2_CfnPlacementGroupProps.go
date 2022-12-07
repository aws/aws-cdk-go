package awsec2


// Properties for defining a `CfnPlacementGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPlacementGroupProps := &cfnPlacementGroupProps{
//   	spreadLevel: jsii.String("spreadLevel"),
//   	strategy: jsii.String("strategy"),
//   }
//
type CfnPlacementGroupProps struct {
	// The spread level for the placement group.
	//
	// *Only* Outpost placement groups can be spread across hosts.
	SpreadLevel *string `field:"optional" json:"spreadLevel" yaml:"spreadLevel"`
	// The placement strategy.
	Strategy *string `field:"optional" json:"strategy" yaml:"strategy"`
}

