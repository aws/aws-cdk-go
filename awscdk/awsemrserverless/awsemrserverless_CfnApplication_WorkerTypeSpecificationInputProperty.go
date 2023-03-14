package awsemrserverless


// The specifications for a worker type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   workerTypeSpecificationInputProperty := &workerTypeSpecificationInputProperty{
//   	imageConfiguration: &imageConfigurationInputProperty{
//   		imageUri: jsii.String("imageUri"),
//   	},
//   }
//
type CfnApplication_WorkerTypeSpecificationInputProperty struct {
	// The image configuration for a worker type.
	ImageConfiguration interface{} `field:"optional" json:"imageConfiguration" yaml:"imageConfiguration"`
}

