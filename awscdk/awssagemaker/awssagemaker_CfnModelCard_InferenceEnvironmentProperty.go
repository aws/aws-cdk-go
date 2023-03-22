package awssagemaker


// An overview of a model's inference environment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inferenceEnvironmentProperty := &inferenceEnvironmentProperty{
//   	containerImage: []*string{
//   		jsii.String("containerImage"),
//   	},
//   }
//
type CfnModelCard_InferenceEnvironmentProperty struct {
	// The container used to run the inference environment.
	ContainerImage *[]*string `field:"optional" json:"containerImage" yaml:"containerImage"`
}

