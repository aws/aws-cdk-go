package awswafv2


// Configures the options for on-source DDoS protection provided by supported resource type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   onSourceDDoSProtectionConfigProperty := &OnSourceDDoSProtectionConfigProperty{
//   	AlbLowReputationMode: jsii.String("albLowReputationMode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-onsourceddosprotectionconfig.html
//
type CfnWebACL_OnSourceDDoSProtectionConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-onsourceddosprotectionconfig.html#cfn-wafv2-webacl-onsourceddosprotectionconfig-alblowreputationmode
	//
	AlbLowReputationMode *string `field:"required" json:"albLowReputationMode" yaml:"albLowReputationMode"`
}

