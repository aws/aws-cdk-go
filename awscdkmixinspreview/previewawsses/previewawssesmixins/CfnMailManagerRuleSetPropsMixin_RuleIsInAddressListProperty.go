package previewawssesmixins


// The structure type for a boolean condition that provides the address lists and address list attribute to evaluate.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   ruleIsInAddressListProperty := &RuleIsInAddressListProperty{
//   	AddressLists: []*string{
//   		jsii.String("addressLists"),
//   	},
//   	Attribute: jsii.String("attribute"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-ruleisinaddresslist.html
//
type CfnMailManagerRuleSetPropsMixin_RuleIsInAddressListProperty struct {
	// The address lists that will be used for evaluation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-ruleisinaddresslist.html#cfn-ses-mailmanagerruleset-ruleisinaddresslist-addresslists
	//
	AddressLists *[]*string `field:"optional" json:"addressLists" yaml:"addressLists"`
	// The email attribute that needs to be evaluated against the address list.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-ruleisinaddresslist.html#cfn-ses-mailmanagerruleset-ruleisinaddresslist-attribute
	//
	Attribute *string `field:"optional" json:"attribute" yaml:"attribute"`
}

