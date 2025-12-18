package previewawsbedrockmixins


// Configuration for detecting and redacting Personally Identifiable Information (PII) entities.
//
// Specify which PII entity types to detect and the redaction mask mode. If not provided, defaults to ALL entity types with ENTITY_TYPE redaction mask mode.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   pIIEntitiesConfigurationProperty := &PIIEntitiesConfigurationProperty{
//   	PiiEntityTypes: []*string{
//   		jsii.String("piiEntityTypes"),
//   	},
//   	RedactionMaskMode: jsii.String("redactionMaskMode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-piientitiesconfiguration.html
//
type CfnDataAutomationProjectPropsMixin_PIIEntitiesConfigurationProperty struct {
	// List of PII entity types to detect/redact in the output.
	//
	// Choose from specific entity types (such as ADDRESS, NAME, EMAIL, PHONE, US_SOCIAL_SECURITY_NUMBER) or specify ALL to detect all supported PII types. If not specified, defaults to ALL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-piientitiesconfiguration.html#cfn-bedrock-dataautomationproject-piientitiesconfiguration-piientitytypes
	//
	PiiEntityTypes *[]*string `field:"optional" json:"piiEntityTypes" yaml:"piiEntityTypes"`
	// Defines how detected PII entities are masked in redacted output files.
	//
	// Set to PII to replace all detected entities with a generic [PII] marker regardless of entity type. Set to ENTITY_TYPE to replace each detected entity with its specific type marker (for example, [NAME], [EMAIL], [ADDRESS]). This setting only applies when detectionMode is set to DETECTION_AND_REDACTION. If not specified, defaults to ENTITY_TYPE.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-piientitiesconfiguration.html#cfn-bedrock-dataautomationproject-piientitiesconfiguration-redactionmaskmode
	//
	RedactionMaskMode *string `field:"optional" json:"redactionMaskMode" yaml:"redactionMaskMode"`
}

