package mixinsawsssmincidents


// The `SsmAutomation` property type specifies details about the Systems Manager Automation runbook that will be used as the runbook during an incident.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   ssmAutomationProperty := &SsmAutomationProperty{
//   	DocumentName: jsii.String("documentName"),
//   	DocumentVersion: jsii.String("documentVersion"),
//   	DynamicParameters: []interface{}{
//   		&DynamicSsmParameterProperty{
//   			Key: jsii.String("key"),
//   			Value: &DynamicSsmParameterValueProperty{
//   				Variable: jsii.String("variable"),
//   			},
//   		},
//   	},
//   	Parameters: []interface{}{
//   		&SsmParameterProperty{
//   			Key: jsii.String("key"),
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   	},
//   	RoleArn: jsii.String("roleArn"),
//   	TargetAccount: jsii.String("targetAccount"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmincidents-responseplan-ssmautomation.html
//
type CfnResponsePlanPropsMixin_SsmAutomationProperty struct {
	// The automation document's name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmincidents-responseplan-ssmautomation.html#cfn-ssmincidents-responseplan-ssmautomation-documentname
	//
	DocumentName *string `field:"optional" json:"documentName" yaml:"documentName"`
	// The version of the runbook to use when running.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmincidents-responseplan-ssmautomation.html#cfn-ssmincidents-responseplan-ssmautomation-documentversion
	//
	DocumentVersion *string `field:"optional" json:"documentVersion" yaml:"documentVersion"`
	// The key-value pairs to resolve dynamic parameter values when processing a Systems Manager Automation runbook.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmincidents-responseplan-ssmautomation.html#cfn-ssmincidents-responseplan-ssmautomation-dynamicparameters
	//
	DynamicParameters interface{} `field:"optional" json:"dynamicParameters" yaml:"dynamicParameters"`
	// The key-value pair parameters to use when running the runbook.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmincidents-responseplan-ssmautomation.html#cfn-ssmincidents-responseplan-ssmautomation-parameters
	//
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// The Amazon Resource Name (ARN) of the role that the automation document will assume when running commands.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmincidents-responseplan-ssmautomation.html#cfn-ssmincidents-responseplan-ssmautomation-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// The account that the automation document will be run in.
	//
	// This can be in either the management account or an application account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmincidents-responseplan-ssmautomation.html#cfn-ssmincidents-responseplan-ssmautomation-targetaccount
	//
	TargetAccount *string `field:"optional" json:"targetAccount" yaml:"targetAccount"`
}

