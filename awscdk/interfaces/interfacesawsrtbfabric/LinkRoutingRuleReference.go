package interfacesawsrtbfabric


// A reference to a LinkRoutingRule resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   linkRoutingRuleReference := &LinkRoutingRuleReference{
//   	LinkRoutingRuleArn: jsii.String("linkRoutingRuleArn"),
//   }
//
type LinkRoutingRuleReference struct {
	// The Arn of the LinkRoutingRule resource.
	LinkRoutingRuleArn *string `field:"required" json:"linkRoutingRuleArn" yaml:"linkRoutingRuleArn"`
}

