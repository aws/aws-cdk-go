package awssagemaker

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnImage`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnImageProps := &CfnImageProps{
//   	ImageName: jsii.String("imageName"),
//   	ImageRoleArn: jsii.String("imageRoleArn"),
//
//   	// the properties below are optional
//   	ImageDescription: jsii.String("imageDescription"),
//   	ImageDisplayName: jsii.String("imageDisplayName"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnImageProps struct {
	// The name of the Image. Must be unique by region in your account.
	//
	// *Length Constraints* : Minimum length of 1. Maximum length of 63.
	//
	// *Pattern* : `^[a-zA-Z0-9]([-.]?[a-zA-Z0-9]){0,62}$`
	ImageName *string `field:"required" json:"imageName" yaml:"imageName"`
	// The Amazon Resource Name (ARN) of an IAM role that enables Amazon SageMaker to perform tasks on your behalf.
	//
	// *Length Constraints* : Minimum length of 20. Maximum length of 2048.
	//
	// *Pattern* : `^arn:aws[a-z\-]*:iam::\d{12}:role/?[a-zA-Z_0-9+=,.@\-_/]+$`
	ImageRoleArn *string `field:"required" json:"imageRoleArn" yaml:"imageRoleArn"`
	// The description of the image.
	//
	// *Length Constraints* : Minimum length of 1. Maximum length of 512.
	//
	// *Pattern* : `.*`
	ImageDescription *string `field:"optional" json:"imageDescription" yaml:"imageDescription"`
	// The display name of the image.
	//
	// *Length Constraints* : Minimum length of 1. Maximum length of 128.
	//
	// *Pattern* : `^\S(.*\S)?$`
	ImageDisplayName *string `field:"optional" json:"imageDisplayName" yaml:"imageDisplayName"`
	// A list of key-value pairs to apply to this resource.
	//
	// *Array Members* : Minimum number of 0 items. Maximum number of 50 items.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

