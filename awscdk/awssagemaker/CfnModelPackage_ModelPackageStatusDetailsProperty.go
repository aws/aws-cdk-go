package awssagemaker


// Specifies the validation and image scan statuses of the model package.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   modelPackageStatusDetailsProperty := &ModelPackageStatusDetailsProperty{
//   	ValidationStatuses: []interface{}{
//   		&ModelPackageStatusItemProperty{
//   			Name: jsii.String("name"),
//   			Status: jsii.String("status"),
//
//   			// the properties below are optional
//   			FailureReason: jsii.String("failureReason"),
//   		},
//   	},
//   }
//
type CfnModelPackage_ModelPackageStatusDetailsProperty struct {
	// The validation status of the model package.
	ValidationStatuses interface{} `field:"optional" json:"validationStatuses" yaml:"validationStatuses"`
}

