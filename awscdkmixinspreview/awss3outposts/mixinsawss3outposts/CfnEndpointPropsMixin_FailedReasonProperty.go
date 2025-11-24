package mixinsawss3outposts


// The failure reason, if any, for a create or delete endpoint operation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   failedReasonProperty := &FailedReasonProperty{
//   	ErrorCode: jsii.String("errorCode"),
//   	Message: jsii.String("message"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3outposts-endpoint-failedreason.html
//
type CfnEndpointPropsMixin_FailedReasonProperty struct {
	// The failure code, if any, for a create or delete endpoint operation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3outposts-endpoint-failedreason.html#cfn-s3outposts-endpoint-failedreason-errorcode
	//
	ErrorCode *string `field:"optional" json:"errorCode" yaml:"errorCode"`
	// Additional error details describing the endpoint failure and recommended action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3outposts-endpoint-failedreason.html#cfn-s3outposts-endpoint-failedreason-message
	//
	Message *string `field:"optional" json:"message" yaml:"message"`
}

