package mixinsawsappmesh


// An object that represents the format for the logs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   loggingFormatProperty := &LoggingFormatProperty{
//   	Json: []interface{}{
//   		&JsonFormatRefProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Text: jsii.String("text"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-loggingformat.html
//
type CfnVirtualGatewayPropsMixin_LoggingFormatProperty struct {
	// The logging format for JSON.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-loggingformat.html#cfn-appmesh-virtualgateway-loggingformat-json
	//
	Json interface{} `field:"optional" json:"json" yaml:"json"`
	// The logging format for text.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-loggingformat.html#cfn-appmesh-virtualgateway-loggingformat-text
	//
	Text *string `field:"optional" json:"text" yaml:"text"`
}

