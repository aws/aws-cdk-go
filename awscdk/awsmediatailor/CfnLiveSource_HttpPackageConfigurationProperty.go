package awsmediatailor


// The HTTP package configuration properties for the requested VOD source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   httpPackageConfigurationProperty := &HttpPackageConfigurationProperty{
//   	Path: jsii.String("path"),
//   	SourceGroup: jsii.String("sourceGroup"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-livesource-httppackageconfiguration.html
//
type CfnLiveSource_HttpPackageConfigurationProperty struct {
	// The relative path to the URL for this VOD source.
	//
	// This is combined with `SourceLocation::HttpConfiguration::BaseUrl` to form a valid URL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-livesource-httppackageconfiguration.html#cfn-mediatailor-livesource-httppackageconfiguration-path
	//
	Path *string `field:"required" json:"path" yaml:"path"`
	// The name of the source group.
	//
	// This has to match one of the `Channel::Outputs::SourceGroup` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-livesource-httppackageconfiguration.html#cfn-mediatailor-livesource-httppackageconfiguration-sourcegroup
	//
	SourceGroup *string `field:"required" json:"sourceGroup" yaml:"sourceGroup"`
	// The streaming protocol for this package configuration.
	//
	// Supported values are `HLS` and `DASH` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-livesource-httppackageconfiguration.html#cfn-mediatailor-livesource-httppackageconfiguration-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
}

