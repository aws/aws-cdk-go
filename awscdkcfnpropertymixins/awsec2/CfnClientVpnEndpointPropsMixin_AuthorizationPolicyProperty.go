package awsec2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   authorizationPolicyProperty := &AuthorizationPolicyProperty{
//   	Description: jsii.String("description"),
//   	PolicyDocument: jsii.String("policyDocument"),
//   	ShadowMode: jsii.String("shadowMode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-clientvpnendpoint-authorizationpolicy.html
//
type CfnClientVpnEndpointPropsMixin_AuthorizationPolicyProperty struct {
	// A brief description of the authorization policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-clientvpnendpoint-authorizationpolicy.html#cfn-ec2-clientvpnendpoint-authorizationpolicy-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The Cedar policy document.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-clientvpnendpoint-authorizationpolicy.html#cfn-ec2-clientvpnendpoint-authorizationpolicy-policydocument
	//
	PolicyDocument *string `field:"optional" json:"policyDocument" yaml:"policyDocument"`
	// Whether the authorization policy runs in shadow (observe-only) mode.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-clientvpnendpoint-authorizationpolicy.html#cfn-ec2-clientvpnendpoint-authorizationpolicy-shadowmode
	//
	ShadowMode *string `field:"optional" json:"shadowMode" yaml:"shadowMode"`
}

