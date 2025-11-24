package mixinsawsbedrock


// This element is used to determine if the modality it is associated with is enabled or disabled.
//
// All modalities are enabled by default.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   modalityProcessingConfigurationProperty := &ModalityProcessingConfigurationProperty{
//   	State: jsii.String("state"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-modalityprocessingconfiguration.html
//
type CfnDataAutomationProjectPropsMixin_ModalityProcessingConfigurationProperty struct {
	// Stores the state of the modality for your project, set to either enabled or disabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-modalityprocessingconfiguration.html#cfn-bedrock-dataautomationproject-modalityprocessingconfiguration-state
	//
	State *string `field:"optional" json:"state" yaml:"state"`
}

