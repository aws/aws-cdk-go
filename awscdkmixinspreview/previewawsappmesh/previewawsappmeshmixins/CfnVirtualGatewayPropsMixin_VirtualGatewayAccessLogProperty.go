package previewawsappmeshmixins


// The access log configuration for a virtual gateway.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   virtualGatewayAccessLogProperty := &VirtualGatewayAccessLogProperty{
//   	File: &VirtualGatewayFileAccessLogProperty{
//   		Format: &LoggingFormatProperty{
//   			Json: []interface{}{
//   				&JsonFormatRefProperty{
//   					Key: jsii.String("key"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			Text: jsii.String("text"),
//   		},
//   		Path: jsii.String("path"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewayaccesslog.html
//
type CfnVirtualGatewayPropsMixin_VirtualGatewayAccessLogProperty struct {
	// The file object to send virtual gateway access logs to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewayaccesslog.html#cfn-appmesh-virtualgateway-virtualgatewayaccesslog-file
	//
	File interface{} `field:"optional" json:"file" yaml:"file"`
}

