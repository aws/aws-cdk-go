package awsworkspacesweb


// The IP rules of the IP access settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ipRuleProperty := &IpRuleProperty{
//   	IpRange: jsii.String("ipRange"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-ipaccesssettings-iprule.html
//
type CfnIpAccessSettings_IpRuleProperty struct {
	// The IP range of the IP rule.
	//
	// This can either be a single IP address or a range using CIDR notation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-ipaccesssettings-iprule.html#cfn-workspacesweb-ipaccesssettings-iprule-iprange
	//
	IpRange *string `field:"required" json:"ipRange" yaml:"ipRange"`
	// The description of the IP rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-ipaccesssettings-iprule.html#cfn-workspacesweb-ipaccesssettings-iprule-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
}

