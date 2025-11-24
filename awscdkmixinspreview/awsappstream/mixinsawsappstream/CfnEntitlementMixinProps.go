package mixinsawsappstream


// Properties for CfnEntitlementPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnEntitlementMixinProps := &CfnEntitlementMixinProps{
//   	AppVisibility: jsii.String("appVisibility"),
//   	Attributes: []interface{}{
//   		&AttributeProperty{
//   			Name: jsii.String("name"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	StackName: jsii.String("stackName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-entitlement.html
//
type CfnEntitlementMixinProps struct {
	// Specifies whether to entitle all apps or only selected apps.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-entitlement.html#cfn-appstream-entitlement-appvisibility
	//
	AppVisibility *string `field:"optional" json:"appVisibility" yaml:"appVisibility"`
	// The attributes of the entitlement.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-entitlement.html#cfn-appstream-entitlement-attributes
	//
	Attributes interface{} `field:"optional" json:"attributes" yaml:"attributes"`
	// The description of the entitlement.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-entitlement.html#cfn-appstream-entitlement-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the entitlement.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-entitlement.html#cfn-appstream-entitlement-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The name of the stack.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-entitlement.html#cfn-appstream-entitlement-stackname
	//
	StackName *string `field:"optional" json:"stackName" yaml:"stackName"`
}

