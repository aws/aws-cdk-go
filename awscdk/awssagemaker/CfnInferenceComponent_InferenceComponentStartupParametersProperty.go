package awssagemaker


// Settings that take effect while the model container starts up.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inferenceComponentStartupParametersProperty := &InferenceComponentStartupParametersProperty{
//   	ContainerStartupHealthCheckTimeoutInSeconds: jsii.Number(123),
//   	ModelDataDownloadTimeoutInSeconds: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-inferencecomponentstartupparameters.html
//
type CfnInferenceComponent_InferenceComponentStartupParametersProperty struct {
	// The timeout value, in seconds, for your inference container to pass health check by Amazon S3 Hosting.
	//
	// For more information about health check, see [How Your Container Should Respond to Health Check (Ping) Requests](https://docs.aws.amazon.com/sagemaker/latest/dg/your-algorithms-inference-code.html#your-algorithms-inference-algo-ping-requests) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-inferencecomponentstartupparameters.html#cfn-sagemaker-inferencecomponent-inferencecomponentstartupparameters-containerstartuphealthchecktimeoutinseconds
	//
	ContainerStartupHealthCheckTimeoutInSeconds *float64 `field:"optional" json:"containerStartupHealthCheckTimeoutInSeconds" yaml:"containerStartupHealthCheckTimeoutInSeconds"`
	// The timeout value, in seconds, to download and extract the model that you want to host from Amazon S3 to the individual inference instance associated with this inference component.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-inferencecomponentstartupparameters.html#cfn-sagemaker-inferencecomponent-inferencecomponentstartupparameters-modeldatadownloadtimeoutinseconds
	//
	ModelDataDownloadTimeoutInSeconds *float64 `field:"optional" json:"modelDataDownloadTimeoutInSeconds" yaml:"modelDataDownloadTimeoutInSeconds"`
}

