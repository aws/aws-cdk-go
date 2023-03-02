package awsquicksight


// Secure Socket Layer (SSL) properties that apply when Amazon QuickSight connects to your underlying data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sslPropertiesProperty := &sslPropertiesProperty{
//   	disableSsl: jsii.Boolean(false),
//   }
//
type CfnDataSource_SslPropertiesProperty struct {
	// A Boolean option to control whether SSL should be disabled.
	DisableSsl interface{} `field:"optional" json:"disableSsl" yaml:"disableSsl"`
}

