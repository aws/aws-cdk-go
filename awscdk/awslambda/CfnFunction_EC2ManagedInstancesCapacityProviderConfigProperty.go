package awslambda


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eC2ManagedInstancesCapacityProviderConfigProperty := &EC2ManagedInstancesCapacityProviderConfigProperty{
//   	CapacityProviderArn: jsii.String("capacityProviderArn"),
//
//   	// the properties below are optional
//   	ExecutionEnvironmentMaxConcurrency: jsii.Number(123),
//   	ExecutionEnvironmentMemoryGiBPerVCpu: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-ec2managedinstancescapacityproviderconfig.html
//
type CfnFunction_EC2ManagedInstancesCapacityProviderConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-ec2managedinstancescapacityproviderconfig.html#cfn-lambda-function-ec2managedinstancescapacityproviderconfig-capacityproviderarn
	//
	CapacityProviderArn *string `field:"required" json:"capacityProviderArn" yaml:"capacityProviderArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-ec2managedinstancescapacityproviderconfig.html#cfn-lambda-function-ec2managedinstancescapacityproviderconfig-executionenvironmentmaxconcurrency
	//
	ExecutionEnvironmentMaxConcurrency *float64 `field:"optional" json:"executionEnvironmentMaxConcurrency" yaml:"executionEnvironmentMaxConcurrency"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-ec2managedinstancescapacityproviderconfig.html#cfn-lambda-function-ec2managedinstancescapacityproviderconfig-executionenvironmentmemorygibpervcpu
	//
	ExecutionEnvironmentMemoryGiBPerVCpu *float64 `field:"optional" json:"executionEnvironmentMemoryGiBPerVCpu" yaml:"executionEnvironmentMemoryGiBPerVCpu"`
}

