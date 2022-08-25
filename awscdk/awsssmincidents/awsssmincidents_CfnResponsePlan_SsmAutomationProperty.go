package awsssmincidents


// The `SsmAutomation` property type specifies details about the Systems Manager automation document that will be used as a runbook during an incident.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ssmAutomationProperty := &ssmAutomationProperty{
//   	documentName: jsii.String("documentName"),
//   	roleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	documentVersion: jsii.String("documentVersion"),
//   	dynamicParameters: []interface{}{
//   		&dynamicSsmParameterProperty{
//   			key: jsii.String("key"),
//   			value: &dynamicSsmParameterValueProperty{
//   				variable: jsii.String("variable"),
//   			},
//   		},
//   	},
//   	parameters: []interface{}{
//   		&ssmParameterProperty{
//   			key: jsii.String("key"),
//   			values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   	},
//   	targetAccount: jsii.String("targetAccount"),
//   }
//
type CfnResponsePlan_SsmAutomationProperty struct {
	// The automation document's name.
	DocumentName *string `field:"required" json:"documentName" yaml:"documentName"`
	// The Amazon Resource Name (ARN) of the role that the automation document will assume when running commands.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The automation document's version to use when running.
	DocumentVersion *string `field:"optional" json:"documentVersion" yaml:"documentVersion"`
	// `CfnResponsePlan.SsmAutomationProperty.DynamicParameters`.
	DynamicParameters interface{} `field:"optional" json:"dynamicParameters" yaml:"dynamicParameters"`
	// The key-value pair parameters to use when running the automation document.
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// The account that the automation document will be run in.
	//
	// This can be in either the management account or an application account.
	TargetAccount *string `field:"optional" json:"targetAccount" yaml:"targetAccount"`
}

