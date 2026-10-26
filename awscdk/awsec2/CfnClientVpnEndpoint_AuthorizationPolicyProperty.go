package awsec2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   authorizationPolicyProperty := &AuthorizationPolicyProperty{
//   	PolicyDocument: jsii.String("policyDocument"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	ShadowMode: jsii.String("shadowMode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-clientvpnendpoint-authorizationpolicy.html
//
type CfnClientVpnEndpoint_AuthorizationPolicyProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-clientvpnendpoint-authorizationpolicy.html#cfn-ec2-clientvpnendpoint-authorizationpolicy-policydocument
	//
	PolicyDocument *string `field:"required" json:"policyDocument" yaml:"policyDocument"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-clientvpnendpoint-authorizationpolicy.html#cfn-ec2-clientvpnendpoint-authorizationpolicy-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-clientvpnendpoint-authorizationpolicy.html#cfn-ec2-clientvpnendpoint-authorizationpolicy-shadowmode
	//
	ShadowMode *string `field:"optional" json:"shadowMode" yaml:"shadowMode"`
}

