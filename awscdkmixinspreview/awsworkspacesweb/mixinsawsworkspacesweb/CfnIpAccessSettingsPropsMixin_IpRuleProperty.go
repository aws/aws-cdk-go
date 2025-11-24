package mixinsawsworkspacesweb


// The IP rules of the IP access settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   ipRuleProperty := &IpRuleProperty{
//   	Description: jsii.String("description"),
//   	IpRange: jsii.String("ipRange"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-ipaccesssettings-iprule.html
//
type CfnIpAccessSettingsPropsMixin_IpRuleProperty struct {
	// The description of the IP rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-ipaccesssettings-iprule.html#cfn-workspacesweb-ipaccesssettings-iprule-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The IP range of the IP rule.
	//
	// This can either be a single IP address or a range using CIDR notation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-ipaccesssettings-iprule.html#cfn-workspacesweb-ipaccesssettings-iprule-iprange
	//
	IpRange *string `field:"optional" json:"ipRange" yaml:"ipRange"`
}

