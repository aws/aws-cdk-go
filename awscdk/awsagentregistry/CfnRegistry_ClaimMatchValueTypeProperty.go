package awsagentregistry


// The expected value used to match a claim.
//
// Exactly one member is set.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   claimMatchValueTypeProperty := &ClaimMatchValueTypeProperty{
//   	MatchValueString: jsii.String("matchValueString"),
//   	MatchValueStringList: []*string{
//   		jsii.String("matchValueStringList"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-agentregistry-registry-claimmatchvaluetype.html
//
type CfnRegistry_ClaimMatchValueTypeProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-agentregistry-registry-claimmatchvaluetype.html#cfn-agentregistry-registry-claimmatchvaluetype-matchvaluestring
	//
	MatchValueString *string `field:"optional" json:"matchValueString" yaml:"matchValueString"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-agentregistry-registry-claimmatchvaluetype.html#cfn-agentregistry-registry-claimmatchvaluetype-matchvaluestringlist
	//
	MatchValueStringList *[]*string `field:"optional" json:"matchValueStringList" yaml:"matchValueStringList"`
}

