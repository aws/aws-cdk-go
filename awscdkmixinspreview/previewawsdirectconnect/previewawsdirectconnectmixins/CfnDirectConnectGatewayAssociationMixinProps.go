package previewawsdirectconnectmixins


// Properties for CfnDirectConnectGatewayAssociationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDirectConnectGatewayAssociationMixinProps := &CfnDirectConnectGatewayAssociationMixinProps{
//   	AcceptDirectConnectGatewayAssociationProposalRoleArn: jsii.String("acceptDirectConnectGatewayAssociationProposalRoleArn"),
//   	AllowedPrefixesToDirectConnectGateway: []*string{
//   		jsii.String("allowedPrefixesToDirectConnectGateway"),
//   	},
//   	AssociatedGatewayId: jsii.String("associatedGatewayId"),
//   	DirectConnectGatewayId: jsii.String("directConnectGatewayId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-directconnectgatewayassociation.html
//
type CfnDirectConnectGatewayAssociationMixinProps struct {
	// The Amazon Resource Name (ARN) of the role to accept the Direct Connect Gateway association proposal.
	//
	// Needs directconnect:AcceptDirectConnectGatewayAssociationProposal permissions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-directconnectgatewayassociation.html#cfn-directconnect-directconnectgatewayassociation-acceptdirectconnectgatewayassociationproposalrolearn
	//
	AcceptDirectConnectGatewayAssociationProposalRoleArn *string `field:"optional" json:"acceptDirectConnectGatewayAssociationProposalRoleArn" yaml:"acceptDirectConnectGatewayAssociationProposalRoleArn"`
	// The Amazon VPC prefixes to advertise to the Direct Connect gateway.
	//
	// This parameter is required when you create an association to a transit gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-directconnectgatewayassociation.html#cfn-directconnect-directconnectgatewayassociation-allowedprefixestodirectconnectgateway
	//
	AllowedPrefixesToDirectConnectGateway *[]*string `field:"optional" json:"allowedPrefixesToDirectConnectGateway" yaml:"allowedPrefixesToDirectConnectGateway"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-directconnectgatewayassociation.html#cfn-directconnect-directconnectgatewayassociation-associatedgatewayid
	//
	AssociatedGatewayId *string `field:"optional" json:"associatedGatewayId" yaml:"associatedGatewayId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-directconnectgatewayassociation.html#cfn-directconnect-directconnectgatewayassociation-directconnectgatewayid
	//
	DirectConnectGatewayId *string `field:"optional" json:"directConnectGatewayId" yaml:"directConnectGatewayId"`
}

