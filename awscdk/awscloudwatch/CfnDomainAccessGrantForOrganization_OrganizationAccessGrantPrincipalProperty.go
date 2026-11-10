package awscloudwatch


// The principal the grant is issued to.
//
// Supply PrincipalId, or PrincipalAttributes for an attribute-based grant.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   organizationAccessGrantPrincipalProperty := &OrganizationAccessGrantPrincipalProperty{
//   	PrincipalType: jsii.String("principalType"),
//
//   	// the properties below are optional
//   	PrincipalAttributes: []interface{}{
//   		&PrincipalAttributeProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	PrincipalId: jsii.String("principalId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-domainaccessgrantfororganization-organizationaccessgrantprincipal.html
//
type CfnDomainAccessGrantForOrganization_OrganizationAccessGrantPrincipalProperty struct {
	// The type of principal the grant is issued to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-domainaccessgrantfororganization-organizationaccessgrantprincipal.html#cfn-cloudwatch-domainaccessgrantfororganization-organizationaccessgrantprincipal-principaltype
	//
	PrincipalType *string `field:"required" json:"principalType" yaml:"principalType"`
	// Attribute key-value pairs that select the principals the grant applies to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-domainaccessgrantfororganization-organizationaccessgrantprincipal.html#cfn-cloudwatch-domainaccessgrantfororganization-organizationaccessgrantprincipal-principalattributes
	//
	PrincipalAttributes interface{} `field:"optional" json:"principalAttributes" yaml:"principalAttributes"`
	// The identifier of the principal, such as an IAM role ARN or an Identity Center user ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-domainaccessgrantfororganization-organizationaccessgrantprincipal.html#cfn-cloudwatch-domainaccessgrantfororganization-organizationaccessgrantprincipal-principalid
	//
	PrincipalId *string `field:"optional" json:"principalId" yaml:"principalId"`
}

