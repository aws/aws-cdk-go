package mixinsawscustomerprofiles


// Configures information about the `AttributeTypesSelector` which rule-based identity resolution uses to match profiles.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   attributeTypesSelectorProperty := &AttributeTypesSelectorProperty{
//   	Address: []*string{
//   		jsii.String("address"),
//   	},
//   	AttributeMatchingModel: jsii.String("attributeMatchingModel"),
//   	EmailAddress: []*string{
//   		jsii.String("emailAddress"),
//   	},
//   	PhoneNumber: []*string{
//   		jsii.String("phoneNumber"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-domain-attributetypesselector.html
//
type CfnDomainPropsMixin_AttributeTypesSelectorProperty struct {
	// The `Address` type.
	//
	// You can choose from `Address` , `BusinessAddress` , `MaillingAddress` , and `ShippingAddress` . You only can use the `Address` type in the `MatchingRule` . For example, if you want to match a profile based on `BusinessAddress.City` or `MaillingAddress.City` , you can choose the `BusinessAddress` and the `MaillingAddress` to represent the `Address` type and specify the `Address.City` on the matching rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-domain-attributetypesselector.html#cfn-customerprofiles-domain-attributetypesselector-address
	//
	Address *[]*string `field:"optional" json:"address" yaml:"address"`
	// Configures the `AttributeMatchingModel` , you can either choose `ONE_TO_ONE` or `MANY_TO_MANY` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-domain-attributetypesselector.html#cfn-customerprofiles-domain-attributetypesselector-attributematchingmodel
	//
	AttributeMatchingModel *string `field:"optional" json:"attributeMatchingModel" yaml:"attributeMatchingModel"`
	// The Email type.
	//
	// You can choose from `EmailAddress` , `BusinessEmailAddress` and `PersonalEmailAddress` . You only can use the `EmailAddress` type in the `MatchingRule` . For example, if you want to match profile based on `PersonalEmailAddress` or `BusinessEmailAddress` , you can choose the `PersonalEmailAddress` and the `BusinessEmailAddress` to represent the `EmailAddress` type and only specify the `EmailAddress` on the matching rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-domain-attributetypesselector.html#cfn-customerprofiles-domain-attributetypesselector-emailaddress
	//
	EmailAddress *[]*string `field:"optional" json:"emailAddress" yaml:"emailAddress"`
	// The `PhoneNumber` type.
	//
	// You can choose from `PhoneNumber` , `HomePhoneNumber` , and `MobilePhoneNumber` . You only can use the `PhoneNumber` type in the `MatchingRule` . For example, if you want to match a profile based on `Phone` or `HomePhone` , you can choose the `Phone` and the `HomePhone` to represent the `PhoneNumber` type and only specify the `PhoneNumber` on the matching rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-domain-attributetypesselector.html#cfn-customerprofiles-domain-attributetypesselector-phonenumber
	//
	PhoneNumber *[]*string `field:"optional" json:"phoneNumber" yaml:"phoneNumber"`
}

