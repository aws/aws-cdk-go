package awsappstream

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnAppBlock`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAppBlockProps := &cfnAppBlockProps{
//   	name: jsii.String("name"),
//   	setupScriptDetails: &scriptDetailsProperty{
//   		executablePath: jsii.String("executablePath"),
//   		scriptS3Location: &s3LocationProperty{
//   			s3Bucket: jsii.String("s3Bucket"),
//   			s3Key: jsii.String("s3Key"),
//   		},
//   		timeoutInSeconds: jsii.Number(123),
//
//   		// the properties below are optional
//   		executableParameters: jsii.String("executableParameters"),
//   	},
//   	sourceS3Location: &s3LocationProperty{
//   		s3Bucket: jsii.String("s3Bucket"),
//   		s3Key: jsii.String("s3Key"),
//   	},
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	displayName: jsii.String("displayName"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnAppBlockProps struct {
	// The name of the app block.
	//
	// *Pattern* : `^[a-zA-Z0-9][a-zA-Z0-9_.-]{0,100}$`
	Name *string `field:"required" json:"name" yaml:"name"`
	// The setup script details of the app block.
	SetupScriptDetails interface{} `field:"required" json:"setupScriptDetails" yaml:"setupScriptDetails"`
	// The source S3 location of the app block.
	SourceS3Location interface{} `field:"required" json:"sourceS3Location" yaml:"sourceS3Location"`
	// The description of the app block.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The display name of the app block.
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// The tags of the app block.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

