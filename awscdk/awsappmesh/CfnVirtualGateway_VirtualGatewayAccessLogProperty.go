package awsappmesh


// The access log configuration for a virtual gateway.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   virtualGatewayAccessLogProperty := &VirtualGatewayAccessLogProperty{
//   	File: &VirtualGatewayFileAccessLogProperty{
//   		Path: jsii.String("path"),
//
//   		// the properties below are optional
//   		Format: &LoggingFormatProperty{
//   			Json: []interface{}{
//   				&JsonFormatRefProperty{
//   					Key: jsii.String("key"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			Text: jsii.String("text"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewayaccesslog.html
//
type CfnVirtualGateway_VirtualGatewayAccessLogProperty struct {
	// The file object to send virtual gateway access logs to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewayaccesslog.html#cfn-appmesh-virtualgateway-virtualgatewayaccesslog-file
	//
	File interface{} `field:"optional" json:"file" yaml:"file"`
}

