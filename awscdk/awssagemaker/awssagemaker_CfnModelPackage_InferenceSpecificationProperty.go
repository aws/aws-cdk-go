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
//   inferenceSpecificationProperty := &inferenceSpecificationProperty{
//   	containers: []interface{}{
//   		&modelPackageContainerDefinitionProperty{
//   			image: jsii.String("image"),
//
//   			// the properties below are optional
//   			containerHostname: jsii.String("containerHostname"),
//   			environment: map[string]*string{
//   				"environmentKey": jsii.String("environment"),
//   			},
//   			framework: jsii.String("framework"),
//   			frameworkVersion: jsii.String("frameworkVersion"),
//   			imageDigest: jsii.String("imageDigest"),
//   			modelDataUrl: jsii.String("modelDataUrl"),
//   			modelInput: modelInput,
//   			nearestModelName: jsii.String("nearestModelName"),
//   			productId: jsii.String("productId"),
//   		},
//   	},
//   	supportedContentTypes: []*string{
//   		jsii.String("supportedContentTypes"),
//   	},
//   	supportedResponseMimeTypes: []*string{
//   		jsii.String("supportedResponseMimeTypes"),
//   	},
//
//   	// the properties below are optional
//   	supportedRealtimeInferenceInstanceTypes: []*string{
//   		jsii.String("supportedRealtimeInferenceInstanceTypes"),
//   	},
//   	supportedTransformInstanceTypes: []*string{
//   		jsii.String("supportedTransformInstanceTypes"),
//   	},
//   }
//
type CfnModelPackage_InferenceSpecificationProperty struct {
	// The Amazon ECR registry path of the Docker image that contains the inference code.
	Containers interface{} `field:"required" json:"containers" yaml:"containers"`
	// The supported MIME types for the input data.
	SupportedContentTypes *[]*string `field:"required" json:"supportedContentTypes" yaml:"supportedContentTypes"`
	// The supported MIME types for the output data.
	SupportedResponseMimeTypes *[]*string `field:"required" json:"supportedResponseMimeTypes" yaml:"supportedResponseMimeTypes"`
	// A list of the instance types that are used to generate inferences in real-time.
	//
	// This parameter is required for unversioned models, and optional for versioned models.
	SupportedRealtimeInferenceInstanceTypes *[]*string `field:"optional" json:"supportedRealtimeInferenceInstanceTypes" yaml:"supportedRealtimeInferenceInstanceTypes"`
	// A list of the instance types on which a transformation job can be run or on which an endpoint can be deployed.
	//
	// This parameter is required for unversioned models, and optional for versioned models.
	SupportedTransformInstanceTypes *[]*string `field:"optional" json:"supportedTransformInstanceTypes" yaml:"supportedTransformInstanceTypes"`
}

