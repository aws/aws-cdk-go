package awsecs


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   serviceConnectContainerConfigurationProperty := &ServiceConnectContainerConfigurationProperty{
//   	LogLevel: jsii.String("logLevel"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-serviceconnectcontainerconfiguration.html
//
type CfnServicePropsMixin_ServiceConnectContainerConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-serviceconnectcontainerconfiguration.html#cfn-ecs-service-serviceconnectcontainerconfiguration-loglevel
	//
	LogLevel *string `field:"optional" json:"logLevel" yaml:"logLevel"`
}

