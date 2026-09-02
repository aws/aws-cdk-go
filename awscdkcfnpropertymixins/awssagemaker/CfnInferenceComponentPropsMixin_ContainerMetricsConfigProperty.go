package awssagemaker


// The configuration for container metrics scraping.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   containerMetricsConfigProperty := &ContainerMetricsConfigProperty{
//   	MetricsEndpoints: []interface{}{
//   		&MetricsEndpointProperty{
//   			MetricPublishFrequencyInSeconds: jsii.Number(123),
//   			MetricsEndpointPath: jsii.String("metricsEndpointPath"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-containermetricsconfig.html
//
type CfnInferenceComponentPropsMixin_ContainerMetricsConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-containermetricsconfig.html#cfn-sagemaker-inferencecomponent-containermetricsconfig-metricsendpoints
	//
	MetricsEndpoints interface{} `field:"optional" json:"metricsEndpoints" yaml:"metricsEndpoints"`
}

