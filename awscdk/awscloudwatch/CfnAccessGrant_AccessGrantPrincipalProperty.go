package awscloudwatch


// The principal receiving the grant.
//
// Specify PrincipalId, PrincipalAttributes, or both.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accessGrantPrincipalProperty := &AccessGrantPrincipalProperty{
//   	PrincipalType: jsii.String("principalType"),
//
//   	// the properties below are optional
//   	PrincipalAttributes: []interface{}{
//   		&AccessGrantPrincipalAttributeProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	PrincipalId: jsii.String("principalId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-accessgrant-accessgrantprincipal.html
//
type CfnAccessGrant_AccessGrantPrincipalProperty struct {
	// The type of principal receiving the grant.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-accessgrant-accessgrantprincipal.html#cfn-cloudwatch-accessgrant-accessgrantprincipal-principaltype
	//
	PrincipalType *string `field:"required" json:"principalType" yaml:"principalType"`
	// Attribute conditions for attribute-based access.
	//
	// When provided, the grant targets any principal matching all specified conditions. Supported only for IDC_USER principals.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-accessgrant-accessgrantprincipal.html#cfn-cloudwatch-accessgrant-accessgrantprincipal-principalattributes
	//
	PrincipalAttributes interface{} `field:"optional" json:"principalAttributes" yaml:"principalAttributes"`
	// The ID of the principal receiving the grant.
	//
	// An IAM principal ARN, an Identity Center UUID, or a service-defined identifier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-accessgrant-accessgrantprincipal.html#cfn-cloudwatch-accessgrant-accessgrantprincipal-principalid
	//
	PrincipalId *string `field:"optional" json:"principalId" yaml:"principalId"`
}

