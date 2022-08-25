package awss3objectlambda


// A configuration used when creating an Object Lambda Access Point transformation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var contentTransformation interface{}
//
//   transformationConfigurationProperty := &transformationConfigurationProperty{
//   	actions: []*string{
//   		jsii.String("actions"),
//   	},
//   	contentTransformation: contentTransformation,
//   }
//
type CfnAccessPoint_TransformationConfigurationProperty struct {
	// A container for the action of an Object Lambda Access Point configuration.
	//
	// Valid input is `GetObject` .
	Actions *[]*string `field:"required" json:"actions" yaml:"actions"`
	// A container for the content transformation of an Object Lambda Access Point configuration.
	//
	// Can include the FunctionArn and FunctionPayload. For more information, see [AwsLambdaTransformation](https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_AwsLambdaTransformation.html) in the *Amazon S3 API Reference* .
	ContentTransformation interface{} `field:"required" json:"contentTransformation" yaml:"contentTransformation"`
}

