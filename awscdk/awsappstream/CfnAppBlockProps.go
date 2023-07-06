package awsappstream

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnAppBlock`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAppBlockProps := &CfnAppBlockProps{
//   	Name: jsii.String("name"),
//   	SourceS3Location: &S3LocationProperty{
//   		S3Bucket: jsii.String("s3Bucket"),
//
//   		// the properties below are optional
//   		S3Key: jsii.String("s3Key"),
//   	},
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	DisplayName: jsii.String("displayName"),
//   	PackagingType: jsii.String("packagingType"),
//   	PostSetupScriptDetails: &ScriptDetailsProperty{
//   		ExecutablePath: jsii.String("executablePath"),
//   		ScriptS3Location: &S3LocationProperty{
//   			S3Bucket: jsii.String("s3Bucket"),
//
//   			// the properties below are optional
//   			S3Key: jsii.String("s3Key"),
//   		},
//   		TimeoutInSeconds: jsii.Number(123),
//
//   		// the properties below are optional
//   		ExecutableParameters: jsii.String("executableParameters"),
//   	},
//   	SetupScriptDetails: &ScriptDetailsProperty{
//   		ExecutablePath: jsii.String("executablePath"),
//   		ScriptS3Location: &S3LocationProperty{
//   			S3Bucket: jsii.String("s3Bucket"),
//
//   			// the properties below are optional
//   			S3Key: jsii.String("s3Key"),
//   		},
//   		TimeoutInSeconds: jsii.Number(123),
//
//   		// the properties below are optional
//   		ExecutableParameters: jsii.String("executableParameters"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnAppBlockProps struct {
	// The name of the app block.
	//
	// *Pattern* : `^[a-zA-Z0-9][a-zA-Z0-9_.-]{0,100}$`
	Name *string `field:"required" json:"name" yaml:"name"`
	// The source S3 location of the app block.
	SourceS3Location interface{} `field:"required" json:"sourceS3Location" yaml:"sourceS3Location"`
	// The description of the app block.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The display name of the app block.
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// The packaging type of the app block.
	PackagingType *string `field:"optional" json:"packagingType" yaml:"packagingType"`
	// The post setup script details of the app block.
	PostSetupScriptDetails interface{} `field:"optional" json:"postSetupScriptDetails" yaml:"postSetupScriptDetails"`
	// The setup script details of the app block.
	SetupScriptDetails interface{} `field:"optional" json:"setupScriptDetails" yaml:"setupScriptDetails"`
	// The tags of the app block.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

