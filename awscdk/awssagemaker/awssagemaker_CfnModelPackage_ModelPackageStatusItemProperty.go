package awssagemaker


// Represents the overall status of a model package.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   modelPackageStatusItemProperty := &modelPackageStatusItemProperty{
//   	name: jsii.String("name"),
//   	status: jsii.String("status"),
//
//   	// the properties below are optional
//   	failureReason: jsii.String("failureReason"),
//   }
//
type CfnModelPackage_ModelPackageStatusItemProperty struct {
	// The name of the model package for which the overall status is being reported.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The current status.
	Status *string `field:"required" json:"status" yaml:"status"`
	// if the overall status is `Failed` , the reason for the failure.
	FailureReason *string `field:"optional" json:"failureReason" yaml:"failureReason"`
}

