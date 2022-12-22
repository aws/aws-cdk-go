package awsapprunner

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnObservabilityConfiguration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnObservabilityConfigurationProps := &cfnObservabilityConfigurationProps{
//   	observabilityConfigurationName: jsii.String("observabilityConfigurationName"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	traceConfiguration: &traceConfigurationProperty{
//   		vendor: jsii.String("vendor"),
//   	},
//   }
//
type CfnObservabilityConfigurationProps struct {
	// A name for the observability configuration.
	//
	// When you use it for the first time in an AWS Region , App Runner creates revision number `1` of this name. When you use the same name in subsequent calls, App Runner creates incremental revisions of the configuration.
	//
	// > The name `DefaultConfiguration` is reserved. You can't use it to create a new observability configuration, and you can't create a revision of it.
	// >
	// > When you want to use your own observability configuration for your App Runner service, *create a configuration with a different name* , and then provide it when you create or update your service.
	//
	// If you don't specify a name, AWS CloudFormation generates a name for your observability configuration.
	ObservabilityConfigurationName *string `field:"optional" json:"observabilityConfigurationName" yaml:"observabilityConfigurationName"`
	// A list of metadata items that you can associate with your observability configuration resource.
	//
	// A tag is a key-value pair.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The configuration of the tracing feature within this observability configuration.
	//
	// If you don't specify it, App Runner doesn't enable tracing.
	TraceConfiguration interface{} `field:"optional" json:"traceConfiguration" yaml:"traceConfiguration"`
}

