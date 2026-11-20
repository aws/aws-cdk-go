package awsbedrockagentcore


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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentmanager-claimmatchvaluetype.html
//
type CfnPaymentManager_ClaimMatchValueTypeProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentmanager-claimmatchvaluetype.html#cfn-bedrockagentcore-paymentmanager-claimmatchvaluetype-matchvaluestring
	//
	MatchValueString *string `field:"optional" json:"matchValueString" yaml:"matchValueString"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentmanager-claimmatchvaluetype.html#cfn-bedrockagentcore-paymentmanager-claimmatchvaluetype-matchvaluestringlist
	//
	MatchValueStringList *[]*string `field:"optional" json:"matchValueStringList" yaml:"matchValueStringList"`
}

