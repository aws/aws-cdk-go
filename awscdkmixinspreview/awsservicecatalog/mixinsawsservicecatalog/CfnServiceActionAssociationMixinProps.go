package mixinsawsservicecatalog


// Properties for CfnServiceActionAssociationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnServiceActionAssociationMixinProps := &CfnServiceActionAssociationMixinProps{
//   	ProductId: jsii.String("productId"),
//   	ProvisioningArtifactId: jsii.String("provisioningArtifactId"),
//   	ServiceActionId: jsii.String("serviceActionId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-serviceactionassociation.html
//
type CfnServiceActionAssociationMixinProps struct {
	// The product identifier.
	//
	// For example, `prod-abcdzk7xy33qa` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-serviceactionassociation.html#cfn-servicecatalog-serviceactionassociation-productid
	//
	ProductId *string `field:"optional" json:"productId" yaml:"productId"`
	// The identifier of the provisioning artifact.
	//
	// For example, `pa-4abcdjnxjj6ne` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-serviceactionassociation.html#cfn-servicecatalog-serviceactionassociation-provisioningartifactid
	//
	ProvisioningArtifactId *string `field:"optional" json:"provisioningArtifactId" yaml:"provisioningArtifactId"`
	// The self-service action identifier.
	//
	// For example, `act-fs7abcd89wxyz` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-serviceactionassociation.html#cfn-servicecatalog-serviceactionassociation-serviceactionid
	//
	ServiceActionId *string `field:"optional" json:"serviceActionId" yaml:"serviceActionId"`
}

