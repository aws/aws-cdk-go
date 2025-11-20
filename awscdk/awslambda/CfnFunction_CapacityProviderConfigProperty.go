package awslambda


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   capacityProviderConfigProperty := &CapacityProviderConfigProperty{
//   	Ec2ManagedInstancesCapacityProviderConfig: &EC2ManagedInstancesCapacityProviderConfigProperty{
//   		CapacityProviderArn: jsii.String("capacityProviderArn"),
//
//   		// the properties below are optional
//   		ExecutionEnvironmentMaxConcurrency: jsii.Number(123),
//   		ExecutionEnvironmentMemoryGiBPerVCpu: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-capacityproviderconfig.html
//
type CfnFunction_CapacityProviderConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-capacityproviderconfig.html#cfn-lambda-function-capacityproviderconfig-ec2managedinstancescapacityproviderconfig
	//
	Ec2ManagedInstancesCapacityProviderConfig interface{} `field:"required" json:"ec2ManagedInstancesCapacityProviderConfig" yaml:"ec2ManagedInstancesCapacityProviderConfig"`
}

