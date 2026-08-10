package awsathena


// Properties for CfnSessionPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnSessionMixinProps := &CfnSessionMixinProps{
//   	EngineConfiguration: &EngineConfigurationProperty{
//   		AdditionalConfigs: map[string]*string{
//   			"additionalConfigsKey": jsii.String("additionalConfigs"),
//   		},
//   		CoordinatorDpuSize: jsii.Number(123),
//   		DefaultExecutorDpuSize: jsii.Number(123),
//   		MaxConcurrentDpus: jsii.Number(123),
//   		SparkProperties: map[string]*string{
//   			"sparkPropertiesKey": jsii.String("sparkProperties"),
//   		},
//   	},
//   	ExecutionRole: jsii.String("executionRole"),
//   	WorkGroup: jsii.String("workGroup"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-session.html
//
type CfnSessionMixinProps struct {
	// Contains engine data processing unit (DPU) configuration settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-session.html#cfn-athena-session-engineconfiguration
	//
	EngineConfiguration interface{} `field:"optional" json:"engineConfiguration" yaml:"engineConfiguration"`
	// The ARN of the execution role used to access user resources for Spark sessions and Identity Center enabled workgroups.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-session.html#cfn-athena-session-executionrole
	//
	ExecutionRole *string `field:"optional" json:"executionRole" yaml:"executionRole"`
	// The workgroup to which the session belongs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-session.html#cfn-athena-session-workgroup
	//
	WorkGroup *string `field:"optional" json:"workGroup" yaml:"workGroup"`
}

