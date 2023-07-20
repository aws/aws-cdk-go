package awssagemaker


// Represents the overall status of a model package.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   modelPackageStatusItemProperty := &ModelPackageStatusItemProperty{
//   	Name: jsii.String("name"),
//   	Status: jsii.String("status"),
//
//   	// the properties below are optional
//   	FailureReason: jsii.String("failureReason"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-modelpackagestatusitem.html
//
type CfnModelPackage_ModelPackageStatusItemProperty struct {
	// The name of the model package for which the overall status is being reported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-modelpackagestatusitem.html#cfn-sagemaker-modelpackage-modelpackagestatusitem-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The current status.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-modelpackagestatusitem.html#cfn-sagemaker-modelpackage-modelpackagestatusitem-status
	//
	Status *string `field:"required" json:"status" yaml:"status"`
	// if the overall status is `Failed` , the reason for the failure.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-modelpackagestatusitem.html#cfn-sagemaker-modelpackage-modelpackagestatusitem-failurereason
	//
	FailureReason *string `field:"optional" json:"failureReason" yaml:"failureReason"`
}

