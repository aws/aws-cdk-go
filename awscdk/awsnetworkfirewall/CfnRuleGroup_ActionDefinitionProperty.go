package awsnetworkfirewall


// A custom action to use in stateless rule actions settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   actionDefinitionProperty := &ActionDefinitionProperty{
//   	PublishMetricAction: &PublishMetricActionProperty{
//   		Dimensions: []interface{}{
//   			&DimensionProperty{
//   				Value: jsii.String("value"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-rulegroup-actiondefinition.html
//
type CfnRuleGroup_ActionDefinitionProperty struct {
	// Stateless inspection criteria that publishes the specified metrics to Amazon CloudWatch for the matching packet.
	//
	// This setting defines a CloudWatch dimension value to be published.
	//
	// You can pair this custom action with any of the standard stateless rule actions. For example, you could pair this in a rule action with the standard action that forwards the packet for stateful inspection. Then, when a packet matches the rule, Network Firewall publishes metrics for the packet and forwards it.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-rulegroup-actiondefinition.html#cfn-networkfirewall-rulegroup-actiondefinition-publishmetricaction
	//
	PublishMetricAction interface{} `field:"optional" json:"publishMetricAction" yaml:"publishMetricAction"`
}

