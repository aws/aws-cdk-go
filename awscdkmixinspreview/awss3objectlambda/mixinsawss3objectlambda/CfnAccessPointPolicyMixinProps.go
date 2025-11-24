package mixinsawss3objectlambda


// Properties for CfnAccessPointPolicyPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var policyDocument interface{}
//
//   cfnAccessPointPolicyMixinProps := &CfnAccessPointPolicyMixinProps{
//   	ObjectLambdaAccessPoint: jsii.String("objectLambdaAccessPoint"),
//   	PolicyDocument: policyDocument,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3objectlambda-accesspointpolicy.html
//
type CfnAccessPointPolicyMixinProps struct {
	// > Amazon S3 Object Lambda will no longer be open to new customers starting on 11/7/2025.
	//
	// If you would like to use the service, please sign up prior to 11/7/2025. For capabilities similar to S3 Object Lambda, learn more here - [Amazon S3 Object Lambda availability change](https://docs.aws.amazon.com/AmazonS3/latest/userguide/amazons3-ol-change.html) .
	//
	// An access point with an attached AWS Lambda function used to access transformed data from an Amazon S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3objectlambda-accesspointpolicy.html#cfn-s3objectlambda-accesspointpolicy-objectlambdaaccesspoint
	//
	ObjectLambdaAccessPoint *string `field:"optional" json:"objectLambdaAccessPoint" yaml:"objectLambdaAccessPoint"`
	// Object Lambda Access Point resource policy document.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3objectlambda-accesspointpolicy.html#cfn-s3objectlambda-accesspointpolicy-policydocument
	//
	PolicyDocument interface{} `field:"optional" json:"policyDocument" yaml:"policyDocument"`
}

