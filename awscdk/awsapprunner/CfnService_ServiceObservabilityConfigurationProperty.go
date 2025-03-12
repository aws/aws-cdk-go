package awsapprunner


// Describes the observability configuration of an AWS App Runner service.
//
// These are additional observability features, like tracing, that you choose to enable. They're configured in a separate resource that you associate with your service.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serviceObservabilityConfigurationProperty := &ServiceObservabilityConfigurationProperty{
//   	ObservabilityEnabled: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	ObservabilityConfigurationArn: jsii.String("observabilityConfigurationArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-serviceobservabilityconfiguration.html
//
type CfnService_ServiceObservabilityConfigurationProperty struct {
	// When `true` , an observability configuration resource is associated with the service, and an `ObservabilityConfigurationArn` is specified.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-serviceobservabilityconfiguration.html#cfn-apprunner-service-serviceobservabilityconfiguration-observabilityenabled
	//
	ObservabilityEnabled interface{} `field:"required" json:"observabilityEnabled" yaml:"observabilityEnabled"`
	// The Amazon Resource Name (ARN) of the observability configuration that is associated with the service.
	//
	// Specified only when `ObservabilityEnabled` is `true` .
	//
	// Specify an ARN with a name and a revision number to associate that revision. For example: `arn:aws:apprunner:us-east-1:123456789012:observabilityconfiguration/xray-tracing/3`
	//
	// Specify just the name to associate the latest revision. For example: `arn:aws:apprunner:us-east-1:123456789012:observabilityconfiguration/xray-tracing`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-serviceobservabilityconfiguration.html#cfn-apprunner-service-serviceobservabilityconfiguration-observabilityconfigurationarn
	//
	ObservabilityConfigurationArn *string `field:"optional" json:"observabilityConfigurationArn" yaml:"observabilityConfigurationArn"`
}

