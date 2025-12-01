package previewawss3objectlambdamixins


// Properties for CfnAccessPointPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var contentTransformation interface{}
//
//   cfnAccessPointMixinProps := &CfnAccessPointMixinProps{
//   	Name: jsii.String("name"),
//   	ObjectLambdaConfiguration: &ObjectLambdaConfigurationProperty{
//   		AllowedFeatures: []*string{
//   			jsii.String("allowedFeatures"),
//   		},
//   		CloudWatchMetricsEnabled: jsii.Boolean(false),
//   		SupportingAccessPoint: jsii.String("supportingAccessPoint"),
//   		TransformationConfigurations: []interface{}{
//   			&TransformationConfigurationProperty{
//   				Actions: []*string{
//   					jsii.String("actions"),
//   				},
//   				ContentTransformation: contentTransformation,
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3objectlambda-accesspoint.html
//
type CfnAccessPointMixinProps struct {
	// The name of this access point.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3objectlambda-accesspoint.html#cfn-s3objectlambda-accesspoint-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A configuration used when creating an Object Lambda Access Point.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3objectlambda-accesspoint.html#cfn-s3objectlambda-accesspoint-objectlambdaconfiguration
	//
	ObjectLambdaConfiguration interface{} `field:"optional" json:"objectLambdaConfiguration" yaml:"objectLambdaConfiguration"`
}

