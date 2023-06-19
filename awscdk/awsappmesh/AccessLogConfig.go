package awsappmesh


// All Properties for Envoy Access logs for mesh endpoints.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accessLogConfig := &AccessLogConfig{
//   	VirtualGatewayAccessLog: &VirtualGatewayAccessLogProperty{
//   		File: &VirtualGatewayFileAccessLogProperty{
//   			Path: jsii.String("path"),
//
//   			// the properties below are optional
//   			Format: &LoggingFormatProperty{
//   				Json: []interface{}{
//   					&JsonFormatRefProperty{
//   						Key: jsii.String("key"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   				Text: jsii.String("text"),
//   			},
//   		},
//   	},
//   	VirtualNodeAccessLog: &AccessLogProperty{
//   		File: &FileAccessLogProperty{
//   			Path: jsii.String("path"),
//
//   			// the properties below are optional
//   			Format: &LoggingFormatProperty{
//   				Json: []interface{}{
//   					&JsonFormatRefProperty{
//   						Key: jsii.String("key"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   				Text: jsii.String("text"),
//   			},
//   		},
//   	},
//   }
//
// Experimental.
type AccessLogConfig struct {
	// VirtualGateway CFN configuration for Access Logging.
	// Experimental.
	VirtualGatewayAccessLog *CfnVirtualGateway_VirtualGatewayAccessLogProperty `field:"optional" json:"virtualGatewayAccessLog" yaml:"virtualGatewayAccessLog"`
	// VirtualNode CFN configuration for Access Logging.
	// Experimental.
	VirtualNodeAccessLog *CfnVirtualNode_AccessLogProperty `field:"optional" json:"virtualNodeAccessLog" yaml:"virtualNodeAccessLog"`
}

