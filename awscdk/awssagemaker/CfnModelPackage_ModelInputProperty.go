package awssagemaker


// Input object for the model.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   modelInputProperty := &ModelInputProperty{
//   	DataInputConfig: jsii.String("dataInputConfig"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-modelinput.html
//
type CfnModelPackage_ModelInputProperty struct {
	// The input configuration object for the model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-modelinput.html#cfn-sagemaker-modelpackage-modelinput-datainputconfig
	//
	DataInputConfig *string `field:"required" json:"dataInputConfig" yaml:"dataInputConfig"`
}

