package mixinsawsopsworks


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
type CfnLayerPropsMixin_LifecycleEventConfigurationProperty struct {
	// The Shutdown event configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-lifecycleeventconfiguration.html#cfn-opsworks-layer-lifecycleeventconfiguration-shutdowneventconfiguration
	//
	ShutdownEventConfiguration interface{} `field:"optional" json:"shutdownEventConfiguration" yaml:"shutdownEventConfiguration"`
}

