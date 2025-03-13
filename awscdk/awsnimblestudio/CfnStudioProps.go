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
	// The IAM role that studio admins assume when logging in to the Nimble Studio portal.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-nimblestudio-studio.html#cfn-nimblestudio-studio-adminrolearn
	//
	AdminRoleArn *string `field:"required" json:"adminRoleArn" yaml:"adminRoleArn"`
	// A friendly name for the studio.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-nimblestudio-studio.html#cfn-nimblestudio-studio-displayname
	//
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The name of the studio, as included in the URL when accessing it in the Nimble Studio portal.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-nimblestudio-studio.html#cfn-nimblestudio-studio-studioname
	//
	StudioName *string `field:"required" json:"studioName" yaml:"studioName"`
	// The IAM role that studio users assume when logging in to the Nimble Studio portal.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-nimblestudio-studio.html#cfn-nimblestudio-studio-userrolearn
	//
	UserRoleArn *string `field:"required" json:"userRoleArn" yaml:"userRoleArn"`
	// Configuration of the encryption method that is used for the studio.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-nimblestudio-studio.html#cfn-nimblestudio-studio-studioencryptionconfiguration
	//
	StudioEncryptionConfiguration interface{} `field:"optional" json:"studioEncryptionConfiguration" yaml:"studioEncryptionConfiguration"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-nimblestudio-studio.html#cfn-nimblestudio-studio-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

