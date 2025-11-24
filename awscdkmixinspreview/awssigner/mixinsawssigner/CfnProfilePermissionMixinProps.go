package mixinsawssigner


// Properties for CfnProfilePermissionPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnProfilePermissionMixinProps := &CfnProfilePermissionMixinProps{
//   	Action: jsii.String("action"),
//   	Principal: jsii.String("principal"),
//   	ProfileName: jsii.String("profileName"),
//   	ProfileVersion: jsii.String("profileVersion"),
//   	StatementId: jsii.String("statementId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-signer-profilepermission.html
//
type CfnProfilePermissionMixinProps struct {
	// The AWS Signer action permitted as part of cross-account permissions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-signer-profilepermission.html#cfn-signer-profilepermission-action
	//
	Action *string `field:"optional" json:"action" yaml:"action"`
	// The AWS principal receiving cross-account permissions.
	//
	// This may be an IAM role or another AWS account ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-signer-profilepermission.html#cfn-signer-profilepermission-principal
	//
	Principal *string `field:"optional" json:"principal" yaml:"principal"`
	// The human-readable name of the signing profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-signer-profilepermission.html#cfn-signer-profilepermission-profilename
	//
	ProfileName *string `field:"optional" json:"profileName" yaml:"profileName"`
	// The version of the signing profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-signer-profilepermission.html#cfn-signer-profilepermission-profileversion
	//
	ProfileVersion *string `field:"optional" json:"profileVersion" yaml:"profileVersion"`
	// A unique identifier for the cross-account permission statement.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-signer-profilepermission.html#cfn-signer-profilepermission-statementid
	//
	StatementId *string `field:"optional" json:"statementId" yaml:"statementId"`
}

