package awscdk


// The configuration details of the generated template.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   templateConfigurationProperty := &TemplateConfigurationProperty{
//   	DeletionPolicy: jsii.String("deletionPolicy"),
//   	UpdateReplacePolicy: jsii.String("updateReplacePolicy"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-generatedtemplate-templateconfiguration.html
//
type CfnGeneratedTemplate_TemplateConfigurationProperty struct {
	// The DeletionPolicy assigned to resources in the generated template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-generatedtemplate-templateconfiguration.html#cfn-cloudformation-generatedtemplate-templateconfiguration-deletionpolicy
	//
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// The UpdateReplacePolicy assigned to resources in the generated template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-generatedtemplate-templateconfiguration.html#cfn-cloudformation-generatedtemplate-templateconfiguration-updatereplacepolicy
	//
	UpdateReplacePolicy *string `field:"optional" json:"updateReplacePolicy" yaml:"updateReplacePolicy"`
}

