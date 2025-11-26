package previewawskinesisanalyticsv2mixins


// Describes parameters for how a Flink-based Kinesis Data Analytics application executes multiple tasks simultaneously.
//
// For more information about parallelism, see [Parallel Execution](https://docs.aws.amazon.com/https://nightlies.apache.org/flink/flink-docs-master/docs/dev/datastream/execution/parallel/) in the [Apache Flink Documentation](https://docs.aws.amazon.com/https://nightlies.apache.org/flink/flink-docs-master) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   parallelismConfigurationProperty := &ParallelismConfigurationProperty{
//   	AutoScalingEnabled: jsii.Boolean(false),
//   	ConfigurationType: jsii.String("configurationType"),
//   	Parallelism: jsii.Number(123),
//   	ParallelismPerKpu: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-parallelismconfiguration.html
//
type CfnApplicationPropsMixin_ParallelismConfigurationProperty struct {
	// Describes whether the Managed Service for Apache Flink service can increase the parallelism of the application in response to increased throughput.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-parallelismconfiguration.html#cfn-kinesisanalyticsv2-application-parallelismconfiguration-autoscalingenabled
	//
	AutoScalingEnabled interface{} `field:"optional" json:"autoScalingEnabled" yaml:"autoScalingEnabled"`
	// Describes whether the application uses the default parallelism for the Managed Service for Apache Flink service.
	//
	// You must set this property to `CUSTOM` in order to change your application's `AutoScalingEnabled` , `Parallelism` , or `ParallelismPerKPU` properties.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-parallelismconfiguration.html#cfn-kinesisanalyticsv2-application-parallelismconfiguration-configurationtype
	//
	ConfigurationType *string `field:"optional" json:"configurationType" yaml:"configurationType"`
	// Describes the initial number of parallel tasks that a Java-based Kinesis Data Analytics application can perform.
	//
	// The Kinesis Data Analytics service can increase this number automatically if [ParallelismConfiguration:AutoScalingEnabled](https://docs.aws.amazon.com/managed-flink/latest/apiv2/API_ParallelismConfiguration.html#kinesisanalytics-Type-ParallelismConfiguration-AutoScalingEnabled.html) is set to `true` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-parallelismconfiguration.html#cfn-kinesisanalyticsv2-application-parallelismconfiguration-parallelism
	//
	Parallelism *float64 `field:"optional" json:"parallelism" yaml:"parallelism"`
	// Describes the number of parallel tasks that a Java-based Kinesis Data Analytics application can perform per Kinesis Processing Unit (KPU) used by the application.
	//
	// For more information about KPUs, see [Amazon Kinesis Data Analytics Pricing](https://docs.aws.amazon.com/kinesis/data-analytics/pricing/) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-parallelismconfiguration.html#cfn-kinesisanalyticsv2-application-parallelismconfiguration-parallelismperkpu
	//
	ParallelismPerKpu *float64 `field:"optional" json:"parallelismPerKpu" yaml:"parallelismPerKpu"`
}

