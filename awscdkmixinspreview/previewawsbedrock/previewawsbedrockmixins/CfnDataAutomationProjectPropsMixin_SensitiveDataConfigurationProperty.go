package previewawsbedrockmixins


// Configuration for detecting and redacting sensitive data in content.
//
// Use this to control whether sensitive data is detected only or both detected and redacted, specify the scope of detection (standard output, custom output, or both), and configure specific PII entity types to detect along with how they should be masked when redacted.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sensitiveDataConfigurationProperty := &SensitiveDataConfigurationProperty{
//   	DetectionMode: jsii.String("detectionMode"),
//   	DetectionScope: []*string{
//   		jsii.String("detectionScope"),
//   	},
//   	PiiEntitiesConfiguration: &PIIEntitiesConfigurationProperty{
//   		PiiEntityTypes: []*string{
//   			jsii.String("piiEntityTypes"),
//   		},
//   		RedactionMaskMode: jsii.String("redactionMaskMode"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-sensitivedataconfiguration.html
//
type CfnDataAutomationProjectPropsMixin_SensitiveDataConfigurationProperty struct {
	// Specifies the mode for handling sensitive data detection.
	//
	// Set to DETECTION to only identify sensitive data without modifying content - this produces one output file per detection scope containing detection information with original unredacted content. Set to DETECTION_AND_REDACTION to both identify and mask sensitive data - this produces two output files per detection scope: one unredacted file with detection information and one redacted file with masking applied to sensitive content. For example, if detectionScope includes both STANDARD and CUSTOM with DETECTION_AND_REDACTION mode, four output files will be generated (two for standard output and two for custom output).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-sensitivedataconfiguration.html#cfn-bedrock-dataautomationproject-sensitivedataconfiguration-detectionmode
	//
	DetectionMode *string `field:"optional" json:"detectionMode" yaml:"detectionMode"`
	// Defines which BDA output types to apply sensitive data detection to.
	//
	// Specify STANDARD to detect sensitive data in standard output, CUSTOM to detect in custom output (blueprint-based extraction), or both to apply detection to both output types. If not specified, defaults to both STANDARD and CUSTOM. The number of output files generated depends on both the detection mode and the scopes selected - each scope specified will produce its own set of output files according to the detection mode configured.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-sensitivedataconfiguration.html#cfn-bedrock-dataautomationproject-sensitivedataconfiguration-detectionscope
	//
	DetectionScope *[]*string `field:"optional" json:"detectionScope" yaml:"detectionScope"`
	// Configuration for detecting and redacting Personally Identifiable Information (PII) entities.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-sensitivedataconfiguration.html#cfn-bedrock-dataautomationproject-sensitivedataconfiguration-piientitiesconfiguration
	//
	PiiEntitiesConfiguration interface{} `field:"optional" json:"piiEntitiesConfiguration" yaml:"piiEntitiesConfiguration"`
}

