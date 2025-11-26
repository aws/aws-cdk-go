package previewawsmediatailormixins


// The optional configuration for a server that serves segments.
//
// Use this if you want the segment delivery server to be different from the source location server. For example, you can configure your source location server to be an origination server, such as MediaPackage, and the segment delivery server to be a content delivery network (CDN), such as CloudFront. If you don't specify a segment delivery server, then the source location server is used.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   defaultSegmentDeliveryConfigurationProperty := &DefaultSegmentDeliveryConfigurationProperty{
//   	BaseUrl: jsii.String("baseUrl"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-sourcelocation-defaultsegmentdeliveryconfiguration.html
//
type CfnSourceLocationPropsMixin_DefaultSegmentDeliveryConfigurationProperty struct {
	// The hostname of the server that will be used to serve segments.
	//
	// This string must include the protocol, such as *https://* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-sourcelocation-defaultsegmentdeliveryconfiguration.html#cfn-mediatailor-sourcelocation-defaultsegmentdeliveryconfiguration-baseurl
	//
	BaseUrl *string `field:"optional" json:"baseUrl" yaml:"baseUrl"`
}

