package awsglue


// Configuration for the validation endpoint.
//
// Only supports RequestMethod and RequestPath.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   validationEndpointConfigurationProperty := &ValidationEndpointConfigurationProperty{
//   	RequestMethod: jsii.String("requestMethod"),
//   	RequestPath: jsii.String("requestPath"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-validationendpointconfiguration.html
//
type CfnConnectionTypePropsMixin_ValidationEndpointConfigurationProperty struct {
	// The HTTP method to use.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-validationendpointconfiguration.html#cfn-glue-connectiontype-validationendpointconfiguration-requestmethod
	//
	RequestMethod *string `field:"optional" json:"requestMethod" yaml:"requestMethod"`
	// The URL path for the REST endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connectiontype-validationendpointconfiguration.html#cfn-glue-connectiontype-validationendpointconfiguration-requestpath
	//
	RequestPath *string `field:"optional" json:"requestPath" yaml:"requestPath"`
}

