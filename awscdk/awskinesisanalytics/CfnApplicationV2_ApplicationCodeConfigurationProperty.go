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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-applicationcodeconfiguration.html
//
type CfnApplicationV2_ApplicationCodeConfigurationProperty struct {
	// The location and type of the application code.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-applicationcodeconfiguration.html#cfn-kinesisanalyticsv2-application-applicationcodeconfiguration-codecontent
	//
	CodeContent interface{} `field:"required" json:"codeContent" yaml:"codeContent"`
	// Specifies whether the code content is in text or zip format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-applicationcodeconfiguration.html#cfn-kinesisanalyticsv2-application-applicationcodeconfiguration-codecontenttype
	//
	CodeContentType *string `field:"required" json:"codeContentType" yaml:"codeContentType"`
}

