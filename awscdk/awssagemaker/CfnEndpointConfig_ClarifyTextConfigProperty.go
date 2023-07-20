package awssagemaker


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clarifyTextConfigProperty := &ClarifyTextConfigProperty{
//   	Granularity: jsii.String("granularity"),
//   	Language: jsii.String("language"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-clarifytextconfig.html
//
type CfnEndpointConfig_ClarifyTextConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-clarifytextconfig.html#cfn-sagemaker-endpointconfig-clarifytextconfig-granularity
	//
	Granularity *string `field:"required" json:"granularity" yaml:"granularity"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-clarifytextconfig.html#cfn-sagemaker-endpointconfig-clarifytextconfig-language
	//
	Language *string `field:"required" json:"language" yaml:"language"`
}

