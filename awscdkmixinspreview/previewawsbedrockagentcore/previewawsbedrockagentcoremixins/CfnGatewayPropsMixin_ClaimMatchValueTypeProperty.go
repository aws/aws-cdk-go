package previewawsbedrockagentcoremixins


// The value or values in the custom claim to match for.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   claimMatchValueTypeProperty := &ClaimMatchValueTypeProperty{
//   	MatchValueString: jsii.String("matchValueString"),
//   	MatchValueStringList: []*string{
//   		jsii.String("matchValueStringList"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gateway-claimmatchvaluetype.html
//
type CfnGatewayPropsMixin_ClaimMatchValueTypeProperty struct {
	// The string value to match for.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gateway-claimmatchvaluetype.html#cfn-bedrockagentcore-gateway-claimmatchvaluetype-matchvaluestring
	//
	MatchValueString *string `field:"optional" json:"matchValueString" yaml:"matchValueString"`
	// The list of strings to check for a match.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gateway-claimmatchvaluetype.html#cfn-bedrockagentcore-gateway-claimmatchvaluetype-matchvaluestringlist
	//
	MatchValueStringList *[]*string `field:"optional" json:"matchValueStringList" yaml:"matchValueStringList"`
}

