package awsappmesh


// All Properties for Envoy Access logs for mesh endpoints.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accessLogConfig := &accessLogConfig{
//   	virtualGatewayAccessLog: &virtualGatewayAccessLogProperty{
//   		file: &virtualGatewayFileAccessLogProperty{
//   			path: jsii.String("path"),
//
//   			// the properties below are optional
//   			format: &loggingFormatProperty{
//   				json: []interface{}{
//   					&jsonFormatRefProperty{
//   						key: jsii.String("key"),
//   						value: jsii.String("value"),
//   					},
//   				},
//   				text: jsii.String("text"),
//   			},
//   		},
//   	},
//   	virtualNodeAccessLog: &accessLogProperty{
//   		file: &fileAccessLogProperty{
//   			path: jsii.String("path"),
//
//   			// the properties below are optional
//   			format: &loggingFormatProperty{
//   				json: []interface{}{
//   					&jsonFormatRefProperty{
//   						key: jsii.String("key"),
//   						value: jsii.String("value"),
//   					},
//   				},
//   				text: jsii.String("text"),
//   			},
//   		},
//   	},
//   }
//
type AccessLogConfig struct {
	// VirtualGateway CFN configuration for Access Logging.
	VirtualGatewayAccessLog *CfnVirtualGateway_VirtualGatewayAccessLogProperty `field:"optional" json:"virtualGatewayAccessLog" yaml:"virtualGatewayAccessLog"`
	// VirtualNode CFN configuration for Access Logging.
	VirtualNodeAccessLog *CfnVirtualNode_AccessLogProperty `field:"optional" json:"virtualNodeAccessLog" yaml:"virtualNodeAccessLog"`
}

