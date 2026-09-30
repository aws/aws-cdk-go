package awsroute53


// Properties for CfnTrafficPolicyPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnTrafficPolicyMixinProps := &CfnTrafficPolicyMixinProps{
//   	Comment: jsii.String("comment"),
//   	Document: jsii.String("document"),
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-trafficpolicy.html
//
type CfnTrafficPolicyMixinProps struct {
	// Any comments to include about the traffic policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-trafficpolicy.html#cfn-route53-trafficpolicy-comment
	//
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// The definition of the traffic policy in JSON format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-trafficpolicy.html#cfn-route53-trafficpolicy-document
	//
	Document *string `field:"optional" json:"document" yaml:"document"`
	// The name of the traffic policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-trafficpolicy.html#cfn-route53-trafficpolicy-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

