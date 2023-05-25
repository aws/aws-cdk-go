package awscdksagemakeralpha


// Describes the container, as part of model definition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import sagemaker_alpha "github.com/aws/aws-cdk-go/awscdksagemakeralpha"
//
//   var containerImage containerImage
//   var modelData modelData
//
//   containerDefinition := &ContainerDefinition{
//   	Image: containerImage,
//
//   	// the properties below are optional
//   	ContainerHostname: jsii.String("containerHostname"),
//   	Environment: map[string]*string{
//   		"environmentKey": jsii.String("environment"),
//   	},
//   	ModelData: modelData,
//   }
//
// Experimental.
type ContainerDefinition struct {
	// The image used to start a container.
	// Experimental.
	Image ContainerImage `field:"required" json:"image" yaml:"image"`
	// Hostname of the container within an inference pipeline.
	//
	// For single container models, this field
	// is ignored. When specifying a hostname for one ContainerDefinition in a pipeline, hostnames
	// must be specified for all other ContainerDefinitions in that pipeline.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-model-containerdefinition.html#cfn-sagemaker-model-containerdefinition-containerhostname
	//
	// Experimental.
	ContainerHostname *string `field:"optional" json:"containerHostname" yaml:"containerHostname"`
	// A map of environment variables to pass into the container.
	// Experimental.
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// S3 path to the model artifacts.
	// Experimental.
	ModelData ModelData `field:"optional" json:"modelData" yaml:"modelData"`
}

