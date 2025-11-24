package mixinsawsiot


// Properties for CfnPolicyPrincipalAttachmentPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnPolicyPrincipalAttachmentMixinProps := &CfnPolicyPrincipalAttachmentMixinProps{
//   	PolicyName: jsii.String("policyName"),
//   	Principal: jsii.String("principal"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-policyprincipalattachment.html
//
type CfnPolicyPrincipalAttachmentMixinProps struct {
	// The name of the AWS IoT policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-policyprincipalattachment.html#cfn-iot-policyprincipalattachment-policyname
	//
	PolicyName *string `field:"optional" json:"policyName" yaml:"policyName"`
	// The principal, which can be a certificate ARN (as returned from the `CreateCertificate` operation) or an Amazon Cognito ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-policyprincipalattachment.html#cfn-iot-policyprincipalattachment-principal
	//
	Principal *string `field:"optional" json:"principal" yaml:"principal"`
}

