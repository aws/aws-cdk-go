package mixinsawsnimblestudio


// Properties for CfnStudioPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnStudioMixinProps := &CfnStudioMixinProps{
//   	AdminRoleArn: jsii.String("adminRoleArn"),
//   	DisplayName: jsii.String("displayName"),
//   	StudioEncryptionConfiguration: &StudioEncryptionConfigurationProperty{
//   		KeyArn: jsii.String("keyArn"),
//   		KeyType: jsii.String("keyType"),
//   	},
//   	StudioName: jsii.String("studioName"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	UserRoleArn: jsii.String("userRoleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-nimblestudio-studio.html
//
type CfnStudioMixinProps struct {
	// <p>The IAM role that Studio Admins will assume when logging in to the Nimble Studio portal.</p>.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-nimblestudio-studio.html#cfn-nimblestudio-studio-adminrolearn
	//
	AdminRoleArn *string `field:"optional" json:"adminRoleArn" yaml:"adminRoleArn"`
	// <p>A friendly name for the studio.</p>.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-nimblestudio-studio.html#cfn-nimblestudio-studio-displayname
	//
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// <p>Configuration of the encryption method that is used for the studio.</p>.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-nimblestudio-studio.html#cfn-nimblestudio-studio-studioencryptionconfiguration
	//
	StudioEncryptionConfiguration interface{} `field:"optional" json:"studioEncryptionConfiguration" yaml:"studioEncryptionConfiguration"`
	// <p>The studio name that is used in the URL of the Nimble Studio portal when accessed by Nimble Studio users.</p>.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-nimblestudio-studio.html#cfn-nimblestudio-studio-studioname
	//
	StudioName *string `field:"optional" json:"studioName" yaml:"studioName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-nimblestudio-studio.html#cfn-nimblestudio-studio-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// <p>The IAM role that Studio Users will assume when logging in to the Nimble Studio portal.</p>.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-nimblestudio-studio.html#cfn-nimblestudio-studio-userrolearn
	//
	UserRoleArn *string `field:"optional" json:"userRoleArn" yaml:"userRoleArn"`
}

