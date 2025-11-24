package mixinsawsnotifications


// Properties for CfnOrganizationalUnitAssociationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnOrganizationalUnitAssociationMixinProps := &CfnOrganizationalUnitAssociationMixinProps{
//   	NotificationConfigurationArn: jsii.String("notificationConfigurationArn"),
//   	OrganizationalUnitId: jsii.String("organizationalUnitId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-notifications-organizationalunitassociation.html
//
type CfnOrganizationalUnitAssociationMixinProps struct {
	// ARN identifier of the NotificationConfiguration.
	//
	// Example: arn:aws:notifications::123456789012:configuration/a01jes88qxwkbj05xv9c967pgm1.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-notifications-organizationalunitassociation.html#cfn-notifications-organizationalunitassociation-notificationconfigurationarn
	//
	NotificationConfigurationArn *string `field:"optional" json:"notificationConfigurationArn" yaml:"notificationConfigurationArn"`
	// The ID of the organizational unit.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-notifications-organizationalunitassociation.html#cfn-notifications-organizationalunitassociation-organizationalunitid
	//
	OrganizationalUnitId *string `field:"optional" json:"organizationalUnitId" yaml:"organizationalUnitId"`
}

