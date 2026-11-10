package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   claimMatchValueProperty := &ClaimMatchValueProperty{
//   	MatchValueString: jsii.String("matchValueString"),
//   	MatchValueStringList: []*string{
//   		jsii.String("matchValueStringList"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentmanager-claimmatchvalue.html
//
type CfnPaymentManagerPropsMixin_ClaimMatchValueProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentmanager-claimmatchvalue.html#cfn-bedrockagentcore-paymentmanager-claimmatchvalue-matchvaluestring
	//
	MatchValueString *string `field:"optional" json:"matchValueString" yaml:"matchValueString"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentmanager-claimmatchvalue.html#cfn-bedrockagentcore-paymentmanager-claimmatchvalue-matchvaluestringlist
	//
	MatchValueStringList *[]*string `field:"optional" json:"matchValueStringList" yaml:"matchValueStringList"`
}

