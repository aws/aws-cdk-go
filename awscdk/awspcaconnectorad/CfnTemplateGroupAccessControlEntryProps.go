package awspcaconnectorad


// Properties for defining a `CfnTemplateGroupAccessControlEntry`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTemplateGroupAccessControlEntryProps := &CfnTemplateGroupAccessControlEntryProps{
//   	AccessRights: &AccessRightsProperty{
//   		AutoEnroll: jsii.String("autoEnroll"),
//   		Enroll: jsii.String("enroll"),
//   	},
//   	GroupDisplayName: jsii.String("groupDisplayName"),
//
//   	// the properties below are optional
//   	GroupSecurityIdentifier: jsii.String("groupSecurityIdentifier"),
//   	TemplateArn: jsii.String("templateArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pcaconnectorad-templategroupaccesscontrolentry.html
//
type CfnTemplateGroupAccessControlEntryProps struct {
	// Permissions to allow or deny an Active Directory group to enroll or autoenroll certificates issued against a template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pcaconnectorad-templategroupaccesscontrolentry.html#cfn-pcaconnectorad-templategroupaccesscontrolentry-accessrights
	//
	AccessRights interface{} `field:"required" json:"accessRights" yaml:"accessRights"`
	// Name of the Active Directory group.
	//
	// This name does not need to match the group name in Active Directory.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pcaconnectorad-templategroupaccesscontrolentry.html#cfn-pcaconnectorad-templategroupaccesscontrolentry-groupdisplayname
	//
	GroupDisplayName *string `field:"required" json:"groupDisplayName" yaml:"groupDisplayName"`
	// Security identifier (SID) of the group object from Active Directory.
	//
	// The SID starts with "S-".
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pcaconnectorad-templategroupaccesscontrolentry.html#cfn-pcaconnectorad-templategroupaccesscontrolentry-groupsecurityidentifier
	//
	GroupSecurityIdentifier *string `field:"optional" json:"groupSecurityIdentifier" yaml:"groupSecurityIdentifier"`
	// The Amazon Resource Name (ARN) that was returned when you called [CreateTemplate](https://docs.aws.amazon.com/pca-connector-ad/latest/APIReference/API_CreateTemplate.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pcaconnectorad-templategroupaccesscontrolentry.html#cfn-pcaconnectorad-templategroupaccesscontrolentry-templatearn
	//
	TemplateArn *string `field:"optional" json:"templateArn" yaml:"templateArn"`
}

