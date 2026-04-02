package awsbedrockagentcore

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnOnlineEvaluationConfig`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnOnlineEvaluationConfigProps := &CfnOnlineEvaluationConfigProps{
//   	DataSourceConfig: &DataSourceConfigProperty{
//   		CloudWatchLogs: &CloudWatchLogsInputConfigProperty{
//   			LogGroupNames: []*string{
//   				jsii.String("logGroupNames"),
//   			},
//   			ServiceNames: []*string{
//   				jsii.String("serviceNames"),
//   			},
//   		},
//   	},
//   	EvaluationExecutionRoleArn: jsii.String("evaluationExecutionRoleArn"),
//   	Evaluators: []interface{}{
//   		&EvaluatorReferenceProperty{
//   			EvaluatorId: jsii.String("evaluatorId"),
//   		},
//   	},
//   	OnlineEvaluationConfigName: jsii.String("onlineEvaluationConfigName"),
//   	Rule: &RuleProperty{
//   		SamplingConfig: &SamplingConfigProperty{
//   			SamplingPercentage: jsii.Number(123),
//   		},
//
//   		// the properties below are optional
//   		Filters: []interface{}{
//   			&FilterProperty{
//   				Key: jsii.String("key"),
//   				Operator: jsii.String("operator"),
//   				Value: &FilterValueProperty{
//   					BooleanValue: jsii.Boolean(false),
//   					DoubleValue: jsii.Number(123),
//   					StringValue: jsii.String("stringValue"),
//   				},
//   			},
//   		},
//   		SessionConfig: &SessionConfigProperty{
//   			SessionTimeoutMinutes: jsii.Number(123),
//   		},
//   	},
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	ExecutionStatus: jsii.String("executionStatus"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-onlineevaluationconfig.html
//
type CfnOnlineEvaluationConfigProps struct {
	// The configuration that specifies where to read agent traces for online evaluation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-onlineevaluationconfig.html#cfn-bedrockagentcore-onlineevaluationconfig-datasourceconfig
	//
	DataSourceConfig interface{} `field:"required" json:"dataSourceConfig" yaml:"dataSourceConfig"`
	// The Amazon Resource Name (ARN) of the IAM role that grants permissions for evaluation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-onlineevaluationconfig.html#cfn-bedrockagentcore-onlineevaluationconfig-evaluationexecutionrolearn
	//
	EvaluationExecutionRoleArn *string `field:"required" json:"evaluationExecutionRoleArn" yaml:"evaluationExecutionRoleArn"`
	// The list of evaluators to apply during online evaluation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-onlineevaluationconfig.html#cfn-bedrockagentcore-onlineevaluationconfig-evaluators
	//
	Evaluators interface{} `field:"required" json:"evaluators" yaml:"evaluators"`
	// The name of the online evaluation configuration.
	//
	// Must be unique within your account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-onlineevaluationconfig.html#cfn-bedrockagentcore-onlineevaluationconfig-onlineevaluationconfigname
	//
	OnlineEvaluationConfigName *string `field:"required" json:"onlineEvaluationConfigName" yaml:"onlineEvaluationConfigName"`
	// The evaluation rule that defines sampling configuration, filtering criteria, and session detection settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-onlineevaluationconfig.html#cfn-bedrockagentcore-onlineevaluationconfig-rule
	//
	Rule interface{} `field:"required" json:"rule" yaml:"rule"`
	// The description of the online evaluation configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-onlineevaluationconfig.html#cfn-bedrockagentcore-onlineevaluationconfig-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-onlineevaluationconfig.html#cfn-bedrockagentcore-onlineevaluationconfig-executionstatus
	//
	ExecutionStatus *string `field:"optional" json:"executionStatus" yaml:"executionStatus"`
	// A list of tags to assign to the online evaluation configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-onlineevaluationconfig.html#cfn-bedrockagentcore-onlineevaluationconfig-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

