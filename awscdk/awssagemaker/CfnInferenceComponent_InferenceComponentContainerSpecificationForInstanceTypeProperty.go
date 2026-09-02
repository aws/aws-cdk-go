package awssagemaker


// Container specification for one Specifications entry.
//
// Distinct from InferenceComponentContainerSpecification: DescribeInferenceComponent returns no per-entry DeployedImage (VERIFIED in us-west-2), so DeployedImage is intentionally omitted here and this definition can never be aggregated into a plural READ response. The singular InferenceComponentContainerSpecification keeps DeployedImage - the service DOES return it there.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inferenceComponentContainerSpecificationForInstanceTypeProperty := &InferenceComponentContainerSpecificationForInstanceTypeProperty{
//   	ArtifactUrl: jsii.String("artifactUrl"),
//   	ContainerMetricsConfig: &ContainerMetricsConfigProperty{
//   		MetricsEndpoints: []interface{}{
//   			&MetricsEndpointProperty{
//   				MetricsEndpointPath: jsii.String("metricsEndpointPath"),
//
//   				// the properties below are optional
//   				MetricPublishFrequencyInSeconds: jsii.Number(123),
//   			},
//   		},
//   	},
//   	Environment: map[string]*string{
//   		"environmentKey": jsii.String("environment"),
//   	},
//   	Image: jsii.String("image"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-inferencecomponentcontainerspecificationforinstancetype.html
//
type CfnInferenceComponent_InferenceComponentContainerSpecificationForInstanceTypeProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-inferencecomponentcontainerspecificationforinstancetype.html#cfn-sagemaker-inferencecomponent-inferencecomponentcontainerspecificationforinstancetype-artifacturl
	//
	ArtifactUrl *string `field:"optional" json:"artifactUrl" yaml:"artifactUrl"`
	// The configuration for container metrics scraping.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-inferencecomponentcontainerspecificationforinstancetype.html#cfn-sagemaker-inferencecomponent-inferencecomponentcontainerspecificationforinstancetype-containermetricsconfig
	//
	ContainerMetricsConfig interface{} `field:"optional" json:"containerMetricsConfig" yaml:"containerMetricsConfig"`
	// Environment variables to specify on the container.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-inferencecomponentcontainerspecificationforinstancetype.html#cfn-sagemaker-inferencecomponent-inferencecomponentcontainerspecificationforinstancetype-environment
	//
	Environment interface{} `field:"optional" json:"environment" yaml:"environment"`
	// The image to use for the container that will be materialized for the inference component.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-inferencecomponentcontainerspecificationforinstancetype.html#cfn-sagemaker-inferencecomponent-inferencecomponentcontainerspecificationforinstancetype-image
	//
	Image *string `field:"optional" json:"image" yaml:"image"`
}

