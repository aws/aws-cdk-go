package awsecs


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serviceConnectConfigurationProperty := &serviceConnectConfigurationProperty{
//   	enabled: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	logConfiguration: &logConfigurationProperty{
//   		logDriver: jsii.String("logDriver"),
//   		options: map[string]*string{
//   			"optionsKey": jsii.String("options"),
//   		},
//   		secretOptions: []interface{}{
//   			&secretProperty{
//   				name: jsii.String("name"),
//   				valueFrom: jsii.String("valueFrom"),
//   			},
//   		},
//   	},
//   	namespace: jsii.String("namespace"),
//   	services: []interface{}{
//   		&serviceConnectServiceProperty{
//   			portName: jsii.String("portName"),
//
//   			// the properties below are optional
//   			clientAliases: []interface{}{
//   				&serviceConnectClientAliasProperty{
//   					port: jsii.Number(123),
//
//   					// the properties below are optional
//   					dnsName: jsii.String("dnsName"),
//   				},
//   			},
//   			discoveryName: jsii.String("discoveryName"),
//   			ingressPortOverride: jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnService_ServiceConnectConfigurationProperty struct {
	// `CfnService.ServiceConnectConfigurationProperty.Enabled`.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// `CfnService.ServiceConnectConfigurationProperty.LogConfiguration`.
	LogConfiguration interface{} `field:"optional" json:"logConfiguration" yaml:"logConfiguration"`
	// `CfnService.ServiceConnectConfigurationProperty.Namespace`.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// `CfnService.ServiceConnectConfigurationProperty.Services`.
	Services interface{} `field:"optional" json:"services" yaml:"services"`
}

