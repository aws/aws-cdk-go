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
//   objectLambdaConfigurationProperty := &objectLambdaConfigurationProperty{
//   	supportingAccessPoint: jsii.String("supportingAccessPoint"),
//   	transformationConfigurations: []interface{}{
//   		&transformationConfigurationProperty{
//   			actions: []*string{
//   				jsii.String("actions"),
//   			},
//   			contentTransformation: contentTransformation,
//   		},
//   	},
//
//   	// the properties below are optional
//   	allowedFeatures: []*string{
//   		jsii.String("allowedFeatures"),
//   	},
//   	cloudWatchMetricsEnabled: jsii.Boolean(false),
//   }
//
type CfnAccessPoint_ObjectLambdaConfigurationProperty struct {
	// Standard access point associated with the Object Lambda Access Point.
	SupportingAccessPoint *string `field:"required" json:"supportingAccessPoint" yaml:"supportingAccessPoint"`
	// A container for transformation configurations for an Object Lambda Access Point.
	TransformationConfigurations interface{} `field:"required" json:"transformationConfigurations" yaml:"transformationConfigurations"`
	// A container for allowed features.
	//
	// Valid inputs are `GetObject-Range` and `GetObject-PartNumber` .
	AllowedFeatures *[]*string `field:"optional" json:"allowedFeatures" yaml:"allowedFeatures"`
	// A container for whether the CloudWatch metrics configuration is enabled.
	CloudWatchMetricsEnabled interface{} `field:"optional" json:"cloudWatchMetricsEnabled" yaml:"cloudWatchMetricsEnabled"`
}

