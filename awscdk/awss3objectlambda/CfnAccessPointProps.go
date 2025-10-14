package awss3objectlambda


// Properties for defining a `CfnAccessPoint`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var contentTransformation interface{}
//
//   cfnAccessPointProps := &CfnAccessPointProps{
//   	ObjectLambdaConfiguration: &ObjectLambdaConfigurationProperty{
//   		SupportingAccessPoint: jsii.String("supportingAccessPoint"),
//   		TransformationConfigurations: []interface{}{
//   			&TransformationConfigurationProperty{
//   				Actions: []*string{
//   					jsii.String("actions"),
//   				},
//   				ContentTransformation: contentTransformation,
//   			},
//   		},
//
//   		// the properties below are optional
//   		AllowedFeatures: []*string{
//   			jsii.String("allowedFeatures"),
//   		},
//   		CloudWatchMetricsEnabled: jsii.Boolean(false),
//   	},
//
//   	// the properties below are optional
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3objectlambda-accesspoint.html
//
type CfnAccessPointProps struct {
	// > Amazon S3 Object Lambda will no longer be open to new customers starting on 11/7/2025.
	//
	// If you would like to use the service, please sign up prior to 11/7/2025. For capabilities similar to S3 Object Lambda, learn more here - [Amazon S3 Object Lambda availability change](https://docs.aws.amazon.com/AmazonS3/latest/userguide/amazons3-ol-change.html) .
	//
	// A configuration used when creating an Object Lambda Access Point.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3objectlambda-accesspoint.html#cfn-s3objectlambda-accesspoint-objectlambdaconfiguration
	//
	ObjectLambdaConfiguration interface{} `field:"required" json:"objectLambdaConfiguration" yaml:"objectLambdaConfiguration"`
	// The name of this access point.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3objectlambda-accesspoint.html#cfn-s3objectlambda-accesspoint-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

