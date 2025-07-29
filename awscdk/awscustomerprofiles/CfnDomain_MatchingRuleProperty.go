package awscustomerprofiles


// Specifies how the rule-based matching process should match profiles.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   matchingRuleProperty := &MatchingRuleProperty{
//   	Rule: []*string{
//   		jsii.String("rule"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-domain-matchingrule.html
//
type CfnDomain_MatchingRuleProperty struct {
	// A single rule level of the `MatchRules` .
	//
	// Configures how the rule-based matching process should match profiles.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-domain-matchingrule.html#cfn-customerprofiles-domain-matchingrule-rule
	//
	Rule *[]*string `field:"required" json:"rule" yaml:"rule"`
}

