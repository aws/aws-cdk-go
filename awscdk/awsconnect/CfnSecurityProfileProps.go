package awsconnect

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnSecurityProfile`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSecurityProfileProps := &CfnSecurityProfileProps{
//   	InstanceArn: jsii.String("instanceArn"),
//   	SecurityProfileName: jsii.String("securityProfileName"),
//
//   	// the properties below are optional
//   	AllowedAccessControlHierarchyGroupId: jsii.String("allowedAccessControlHierarchyGroupId"),
//   	AllowedAccessControlTags: []interface{}{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Applications: []interface{}{
//   		&ApplicationProperty{
//   			ApplicationPermissions: []*string{
//   				jsii.String("applicationPermissions"),
//   			},
//   			Namespace: jsii.String("namespace"),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	HierarchyRestrictedResources: []*string{
//   		jsii.String("hierarchyRestrictedResources"),
//   	},
//   	Permissions: []*string{
//   		jsii.String("permissions"),
//   	},
//   	TagRestrictedResources: []*string{
//   		jsii.String("tagRestrictedResources"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-securityprofile.html
//
type CfnSecurityProfileProps struct {
	// The identifier of the Amazon Connect instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-securityprofile.html#cfn-connect-securityprofile-instancearn
	//
	InstanceArn *string `field:"required" json:"instanceArn" yaml:"instanceArn"`
	// The name for the security profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-securityprofile.html#cfn-connect-securityprofile-securityprofilename
	//
	SecurityProfileName *string `field:"required" json:"securityProfileName" yaml:"securityProfileName"`
	// The identifier of the hierarchy group that a security profile uses to restrict access to resources in Amazon Connect.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-securityprofile.html#cfn-connect-securityprofile-allowedaccesscontrolhierarchygroupid
	//
	AllowedAccessControlHierarchyGroupId *string `field:"optional" json:"allowedAccessControlHierarchyGroupId" yaml:"allowedAccessControlHierarchyGroupId"`
	// The list of tags that a security profile uses to restrict access to resources in Amazon Connect.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-securityprofile.html#cfn-connect-securityprofile-allowedaccesscontroltags
	//
	AllowedAccessControlTags interface{} `field:"optional" json:"allowedAccessControlTags" yaml:"allowedAccessControlTags"`
	// A list of third-party applications that the security profile will give access to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-securityprofile.html#cfn-connect-securityprofile-applications
	//
	Applications interface{} `field:"optional" json:"applications" yaml:"applications"`
	// The description of the security profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-securityprofile.html#cfn-connect-securityprofile-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The list of resources that a security profile applies hierarchy restrictions to in Amazon Connect.
	//
	// Following are acceptable ResourceNames: `User` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-securityprofile.html#cfn-connect-securityprofile-hierarchyrestrictedresources
	//
	HierarchyRestrictedResources *[]*string `field:"optional" json:"hierarchyRestrictedResources" yaml:"hierarchyRestrictedResources"`
	// Permissions assigned to the security profile.
	//
	// For a list of valid permissions, see [List of security profile permissions](https://docs.aws.amazon.com/connect/latest/adminguide/security-profile-list.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-securityprofile.html#cfn-connect-securityprofile-permissions
	//
	Permissions *[]*string `field:"optional" json:"permissions" yaml:"permissions"`
	// The list of resources that a security profile applies tag restrictions to in Amazon Connect.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-securityprofile.html#cfn-connect-securityprofile-tagrestrictedresources
	//
	TagRestrictedResources *[]*string `field:"optional" json:"tagRestrictedResources" yaml:"tagRestrictedResources"`
	// The tags used to organize, track, or control access for this resource.
	//
	// For example, { "Tags": {"key1":"value1", "key2":"value2"} }.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-securityprofile.html#cfn-connect-securityprofile-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

