package awsarcregionswitch


// The AWS EKS resource scaling configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eksResourceScalingConfigurationProperty := &EksResourceScalingConfigurationProperty{
//   	KubernetesResourceType: &KubernetesResourceTypeProperty{
//   		ApiVersion: jsii.String("apiVersion"),
//   		Kind: jsii.String("kind"),
//   	},
//
//   	// the properties below are optional
//   	CapacityMonitoringApproach: jsii.String("capacityMonitoringApproach"),
//   	EksClusters: []interface{}{
//   		&EksClusterProperty{
//   			ClusterArn: jsii.String("clusterArn"),
//
//   			// the properties below are optional
//   			CrossAccountRole: jsii.String("crossAccountRole"),
//   			ExternalId: jsii.String("externalId"),
//   		},
//   	},
//   	ScalingResources: []interface{}{
//   		map[string]interface{}{
//   			"scalingResourcesKey": map[string]interface{}{
//   				"scalingResourcesKey": &KubernetesScalingResourceProperty{
//   					"name": jsii.String("name"),
//   					"namespace": jsii.String("namespace"),
//
//   					// the properties below are optional
//   					"hpaName": jsii.String("hpaName"),
//   				},
//   			},
//   		},
//   	},
//   	TargetPercent: jsii.Number(123),
//   	TimeoutMinutes: jsii.Number(123),
//   	Ungraceful: &EksResourceScalingUngracefulProperty{
//   		MinimumSuccessPercentage: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-eksresourcescalingconfiguration.html
//
type CfnPlan_EksResourceScalingConfigurationProperty struct {
	// The Kubernetes resource type for the configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-eksresourcescalingconfiguration.html#cfn-arcregionswitch-plan-eksresourcescalingconfiguration-kubernetesresourcetype
	//
	KubernetesResourceType interface{} `field:"required" json:"kubernetesResourceType" yaml:"kubernetesResourceType"`
	// The monitoring approach for the configuration, that is, whether it was sampled in the last 24 hours or autoscaled in the last 24 hours.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-eksresourcescalingconfiguration.html#cfn-arcregionswitch-plan-eksresourcescalingconfiguration-capacitymonitoringapproach
	//
	CapacityMonitoringApproach *string `field:"optional" json:"capacityMonitoringApproach" yaml:"capacityMonitoringApproach"`
	// The clusters for the configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-eksresourcescalingconfiguration.html#cfn-arcregionswitch-plan-eksresourcescalingconfiguration-eksclusters
	//
	EksClusters interface{} `field:"optional" json:"eksClusters" yaml:"eksClusters"`
	// The scaling resources for the configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-eksresourcescalingconfiguration.html#cfn-arcregionswitch-plan-eksresourcescalingconfiguration-scalingresources
	//
	ScalingResources interface{} `field:"optional" json:"scalingResources" yaml:"scalingResources"`
	// The target percentage for the configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-eksresourcescalingconfiguration.html#cfn-arcregionswitch-plan-eksresourcescalingconfiguration-targetpercent
	//
	// Default: - 100.
	//
	TargetPercent *float64 `field:"optional" json:"targetPercent" yaml:"targetPercent"`
	// The timeout value specified for the configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-eksresourcescalingconfiguration.html#cfn-arcregionswitch-plan-eksresourcescalingconfiguration-timeoutminutes
	//
	// Default: - 60.
	//
	TimeoutMinutes *float64 `field:"optional" json:"timeoutMinutes" yaml:"timeoutMinutes"`
	// The settings for ungraceful execution.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-eksresourcescalingconfiguration.html#cfn-arcregionswitch-plan-eksresourcescalingconfiguration-ungraceful
	//
	Ungraceful interface{} `field:"optional" json:"ungraceful" yaml:"ungraceful"`
}

