package mixinsawsmediatailor


// The HTTP package configuration properties for the requested VOD source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   httpPackageConfigurationProperty := &HttpPackageConfigurationProperty{
//   	Path: jsii.String("path"),
//   	SourceGroup: jsii.String("sourceGroup"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-vodsource-httppackageconfiguration.html
//
type CfnVodSourcePropsMixin_HttpPackageConfigurationProperty struct {
	// The relative path to the URL for this VOD source.
	//
	// This is combined with `SourceLocation::HttpConfiguration::BaseUrl` to form a valid URL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-vodsource-httppackageconfiguration.html#cfn-mediatailor-vodsource-httppackageconfiguration-path
	//
	Path *string `field:"optional" json:"path" yaml:"path"`
	// The name of the source group.
	//
	// This has to match one of the `Channel::Outputs::SourceGroup` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-vodsource-httppackageconfiguration.html#cfn-mediatailor-vodsource-httppackageconfiguration-sourcegroup
	//
	SourceGroup *string `field:"optional" json:"sourceGroup" yaml:"sourceGroup"`
	// The streaming protocol for this package configuration.
	//
	// Supported values are `HLS` and `DASH` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-vodsource-httppackageconfiguration.html#cfn-mediatailor-vodsource-httppackageconfiguration-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

