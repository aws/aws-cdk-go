package mixinsawss3objectlambda


// > Amazon S3 Object Lambda will no longer be open to new customers starting on 11/7/2025.
//
// If you would like to use the service, please sign up prior to 11/7/2025. For capabilities similar to S3 Object Lambda, learn more here - [Amazon S3 Object Lambda availability change](https://docs.aws.amazon.com/AmazonS3/latest/userguide/amazons3-ol-change.html) .
//
// A configuration used when creating an Object Lambda Access Point transformation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var contentTransformation interface{}
//
//   transformationConfigurationProperty := &TransformationConfigurationProperty{
//   	Actions: []*string{
//   		jsii.String("actions"),
//   	},
//   	ContentTransformation: contentTransformation,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3objectlambda-accesspoint-transformationconfiguration.html
//
type CfnAccessPointPropsMixin_TransformationConfigurationProperty struct {
	// A container for the action of an Object Lambda Access Point configuration.
	//
	// Valid inputs are `GetObject` , `HeadObject` , `ListObject` , and `ListObjectV2` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3objectlambda-accesspoint-transformationconfiguration.html#cfn-s3objectlambda-accesspoint-transformationconfiguration-actions
	//
	Actions *[]*string `field:"optional" json:"actions" yaml:"actions"`
	// A container for the content transformation of an Object Lambda Access Point configuration.
	//
	// Can include the FunctionArn and FunctionPayload. For more information, see [AwsLambdaTransformation](https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_AwsLambdaTransformation.html) in the *Amazon S3 API Reference* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3objectlambda-accesspoint-transformationconfiguration.html#cfn-s3objectlambda-accesspoint-transformationconfiguration-contenttransformation
	//
	ContentTransformation interface{} `field:"optional" json:"contentTransformation" yaml:"contentTransformation"`
}

