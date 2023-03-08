package awssagemaker


// Specifies the validation and image scan statuses of the model package.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   modelPackageStatusDetailsProperty := &modelPackageStatusDetailsProperty{
//   	validationStatuses: []interface{}{
//   		&modelPackageStatusItemProperty{
//   			name: jsii.String("name"),
//   			status: jsii.String("status"),
//
//   			// the properties below are optional
//   			failureReason: jsii.String("failureReason"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	imageScanStatuses: []interface{}{
//   		&modelPackageStatusItemProperty{
//   			name: jsii.String("name"),
//   			status: jsii.String("status"),
//
//   			// the properties below are optional
//   			failureReason: jsii.String("failureReason"),
//   		},
//   	},
//   }
//
type CfnModelPackage_ModelPackageStatusDetailsProperty struct {
	// The validation status of the model package.
	ValidationStatuses interface{} `field:"required" json:"validationStatuses" yaml:"validationStatuses"`
	// The status of the scan of the Docker image container for the model package.
	ImageScanStatuses interface{} `field:"optional" json:"imageScanStatuses" yaml:"imageScanStatuses"`
}

