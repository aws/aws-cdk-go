package mixinsawscleanrooms

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnIdNamespaceAssociationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnIdNamespaceAssociationMixinProps := &CfnIdNamespaceAssociationMixinProps{
//   	Description: jsii.String("description"),
//   	IdMappingConfig: &IdMappingConfigProperty{
//   		AllowUseAsDimensionColumn: jsii.Boolean(false),
//   	},
//   	InputReferenceConfig: &IdNamespaceAssociationInputReferenceConfigProperty{
//   		InputReferenceArn: jsii.String("inputReferenceArn"),
//   		ManageResourcePolicies: jsii.Boolean(false),
//   	},
//   	MembershipIdentifier: jsii.String("membershipIdentifier"),
//   	Name: jsii.String("name"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-idnamespaceassociation.html
//
type CfnIdNamespaceAssociationMixinProps struct {
	// The description of the ID namespace association.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-idnamespaceassociation.html#cfn-cleanrooms-idnamespaceassociation-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The configuration settings for the ID mapping table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-idnamespaceassociation.html#cfn-cleanrooms-idnamespaceassociation-idmappingconfig
	//
	IdMappingConfig interface{} `field:"optional" json:"idMappingConfig" yaml:"idMappingConfig"`
	// The input reference configuration for the ID namespace association.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-idnamespaceassociation.html#cfn-cleanrooms-idnamespaceassociation-inputreferenceconfig
	//
	InputReferenceConfig interface{} `field:"optional" json:"inputReferenceConfig" yaml:"inputReferenceConfig"`
	// The unique identifier of the membership that contains the ID namespace association.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-idnamespaceassociation.html#cfn-cleanrooms-idnamespaceassociation-membershipidentifier
	//
	MembershipIdentifier *string `field:"optional" json:"membershipIdentifier" yaml:"membershipIdentifier"`
	// The name of this ID namespace association.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-idnamespaceassociation.html#cfn-cleanrooms-idnamespaceassociation-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// An optional label that you can assign to a resource when you create it.
	//
	// Each tag consists of a key and an optional value, both of which you define. When you use tagging, you can also use tag-based access control in IAM policies to control access to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-idnamespaceassociation.html#cfn-cleanrooms-idnamespaceassociation-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

