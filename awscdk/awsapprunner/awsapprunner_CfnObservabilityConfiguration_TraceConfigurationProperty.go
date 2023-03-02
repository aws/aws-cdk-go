package awsapprunner


// Describes the configuration of the tracing feature within an AWS App Runner observability configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   traceConfigurationProperty := &traceConfigurationProperty{
//   	vendor: jsii.String("vendor"),
//   }
//
type CfnObservabilityConfiguration_TraceConfigurationProperty struct {
	// The implementation provider chosen for tracing App Runner services.
	Vendor *string `field:"required" json:"vendor" yaml:"vendor"`
}

