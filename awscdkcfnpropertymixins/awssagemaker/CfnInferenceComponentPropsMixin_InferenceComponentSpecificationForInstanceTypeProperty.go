package awssagemaker


// A specification for one instance type, for use in Specifications.
//
// InstanceType is required here, and is not accepted on the singular Specification. BaseInferenceComponentName is not accepted here either: adapter inference components are supported only on the singular Specification.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   inferenceComponentSpecificationForInstanceTypeProperty := &InferenceComponentSpecificationForInstanceTypeProperty{
//   	ComputeResourceRequirements: &InferenceComponentComputeResourceRequirementsProperty{
//   		MaxMemoryRequiredInMb: jsii.Number(123),
//   		MinMemoryRequiredInMb: jsii.Number(123),
//   		NumberOfAcceleratorDevicesRequired: jsii.Number(123),
//   		NumberOfCpuCoresRequired: jsii.Number(123),
//   	},
//   	Container: &InferenceComponentContainerSpecificationForInstanceTypeProperty{
//   		ArtifactUrl: jsii.String("artifactUrl"),
//   		ContainerMetricsConfig: &ContainerMetricsConfigProperty{
//   			MetricsEndpoints: []interface{}{
//   				&MetricsEndpointProperty{
//   					MetricPublishFrequencyInSeconds: jsii.Number(123),
//   					MetricsEndpointPath: jsii.String("metricsEndpointPath"),
//   				},
//   			},
//   		},
//   		Environment: map[string]*string{
//   			"environmentKey": jsii.String("environment"),
//   		},
//   		Image: jsii.String("image"),
//   	},
//   	CurrentDataCacheConfig: &InferenceComponentDataCacheConfigProperty{
//   		EnableCaching: jsii.Boolean(false),
//   	},
//   	DataCacheConfig: &InferenceComponentDataCacheConfigProperty{
//   		EnableCaching: jsii.Boolean(false),
//   	},
//   	InstanceType: jsii.String("instanceType"),
//   	ModelName: jsii.String("modelName"),
//   	SchedulingConfig: &InferenceComponentSchedulingConfigProperty{
//   		AvailabilityZoneBalance: &InferenceComponentAvailabilityZoneBalanceProperty{
//   			EnforcementMode: jsii.String("enforcementMode"),
//   			MaxImbalance: jsii.Number(123),
//   		},
//   		PlacementStrategy: jsii.String("placementStrategy"),
//   	},
//   	StartupParameters: &InferenceComponentStartupParametersProperty{
//   		ContainerStartupHealthCheckTimeoutInSeconds: jsii.Number(123),
//   		ModelDataDownloadTimeoutInSeconds: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-inferencecomponentspecificationforinstancetype.html
//
type CfnInferenceComponentPropsMixin_InferenceComponentSpecificationForInstanceTypeProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-inferencecomponentspecificationforinstancetype.html#cfn-sagemaker-inferencecomponent-inferencecomponentspecificationforinstancetype-computeresourcerequirements
	//
	ComputeResourceRequirements interface{} `field:"optional" json:"computeResourceRequirements" yaml:"computeResourceRequirements"`
	// Container specification for one Specifications entry.
	//
	// Distinct from InferenceComponentContainerSpecification: DescribeInferenceComponent returns no per-entry DeployedImage (VERIFIED in us-west-2), so DeployedImage is intentionally omitted here and this definition can never be aggregated into a plural READ response. The singular InferenceComponentContainerSpecification keeps DeployedImage - the service DOES return it there.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-inferencecomponentspecificationforinstancetype.html#cfn-sagemaker-inferencecomponent-inferencecomponentspecificationforinstancetype-container
	//
	Container interface{} `field:"optional" json:"container" yaml:"container"`
	// Settings that affect how the inference component caches data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-inferencecomponentspecificationforinstancetype.html#cfn-sagemaker-inferencecomponent-inferencecomponentspecificationforinstancetype-currentdatacacheconfig
	//
	CurrentDataCacheConfig interface{} `field:"optional" json:"currentDataCacheConfig" yaml:"currentDataCacheConfig"`
	// Settings that affect how the inference component caches data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-inferencecomponentspecificationforinstancetype.html#cfn-sagemaker-inferencecomponent-inferencecomponentspecificationforinstancetype-datacacheconfig
	//
	DataCacheConfig interface{} `field:"optional" json:"dataCacheConfig" yaml:"dataCacheConfig"`
	// An ML compute instance type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-inferencecomponentspecificationforinstancetype.html#cfn-sagemaker-inferencecomponent-inferencecomponentspecificationforinstancetype-instancetype
	//
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// The name of the model to use with the inference component.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-inferencecomponentspecificationforinstancetype.html#cfn-sagemaker-inferencecomponent-inferencecomponentspecificationforinstancetype-modelname
	//
	ModelName *string `field:"optional" json:"modelName" yaml:"modelName"`
	// The scheduling configuration that determines how inference component copies are placed across available instances.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-inferencecomponentspecificationforinstancetype.html#cfn-sagemaker-inferencecomponent-inferencecomponentspecificationforinstancetype-schedulingconfig
	//
	SchedulingConfig interface{} `field:"optional" json:"schedulingConfig" yaml:"schedulingConfig"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-inferencecomponentspecificationforinstancetype.html#cfn-sagemaker-inferencecomponent-inferencecomponentspecificationforinstancetype-startupparameters
	//
	StartupParameters interface{} `field:"optional" json:"startupParameters" yaml:"startupParameters"`
}

