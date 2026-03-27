package awscleanroomsml


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   trainedModelsConfigurationPolicyProperty := &TrainedModelsConfigurationPolicyProperty{
//   	ContainerLogs: []interface{}{
//   		&LogsConfigurationPolicyProperty{
//   			AllowedAccountIds: []*string{
//   				jsii.String("allowedAccountIds"),
//   			},
//
//   			// the properties below are optional
//   			FilterPattern: jsii.String("filterPattern"),
//   			LogRedactionConfiguration: &LogRedactionConfigurationProperty{
//   				EntitiesToRedact: []*string{
//   					jsii.String("entitiesToRedact"),
//   				},
//
//   				// the properties below are optional
//   				CustomEntityConfig: &CustomEntityConfigProperty{
//   					CustomDataIdentifiers: []*string{
//   						jsii.String("customDataIdentifiers"),
//   					},
//   				},
//   			},
//   			LogType: jsii.String("logType"),
//   		},
//   	},
//   	ContainerMetrics: &MetricsConfigurationPolicyProperty{
//   		NoiseLevel: jsii.String("noiseLevel"),
//   	},
//   	MaxArtifactSize: &TrainedModelArtifactMaxSizeProperty{
//   		Unit: jsii.String("unit"),
//   		Value: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanroomsml-configuredmodelalgorithmassociation-trainedmodelsconfigurationpolicy.html
//
type CfnConfiguredModelAlgorithmAssociation_TrainedModelsConfigurationPolicyProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanroomsml-configuredmodelalgorithmassociation-trainedmodelsconfigurationpolicy.html#cfn-cleanroomsml-configuredmodelalgorithmassociation-trainedmodelsconfigurationpolicy-containerlogs
	//
	ContainerLogs interface{} `field:"optional" json:"containerLogs" yaml:"containerLogs"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanroomsml-configuredmodelalgorithmassociation-trainedmodelsconfigurationpolicy.html#cfn-cleanroomsml-configuredmodelalgorithmassociation-trainedmodelsconfigurationpolicy-containermetrics
	//
	ContainerMetrics interface{} `field:"optional" json:"containerMetrics" yaml:"containerMetrics"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanroomsml-configuredmodelalgorithmassociation-trainedmodelsconfigurationpolicy.html#cfn-cleanroomsml-configuredmodelalgorithmassociation-trainedmodelsconfigurationpolicy-maxartifactsize
	//
	MaxArtifactSize interface{} `field:"optional" json:"maxArtifactSize" yaml:"maxArtifactSize"`
}

