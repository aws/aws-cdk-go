package awssagemaker


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   modelInputProperty := &ModelInputProperty{
//   	DataInputConfig: jsii.String("dataInputConfig"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-algorithm-modelinput.html
//
type CfnAlgorithm_ModelInputProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-algorithm-modelinput.html#cfn-sagemaker-algorithm-modelinput-datainputconfig
	//
	DataInputConfig *string `field:"required" json:"dataInputConfig" yaml:"dataInputConfig"`
}

