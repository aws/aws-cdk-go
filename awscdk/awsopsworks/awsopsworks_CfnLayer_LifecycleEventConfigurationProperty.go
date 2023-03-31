package awsopsworks


// Specifies the lifecycle event configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lifecycleEventConfigurationProperty := &lifecycleEventConfigurationProperty{
//   	shutdownEventConfiguration: &shutdownEventConfigurationProperty{
//   		delayUntilElbConnectionsDrained: jsii.Boolean(false),
//   		executionTimeout: jsii.Number(123),
//   	},
//   }
//
type CfnLayer_LifecycleEventConfigurationProperty struct {
	// The Shutdown event configuration.
	ShutdownEventConfiguration interface{} `field:"optional" json:"shutdownEventConfiguration" yaml:"shutdownEventConfiguration"`
}

