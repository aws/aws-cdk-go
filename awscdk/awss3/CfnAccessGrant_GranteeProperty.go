package awss3


// The user, group, or role to which you are granting access.
//
// You can grant access to an IAM user or role. If you have added your corporate directory to AWS IAM Identity Center and associated your Identity Center instance with your S3 Access Grants instance, the grantee can also be a corporate directory user or group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   granteeProperty := &GranteeProperty{
//   	GranteeIdentifier: jsii.String("granteeIdentifier"),
//   	GranteeType: jsii.String("granteeType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-accessgrant-grantee.html
//
type CfnAccessGrant_GranteeProperty struct {
	// The unique identifier of the `Grantee` .
	//
	// If the grantee type is `IAM` , the identifier is the IAM Amazon Resource Name (ARN) of the user or role. If the grantee type is a directory user or group, the identifier is 128-bit universally unique identifier (UUID) in the format `a1b2c3d4-5678-90ab-cdef-EXAMPLE11111` . You can obtain this UUID from your AWS IAM Identity Center instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-accessgrant-grantee.html#cfn-s3-accessgrant-grantee-granteeidentifier
	//
	GranteeIdentifier *string `field:"required" json:"granteeIdentifier" yaml:"granteeIdentifier"`
	// The type of the grantee to which access has been granted. It can be one of the following values:.
	//
	// - `IAM` - An IAM user or role.
	// - `DIRECTORY_USER` - Your corporate directory user. You can use this option if you have added your corporate identity directory to IAM Identity Center and associated the IAM Identity Center instance with your S3 Access Grants instance.
	// - `DIRECTORY_GROUP` - Your corporate directory group. You can use this option if you have added your corporate identity directory to IAM Identity Center and associated the IAM Identity Center instance with your S3 Access Grants instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-accessgrant-grantee.html#cfn-s3-accessgrant-grantee-granteetype
	//
	GranteeType *string `field:"required" json:"granteeType" yaml:"granteeType"`
}

