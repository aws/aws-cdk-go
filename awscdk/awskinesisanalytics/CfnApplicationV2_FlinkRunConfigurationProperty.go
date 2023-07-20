package awskinesisanalytics


// Describes the starting parameters for a Flink-based Kinesis Data Analytics application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   flinkRunConfigurationProperty := &FlinkRunConfigurationProperty{
//   	AllowNonRestoredState: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-flinkrunconfiguration.html
//
type CfnApplicationV2_FlinkRunConfigurationProperty struct {
	// When restoring from a snapshot, specifies whether the runtime is allowed to skip a state that cannot be mapped to the new program.
	//
	// This will happen if the program is updated between snapshots to remove stateful parameters, and state data in the snapshot no longer corresponds to valid application data. For more information, see [Allowing Non-Restored State](https://docs.aws.amazon.com/https://ci.apache.org/projects/flink/flink-docs-release-1.8/ops/state/savepoints.html#allowing-non-restored-state) in the [Apache Flink documentation](https://docs.aws.amazon.com/https://ci.apache.org/projects/flink/flink-docs-release-1.8/) .
	//
	// > This value defaults to `false` . If you update your application without specifying this parameter, `AllowNonRestoredState` will be set to `false` , even if it was previously set to `true` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-flinkrunconfiguration.html#cfn-kinesisanalyticsv2-application-flinkrunconfiguration-allownonrestoredstate
	//
	AllowNonRestoredState interface{} `field:"optional" json:"allowNonRestoredState" yaml:"allowNonRestoredState"`
}

