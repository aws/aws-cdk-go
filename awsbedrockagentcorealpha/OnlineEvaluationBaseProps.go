package awsbedrockagentcorealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Base properties for creating an OnlineEvaluationConfig.
//
// The actual OnlineEvaluationProps is defined in online-evaluation-config.ts
// to avoid circular dependencies.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_agentcore_alpha "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var executionStatus ExecutionStatus
//   var filterOperator FilterOperator
//   var filterValue FilterValue
//   var role Role
//
//   onlineEvaluationBaseProps := &OnlineEvaluationBaseProps{
//   	OnlineEvaluationConfigName: jsii.String("onlineEvaluationConfigName"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	ExecutionRole: role,
//   	ExecutionStatus: executionStatus,
//   	Filters: []FilterConfig{
//   		&FilterConfig{
//   			Key: jsii.String("key"),
//   			Operator: filterOperator,
//   			Value: filterValue,
//   		},
//   	},
//   	SamplingPercentage: jsii.Number(123),
//   	SessionTimeout: cdk.Duration_Minutes(jsii.Number(30)),
//   }
//
// Experimental.
type OnlineEvaluationBaseProps struct {
	// The name of the online evaluation configuration.
	//
	// Must be unique within your account. Valid characters are a-z, A-Z, 0-9, _ (underscore).
	// Must start with a letter and can be up to 48 characters long.
	// Experimental.
	OnlineEvaluationConfigName *string `field:"required" json:"onlineEvaluationConfigName" yaml:"onlineEvaluationConfigName"`
	// The description of the online evaluation configuration.
	// Default: - No description.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The IAM role that provides permissions for the evaluation to access AWS services.
	//
	// If not provided, a role will be created automatically with the required permissions
	// including cross-region Bedrock model invocation (to support cross-region inference
	// profiles). For strict cost controls or data residency compliance, provide a custom
	// role with region-scoped permissions.
	// Default: - A new role will be created.
	//
	// Experimental.
	ExecutionRole awsiam.IRole `field:"optional" json:"executionRole" yaml:"executionRole"`
	// The execution status of the online evaluation configuration.
	//
	// Controls whether the evaluation actively processes agent traces.
	// Default: ExecutionStatus.ENABLED
	//
	// Experimental.
	ExecutionStatus ExecutionStatus `field:"optional" json:"executionStatus" yaml:"executionStatus"`
	// The list of filters that determine which agent traces should be evaluated.
	// Default: - No filters (evaluate all sampled traces).
	//
	// Experimental.
	Filters *[]*FilterConfig `field:"optional" json:"filters" yaml:"filters"`
	// The percentage of agent traces to sample for evaluation.
	// Default: 10.
	//
	// Experimental.
	SamplingPercentage *float64 `field:"optional" json:"samplingPercentage" yaml:"samplingPercentage"`
	// The duration of inactivity after which an agent session is considered complete and ready for evaluation.
	//
	// Must be between 1 minute and 1440 minutes (24 hours).
	// Default: Duration.minutes(15)
	//
	// Experimental.
	SessionTimeout awscdk.Duration `field:"optional" json:"sessionTimeout" yaml:"sessionTimeout"`
}

