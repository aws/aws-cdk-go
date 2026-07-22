package awsroute53resolver


// Configuration for a partner threat protection rule type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   partnerThreatProtectionConfigProperty := &PartnerThreatProtectionConfigProperty{
//   	Partner: jsii.String("partner"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53resolver-firewallrulegroup-partnerthreatprotectionconfig.html
//
type CfnFirewallRuleGroup_PartnerThreatProtectionConfigProperty struct {
	// The partner identifier value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53resolver-firewallrulegroup-partnerthreatprotectionconfig.html#cfn-route53resolver-firewallrulegroup-partnerthreatprotectionconfig-partner
	//
	Partner *string `field:"required" json:"partner" yaml:"partner"`
}

