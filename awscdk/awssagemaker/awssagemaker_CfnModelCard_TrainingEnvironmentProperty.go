package awssagemaker


// SageMaker training image.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   trainingEnvironmentProperty := &trainingEnvironmentProperty{
//   	containerImage: []*string{
//   		jsii.String("containerImage"),
//   	},
//   }
//
type CfnModelCard_TrainingEnvironmentProperty struct {
	// SageMaker inference image URI.
	ContainerImage *[]*string `field:"optional" json:"containerImage" yaml:"containerImage"`
}

