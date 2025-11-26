package previewawsbedrockmixins


// This element allows you to set up where JPEG, PNG, MOV, and MP4 files get routed to for processing.
//
// JPEG routing applies to both "JPEG" and "JPG" file extensions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   modalityRoutingConfigurationProperty := &ModalityRoutingConfigurationProperty{
//   	Jpeg: jsii.String("jpeg"),
//   	Mov: jsii.String("mov"),
//   	Mp4: jsii.String("mp4"),
//   	Png: jsii.String("png"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-modalityroutingconfiguration.html
//
type CfnDataAutomationProjectPropsMixin_ModalityRoutingConfigurationProperty struct {
	// Sets whether JPEG files are routed to document or image processing.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-modalityroutingconfiguration.html#cfn-bedrock-dataautomationproject-modalityroutingconfiguration-jpeg
	//
	Jpeg *string `field:"optional" json:"jpeg" yaml:"jpeg"`
	// Sets whether MOV files are routed to audio or video processing.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-modalityroutingconfiguration.html#cfn-bedrock-dataautomationproject-modalityroutingconfiguration-mov
	//
	Mov *string `field:"optional" json:"mov" yaml:"mov"`
	// Sets whether MP4 files are routed to audio or video processing.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-modalityroutingconfiguration.html#cfn-bedrock-dataautomationproject-modalityroutingconfiguration-mp4
	//
	Mp4 *string `field:"optional" json:"mp4" yaml:"mp4"`
	// Sets whether PNG files are routed to document or image processing.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-modalityroutingconfiguration.html#cfn-bedrock-dataautomationproject-modalityroutingconfiguration-png
	//
	Png *string `field:"optional" json:"png" yaml:"png"`
}

