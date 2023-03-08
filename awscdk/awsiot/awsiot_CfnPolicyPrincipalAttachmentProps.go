package awsiot


// Properties for defining a `CfnPolicyPrincipalAttachment`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPolicyPrincipalAttachmentProps := &cfnPolicyPrincipalAttachmentProps{
//   	policyName: jsii.String("policyName"),
//   	principal: jsii.String("principal"),
//   }
//
type CfnPolicyPrincipalAttachmentProps struct {
	// The name of the AWS IoT policy.
	PolicyName *string `field:"required" json:"policyName" yaml:"policyName"`
	// The principal, which can be a certificate ARN (as returned from the `CreateCertificate` operation) or an Amazon Cognito ID.
	Principal *string `field:"required" json:"principal" yaml:"principal"`
}

