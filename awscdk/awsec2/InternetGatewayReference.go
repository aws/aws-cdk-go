package awsec2


// A reference to a InternetGateway resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   internetGatewayReference := &InternetGatewayReference{
//   	InternetGatewayId: jsii.String("internetGatewayId"),
//   }
//
type InternetGatewayReference struct {
	// The InternetGatewayId of the InternetGateway resource.
	InternetGatewayId *string `field:"required" json:"internetGatewayId" yaml:"internetGatewayId"`
}

