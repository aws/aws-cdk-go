package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   managedResourceDetailsProperty := &ManagedResourceDetailsProperty{
//   	Domain: jsii.String("domain"),
//   	ResourceAssociationArn: jsii.String("resourceAssociationArn"),
//   	ResourceGatewayArn: jsii.String("resourceGatewayArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-managedresourcedetails.html
//
type CfnGatewayTargetPropsMixin_ManagedResourceDetailsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-managedresourcedetails.html#cfn-bedrockagentcore-gatewaytarget-managedresourcedetails-domain
	//
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-managedresourcedetails.html#cfn-bedrockagentcore-gatewaytarget-managedresourcedetails-resourceassociationarn
	//
	ResourceAssociationArn *string `field:"optional" json:"resourceAssociationArn" yaml:"resourceAssociationArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-managedresourcedetails.html#cfn-bedrockagentcore-gatewaytarget-managedresourcedetails-resourcegatewayarn
	//
	ResourceGatewayArn *string `field:"optional" json:"resourceGatewayArn" yaml:"resourceGatewayArn"`
}

