package awsopsworks


// Specifies the lifecycle event configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lifecycleEventConfigurationProperty := &LifecycleEventConfigurationProperty{
//   	ShutdownEventConfiguration: &ShutdownEventConfigurationProperty{
//   		DelayUntilElbConnectionsDrained: jsii.Boolean(false),
//   		ExecutionTimeout: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-lifecycleeventconfiguration.html
//
type CfnLayer_LifecycleEventConfigurationProperty struct {
	// The Shutdown event configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-lifecycleeventconfiguration.html#cfn-opsworks-layer-lifecycleeventconfiguration-shutdowneventconfiguration
	//
	ShutdownEventConfiguration interface{} `field:"optional" json:"shutdownEventConfiguration" yaml:"shutdownEventConfiguration"`
}

