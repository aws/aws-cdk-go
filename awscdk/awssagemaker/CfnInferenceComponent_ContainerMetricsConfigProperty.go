package awssagemaker


// The configuration for container metrics scraping.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   containerMetricsConfigProperty := &ContainerMetricsConfigProperty{
//   	MetricsEndpoints: []interface{}{
//   		&MetricsEndpointProperty{
//   			MetricsEndpointPath: jsii.String("metricsEndpointPath"),
//
//   			// the properties below are optional
//   			MetricPublishFrequencyInSeconds: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-containermetricsconfig.html
//
type CfnInferenceComponent_ContainerMetricsConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-containermetricsconfig.html#cfn-sagemaker-inferencecomponent-containermetricsconfig-metricsendpoints
	//
	MetricsEndpoints interface{} `field:"required" json:"metricsEndpoints" yaml:"metricsEndpoints"`
}

