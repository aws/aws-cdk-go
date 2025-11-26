package previewawsbedrockmixins


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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-sensitivedataconfiguration.html#cfn-bedrock-dataautomationproject-sensitivedataconfiguration-detectionmode
	//
	DetectionMode *string `field:"optional" json:"detectionMode" yaml:"detectionMode"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-sensitivedataconfiguration.html#cfn-bedrock-dataautomationproject-sensitivedataconfiguration-detectionscope
	//
	DetectionScope *[]*string `field:"optional" json:"detectionScope" yaml:"detectionScope"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-sensitivedataconfiguration.html#cfn-bedrock-dataautomationproject-sensitivedataconfiguration-piientitiesconfiguration
	//
	PiiEntitiesConfiguration interface{} `field:"optional" json:"piiEntitiesConfiguration" yaml:"piiEntitiesConfiguration"`
}

