package awssigner


// Properties for defining a `CfnProfilePermission`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnProfilePermissionProps := &cfnProfilePermissionProps{
//   	action: jsii.String("action"),
//   	principal: jsii.String("principal"),
//   	profileName: jsii.String("profileName"),
//   	statementId: jsii.String("statementId"),
//
//   	// the properties below are optional
//   	profileVersion: jsii.String("profileVersion"),
//   }
//
type CfnProfilePermissionProps struct {
	// The AWS Signer action permitted as part of cross-account permissions.
	Action *string `field:"required" json:"action" yaml:"action"`
	// The AWS principal receiving cross-account permissions.
	//
	// This may be an IAM role or another AWS account ID.
	Principal *string `field:"required" json:"principal" yaml:"principal"`
	// The human-readable name of the signing profile.
	ProfileName *string `field:"required" json:"profileName" yaml:"profileName"`
	// A unique identifier for the cross-account permission statement.
	StatementId *string `field:"required" json:"statementId" yaml:"statementId"`
	// The version of the signing profile.
	ProfileVersion *string `field:"optional" json:"profileVersion" yaml:"profileVersion"`
}

