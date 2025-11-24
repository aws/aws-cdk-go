package mixinsawsnetworkfirewall


// A set of port ranges for use in the rules in a rule group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   portSetProperty := &PortSetProperty{
//   	Definition: []*string{
//   		jsii.String("definition"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-rulegroup-portset.html
//
type CfnRuleGroupPropsMixin_PortSetProperty struct {
	// The set of port ranges.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-rulegroup-portset.html#cfn-networkfirewall-rulegroup-portset-definition
	//
	Definition *[]*string `field:"optional" json:"definition" yaml:"definition"`
}

