package awslicensemanager


// AND rule statement.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   andRuleStatementProperty := &AndRuleStatementProperty{
//   	MatchingRuleStatements: []interface{}{
//   		&MatchingRuleStatementProperty{
//   			Constraint: jsii.String("constraint"),
//   			KeyToMatch: jsii.String("keyToMatch"),
//   			ValueToMatch: []*string{
//   				jsii.String("valueToMatch"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-licenseassetruleset-andrulestatement.html
//
type CfnLicenseAssetRuleSet_AndRuleStatementProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-licenseassetruleset-andrulestatement.html#cfn-licensemanager-licenseassetruleset-andrulestatement-matchingrulestatements
	//
	MatchingRuleStatements interface{} `field:"optional" json:"matchingRuleStatements" yaml:"matchingRuleStatements"`
}

