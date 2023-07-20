package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   parameterNameSAMPTProperty := &ParameterNameSAMPTProperty{
//   	ParameterName: jsii.String("parameterName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-parameternamesampt.html
//
type CfnFunction_ParameterNameSAMPTProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-parameternamesampt.html#cfn-serverless-function-parameternamesampt-parametername
	//
	ParameterName *string `field:"required" json:"parameterName" yaml:"parameterName"`
}

