package previewawsmpamixins


// AWS IAM Identity Center credentials.
//
// For more information see, [AWS IAM Identity Center](https://docs.aws.amazon.com/identity-center/) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   iamIdentityCenterProperty := &IamIdentityCenterProperty{
//   	ApprovalPortalUrl: jsii.String("approvalPortalUrl"),
//   	InstanceArn: jsii.String("instanceArn"),
//   	Region: jsii.String("region"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mpa-identitysource-iamidentitycenter.html
//
type CfnIdentitySourcePropsMixin_IamIdentityCenterProperty struct {
	// URL for the approval portal associated with the IAM Identity Center instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mpa-identitysource-iamidentitycenter.html#cfn-mpa-identitysource-iamidentitycenter-approvalportalurl
	//
	ApprovalPortalUrl *string `field:"optional" json:"approvalPortalUrl" yaml:"approvalPortalUrl"`
	// Amazon Resource Name (ARN) for the IAM Identity Center instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mpa-identitysource-iamidentitycenter.html#cfn-mpa-identitysource-iamidentitycenter-instancearn
	//
	InstanceArn *string `field:"optional" json:"instanceArn" yaml:"instanceArn"`
	// AWS Region where the IAM Identity Center instance is located.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mpa-identitysource-iamidentitycenter.html#cfn-mpa-identitysource-iamidentitycenter-region
	//
	Region *string `field:"optional" json:"region" yaml:"region"`
}

