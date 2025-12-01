package awslambda


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   capacityProviderConfigProperty := &CapacityProviderConfigProperty{
//   	LambdaManagedInstancesCapacityProviderConfig: &LambdaManagedInstancesCapacityProviderConfigProperty{
//   		CapacityProviderArn: jsii.String("capacityProviderArn"),
//
//   		// the properties below are optional
//   		ExecutionEnvironmentMemoryGiBPerVCpu: jsii.Number(123),
//   		PerExecutionEnvironmentMaxConcurrency: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-capacityproviderconfig.html
//
type CfnFunction_CapacityProviderConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-capacityproviderconfig.html#cfn-lambda-function-capacityproviderconfig-lambdamanagedinstancescapacityproviderconfig
	//
	LambdaManagedInstancesCapacityProviderConfig interface{} `field:"required" json:"lambdaManagedInstancesCapacityProviderConfig" yaml:"lambdaManagedInstancesCapacityProviderConfig"`
}

