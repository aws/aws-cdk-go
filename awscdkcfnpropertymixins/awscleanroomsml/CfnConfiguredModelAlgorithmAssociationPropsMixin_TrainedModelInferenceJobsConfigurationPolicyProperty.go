package awscleanroomsml


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   trainedModelInferenceJobsConfigurationPolicyProperty := &TrainedModelInferenceJobsConfigurationPolicyProperty{
//   	ContainerLogs: []interface{}{
//   		&LogsConfigurationPolicyProperty{
//   			AllowedAccountIds: []*string{
//   				jsii.String("allowedAccountIds"),
//   			},
//   			FilterPattern: jsii.String("filterPattern"),
//   			LogRedactionConfiguration: &LogRedactionConfigurationProperty{
//   				CustomEntityConfig: &CustomEntityConfigProperty{
//   					CustomDataIdentifiers: []*string{
//   						jsii.String("customDataIdentifiers"),
//   					},
//   				},
//   				EntitiesToRedact: []*string{
//   					jsii.String("entitiesToRedact"),
//   				},
//   			},
//   			LogType: jsii.String("logType"),
//   		},
//   	},
//   	MaxOutputSize: &TrainedModelInferenceMaxOutputSizeProperty{
//   		Unit: jsii.String("unit"),
//   		Value: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanroomsml-configuredmodelalgorithmassociation-trainedmodelinferencejobsconfigurationpolicy.html
//
type CfnConfiguredModelAlgorithmAssociationPropsMixin_TrainedModelInferenceJobsConfigurationPolicyProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanroomsml-configuredmodelalgorithmassociation-trainedmodelinferencejobsconfigurationpolicy.html#cfn-cleanroomsml-configuredmodelalgorithmassociation-trainedmodelinferencejobsconfigurationpolicy-containerlogs
	//
	ContainerLogs interface{} `field:"optional" json:"containerLogs" yaml:"containerLogs"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanroomsml-configuredmodelalgorithmassociation-trainedmodelinferencejobsconfigurationpolicy.html#cfn-cleanroomsml-configuredmodelalgorithmassociation-trainedmodelinferencejobsconfigurationpolicy-maxoutputsize
	//
	MaxOutputSize interface{} `field:"optional" json:"maxOutputSize" yaml:"maxOutputSize"`
}

