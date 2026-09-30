package awsroute53


// Properties for defining a `CfnTrafficPolicy`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTrafficPolicyProps := &CfnTrafficPolicyProps{
//   	Document: jsii.String("document"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Comment: jsii.String("comment"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-trafficpolicy.html
//
type CfnTrafficPolicyProps struct {
	// The definition of the traffic policy in JSON format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-trafficpolicy.html#cfn-route53-trafficpolicy-document
	//
	Document *string `field:"required" json:"document" yaml:"document"`
	// The name of the traffic policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-trafficpolicy.html#cfn-route53-trafficpolicy-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// Any comments to include about the traffic policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-trafficpolicy.html#cfn-route53-trafficpolicy-comment
	//
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
}

