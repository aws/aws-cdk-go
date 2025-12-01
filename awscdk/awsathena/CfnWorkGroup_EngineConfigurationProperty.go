package awsathena


// The engine configuration for the workgroup, which includes the minimum/maximum number of Data Processing Units (DPU) that queries should use when running in provisioned capacity.
//
// If not specified, Athena uses default values (Default value for min is 4 and for max is Minimum of 124 and allocated DPUs).
//
// To specify DPU values for PC queries the WG containing EngineConfiguration should have the following values: The name of the Classifications should be `athena-query-engine-properties` , with the only allowed properties as `max-dpu-count` and `min-dpu-count` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   engineConfigurationProperty := &EngineConfigurationProperty{
//   	AdditionalConfigs: map[string]*string{
//   		"additionalConfigsKey": jsii.String("additionalConfigs"),
//   	},
//   	Classifications: []interface{}{
//   		&ClassificationProperty{
//   			Name: jsii.String("name"),
//   			Properties: map[string]*string{
//   				"propertiesKey": jsii.String("properties"),
//   			},
//   		},
//   	},
//   	CoordinatorDpuSize: jsii.Number(123),
//   	DefaultExecutorDpuSize: jsii.Number(123),
//   	MaxConcurrentDpus: jsii.Number(123),
//   	SparkProperties: map[string]*string{
//   		"sparkPropertiesKey": jsii.String("sparkProperties"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-engineconfiguration.html
//
type CfnWorkGroup_EngineConfigurationProperty struct {
	// Contains additional notebook engine `MAP<string, string>` parameter mappings in the form of key-value pairs.
	//
	// To specify an Athena notebook that the Jupyter server will download and serve, specify a value for the `StartSessionRequest$NotebookVersion` field, and then add a key named `NotebookId` to `AdditionalConfigs` that has the value of the Athena notebook ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-engineconfiguration.html#cfn-athena-workgroup-engineconfiguration-additionalconfigs
	//
	AdditionalConfigs interface{} `field:"optional" json:"additionalConfigs" yaml:"additionalConfigs"`
	// The configuration classifications that can be specified for the engine.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-engineconfiguration.html#cfn-athena-workgroup-engineconfiguration-classifications
	//
	Classifications interface{} `field:"optional" json:"classifications" yaml:"classifications"`
	// The number of DPUs to use for the coordinator.
	//
	// A coordinator is a special executor that orchestrates processing work and manages other executors in a notebook session. The default is 1.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-engineconfiguration.html#cfn-athena-workgroup-engineconfiguration-coordinatordpusize
	//
	CoordinatorDpuSize *float64 `field:"optional" json:"coordinatorDpuSize" yaml:"coordinatorDpuSize"`
	// The default number of DPUs to use for executors.
	//
	// An executor is the smallest unit of compute that a notebook session can request from Athena. The default is 1.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-engineconfiguration.html#cfn-athena-workgroup-engineconfiguration-defaultexecutordpusize
	//
	DefaultExecutorDpuSize *float64 `field:"optional" json:"defaultExecutorDpuSize" yaml:"defaultExecutorDpuSize"`
	// The maximum number of DPUs that can run concurrently.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-engineconfiguration.html#cfn-athena-workgroup-engineconfiguration-maxconcurrentdpus
	//
	MaxConcurrentDpus *float64 `field:"optional" json:"maxConcurrentDpus" yaml:"maxConcurrentDpus"`
	// Specifies custom jar files and Spark properties for use cases like cluster encryption, table formats, and general Spark tuning.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-engineconfiguration.html#cfn-athena-workgroup-engineconfiguration-sparkproperties
	//
	SparkProperties interface{} `field:"optional" json:"sparkProperties" yaml:"sparkProperties"`
}

