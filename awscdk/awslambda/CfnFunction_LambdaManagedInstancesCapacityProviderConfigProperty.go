package awslambda


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lambdaManagedInstancesCapacityProviderConfigProperty := &LambdaManagedInstancesCapacityProviderConfigProperty{
//   	CapacityProviderArn: jsii.String("capacityProviderArn"),
//
//   	// the properties below are optional
//   	ExecutionEnvironmentMemoryGiBPerVCpu: jsii.Number(123),
//   	PerExecutionEnvironmentMaxConcurrency: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-lambdamanagedinstancescapacityproviderconfig.html
//
type CfnFunction_LambdaManagedInstancesCapacityProviderConfigProperty struct {
	// The Amazon Resource Name (ARN) of the capacity provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-lambdamanagedinstancescapacityproviderconfig.html#cfn-lambda-function-lambdamanagedinstancescapacityproviderconfig-capacityproviderarn
	//
	CapacityProviderArn *string `field:"required" json:"capacityProviderArn" yaml:"capacityProviderArn"`
	// The amount of memory in GiB allocated per vCPU for execution environments.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-lambdamanagedinstancescapacityproviderconfig.html#cfn-lambda-function-lambdamanagedinstancescapacityproviderconfig-executionenvironmentmemorygibpervcpu
	//
	ExecutionEnvironmentMemoryGiBPerVCpu *float64 `field:"optional" json:"executionEnvironmentMemoryGiBPerVCpu" yaml:"executionEnvironmentMemoryGiBPerVCpu"`
	// The maximum number of concurrent execution environments that can run on each compute instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-lambdamanagedinstancescapacityproviderconfig.html#cfn-lambda-function-lambdamanagedinstancescapacityproviderconfig-perexecutionenvironmentmaxconcurrency
	//
	PerExecutionEnvironmentMaxConcurrency *float64 `field:"optional" json:"perExecutionEnvironmentMaxConcurrency" yaml:"perExecutionEnvironmentMaxConcurrency"`
}

