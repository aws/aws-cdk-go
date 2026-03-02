package previewawsdirectconnectmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnDirectConnectGatewayPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDirectConnectGatewayMixinProps := &CfnDirectConnectGatewayMixinProps{
//   	AmazonSideAsn: jsii.String("amazonSideAsn"),
//   	DirectConnectGatewayName: jsii.String("directConnectGatewayName"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-directconnectgateway.html
//
type CfnDirectConnectGatewayMixinProps struct {
	// The autonomous system number (ASN) for the Amazon side of the connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-directconnectgateway.html#cfn-directconnect-directconnectgateway-amazonsideasn
	//
	AmazonSideAsn *string `field:"optional" json:"amazonSideAsn" yaml:"amazonSideAsn"`
	// The name of the Direct Connect gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-directconnectgateway.html#cfn-directconnect-directconnectgateway-directconnectgatewayname
	//
	DirectConnectGatewayName *string `field:"optional" json:"directConnectGatewayName" yaml:"directConnectGatewayName"`
	// The tags associated with the Direct Connect gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-directconnectgateway.html#cfn-directconnect-directconnectgateway-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

