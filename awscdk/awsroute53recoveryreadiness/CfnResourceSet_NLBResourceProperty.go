package awsroute53recoveryreadiness


// The Network Load Balancer resource that a DNS target resource points to.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   nLBResourceProperty := &NLBResourceProperty{
//   	Arn: jsii.String("arn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53recoveryreadiness-resourceset-nlbresource.html
//
type CfnResourceSet_NLBResourceProperty struct {
	// The Network Load Balancer resource Amazon Resource Name (ARN).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53recoveryreadiness-resourceset-nlbresource.html#cfn-route53recoveryreadiness-resourceset-nlbresource-arn
	//
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
}

