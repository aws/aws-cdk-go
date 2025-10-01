package awssagemaker


// Defines how to perform inference generation after a training job is run.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inferenceSpecificationProperty := &InferenceSpecificationProperty{
//   	Containers: []interface{}{
//   		&ContainerProperty{
//   			Image: jsii.String("image"),
//
//   			// the properties below are optional
//   			ModelDataUrl: jsii.String("modelDataUrl"),
//   			NearestModelName: jsii.String("nearestModelName"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-inferencespecification.html
//
type CfnModelCard_InferenceSpecificationProperty struct {
	// The Amazon ECR registry path of the Docker image that contains the inference code.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-inferencespecification.html#cfn-sagemaker-modelcard-inferencespecification-containers
	//
	Containers interface{} `field:"required" json:"containers" yaml:"containers"`
}

