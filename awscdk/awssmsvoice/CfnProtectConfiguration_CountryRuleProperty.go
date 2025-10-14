package awssmsvoice


// Specifies the type of protection to use for a country.
//
// For example, to set Canada as allowed, the `CountryRule` would be formatted as follows:
//
// `{ "CountryCode": "CA", "ProtectStatus": "ALLOW" }`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   countryRuleProperty := &CountryRuleProperty{
//   	CountryCode: jsii.String("countryCode"),
//   	ProtectStatus: jsii.String("protectStatus"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-smsvoice-protectconfiguration-countryrule.html
//
type CfnProtectConfiguration_CountryRuleProperty struct {
	// The two-character code, in ISO 3166-1 alpha-2 format, for the country or region.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-smsvoice-protectconfiguration-countryrule.html#cfn-smsvoice-protectconfiguration-countryrule-countrycode
	//
	CountryCode *string `field:"required" json:"countryCode" yaml:"countryCode"`
	// The types of protection that can be used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-smsvoice-protectconfiguration-countryrule.html#cfn-smsvoice-protectconfiguration-countryrule-protectstatus
	//
	ProtectStatus *string `field:"required" json:"protectStatus" yaml:"protectStatus"`
}

