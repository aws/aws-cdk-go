package awstransfer


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   connectorAsyncMdnConfigProperty := &ConnectorAsyncMdnConfigProperty{
//   	ServerIds: []*string{
//   		jsii.String("serverIds"),
//   	},
//   	Url: jsii.String("url"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-connector-connectorasyncmdnconfig.html
//
type CfnConnector_ConnectorAsyncMdnConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-connector-connectorasyncmdnconfig.html#cfn-transfer-connector-connectorasyncmdnconfig-serverids
	//
	ServerIds *[]*string `field:"required" json:"serverIds" yaml:"serverIds"`
	// URL of the server to receive the MDN response on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-connector-connectorasyncmdnconfig.html#cfn-transfer-connector-connectorasyncmdnconfig-url
	//
	Url *string `field:"required" json:"url" yaml:"url"`
}

