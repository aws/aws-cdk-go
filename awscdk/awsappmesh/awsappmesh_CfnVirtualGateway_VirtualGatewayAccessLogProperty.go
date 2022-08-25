package awsappmesh


// The access log configuration for a virtual gateway.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   virtualGatewayAccessLogProperty := &virtualGatewayAccessLogProperty{
//   	file: &virtualGatewayFileAccessLogProperty{
//   		path: jsii.String("path"),
//   	},
//   }
//
type CfnVirtualGateway_VirtualGatewayAccessLogProperty struct {
	// The file object to send virtual gateway access logs to.
	File interface{} `field:"optional" json:"file" yaml:"file"`
}

