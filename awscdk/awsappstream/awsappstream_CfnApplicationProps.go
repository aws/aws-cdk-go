package awsappstream

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnApplication`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApplicationProps := &cfnApplicationProps{
//   	appBlockArn: jsii.String("appBlockArn"),
//   	iconS3Location: &s3LocationProperty{
//   		s3Bucket: jsii.String("s3Bucket"),
//   		s3Key: jsii.String("s3Key"),
//   	},
//   	instanceFamilies: []*string{
//   		jsii.String("instanceFamilies"),
//   	},
//   	launchPath: jsii.String("launchPath"),
//   	name: jsii.String("name"),
//   	platforms: []*string{
//   		jsii.String("platforms"),
//   	},
//
//   	// the properties below are optional
//   	attributesToDelete: []*string{
//   		jsii.String("attributesToDelete"),
//   	},
//   	description: jsii.String("description"),
//   	displayName: jsii.String("displayName"),
//   	launchParameters: jsii.String("launchParameters"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	workingDirectory: jsii.String("workingDirectory"),
//   }
//
type CfnApplicationProps struct {
	// The app block ARN with which the application should be associated.
	AppBlockArn *string `field:"required" json:"appBlockArn" yaml:"appBlockArn"`
	// The icon S3 location of the application.
	IconS3Location interface{} `field:"required" json:"iconS3Location" yaml:"iconS3Location"`
	// The instance families the application supports.
	//
	// *Allowed Values* : `GENERAL_PURPOSE` | `GRAPHICS_G4`.
	InstanceFamilies *[]*string `field:"required" json:"instanceFamilies" yaml:"instanceFamilies"`
	// The launch path of the application.
	LaunchPath *string `field:"required" json:"launchPath" yaml:"launchPath"`
	// The name of the application.
	//
	// This name is visible to users when a name is not specified in the DisplayName property.
	//
	// *Pattern* : `^[a-zA-Z0-9][a-zA-Z0-9_.-]{0,100}$`
	Name *string `field:"required" json:"name" yaml:"name"`
	// The platforms the application supports.
	//
	// *Allowed Values* : `WINDOWS_SERVER_2019` | `AMAZON_LINUX2`.
	Platforms *[]*string `field:"required" json:"platforms" yaml:"platforms"`
	// A list of attributes to delete from an application.
	AttributesToDelete *[]*string `field:"optional" json:"attributesToDelete" yaml:"attributesToDelete"`
	// The description of the application.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The display name of the application.
	//
	// This name is visible to users in the application catalog.
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// The launch parameters of the application.
	LaunchParameters *string `field:"optional" json:"launchParameters" yaml:"launchParameters"`
	// The tags of the application.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The working directory of the application.
	WorkingDirectory *string `field:"optional" json:"workingDirectory" yaml:"workingDirectory"`
}

