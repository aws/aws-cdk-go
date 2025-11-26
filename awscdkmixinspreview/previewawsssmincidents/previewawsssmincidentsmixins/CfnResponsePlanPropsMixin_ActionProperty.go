package previewawsssmincidentsmixins


// The `Action` property type specifies the configuration to launch.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   actionProperty := &ActionProperty{
//   	SsmAutomation: &SsmAutomationProperty{
//   		DocumentName: jsii.String("documentName"),
//   		DocumentVersion: jsii.String("documentVersion"),
//   		DynamicParameters: []interface{}{
//   			&DynamicSsmParameterProperty{
//   				Key: jsii.String("key"),
//   				Value: &DynamicSsmParameterValueProperty{
//   					Variable: jsii.String("variable"),
//   				},
//   			},
//   		},
//   		Parameters: []interface{}{
//   			&SsmParameterProperty{
//   				Key: jsii.String("key"),
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   		},
//   		RoleArn: jsii.String("roleArn"),
//   		TargetAccount: jsii.String("targetAccount"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmincidents-responseplan-action.html
//
type CfnResponsePlanPropsMixin_ActionProperty struct {
	// Details about the Systems Manager automation document that will be used as a runbook during an incident.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmincidents-responseplan-action.html#cfn-ssmincidents-responseplan-action-ssmautomation
	//
	SsmAutomation interface{} `field:"optional" json:"ssmAutomation" yaml:"ssmAutomation"`
}

