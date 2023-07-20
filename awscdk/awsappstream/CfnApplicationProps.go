package awsappstream

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnApplication`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApplicationProps := &CfnApplicationProps{
//   	AppBlockArn: jsii.String("appBlockArn"),
//   	IconS3Location: &S3LocationProperty{
//   		S3Bucket: jsii.String("s3Bucket"),
//   		S3Key: jsii.String("s3Key"),
//   	},
//   	InstanceFamilies: []*string{
//   		jsii.String("instanceFamilies"),
//   	},
//   	LaunchPath: jsii.String("launchPath"),
//   	Name: jsii.String("name"),
//   	Platforms: []*string{
//   		jsii.String("platforms"),
//   	},
//
//   	// the properties below are optional
//   	AttributesToDelete: []*string{
//   		jsii.String("attributesToDelete"),
//   	},
//   	Description: jsii.String("description"),
//   	DisplayName: jsii.String("displayName"),
//   	LaunchParameters: jsii.String("launchParameters"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	WorkingDirectory: jsii.String("workingDirectory"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-application.html
//
type CfnApplicationProps struct {
	// The app block ARN with which the application should be associated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-application.html#cfn-appstream-application-appblockarn
	//
	AppBlockArn *string `field:"required" json:"appBlockArn" yaml:"appBlockArn"`
	// The icon S3 location of the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-application.html#cfn-appstream-application-icons3location
	//
	IconS3Location interface{} `field:"required" json:"iconS3Location" yaml:"iconS3Location"`
	// The instance families the application supports.
	//
	// *Allowed Values* : `GENERAL_PURPOSE` | `GRAPHICS_G4`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-application.html#cfn-appstream-application-instancefamilies
	//
	InstanceFamilies *[]*string `field:"required" json:"instanceFamilies" yaml:"instanceFamilies"`
	// The launch path of the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-application.html#cfn-appstream-application-launchpath
	//
	LaunchPath *string `field:"required" json:"launchPath" yaml:"launchPath"`
	// The name of the application.
	//
	// This name is visible to users when a name is not specified in the DisplayName property.
	//
	// *Pattern* : `^[a-zA-Z0-9][a-zA-Z0-9_.-]{0,100}$`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-application.html#cfn-appstream-application-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The platforms the application supports.
	//
	// *Allowed Values* : `WINDOWS_SERVER_2019` | `AMAZON_LINUX2`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-application.html#cfn-appstream-application-platforms
	//
	Platforms *[]*string `field:"required" json:"platforms" yaml:"platforms"`
	// A list of attributes to delete from an application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-application.html#cfn-appstream-application-attributestodelete
	//
	AttributesToDelete *[]*string `field:"optional" json:"attributesToDelete" yaml:"attributesToDelete"`
	// The description of the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-application.html#cfn-appstream-application-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The display name of the application.
	//
	// This name is visible to users in the application catalog.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-application.html#cfn-appstream-application-displayname
	//
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// The launch parameters of the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-application.html#cfn-appstream-application-launchparameters
	//
	LaunchParameters *string `field:"optional" json:"launchParameters" yaml:"launchParameters"`
	// The tags of the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-application.html#cfn-appstream-application-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The working directory of the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-application.html#cfn-appstream-application-workingdirectory
	//
	WorkingDirectory *string `field:"optional" json:"workingDirectory" yaml:"workingDirectory"`
}

