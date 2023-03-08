package awsiotanalytics


// Information about the versioning of dataset contents.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   versioningConfigurationProperty := &versioningConfigurationProperty{
//   	maxVersions: jsii.Number(123),
//   	unlimited: jsii.Boolean(false),
//   }
//
type CfnDataset_VersioningConfigurationProperty struct {
	// How many versions of dataset contents are kept.
	//
	// The `unlimited` parameter must be `false` .
	MaxVersions *float64 `field:"optional" json:"maxVersions" yaml:"maxVersions"`
	// If true, unlimited versions of dataset contents are kept.
	Unlimited interface{} `field:"optional" json:"unlimited" yaml:"unlimited"`
}

