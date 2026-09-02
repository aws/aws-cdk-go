package awslicensemanager


// OR rule statement.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   orRuleStatementProperty := &OrRuleStatementProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-licenseassetruleset-orrulestatement.html
//
type CfnLicenseAssetRuleSet_OrRuleStatementProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-licenseassetruleset-orrulestatement.html#cfn-licensemanager-licenseassetruleset-orrulestatement-matchingrulestatements
	//
	MatchingRuleStatements interface{} `field:"optional" json:"matchingRuleStatements" yaml:"matchingRuleStatements"`
}

