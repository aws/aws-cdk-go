package awskinesisanalyticsv2


// Describes an application's checkpointing configuration.
//
// Checkpointing is the process of persisting application state for fault tolerance. For more information, see [Checkpoints for Fault Tolerance](https://docs.aws.amazon.com/https://nightlies.apache.org/flink/flink-docs-master/docs/dev/datastream/fault-tolerance/checkpointing/) in the [Apache Flink Documentation](https://docs.aws.amazon.com/https://nightlies.apache.org/flink/flink-docs-master) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   checkpointConfigurationProperty := &CheckpointConfigurationProperty{
//   	ConfigurationType: jsii.String("configurationType"),
//
//   	// the properties below are optional
//   	CheckpointingEnabled: jsii.Boolean(false),
//   	CheckpointInterval: jsii.Number(123),
//   	MinPauseBetweenCheckpoints: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-checkpointconfiguration.html
//
type CfnApplication_CheckpointConfigurationProperty struct {
	// Describes whether the application uses Managed Service for Apache Flink' default checkpointing behavior.
	//
	// You must set this property to `CUSTOM` in order to set the `CheckpointingEnabled` , `CheckpointInterval` , or `MinPauseBetweenCheckpoints` parameters.
	//
	// > If this value is set to `DEFAULT` , the application will use the following values, even if they are set to other values using APIs or application code:
	// >
	// > - *CheckpointingEnabled:* true
	// > - *CheckpointInterval:* 60000
	// > - *MinPauseBetweenCheckpoints:* 5000.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-checkpointconfiguration.html#cfn-kinesisanalyticsv2-application-checkpointconfiguration-configurationtype
	//
	ConfigurationType *string `field:"required" json:"configurationType" yaml:"configurationType"`
	// Describes whether checkpointing is enabled for a Managed Service for Apache Flink application.
	//
	// > If `CheckpointConfiguration.ConfigurationType` is `DEFAULT` , the application will use a `CheckpointingEnabled` value of `true` , even if this value is set to another value using this API or in application code.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-checkpointconfiguration.html#cfn-kinesisanalyticsv2-application-checkpointconfiguration-checkpointingenabled
	//
	CheckpointingEnabled interface{} `field:"optional" json:"checkpointingEnabled" yaml:"checkpointingEnabled"`
	// Describes the interval in milliseconds between checkpoint operations.
	//
	// > If `CheckpointConfiguration.ConfigurationType` is `DEFAULT` , the application will use a `CheckpointInterval` value of 60000, even if this value is set to another value using this API or in application code.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-checkpointconfiguration.html#cfn-kinesisanalyticsv2-application-checkpointconfiguration-checkpointinterval
	//
	CheckpointInterval *float64 `field:"optional" json:"checkpointInterval" yaml:"checkpointInterval"`
	// Describes the minimum time in milliseconds after a checkpoint operation completes that a new checkpoint operation can start.
	//
	// If a checkpoint operation takes longer than the `CheckpointInterval` , the application otherwise performs continual checkpoint operations. For more information, see [Tuning Checkpointing](https://docs.aws.amazon.com/https://nightlies.apache.org/flink/flink-docs-master/docs/ops/state/large_state_tuning/#tuning-checkpointing) in the [Apache Flink Documentation](https://docs.aws.amazon.com/https://nightlies.apache.org/flink/flink-docs-master) .
	//
	// > If `CheckpointConfiguration.ConfigurationType` is `DEFAULT` , the application will use a `MinPauseBetweenCheckpoints` value of 5000, even if this value is set using this API or in application code.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-checkpointconfiguration.html#cfn-kinesisanalyticsv2-application-checkpointconfiguration-minpausebetweencheckpoints
	//
	MinPauseBetweenCheckpoints *float64 `field:"optional" json:"minPauseBetweenCheckpoints" yaml:"minPauseBetweenCheckpoints"`
}

