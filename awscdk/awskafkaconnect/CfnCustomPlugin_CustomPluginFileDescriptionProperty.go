package awskafkaconnect


// Details about the custom plugin file.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customPluginFileDescriptionProperty := &CustomPluginFileDescriptionProperty{
//   	FileMd5: jsii.String("fileMd5"),
//   	FileSize: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kafkaconnect-customplugin-custompluginfiledescription.html
//
type CfnCustomPlugin_CustomPluginFileDescriptionProperty struct {
	// The hex-encoded MD5 checksum of the custom plugin file.
	//
	// You can use it to validate the file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kafkaconnect-customplugin-custompluginfiledescription.html#cfn-kafkaconnect-customplugin-custompluginfiledescription-filemd5
	//
	FileMd5 *string `field:"optional" json:"fileMd5" yaml:"fileMd5"`
	// The size in bytes of the custom plugin file.
	//
	// You can use it to validate the file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kafkaconnect-customplugin-custompluginfiledescription.html#cfn-kafkaconnect-customplugin-custompluginfiledescription-filesize
	//
	FileSize *float64 `field:"optional" json:"fileSize" yaml:"fileSize"`
}

