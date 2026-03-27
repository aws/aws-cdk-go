package awscleanroomsml


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   privacyConfigurationProperty := &PrivacyConfigurationProperty{
//   	Policies: &PrivacyConfigurationPoliciesProperty{
//   		TrainedModelExports: &TrainedModelExportsConfigurationPolicyProperty{
//   			FilesToExport: []*string{
//   				jsii.String("filesToExport"),
//   			},
//   			MaxSize: &TrainedModelExportsMaxSizeProperty{
//   				Unit: jsii.String("unit"),
//   				Value: jsii.Number(123),
//   			},
//   		},
//   		TrainedModelInferenceJobs: &TrainedModelInferenceJobsConfigurationPolicyProperty{
//   			ContainerLogs: []interface{}{
//   				&LogsConfigurationPolicyProperty{
//   					AllowedAccountIds: []*string{
//   						jsii.String("allowedAccountIds"),
//   					},
//
//   					// the properties below are optional
//   					FilterPattern: jsii.String("filterPattern"),
//   					LogRedactionConfiguration: &LogRedactionConfigurationProperty{
//   						EntitiesToRedact: []*string{
//   							jsii.String("entitiesToRedact"),
//   						},
//
//   						// the properties below are optional
//   						CustomEntityConfig: &CustomEntityConfigProperty{
//   							CustomDataIdentifiers: []*string{
//   								jsii.String("customDataIdentifiers"),
//   							},
//   						},
//   					},
//   					LogType: jsii.String("logType"),
//   				},
//   			},
//   			MaxOutputSize: &TrainedModelInferenceMaxOutputSizeProperty{
//   				Unit: jsii.String("unit"),
//   				Value: jsii.Number(123),
//   			},
//   		},
//   		TrainedModels: &TrainedModelsConfigurationPolicyProperty{
//   			ContainerLogs: []interface{}{
//   				&LogsConfigurationPolicyProperty{
//   					AllowedAccountIds: []*string{
//   						jsii.String("allowedAccountIds"),
//   					},
//
//   					// the properties below are optional
//   					FilterPattern: jsii.String("filterPattern"),
//   					LogRedactionConfiguration: &LogRedactionConfigurationProperty{
//   						EntitiesToRedact: []*string{
//   							jsii.String("entitiesToRedact"),
//   						},
//
//   						// the properties below are optional
//   						CustomEntityConfig: &CustomEntityConfigProperty{
//   							CustomDataIdentifiers: []*string{
//   								jsii.String("customDataIdentifiers"),
//   							},
//   						},
//   					},
//   					LogType: jsii.String("logType"),
//   				},
//   			},
//   			ContainerMetrics: &MetricsConfigurationPolicyProperty{
//   				NoiseLevel: jsii.String("noiseLevel"),
//   			},
//   			MaxArtifactSize: &TrainedModelArtifactMaxSizeProperty{
//   				Unit: jsii.String("unit"),
//   				Value: jsii.Number(123),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanroomsml-configuredmodelalgorithmassociation-privacyconfiguration.html
//
type CfnConfiguredModelAlgorithmAssociation_PrivacyConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanroomsml-configuredmodelalgorithmassociation-privacyconfiguration.html#cfn-cleanroomsml-configuredmodelalgorithmassociation-privacyconfiguration-policies
	//
	Policies interface{} `field:"required" json:"policies" yaml:"policies"`
}

