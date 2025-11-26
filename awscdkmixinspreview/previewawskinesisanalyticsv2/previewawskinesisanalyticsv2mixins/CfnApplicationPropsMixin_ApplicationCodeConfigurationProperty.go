package previewawskinesisanalyticsv2mixins


// Describes code configuration for an application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   applicationCodeConfigurationProperty := &ApplicationCodeConfigurationProperty{
//   	CodeContent: &CodeContentProperty{
//   		S3ContentLocation: &S3ContentLocationProperty{
//   			BucketArn: jsii.String("bucketArn"),
//   			FileKey: jsii.String("fileKey"),
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
type CfnApplicationPropsMixin_ApplicationCodeConfigurationProperty struct {
	// The location and type of the application code.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-applicationcodeconfiguration.html#cfn-kinesisanalyticsv2-application-applicationcodeconfiguration-codecontent
	//
	CodeContent interface{} `field:"optional" json:"codeContent" yaml:"codeContent"`
	// Specifies whether the code content is in text or zip format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-applicationcodeconfiguration.html#cfn-kinesisanalyticsv2-application-applicationcodeconfiguration-codecontenttype
	//
	CodeContentType *string `field:"optional" json:"codeContentType" yaml:"codeContentType"`
}

