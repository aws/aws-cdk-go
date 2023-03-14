package awsopsworks


// The Shutdown event configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   shutdownEventConfigurationProperty := &shutdownEventConfigurationProperty{
//   	delayUntilElbConnectionsDrained: jsii.Boolean(false),
//   	executionTimeout: jsii.Number(123),
//   }
//
type CfnLayer_ShutdownEventConfigurationProperty struct {
	// Whether to enable Elastic Load Balancing connection draining.
	//
	// For more information, see [Connection Draining](https://docs.aws.amazon.com/ElasticLoadBalancing/latest/DeveloperGuide/TerminologyandKeyConcepts.html#conn-drain)
	DelayUntilElbConnectionsDrained interface{} `field:"optional" json:"delayUntilElbConnectionsDrained" yaml:"delayUntilElbConnectionsDrained"`
	// The time, in seconds, that AWS OpsWorks Stacks waits after triggering a Shutdown event before shutting down an instance.
	ExecutionTimeout *float64 `field:"optional" json:"executionTimeout" yaml:"executionTimeout"`
}

