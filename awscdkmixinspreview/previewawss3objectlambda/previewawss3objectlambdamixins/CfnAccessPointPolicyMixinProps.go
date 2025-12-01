package previewawss3objectlambdamixins


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
	// An access point with an attached AWS Lambda function used to access transformed data from an Amazon S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3objectlambda-accesspointpolicy.html#cfn-s3objectlambda-accesspointpolicy-objectlambdaaccesspoint
	//
	ObjectLambdaAccessPoint *string `field:"optional" json:"objectLambdaAccessPoint" yaml:"objectLambdaAccessPoint"`
	// Object Lambda Access Point resource policy document.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3objectlambda-accesspointpolicy.html#cfn-s3objectlambda-accesspointpolicy-policydocument
	//
	PolicyDocument interface{} `field:"optional" json:"policyDocument" yaml:"policyDocument"`
}

