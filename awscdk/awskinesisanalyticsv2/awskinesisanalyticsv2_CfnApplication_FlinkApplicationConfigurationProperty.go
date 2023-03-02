package awskinesisanalyticsv2


// Describes configuration parameters for a Flink-based Kinesis Data Analytics application or a Studio notebook.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   flinkApplicationConfigurationProperty := &flinkApplicationConfigurationProperty{
//   	checkpointConfiguration: &checkpointConfigurationProperty{
//   		configurationType: jsii.String("configurationType"),
//
//   		// the properties below are optional
//   		checkpointingEnabled: jsii.Boolean(false),
//   		checkpointInterval: jsii.Number(123),
//   		minPauseBetweenCheckpoints: jsii.Number(123),
//   	},
//   	monitoringConfiguration: &monitoringConfigurationProperty{
//   		configurationType: jsii.String("configurationType"),
//
//   		// the properties below are optional
//   		logLevel: jsii.String("logLevel"),
//   		metricsLevel: jsii.String("metricsLevel"),
//   	},
//   	parallelismConfiguration: &parallelismConfigurationProperty{
//   		configurationType: jsii.String("configurationType"),
//
//   		// the properties below are optional
//   		autoScalingEnabled: jsii.Boolean(false),
//   		parallelism: jsii.Number(123),
//   		parallelismPerKpu: jsii.Number(123),
//   	},
//   }
//
type CfnApplication_FlinkApplicationConfigurationProperty struct {
	// Describes an application's checkpointing configuration.
	//
	// Checkpointing is the process of persisting application state for fault tolerance. For more information, see [Checkpoints for Fault Tolerance](https://docs.aws.amazon.com/https://ci.apache.org/projects/flink/flink-docs-release-1.8/concepts/programming-model.html#checkpoints-for-fault-tolerance) in the [Apache Flink Documentation](https://docs.aws.amazon.com/https://ci.apache.org/projects/flink/flink-docs-release-1.8/) .
	CheckpointConfiguration interface{} `field:"optional" json:"checkpointConfiguration" yaml:"checkpointConfiguration"`
	// Describes configuration parameters for Amazon CloudWatch logging for an application.
	MonitoringConfiguration interface{} `field:"optional" json:"monitoringConfiguration" yaml:"monitoringConfiguration"`
	// Describes parameters for how an application executes multiple tasks simultaneously.
	ParallelismConfiguration interface{} `field:"optional" json:"parallelismConfiguration" yaml:"parallelismConfiguration"`
}

