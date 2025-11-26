package previewawsappstreammixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnAppBlockPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAppBlockMixinProps := &CfnAppBlockMixinProps{
//   	Description: jsii.String("description"),
//   	DisplayName: jsii.String("displayName"),
//   	Name: jsii.String("name"),
//   	PackagingType: jsii.String("packagingType"),
//   	PostSetupScriptDetails: &ScriptDetailsProperty{
//   		ExecutableParameters: jsii.String("executableParameters"),
//   		ExecutablePath: jsii.String("executablePath"),
//   		ScriptS3Location: &S3LocationProperty{
//   			S3Bucket: jsii.String("s3Bucket"),
//   			S3Key: jsii.String("s3Key"),
//   		},
//   		TimeoutInSeconds: jsii.Number(123),
//   	},
//   	SetupScriptDetails: &ScriptDetailsProperty{
//   		ExecutableParameters: jsii.String("executableParameters"),
//   		ExecutablePath: jsii.String("executablePath"),
//   		ScriptS3Location: &S3LocationProperty{
//   			S3Bucket: jsii.String("s3Bucket"),
//   			S3Key: jsii.String("s3Key"),
//   		},
//   		TimeoutInSeconds: jsii.Number(123),
//   	},
//   	SourceS3Location: &S3LocationProperty{
//   		S3Bucket: jsii.String("s3Bucket"),
//   		S3Key: jsii.String("s3Key"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-appblock.html
//
type CfnAppBlockMixinProps struct {
	// The description of the app block.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-appblock.html#cfn-appstream-appblock-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The display name of the app block.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-appblock.html#cfn-appstream-appblock-displayname
	//
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// The name of the app block.
	//
	// *Pattern* : `^[a-zA-Z0-9][a-zA-Z0-9_.-]{0,100}$`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-appblock.html#cfn-appstream-appblock-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The packaging type of the app block.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-appblock.html#cfn-appstream-appblock-packagingtype
	//
	PackagingType *string `field:"optional" json:"packagingType" yaml:"packagingType"`
	// The post setup script details of the app block.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-appblock.html#cfn-appstream-appblock-postsetupscriptdetails
	//
	PostSetupScriptDetails interface{} `field:"optional" json:"postSetupScriptDetails" yaml:"postSetupScriptDetails"`
	// The setup script details of the app block.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-appblock.html#cfn-appstream-appblock-setupscriptdetails
	//
	SetupScriptDetails interface{} `field:"optional" json:"setupScriptDetails" yaml:"setupScriptDetails"`
	// The source S3 location of the app block.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-appblock.html#cfn-appstream-appblock-sources3location
	//
	SourceS3Location interface{} `field:"optional" json:"sourceS3Location" yaml:"sourceS3Location"`
	// The tags of the app block.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-appblock.html#cfn-appstream-appblock-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

