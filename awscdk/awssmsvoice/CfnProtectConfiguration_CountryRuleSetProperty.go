package awssmsvoice


// The set of `CountryRules` you specify to control which countries End User Messaging  can send your messages to.
//
// > If you don't specify all available ISO country codes in the `CountryRuleSet` for each number capability, the CloudFormation drift detection feature will detect drift. This is because End User Messaging  always returns all country codes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   countryRuleSetProperty := &CountryRuleSetProperty{
//   	Mms: []interface{}{
//   		&CountryRuleProperty{
//   			CountryCode: jsii.String("countryCode"),
//   			ProtectStatus: jsii.String("protectStatus"),
//   		},
//   	},
//   	Sms: []interface{}{
//   		&CountryRuleProperty{
//   			CountryCode: jsii.String("countryCode"),
//   			ProtectStatus: jsii.String("protectStatus"),
//   		},
//   	},
//   	Voice: []interface{}{
//   		&CountryRuleProperty{
//   			CountryCode: jsii.String("countryCode"),
//   			ProtectStatus: jsii.String("protectStatus"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-smsvoice-protectconfiguration-countryruleset.html
//
type CfnProtectConfiguration_CountryRuleSetProperty struct {
	// The set of `CountryRule` s to control which destination countries End User Messaging  can send your MMS messages to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-smsvoice-protectconfiguration-countryruleset.html#cfn-smsvoice-protectconfiguration-countryruleset-mms
	//
	Mms interface{} `field:"optional" json:"mms" yaml:"mms"`
	// The set of `CountryRule` s to control which destination countries End User Messaging  can send your SMS messages to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-smsvoice-protectconfiguration-countryruleset.html#cfn-smsvoice-protectconfiguration-countryruleset-sms
	//
	Sms interface{} `field:"optional" json:"sms" yaml:"sms"`
	// The set of `CountryRule` s to control which destination countries End User Messaging  can send your VOICE messages to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-smsvoice-protectconfiguration-countryruleset.html#cfn-smsvoice-protectconfiguration-countryruleset-voice
	//
	Voice interface{} `field:"optional" json:"voice" yaml:"voice"`
}

