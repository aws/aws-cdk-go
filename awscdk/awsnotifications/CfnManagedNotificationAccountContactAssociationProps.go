package awsnotifications


// Properties for defining a `CfnManagedNotificationAccountContactAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnManagedNotificationAccountContactAssociationProps := &CfnManagedNotificationAccountContactAssociationProps{
//   	ContactIdentifier: jsii.String("contactIdentifier"),
//   	ManagedNotificationConfigurationArn: jsii.String("managedNotificationConfigurationArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-notifications-managednotificationaccountcontactassociation.html
//
type CfnManagedNotificationAccountContactAssociationProps struct {
	// The unique identifier of the notification contact associated with the AWS account.
	//
	// For more information about the contact types associated with an account, see the [AWS Account Management Reference Guide](https://docs.aws.amazon.com/accounts/latest/reference/manage-acct-update-contact-alternate.html#manage-acct-update-contact-alternate-orgs) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-notifications-managednotificationaccountcontactassociation.html#cfn-notifications-managednotificationaccountcontactassociation-contactidentifier
	//
	ContactIdentifier *string `field:"required" json:"contactIdentifier" yaml:"contactIdentifier"`
	// The ARN of the `ManagedNotificationConfiguration` to be associated with the `Channel` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-notifications-managednotificationaccountcontactassociation.html#cfn-notifications-managednotificationaccountcontactassociation-managednotificationconfigurationarn
	//
	ManagedNotificationConfigurationArn *string `field:"required" json:"managedNotificationConfigurationArn" yaml:"managedNotificationConfigurationArn"`
}

