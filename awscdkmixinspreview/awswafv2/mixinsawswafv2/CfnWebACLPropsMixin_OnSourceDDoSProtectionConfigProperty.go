package mixinsawswafv2


// Configures the level of DDoS protection that applies to web ACLs associated with Application Load Balancers.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   onSourceDDoSProtectionConfigProperty := &OnSourceDDoSProtectionConfigProperty{
//   	AlbLowReputationMode: jsii.String("albLowReputationMode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-onsourceddosprotectionconfig.html
//
type CfnWebACLPropsMixin_OnSourceDDoSProtectionConfigProperty struct {
	// The level of DDoS protection that applies to web ACLs associated with Application Load Balancers.
	//
	// `ACTIVE_UNDER_DDOS` protection is enabled by default whenever a web ACL is associated with an Application Load Balancer. In the event that an Application Load Balancer experiences high-load conditions or suspected DDoS attacks, the `ACTIVE_UNDER_DDOS` protection automatically rate limits traffic from known low reputation sources without disrupting Application Load Balancer availability. `ALWAYS_ON` protection provides constant, always-on monitoring of known low reputation sources for suspected DDoS attacks. While this provides a higher level of protection, there may be potential impacts on legitimate traffic.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-onsourceddosprotectionconfig.html#cfn-wafv2-webacl-onsourceddosprotectionconfig-alblowreputationmode
	//
	AlbLowReputationMode *string `field:"optional" json:"albLowReputationMode" yaml:"albLowReputationMode"`
}

