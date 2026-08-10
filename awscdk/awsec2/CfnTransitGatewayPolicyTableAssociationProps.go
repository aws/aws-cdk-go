package awsec2


// Properties for defining a `CfnTransitGatewayPolicyTableAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTransitGatewayPolicyTableAssociationProps := &CfnTransitGatewayPolicyTableAssociationProps{
//   	TransitGatewayAttachmentId: jsii.String("transitGatewayAttachmentId"),
//   	TransitGatewayPolicyTableId: jsii.String("transitGatewayPolicyTableId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewaypolicytableassociation.html
//
type CfnTransitGatewayPolicyTableAssociationProps struct {
	// The ID of transit gateway attachment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewaypolicytableassociation.html#cfn-ec2-transitgatewaypolicytableassociation-transitgatewayattachmentid
	//
	TransitGatewayAttachmentId *string `field:"required" json:"transitGatewayAttachmentId" yaml:"transitGatewayAttachmentId"`
	// The ID of transit gateway policy table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewaypolicytableassociation.html#cfn-ec2-transitgatewaypolicytableassociation-transitgatewaypolicytableid
	//
	TransitGatewayPolicyTableId *string `field:"required" json:"transitGatewayPolicyTableId" yaml:"transitGatewayPolicyTableId"`
}

