package awsappmesh


// An object that represents logging information.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   virtualGatewayLoggingProperty := &VirtualGatewayLoggingProperty{
//   	AccessLog: &VirtualGatewayAccessLogProperty{
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
//   }
//
type CfnVirtualGateway_VirtualGatewayLoggingProperty struct {
	// The access log configuration.
	AccessLog interface{} `field:"optional" json:"accessLog" yaml:"accessLog"`
}

