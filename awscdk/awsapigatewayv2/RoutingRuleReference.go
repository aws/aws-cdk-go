package awsapigatewayv2


// A reference to a RoutingRule resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   routingRuleReference := &RoutingRuleReference{
//   	RoutingRuleArn: jsii.String("routingRuleArn"),
//   }
//
type RoutingRuleReference struct {
	// The RoutingRuleArn of the RoutingRule resource.
	RoutingRuleArn *string `field:"required" json:"routingRuleArn" yaml:"routingRuleArn"`
}

