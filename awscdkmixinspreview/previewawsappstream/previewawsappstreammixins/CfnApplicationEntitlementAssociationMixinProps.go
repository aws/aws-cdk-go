package previewawsappstreammixins


// Properties for CfnApplicationEntitlementAssociationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnApplicationEntitlementAssociationMixinProps := &CfnApplicationEntitlementAssociationMixinProps{
//   	ApplicationIdentifier: jsii.String("applicationIdentifier"),
//   	EntitlementName: jsii.String("entitlementName"),
//   	StackName: jsii.String("stackName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-applicationentitlementassociation.html
//
type CfnApplicationEntitlementAssociationMixinProps struct {
	// The identifier of the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-applicationentitlementassociation.html#cfn-appstream-applicationentitlementassociation-applicationidentifier
	//
	ApplicationIdentifier *string `field:"optional" json:"applicationIdentifier" yaml:"applicationIdentifier"`
	// The name of the entitlement.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-applicationentitlementassociation.html#cfn-appstream-applicationentitlementassociation-entitlementname
	//
	EntitlementName *string `field:"optional" json:"entitlementName" yaml:"entitlementName"`
	// The name of the stack.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-applicationentitlementassociation.html#cfn-appstream-applicationentitlementassociation-stackname
	//
	StackName *string `field:"optional" json:"stackName" yaml:"stackName"`
}

