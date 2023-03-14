package awsnetworkfirewall


// Settings that are available for use in the rules in the `RuleGroup` where this is defined.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ruleVariablesProperty := &ruleVariablesProperty{
//   	ipSets: map[string]interface{}{
//   		"ipSetsKey": map[string][]*string{
//   			"definition": []*string{
//   				jsii.String("definition"),
//   			},
//   		},
//   	},
//   	portSets: map[string]interface{}{
//   		"portSetsKey": &PortSetProperty{
//   			"definition": []*string{
//   				jsii.String("definition"),
//   			},
//   		},
//   	},
//   }
//
type CfnRuleGroup_RuleVariablesProperty struct {
	// A list of IP addresses and address ranges, in CIDR notation.
	IpSets interface{} `field:"optional" json:"ipSets" yaml:"ipSets"`
	// A list of port ranges.
	PortSets interface{} `field:"optional" json:"portSets" yaml:"portSets"`
}

