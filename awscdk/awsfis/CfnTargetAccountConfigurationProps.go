package awsfis


// Properties for defining a `CfnTargetAccountConfiguration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTargetAccountConfigurationProps := &CfnTargetAccountConfigurationProps{
//   	AccountId: jsii.String("accountId"),
//   	ExperimentTemplateId: jsii.String("experimentTemplateId"),
//   	RoleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fis-targetaccountconfiguration.html
//
type CfnTargetAccountConfigurationProps struct {
	// The AWS account ID of the target account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fis-targetaccountconfiguration.html#cfn-fis-targetaccountconfiguration-accountid
	//
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// The ID of the experiment template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fis-targetaccountconfiguration.html#cfn-fis-targetaccountconfiguration-experimenttemplateid
	//
	ExperimentTemplateId *string `field:"required" json:"experimentTemplateId" yaml:"experimentTemplateId"`
	// The Amazon Resource Name (ARN) of an IAM role for the target account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fis-targetaccountconfiguration.html#cfn-fis-targetaccountconfiguration-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The description of the target account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fis-targetaccountconfiguration.html#cfn-fis-targetaccountconfiguration-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
}

