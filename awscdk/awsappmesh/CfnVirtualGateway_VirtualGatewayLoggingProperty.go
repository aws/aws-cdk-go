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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewaylogging.html
//
type CfnVirtualGateway_VirtualGatewayLoggingProperty struct {
	// The access log configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewaylogging.html#cfn-appmesh-virtualgateway-virtualgatewaylogging-accesslog
	//
	AccessLog interface{} `field:"optional" json:"accessLog" yaml:"accessLog"`
}

