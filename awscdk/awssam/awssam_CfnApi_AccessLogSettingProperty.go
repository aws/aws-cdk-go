package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accessLogSettingProperty := &accessLogSettingProperty{
//   	destinationArn: jsii.String("destinationArn"),
//   	format: jsii.String("format"),
//   }
//
type CfnApi_AccessLogSettingProperty struct {
	// `CfnApi.AccessLogSettingProperty.DestinationArn`.
	DestinationArn *string `field:"optional" json:"destinationArn" yaml:"destinationArn"`
	// `CfnApi.AccessLogSettingProperty.Format`.
	Format *string `field:"optional" json:"format" yaml:"format"`
}

