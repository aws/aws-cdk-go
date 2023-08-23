package awsmediatailor


// <p>The HTTP package configuration properties for the requested VOD source.</p>.
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-vodsource-httppackageconfiguration.html
//
type CfnVodSource_HttpPackageConfigurationProperty struct {
	// <p>The relative path to the URL for this VOD source.
	//
	// This is combined with <code>SourceLocation::HttpConfiguration::BaseUrl</code> to form a valid URL.</p>
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-vodsource-httppackageconfiguration.html#cfn-mediatailor-vodsource-httppackageconfiguration-path
	//
	Path *string `field:"required" json:"path" yaml:"path"`
	// <p>The name of the source group.
	//
	// This has to match one of the <code>Channel::Outputs::SourceGroup</code>.</p>
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-vodsource-httppackageconfiguration.html#cfn-mediatailor-vodsource-httppackageconfiguration-sourcegroup
	//
	SourceGroup *string `field:"required" json:"sourceGroup" yaml:"sourceGroup"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-vodsource-httppackageconfiguration.html#cfn-mediatailor-vodsource-httppackageconfiguration-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
}

