package previewawscleanroomsmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnConfiguredTableAssociationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnConfiguredTableAssociationMixinProps := &CfnConfiguredTableAssociationMixinProps{
//   	ConfiguredTableAssociationAnalysisRules: []interface{}{
//   		&ConfiguredTableAssociationAnalysisRuleProperty{
//   			Policy: &ConfiguredTableAssociationAnalysisRulePolicyProperty{
//   				V1: &ConfiguredTableAssociationAnalysisRulePolicyV1Property{
//   					Aggregation: &ConfiguredTableAssociationAnalysisRuleAggregationProperty{
//   						AllowedAdditionalAnalyses: []*string{
//   							jsii.String("allowedAdditionalAnalyses"),
//   						},
//   						AllowedResultReceivers: []*string{
//   							jsii.String("allowedResultReceivers"),
//   						},
//   					},
//   					Custom: &ConfiguredTableAssociationAnalysisRuleCustomProperty{
//   						AllowedAdditionalAnalyses: []*string{
//   							jsii.String("allowedAdditionalAnalyses"),
//   						},
//   						AllowedResultReceivers: []*string{
//   							jsii.String("allowedResultReceivers"),
//   						},
//   					},
//   					List: &ConfiguredTableAssociationAnalysisRuleListProperty{
//   						AllowedAdditionalAnalyses: []*string{
//   							jsii.String("allowedAdditionalAnalyses"),
//   						},
//   						AllowedResultReceivers: []*string{
//   							jsii.String("allowedResultReceivers"),
//   						},
//   					},
//   				},
//   			},
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	ConfiguredTableIdentifier: jsii.String("configuredTableIdentifier"),
//   	Description: jsii.String("description"),
//   	MembershipIdentifier: jsii.String("membershipIdentifier"),
//   	Name: jsii.String("name"),
//   	RoleArn: jsii.String("roleArn"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-configuredtableassociation.html
//
type CfnConfiguredTableAssociationMixinProps struct {
	// An analysis rule for a configured table association.
	//
	// This analysis rule specifies how data from the table can be used within its associated collaboration. In the console, the `ConfiguredTableAssociationAnalysisRule` is referred to as the *collaboration analysis rule* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-configuredtableassociation.html#cfn-cleanrooms-configuredtableassociation-configuredtableassociationanalysisrules
	//
	ConfiguredTableAssociationAnalysisRules interface{} `field:"optional" json:"configuredTableAssociationAnalysisRules" yaml:"configuredTableAssociationAnalysisRules"`
	// A unique identifier for the configured table to be associated to.
	//
	// Currently accepts a configured table ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-configuredtableassociation.html#cfn-cleanrooms-configuredtableassociation-configuredtableidentifier
	//
	ConfiguredTableIdentifier *string `field:"optional" json:"configuredTableIdentifier" yaml:"configuredTableIdentifier"`
	// A description of the configured table association.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-configuredtableassociation.html#cfn-cleanrooms-configuredtableassociation-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The unique ID for the membership this configured table association belongs to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-configuredtableassociation.html#cfn-cleanrooms-configuredtableassociation-membershipidentifier
	//
	MembershipIdentifier *string `field:"optional" json:"membershipIdentifier" yaml:"membershipIdentifier"`
	// The name of the configured table association, in lowercase.
	//
	// The table is identified by this name when running protected queries against the underlying data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-configuredtableassociation.html#cfn-cleanrooms-configuredtableassociation-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The service will assume this role to access catalog metadata and query the table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-configuredtableassociation.html#cfn-cleanrooms-configuredtableassociation-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// An optional label that you can assign to a resource when you create it.
	//
	// Each tag consists of a key and an optional value, both of which you define. When you use tagging, you can also use tag-based access control in IAM policies to control access to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanrooms-configuredtableassociation.html#cfn-cleanrooms-configuredtableassociation-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

