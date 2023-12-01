package awss3


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
	// The unique identifier of the Grantee.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-accessgrant-grantee.html#cfn-s3-accessgrant-grantee-granteeidentifier
	//
	GranteeIdentifier *string `field:"required" json:"granteeIdentifier" yaml:"granteeIdentifier"`
	// Configures the transfer acceleration state for an Amazon S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-accessgrant-grantee.html#cfn-s3-accessgrant-grantee-granteetype
	//
	GranteeType *string `field:"required" json:"granteeType" yaml:"granteeType"`
}

