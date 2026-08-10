package awsathena


// Properties for defining a `CfnSession`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSessionProps := &CfnSessionProps{
//   	EngineConfiguration: &EngineConfigurationProperty{
//   		MaxConcurrentDpus: jsii.Number(123),
//
//   		// the properties below are optional
//   		AdditionalConfigs: map[string]*string{
//   			"additionalConfigsKey": jsii.String("additionalConfigs"),
//   		},
//   		CoordinatorDpuSize: jsii.Number(123),
//   		DefaultExecutorDpuSize: jsii.Number(123),
//   		SparkProperties: map[string]*string{
//   			"sparkPropertiesKey": jsii.String("sparkProperties"),
//   		},
//   	},
//   	WorkGroup: jsii.String("workGroup"),
//
//   	// the properties below are optional
//   	ExecutionRole: jsii.String("executionRole"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-session.html
//
type CfnSessionProps struct {
	// Contains engine data processing unit (DPU) configuration settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-session.html#cfn-athena-session-engineconfiguration
	//
	EngineConfiguration interface{} `field:"required" json:"engineConfiguration" yaml:"engineConfiguration"`
	// The workgroup to which the session belongs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-session.html#cfn-athena-session-workgroup
	//
	WorkGroup *string `field:"required" json:"workGroup" yaml:"workGroup"`
	// The ARN of the execution role used to access user resources for Spark sessions and Identity Center enabled workgroups.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-session.html#cfn-athena-session-executionrole
	//
	ExecutionRole *string `field:"optional" json:"executionRole" yaml:"executionRole"`
}

