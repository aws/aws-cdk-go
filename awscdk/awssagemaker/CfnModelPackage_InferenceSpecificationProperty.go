package awssagemaker


// Defines how to perform inference generation after a training job is run.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var modelInput interface{}
//
//   inferenceSpecificationProperty := &InferenceSpecificationProperty{
//   	Containers: []interface{}{
//   		&ModelPackageContainerDefinitionProperty{
//   			Image: jsii.String("image"),
//
//   			// the properties below are optional
//   			ContainerHostname: jsii.String("containerHostname"),
//   			Environment: map[string]*string{
//   				"environmentKey": jsii.String("environment"),
//   			},
//   			Framework: jsii.String("framework"),
//   			FrameworkVersion: jsii.String("frameworkVersion"),
//   			ImageDigest: jsii.String("imageDigest"),
//   			ModelDataUrl: jsii.String("modelDataUrl"),
//   			ModelInput: modelInput,
//   			NearestModelName: jsii.String("nearestModelName"),
//   		},
//   	},
//   	SupportedContentTypes: []*string{
//   		jsii.String("supportedContentTypes"),
//   	},
//   	SupportedResponseMimeTypes: []*string{
//   		jsii.String("supportedResponseMimeTypes"),
//   	},
//
//   	// the properties below are optional
//   	SupportedRealtimeInferenceInstanceTypes: []*string{
//   		jsii.String("supportedRealtimeInferenceInstanceTypes"),
//   	},
//   	SupportedTransformInstanceTypes: []*string{
//   		jsii.String("supportedTransformInstanceTypes"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-inferencespecification.html
//
type CfnModelPackage_InferenceSpecificationProperty struct {
	// The Amazon ECR registry path of the Docker image that contains the inference code.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-inferencespecification.html#cfn-sagemaker-modelpackage-inferencespecification-containers
	//
	Containers interface{} `field:"required" json:"containers" yaml:"containers"`
	// The supported MIME types for the input data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-inferencespecification.html#cfn-sagemaker-modelpackage-inferencespecification-supportedcontenttypes
	//
	SupportedContentTypes *[]*string `field:"required" json:"supportedContentTypes" yaml:"supportedContentTypes"`
	// The supported MIME types for the output data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-inferencespecification.html#cfn-sagemaker-modelpackage-inferencespecification-supportedresponsemimetypes
	//
	SupportedResponseMimeTypes *[]*string `field:"required" json:"supportedResponseMimeTypes" yaml:"supportedResponseMimeTypes"`
	// A list of the instance types that are used to generate inferences in real-time.
	//
	// This parameter is required for unversioned models, and optional for versioned models.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-inferencespecification.html#cfn-sagemaker-modelpackage-inferencespecification-supportedrealtimeinferenceinstancetypes
	//
	SupportedRealtimeInferenceInstanceTypes *[]*string `field:"optional" json:"supportedRealtimeInferenceInstanceTypes" yaml:"supportedRealtimeInferenceInstanceTypes"`
	// A list of the instance types on which a transformation job can be run or on which an endpoint can be deployed.
	//
	// This parameter is required for unversioned models, and optional for versioned models.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-inferencespecification.html#cfn-sagemaker-modelpackage-inferencespecification-supportedtransforminstancetypes
	//
	SupportedTransformInstanceTypes *[]*string `field:"optional" json:"supportedTransformInstanceTypes" yaml:"supportedTransformInstanceTypes"`
}

