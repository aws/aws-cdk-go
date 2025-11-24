package mixinsawsmediatailor


// The segment delivery configuration settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   segmentDeliveryConfigurationProperty := &SegmentDeliveryConfigurationProperty{
//   	BaseUrl: jsii.String("baseUrl"),
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-sourcelocation-segmentdeliveryconfiguration.html
//
type CfnSourceLocationPropsMixin_SegmentDeliveryConfigurationProperty struct {
	// The base URL of the host or path of the segment delivery server that you're using to serve segments.
	//
	// This is typically a content delivery network (CDN). The URL can be absolute or relative. To use an absolute URL include the protocol, such as `https://example.com/some/path` . To use a relative URL specify the relative path, such as `/some/path*` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-sourcelocation-segmentdeliveryconfiguration.html#cfn-mediatailor-sourcelocation-segmentdeliveryconfiguration-baseurl
	//
	BaseUrl *string `field:"optional" json:"baseUrl" yaml:"baseUrl"`
	// A unique identifier used to distinguish between multiple segment delivery configurations in a source location.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-sourcelocation-segmentdeliveryconfiguration.html#cfn-mediatailor-sourcelocation-segmentdeliveryconfiguration-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

