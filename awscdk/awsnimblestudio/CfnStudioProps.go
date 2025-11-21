package awsnimblestudio


// Properties for defining a `CfnStudio`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnStudioProps := &CfnStudioProps{
//   	AdminRoleArn: jsii.String("adminRoleArn"),
//   	DisplayName: jsii.String("displayName"),
//   	StudioName: jsii.String("studioName"),
//   	UserRoleArn: jsii.String("userRoleArn"),
//
//   	// the properties below are optional
//   	StudioEncryptionConfiguration: &StudioEncryptionConfigurationProperty{
//   		KeyType: jsii.String("keyType"),
//
//   		// the properties below are optional
//   		KeyArn: jsii.String("keyArn"),
//   	},
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-nimblestudio-studio.html
//
type CfnStudioProps struct {
	// <p>The IAM role that Studio Admins will assume when logging in to the Nimble Studio portal.</p>.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-nimblestudio-studio.html#cfn-nimblestudio-studio-adminrolearn
	//
	AdminRoleArn interface{} `field:"required" json:"adminRoleArn" yaml:"adminRoleArn"`
	// <p>A friendly name for the studio.</p>.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-nimblestudio-studio.html#cfn-nimblestudio-studio-displayname
	//
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// <p>The studio name that is used in the URL of the Nimble Studio portal when accessed by Nimble Studio users.</p>.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-nimblestudio-studio.html#cfn-nimblestudio-studio-studioname
	//
	StudioName *string `field:"required" json:"studioName" yaml:"studioName"`
	// <p>The IAM role that Studio Users will assume when logging in to the Nimble Studio portal.</p>.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-nimblestudio-studio.html#cfn-nimblestudio-studio-userrolearn
	//
	UserRoleArn interface{} `field:"required" json:"userRoleArn" yaml:"userRoleArn"`
	// <p>Configuration of the encryption method that is used for the studio.</p>.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-nimblestudio-studio.html#cfn-nimblestudio-studio-studioencryptionconfiguration
	//
	StudioEncryptionConfiguration interface{} `field:"optional" json:"studioEncryptionConfiguration" yaml:"studioEncryptionConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-nimblestudio-studio.html#cfn-nimblestudio-studio-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

