package awscleanroomsml

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnConfiguredModelAlgorithmAssociationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnConfiguredModelAlgorithmAssociationMixinProps := &CfnConfiguredModelAlgorithmAssociationMixinProps{
//   	ConfiguredModelAlgorithmArn: jsii.String("configuredModelAlgorithmArn"),
//   	Description: jsii.String("description"),
//   	MembershipIdentifier: jsii.String("membershipIdentifier"),
//   	Name: jsii.String("name"),
//   	PrivacyConfiguration: &PrivacyConfigurationProperty{
//   		Policies: &PrivacyConfigurationPoliciesProperty{
//   			TrainedModelExports: &TrainedModelExportsConfigurationPolicyProperty{
//   				FilesToExport: []*string{
//   					jsii.String("filesToExport"),
//   				},
//   				MaxSize: &TrainedModelExportsMaxSizeProperty{
//   					Unit: jsii.String("unit"),
//   					Value: jsii.Number(123),
//   				},
//   			},
//   			TrainedModelInferenceJobs: &TrainedModelInferenceJobsConfigurationPolicyProperty{
//   				ContainerLogs: []interface{}{
//   					&LogsConfigurationPolicyProperty{
//   						AllowedAccountIds: []*string{
//   							jsii.String("allowedAccountIds"),
//   						},
//   						FilterPattern: jsii.String("filterPattern"),
//   						LogRedactionConfiguration: &LogRedactionConfigurationProperty{
//   							CustomEntityConfig: &CustomEntityConfigProperty{
//   								CustomDataIdentifiers: []*string{
//   									jsii.String("customDataIdentifiers"),
//   								},
//   							},
//   							EntitiesToRedact: []*string{
//   								jsii.String("entitiesToRedact"),
//   							},
//   						},
//   						LogType: jsii.String("logType"),
//   					},
//   				},
//   				MaxOutputSize: &TrainedModelInferenceMaxOutputSizeProperty{
//   					Unit: jsii.String("unit"),
//   					Value: jsii.Number(123),
//   				},
//   			},
//   			TrainedModels: &TrainedModelsConfigurationPolicyProperty{
//   				ContainerLogs: []interface{}{
//   					&LogsConfigurationPolicyProperty{
//   						AllowedAccountIds: []*string{
//   							jsii.String("allowedAccountIds"),
//   						},
//   						FilterPattern: jsii.String("filterPattern"),
//   						LogRedactionConfiguration: &LogRedactionConfigurationProperty{
//   							CustomEntityConfig: &CustomEntityConfigProperty{
//   								CustomDataIdentifiers: []*string{
//   									jsii.String("customDataIdentifiers"),
//   								},
//   							},
//   							EntitiesToRedact: []*string{
//   								jsii.String("entitiesToRedact"),
//   							},
//   						},
//   						LogType: jsii.String("logType"),
//   					},
//   				},
//   				ContainerMetrics: &MetricsConfigurationPolicyProperty{
//   					NoiseLevel: jsii.String("noiseLevel"),
//   				},
//   				MaxArtifactSize: &TrainedModelArtifactMaxSizeProperty{
//   					Unit: jsii.String("unit"),
//   					Value: jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanroomsml-configuredmodelalgorithmassociation.html
//
type CfnConfiguredModelAlgorithmAssociationMixinProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanroomsml-configuredmodelalgorithmassociation.html#cfn-cleanroomsml-configuredmodelalgorithmassociation-configuredmodelalgorithmarn
	//
	ConfiguredModelAlgorithmArn *string `field:"optional" json:"configuredModelAlgorithmArn" yaml:"configuredModelAlgorithmArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanroomsml-configuredmodelalgorithmassociation.html#cfn-cleanroomsml-configuredmodelalgorithmassociation-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanroomsml-configuredmodelalgorithmassociation.html#cfn-cleanroomsml-configuredmodelalgorithmassociation-membershipidentifier
	//
	MembershipIdentifier *string `field:"optional" json:"membershipIdentifier" yaml:"membershipIdentifier"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanroomsml-configuredmodelalgorithmassociation.html#cfn-cleanroomsml-configuredmodelalgorithmassociation-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanroomsml-configuredmodelalgorithmassociation.html#cfn-cleanroomsml-configuredmodelalgorithmassociation-privacyconfiguration
	//
	PrivacyConfiguration interface{} `field:"optional" json:"privacyConfiguration" yaml:"privacyConfiguration"`
	// An arbitrary set of tags (key-value pairs) for this cleanrooms-ml configured model algorithm association.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanroomsml-configuredmodelalgorithmassociation.html#cfn-cleanroomsml-configuredmodelalgorithmassociation-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

