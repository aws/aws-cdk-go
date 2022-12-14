package awsappmesh


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loggingFormatProperty := &loggingFormatProperty{
//   	json: []interface{}{
//   		&jsonFormatRefProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	text: jsii.String("text"),
//   }
//
type CfnVirtualGateway_LoggingFormatProperty struct {
	// `CfnVirtualGateway.LoggingFormatProperty.Json`.
	Json interface{} `field:"optional" json:"json" yaml:"json"`
	// `CfnVirtualGateway.LoggingFormatProperty.Text`.
	Text *string `field:"optional" json:"text" yaml:"text"`
}

