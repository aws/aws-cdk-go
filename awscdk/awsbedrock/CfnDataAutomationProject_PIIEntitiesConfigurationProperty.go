package awsbedrock


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
type CfnDataAutomationProject_PIIEntitiesConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-piientitiesconfiguration.html#cfn-bedrock-dataautomationproject-piientitiesconfiguration-piientitytypes
	//
	PiiEntityTypes *[]*string `field:"optional" json:"piiEntityTypes" yaml:"piiEntityTypes"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-piientitiesconfiguration.html#cfn-bedrock-dataautomationproject-piientitiesconfiguration-redactionmaskmode
	//
	RedactionMaskMode *string `field:"optional" json:"redactionMaskMode" yaml:"redactionMaskMode"`
}

