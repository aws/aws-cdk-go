package previewawsshieldmixins


// The automatic application layer DDoS mitigation settings for a `Protection` .
//
// This configuration determines whether Shield Advanced automatically manages rules in the web ACL in order to respond to application layer events that Shield Advanced determines to be DDoS attacks.
//
// If you use CloudFormation to manage the web ACLs that you use with Shield Advanced automatic mitigation, see the guidance for the `AWS::WAFv2::WebACL` resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var block interface{}
//   var count interface{}
//
//   applicationLayerAutomaticResponseConfigurationProperty := &ApplicationLayerAutomaticResponseConfigurationProperty{
//   	Action: &ActionProperty{
//   		Block: block,
//   		Count: count,
//   	},
//   	Status: jsii.String("status"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-shield-protection-applicationlayerautomaticresponseconfiguration.html
//
type CfnProtectionPropsMixin_ApplicationLayerAutomaticResponseConfigurationProperty struct {
	// Specifies the action setting that Shield Advanced should use in the AWS WAF rules that it creates on behalf of the protected resource in response to DDoS attacks.
	//
	// You specify this as part of the configuration for the automatic application layer DDoS mitigation feature, when you enable or update automatic mitigation. Shield Advanced creates the AWS WAF rules in a Shield Advanced-managed rule group, inside the web ACL that you have associated with the resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-shield-protection-applicationlayerautomaticresponseconfiguration.html#cfn-shield-protection-applicationlayerautomaticresponseconfiguration-action
	//
	Action interface{} `field:"optional" json:"action" yaml:"action"`
	// Indicates whether automatic application layer DDoS mitigation is enabled for the protection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-shield-protection-applicationlayerautomaticresponseconfiguration.html#cfn-shield-protection-applicationlayerautomaticresponseconfiguration-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
}

