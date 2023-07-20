package awssagemaker


// A structure of additional Inference Specification.
//
// Additional Inference Specification specifies details about inference jobs that can be run with models based on this model package.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var modelInput interface{}
//
//   additionalInferenceSpecificationDefinitionProperty := &AdditionalInferenceSpecificationDefinitionProperty{
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
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	SupportedContentTypes: []*string{
//   		jsii.String("supportedContentTypes"),
//   	},
//   	SupportedRealtimeInferenceInstanceTypes: []*string{
//   		jsii.String("supportedRealtimeInferenceInstanceTypes"),
//   	},
//   	SupportedResponseMimeTypes: []*string{
//   		jsii.String("supportedResponseMimeTypes"),
//   	},
//   	SupportedTransformInstanceTypes: []*string{
//   		jsii.String("supportedTransformInstanceTypes"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-additionalinferencespecificationdefinition.html
//
type CfnModelPackage_AdditionalInferenceSpecificationDefinitionProperty struct {
	// The Amazon ECR registry path of the Docker image that contains the inference code.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-additionalinferencespecificationdefinition.html#cfn-sagemaker-modelpackage-additionalinferencespecificationdefinition-containers
	//
	Containers interface{} `field:"required" json:"containers" yaml:"containers"`
	// A unique name to identify the additional inference specification.
	//
	// The name must be unique within the list of your additional inference specifications for a particular model package.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-additionalinferencespecificationdefinition.html#cfn-sagemaker-modelpackage-additionalinferencespecificationdefinition-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// A description of the additional Inference specification.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-additionalinferencespecificationdefinition.html#cfn-sagemaker-modelpackage-additionalinferencespecificationdefinition-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The supported MIME types for the input data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-additionalinferencespecificationdefinition.html#cfn-sagemaker-modelpackage-additionalinferencespecificationdefinition-supportedcontenttypes
	//
	SupportedContentTypes *[]*string `field:"optional" json:"supportedContentTypes" yaml:"supportedContentTypes"`
	// A list of the instance types that are used to generate inferences in real-time.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-additionalinferencespecificationdefinition.html#cfn-sagemaker-modelpackage-additionalinferencespecificationdefinition-supportedrealtimeinferenceinstancetypes
	//
	SupportedRealtimeInferenceInstanceTypes *[]*string `field:"optional" json:"supportedRealtimeInferenceInstanceTypes" yaml:"supportedRealtimeInferenceInstanceTypes"`
	// The supported MIME types for the output data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-additionalinferencespecificationdefinition.html#cfn-sagemaker-modelpackage-additionalinferencespecificationdefinition-supportedresponsemimetypes
	//
	SupportedResponseMimeTypes *[]*string `field:"optional" json:"supportedResponseMimeTypes" yaml:"supportedResponseMimeTypes"`
	// A list of the instance types on which a transformation job can be run or on which an endpoint can be deployed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-additionalinferencespecificationdefinition.html#cfn-sagemaker-modelpackage-additionalinferencespecificationdefinition-supportedtransforminstancetypes
	//
	SupportedTransformInstanceTypes *[]*string `field:"optional" json:"supportedTransformInstanceTypes" yaml:"supportedTransformInstanceTypes"`
}

