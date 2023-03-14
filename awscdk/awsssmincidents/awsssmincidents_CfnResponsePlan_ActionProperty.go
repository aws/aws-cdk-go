package awsssmincidents


// The `Action` property type specifies the configuration to launch.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   actionProperty := &actionProperty{
//   	ssmAutomation: &ssmAutomationProperty{
//   		documentName: jsii.String("documentName"),
//   		roleArn: jsii.String("roleArn"),
//
//   		// the properties below are optional
//   		documentVersion: jsii.String("documentVersion"),
//   		dynamicParameters: []interface{}{
//   			&dynamicSsmParameterProperty{
//   				key: jsii.String("key"),
//   				value: &dynamicSsmParameterValueProperty{
//   					variable: jsii.String("variable"),
//   				},
//   			},
//   		},
//   		parameters: []interface{}{
//   			&ssmParameterProperty{
//   				key: jsii.String("key"),
//   				values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   		},
//   		targetAccount: jsii.String("targetAccount"),
//   	},
//   }
//
type CfnResponsePlan_ActionProperty struct {
	// Details about the Systems Manager automation document that will be used as a runbook during an incident.
	SsmAutomation interface{} `field:"optional" json:"ssmAutomation" yaml:"ssmAutomation"`
}

