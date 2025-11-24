package mixinsawsfis


// Properties for CfnTargetAccountConfigurationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnTargetAccountConfigurationMixinProps := &CfnTargetAccountConfigurationMixinProps{
//   	AccountId: jsii.String("accountId"),
//   	Description: jsii.String("description"),
//   	ExperimentTemplateId: jsii.String("experimentTemplateId"),
//   	RoleArn: jsii.String("roleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fis-targetaccountconfiguration.html
//
type CfnTargetAccountConfigurationMixinProps struct {
	// The AWS account ID of the target account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fis-targetaccountconfiguration.html#cfn-fis-targetaccountconfiguration-accountid
	//
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// The description of the target account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fis-targetaccountconfiguration.html#cfn-fis-targetaccountconfiguration-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The ID of the experiment template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fis-targetaccountconfiguration.html#cfn-fis-targetaccountconfiguration-experimenttemplateid
	//
	ExperimentTemplateId *string `field:"optional" json:"experimentTemplateId" yaml:"experimentTemplateId"`
	// The Amazon Resource Name (ARN) of an IAM role for the target account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fis-targetaccountconfiguration.html#cfn-fis-targetaccountconfiguration-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

