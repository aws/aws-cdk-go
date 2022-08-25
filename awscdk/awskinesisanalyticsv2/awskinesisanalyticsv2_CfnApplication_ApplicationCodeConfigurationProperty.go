package awskinesisanalyticsv2


// Describes code configuration for an application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   applicationCodeConfigurationProperty := &applicationCodeConfigurationProperty{
//   	codeContent: &codeContentProperty{
//   		s3ContentLocation: &s3ContentLocationProperty{
//   			bucketArn: jsii.String("bucketArn"),
//   			fileKey: jsii.String("fileKey"),
//
//   			// the properties below are optional
//   			objectVersion: jsii.String("objectVersion"),
//   		},
//   		textContent: jsii.String("textContent"),
//   		zipFileContent: jsii.String("zipFileContent"),
//   	},
//   	codeContentType: jsii.String("codeContentType"),
//   }
//
type CfnApplication_ApplicationCodeConfigurationProperty struct {
	// The location and type of the application code.
	CodeContent interface{} `field:"required" json:"codeContent" yaml:"codeContent"`
	// Specifies whether the code content is in text or zip format.
	CodeContentType *string `field:"required" json:"codeContentType" yaml:"codeContentType"`
}

