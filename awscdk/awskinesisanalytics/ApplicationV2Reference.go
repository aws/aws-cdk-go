package awskinesisanalytics


// A reference to a Application resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   applicationV2Reference := &ApplicationV2Reference{
//   	ApplicationName: jsii.String("applicationName"),
//   }
//
// Deprecated: use `aws-kinesisanalyticsv2` instead.
type ApplicationV2Reference struct {
	// The ApplicationName of the Application resource.
	// Deprecated: use `aws-kinesisanalyticsv2` instead.
	ApplicationName *string `field:"required" json:"applicationName" yaml:"applicationName"`
}

