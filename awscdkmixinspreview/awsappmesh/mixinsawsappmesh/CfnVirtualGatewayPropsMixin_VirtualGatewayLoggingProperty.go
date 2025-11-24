package mixinsawsappmesh


// An object that represents logging information.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   virtualGatewayLoggingProperty := &VirtualGatewayLoggingProperty{
//   	AccessLog: &VirtualGatewayAccessLogProperty{
//   		File: &VirtualGatewayFileAccessLogProperty{
//   			Format: &LoggingFormatProperty{
//   				Json: []interface{}{
//   					&JsonFormatRefProperty{
//   						Key: jsii.String("key"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   				Text: jsii.String("text"),
//   			},
//   			Path: jsii.String("path"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewaylogging.html
//
type CfnVirtualGatewayPropsMixin_VirtualGatewayLoggingProperty struct {
	// The access log configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewaylogging.html#cfn-appmesh-virtualgateway-virtualgatewaylogging-accesslog
	//
	AccessLog interface{} `field:"optional" json:"accessLog" yaml:"accessLog"`
}

