package awskinesisanalytics


// Describes code configuration for an application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   applicationCodeConfigurationProperty := &ApplicationCodeConfigurationProperty{
//   	CodeContent: &CodeContentProperty{
//   		S3ContentLocation: &S3ContentLocationProperty{
//   			BucketArn: jsii.String("bucketArn"),
//   			FileKey: jsii.String("fileKey"),
//
//   			// the properties below are optional
//   			ObjectVersion: jsii.String("objectVersion"),
//   		},
//   		TextContent: jsii.String("textContent"),
//   		ZipFileContent: jsii.String("zipFileContent"),
//   	},
//   	CodeContentType: jsii.String("codeContentType"),
//   }
//
type CfnApplicationV2_ApplicationCodeConfigurationProperty struct {
	// The location and type of the application code.
	CodeContent interface{} `field:"required" json:"codeContent" yaml:"codeContent"`
	// Specifies whether the code content is in text or zip format.
	CodeContentType *string `field:"required" json:"codeContentType" yaml:"codeContentType"`
}

