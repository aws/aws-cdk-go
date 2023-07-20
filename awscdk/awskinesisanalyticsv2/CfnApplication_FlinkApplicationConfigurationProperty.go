package awskinesisanalyticsv2


// Describes configuration parameters for a Flink-based Kinesis Data Analytics application or a Studio notebook.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   flinkApplicationConfigurationProperty := &FlinkApplicationConfigurationProperty{
//   	CheckpointConfiguration: &CheckpointConfigurationProperty{
//   		ConfigurationType: jsii.String("configurationType"),
//
//   		// the properties below are optional
//   		CheckpointingEnabled: jsii.Boolean(false),
//   		CheckpointInterval: jsii.Number(123),
//   		MinPauseBetweenCheckpoints: jsii.Number(123),
//   	},
//   	MonitoringConfiguration: &MonitoringConfigurationProperty{
//   		ConfigurationType: jsii.String("configurationType"),
//
//   		// the properties below are optional
//   		LogLevel: jsii.String("logLevel"),
//   		MetricsLevel: jsii.String("metricsLevel"),
//   	},
//   	ParallelismConfiguration: &ParallelismConfigurationProperty{
//   		ConfigurationType: jsii.String("configurationType"),
//
//   		// the properties below are optional
//   		AutoScalingEnabled: jsii.Boolean(false),
//   		Parallelism: jsii.Number(123),
//   		ParallelismPerKpu: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-flinkapplicationconfiguration.html
//
type CfnApplication_FlinkApplicationConfigurationProperty struct {
	// Describes an application's checkpointing configuration.
	//
	// Checkpointing is the process of persisting application state for fault tolerance. For more information, see [Checkpoints for Fault Tolerance](https://docs.aws.amazon.com/https://ci.apache.org/projects/flink/flink-docs-release-1.8/concepts/programming-model.html#checkpoints-for-fault-tolerance) in the [Apache Flink Documentation](https://docs.aws.amazon.com/https://ci.apache.org/projects/flink/flink-docs-release-1.8/) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-flinkapplicationconfiguration.html#cfn-kinesisanalyticsv2-application-flinkapplicationconfiguration-checkpointconfiguration
	//
	CheckpointConfiguration interface{} `field:"optional" json:"checkpointConfiguration" yaml:"checkpointConfiguration"`
	// Describes configuration parameters for Amazon CloudWatch logging for an application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-flinkapplicationconfiguration.html#cfn-kinesisanalyticsv2-application-flinkapplicationconfiguration-monitoringconfiguration
	//
	MonitoringConfiguration interface{} `field:"optional" json:"monitoringConfiguration" yaml:"monitoringConfiguration"`
	// Describes parameters for how an application executes multiple tasks simultaneously.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-flinkapplicationconfiguration.html#cfn-kinesisanalyticsv2-application-flinkapplicationconfiguration-parallelismconfiguration
	//
	ParallelismConfiguration interface{} `field:"optional" json:"parallelismConfiguration" yaml:"parallelismConfiguration"`
}

