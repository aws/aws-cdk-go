package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   iamPrincipalProperty := &IamPrincipalProperty{
//   	Arn: jsii.String("arn"),
//
//   	// the properties below are optional
//   	Operator: jsii.String("operator"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayrule-iamprincipal.html
//
type CfnGatewayRule_IamPrincipalProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayrule-iamprincipal.html#cfn-bedrockagentcore-gatewayrule-iamprincipal-arn
	//
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayrule-iamprincipal.html#cfn-bedrockagentcore-gatewayrule-iamprincipal-operator
	//
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
}

