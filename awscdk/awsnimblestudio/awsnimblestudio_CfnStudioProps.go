package awsnimblestudio


// Properties for defining a `CfnStudio`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnStudioProps := &cfnStudioProps{
//   	adminRoleArn: jsii.String("adminRoleArn"),
//   	displayName: jsii.String("displayName"),
//   	studioName: jsii.String("studioName"),
//   	userRoleArn: jsii.String("userRoleArn"),
//
//   	// the properties below are optional
//   	studioEncryptionConfiguration: &studioEncryptionConfigurationProperty{
//   		keyType: jsii.String("keyType"),
//
//   		// the properties below are optional
//   		keyArn: jsii.String("keyArn"),
//   	},
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
type CfnStudioProps struct {
	// The IAM role that studio admins assume when logging in to the Nimble Studio portal.
	AdminRoleArn *string `field:"required" json:"adminRoleArn" yaml:"adminRoleArn"`
	// A friendly name for the studio.
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The name of the studio, as included in the URL when accessing it in the Nimble Studio portal.
	StudioName *string `field:"required" json:"studioName" yaml:"studioName"`
	// The IAM role that studio users assume when logging in to the Nimble Studio portal.
	UserRoleArn *string `field:"required" json:"userRoleArn" yaml:"userRoleArn"`
	// Configuration of the encryption method that is used for the studio.
	StudioEncryptionConfiguration interface{} `field:"optional" json:"studioEncryptionConfiguration" yaml:"studioEncryptionConfiguration"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

