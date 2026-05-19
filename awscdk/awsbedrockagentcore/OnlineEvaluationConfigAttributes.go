package awsbedrockagentcore


// Attributes for importing an existing OnlineEvaluationConfig.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   onlineEvaluationConfigAttributes := &OnlineEvaluationConfigAttributes{
//   	OnlineEvaluationConfigArn: jsii.String("onlineEvaluationConfigArn"),
//   	OnlineEvaluationConfigId: jsii.String("onlineEvaluationConfigId"),
//   	OnlineEvaluationConfigName: jsii.String("onlineEvaluationConfigName"),
//
//   	// the properties below are optional
//   	ExecutionRoleArn: jsii.String("executionRoleArn"),
//   }
//
type OnlineEvaluationConfigAttributes struct {
	// The ARN of the online evaluation configuration.
	OnlineEvaluationConfigArn *string `field:"required" json:"onlineEvaluationConfigArn" yaml:"onlineEvaluationConfigArn"`
	// The ID of the online evaluation configuration.
	OnlineEvaluationConfigId *string `field:"required" json:"onlineEvaluationConfigId" yaml:"onlineEvaluationConfigId"`
	// The name of the online evaluation configuration.
	OnlineEvaluationConfigName *string `field:"required" json:"onlineEvaluationConfigName" yaml:"onlineEvaluationConfigName"`
	// The ARN of the IAM execution role.
	// Default: - No role ARN provided.
	//
	ExecutionRoleArn *string `field:"optional" json:"executionRoleArn" yaml:"executionRoleArn"`
}

