package awsmediatailor


// <p>The HTTP configuration for the source location.</p>.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   httpConfigurationProperty := &HttpConfigurationProperty{
//   	BaseUrl: jsii.String("baseUrl"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-sourcelocation-httpconfiguration.html
//
type CfnSourceLocation_HttpConfigurationProperty struct {
	// <p>The base URL for the source location host server.
	//
	// This string must include the protocol, such as <b>https://</b>.</p>
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-sourcelocation-httpconfiguration.html#cfn-mediatailor-sourcelocation-httpconfiguration-baseurl
	//
	BaseUrl *string `field:"required" json:"baseUrl" yaml:"baseUrl"`
}

