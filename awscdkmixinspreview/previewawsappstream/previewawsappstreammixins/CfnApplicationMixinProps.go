package previewawsappstreammixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnApplicationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnApplicationMixinProps := &CfnApplicationMixinProps{
//   	AppBlockArn: jsii.String("appBlockArn"),
//   	AttributesToDelete: []*string{
//   		jsii.String("attributesToDelete"),
//   	},
//   	Description: jsii.String("description"),
//   	DisplayName: jsii.String("displayName"),
//   	IconS3Location: &S3LocationProperty{
//   		S3Bucket: jsii.String("s3Bucket"),
//   		S3Key: jsii.String("s3Key"),
//   	},
//   	InstanceFamilies: []*string{
//   		jsii.String("instanceFamilies"),
//   	},
//   	LaunchParameters: jsii.String("launchParameters"),
//   	LaunchPath: jsii.String("launchPath"),
//   	Name: jsii.String("name"),
//   	Platforms: []*string{
//   		jsii.String("platforms"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	WorkingDirectory: jsii.String("workingDirectory"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-application.html
//
type CfnApplicationMixinProps struct {
	// The app block ARN with which the application should be associated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-application.html#cfn-appstream-application-appblockarn
	//
	AppBlockArn *string `field:"optional" json:"appBlockArn" yaml:"appBlockArn"`
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
	// The icon S3 location of the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-application.html#cfn-appstream-application-icons3location
	//
	IconS3Location interface{} `field:"optional" json:"iconS3Location" yaml:"iconS3Location"`
	// The instance families the application supports.
	//
	// *Allowed Values* : `GENERAL_PURPOSE` | `GRAPHICS_G4`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-application.html#cfn-appstream-application-instancefamilies
	//
	InstanceFamilies *[]*string `field:"optional" json:"instanceFamilies" yaml:"instanceFamilies"`
	// The launch parameters of the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-application.html#cfn-appstream-application-launchparameters
	//
	LaunchParameters *string `field:"optional" json:"launchParameters" yaml:"launchParameters"`
	// The launch path of the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-application.html#cfn-appstream-application-launchpath
	//
	LaunchPath *string `field:"optional" json:"launchPath" yaml:"launchPath"`
	// The name of the application.
	//
	// This name is visible to users when a name is not specified in the DisplayName property.
	//
	// *Pattern* : `^[a-zA-Z0-9][a-zA-Z0-9_.-]{0,100}$`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-application.html#cfn-appstream-application-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The platforms the application supports.
	//
	// *Allowed Values* : `WINDOWS_SERVER_2019` | `AMAZON_LINUX2`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-application.html#cfn-appstream-application-platforms
	//
	Platforms *[]*string `field:"optional" json:"platforms" yaml:"platforms"`
	// The tags of the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-application.html#cfn-appstream-application-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The working directory of the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-application.html#cfn-appstream-application-workingdirectory
	//
	WorkingDirectory *string `field:"optional" json:"workingDirectory" yaml:"workingDirectory"`
}

