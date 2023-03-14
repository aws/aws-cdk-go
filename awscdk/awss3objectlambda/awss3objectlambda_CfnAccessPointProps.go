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
//   cfnAccessPointProps := &cfnAccessPointProps{
//   	objectLambdaConfiguration: &objectLambdaConfigurationProperty{
//   		supportingAccessPoint: jsii.String("supportingAccessPoint"),
//   		transformationConfigurations: []interface{}{
//   			&transformationConfigurationProperty{
//   				actions: []*string{
//   					jsii.String("actions"),
//   				},
//   				contentTransformation: contentTransformation,
//   			},
//   		},
//
//   		// the properties below are optional
//   		allowedFeatures: []*string{
//   			jsii.String("allowedFeatures"),
//   		},
//   		cloudWatchMetricsEnabled: jsii.Boolean(false),
//   	},
//
//   	// the properties below are optional
//   	name: jsii.String("name"),
//   }
//
type CfnAccessPointProps struct {
	// A configuration used when creating an Object Lambda Access Point.
	ObjectLambdaConfiguration interface{} `field:"required" json:"objectLambdaConfiguration" yaml:"objectLambdaConfiguration"`
	// The name of this access point.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

