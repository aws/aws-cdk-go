package mixinsawsmediatailor


// The HTTP configuration for the source location.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   httpConfigurationProperty := &HttpConfigurationProperty{
//   	BaseUrl: jsii.String("baseUrl"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-sourcelocation-httpconfiguration.html
//
type CfnSourceLocationPropsMixin_HttpConfigurationProperty struct {
	// The base URL for the source location host server.
	//
	// This string must include the protocol, such as *https://* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-sourcelocation-httpconfiguration.html#cfn-mediatailor-sourcelocation-httpconfiguration-baseurl
	//
	BaseUrl *string `field:"optional" json:"baseUrl" yaml:"baseUrl"`
}

