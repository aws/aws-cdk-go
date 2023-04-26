package awsemrserverless


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   workerTypeSpecificationInputProperty := &WorkerTypeSpecificationInputProperty{
//   	ImageConfiguration: &ImageConfigurationInputProperty{
//   		ImageUri: jsii.String("imageUri"),
//   	},
//   }
//
type CfnApplication_WorkerTypeSpecificationInputProperty struct {
	// `CfnApplication.WorkerTypeSpecificationInputProperty.ImageConfiguration`.
	ImageConfiguration interface{} `field:"optional" json:"imageConfiguration" yaml:"imageConfiguration"`
}

