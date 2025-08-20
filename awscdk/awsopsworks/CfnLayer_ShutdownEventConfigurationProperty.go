package awsopsworks


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   shutdownEventConfigurationProperty := &ShutdownEventConfigurationProperty{
//   	DelayUntilElbConnectionsDrained: jsii.Boolean(false),
//   	ExecutionTimeout: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-shutdowneventconfiguration.html
//
type CfnLayer_ShutdownEventConfigurationProperty struct {
	// Whether to enable Elastic Load Balancing connection draining.
	//
	// For more information, see [Connection Draining](https://docs.aws.amazon.com/ElasticLoadBalancing/latest/DeveloperGuide/TerminologyandKeyConcepts.html#conn-drain)
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-shutdowneventconfiguration.html#cfn-opsworks-layer-shutdowneventconfiguration-delayuntilelbconnectionsdrained
	//
	DelayUntilElbConnectionsDrained interface{} `field:"optional" json:"delayUntilElbConnectionsDrained" yaml:"delayUntilElbConnectionsDrained"`
	// The time, in seconds, that OpsWorks Stacks waits after triggering a Shutdown event before shutting down an instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-shutdowneventconfiguration.html#cfn-opsworks-layer-shutdowneventconfiguration-executiontimeout
	//
	ExecutionTimeout *float64 `field:"optional" json:"executionTimeout" yaml:"executionTimeout"`
}

