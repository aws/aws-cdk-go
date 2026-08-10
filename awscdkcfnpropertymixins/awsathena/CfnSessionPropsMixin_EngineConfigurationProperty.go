package awsathena


// Contains engine data processing unit (DPU) configuration settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   engineConfigurationProperty := &EngineConfigurationProperty{
//   	AdditionalConfigs: map[string]*string{
//   		"additionalConfigsKey": jsii.String("additionalConfigs"),
//   	},
//   	CoordinatorDpuSize: jsii.Number(123),
//   	DefaultExecutorDpuSize: jsii.Number(123),
//   	MaxConcurrentDpus: jsii.Number(123),
//   	SparkProperties: map[string]*string{
//   		"sparkPropertiesKey": jsii.String("sparkProperties"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-session-engineconfiguration.html
//
type CfnSessionPropsMixin_EngineConfigurationProperty struct {
	// Contains additional notebook engine parameter mappings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-session-engineconfiguration.html#cfn-athena-session-engineconfiguration-additionalconfigs
	//
	AdditionalConfigs interface{} `field:"optional" json:"additionalConfigs" yaml:"additionalConfigs"`
	// The number of DPUs to use for the coordinator.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-session-engineconfiguration.html#cfn-athena-session-engineconfiguration-coordinatordpusize
	//
	CoordinatorDpuSize *float64 `field:"optional" json:"coordinatorDpuSize" yaml:"coordinatorDpuSize"`
	// The default number of DPUs to use for executors.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-session-engineconfiguration.html#cfn-athena-session-engineconfiguration-defaultexecutordpusize
	//
	DefaultExecutorDpuSize *float64 `field:"optional" json:"defaultExecutorDpuSize" yaml:"defaultExecutorDpuSize"`
	// The maximum number of DPUs that can run concurrently.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-session-engineconfiguration.html#cfn-athena-session-engineconfiguration-maxconcurrentdpus
	//
	MaxConcurrentDpus *float64 `field:"optional" json:"maxConcurrentDpus" yaml:"maxConcurrentDpus"`
	// Specifies custom jar files and Spark properties.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-session-engineconfiguration.html#cfn-athena-session-engineconfiguration-sparkproperties
	//
	SparkProperties interface{} `field:"optional" json:"sparkProperties" yaml:"sparkProperties"`
}

