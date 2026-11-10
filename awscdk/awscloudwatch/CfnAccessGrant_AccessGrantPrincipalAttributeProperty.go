package awscloudwatch


// A single attribute condition used to match principals for attribute-based access.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accessGrantPrincipalAttributeProperty := &AccessGrantPrincipalAttributeProperty{
//   	Key: jsii.String("key"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-accessgrant-accessgrantprincipalattribute.html
//
type CfnAccessGrant_AccessGrantPrincipalAttributeProperty struct {
	// The Identity Center user attribute to match on.
	//
	// One of userName, active, userStatus, displayName, email, name.givenName, name.familyName, enterprise.department, enterprise.division, enterprise.organization, enterprise.costCenter, or enterprise.employeeNumber.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-accessgrant-accessgrantprincipalattribute.html#cfn-cloudwatch-accessgrant-accessgrantprincipalattribute-key
	//
	Key *string `field:"required" json:"key" yaml:"key"`
	// The attribute value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-accessgrant-accessgrantprincipalattribute.html#cfn-cloudwatch-accessgrant-accessgrantprincipalattribute-value
	//
	Value *string `field:"required" json:"value" yaml:"value"`
}

