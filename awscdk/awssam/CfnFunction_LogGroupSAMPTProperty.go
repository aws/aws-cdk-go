package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   logGroupSAMPTProperty := &LogGroupSAMPTProperty{
//   	LogGroupName: jsii.String("logGroupName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-loggroupsampt.html
//
type CfnFunction_LogGroupSAMPTProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-loggroupsampt.html#cfn-serverless-function-loggroupsampt-loggroupname
	//
	LogGroupName *string `field:"required" json:"logGroupName" yaml:"logGroupName"`
}

