package awss3objectlambda


// A configuration used when creating an Object Lambda Access Point.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var contentTransformation interface{}
//
//   objectLambdaConfigurationProperty := &ObjectLambdaConfigurationProperty{
//   	SupportingAccessPoint: jsii.String("supportingAccessPoint"),
//   	TransformationConfigurations: []interface{}{
//   		&TransformationConfigurationProperty{
//   			Actions: []*string{
//   				jsii.String("actions"),
//   			},
//   			ContentTransformation: contentTransformation,
//   		},
//   	},
//
//   	// the properties below are optional
//   	AllowedFeatures: []*string{
//   		jsii.String("allowedFeatures"),
//   	},
//   	CloudWatchMetricsEnabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3objectlambda-accesspoint-objectlambdaconfiguration.html
//
type CfnAccessPoint_ObjectLambdaConfigurationProperty struct {
	// Standard access point associated with the Object Lambda Access Point.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3objectlambda-accesspoint-objectlambdaconfiguration.html#cfn-s3objectlambda-accesspoint-objectlambdaconfiguration-supportingaccesspoint
	//
	SupportingAccessPoint *string `field:"required" json:"supportingAccessPoint" yaml:"supportingAccessPoint"`
	// A container for transformation configurations for an Object Lambda Access Point.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3objectlambda-accesspoint-objectlambdaconfiguration.html#cfn-s3objectlambda-accesspoint-objectlambdaconfiguration-transformationconfigurations
	//
	TransformationConfigurations interface{} `field:"required" json:"transformationConfigurations" yaml:"transformationConfigurations"`
	// A container for allowed features.
	//
	// Valid inputs are `GetObject-Range` , `GetObject-PartNumber` , `HeadObject-Range` , and `HeadObject-PartNumber` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3objectlambda-accesspoint-objectlambdaconfiguration.html#cfn-s3objectlambda-accesspoint-objectlambdaconfiguration-allowedfeatures
	//
	AllowedFeatures *[]*string `field:"optional" json:"allowedFeatures" yaml:"allowedFeatures"`
	// A container for whether the CloudWatch metrics configuration is enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3objectlambda-accesspoint-objectlambdaconfiguration.html#cfn-s3objectlambda-accesspoint-objectlambdaconfiguration-cloudwatchmetricsenabled
	//
	CloudWatchMetricsEnabled interface{} `field:"optional" json:"cloudWatchMetricsEnabled" yaml:"cloudWatchMetricsEnabled"`
}

