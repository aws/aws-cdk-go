package awsappflow


// Determines the prefix that Amazon AppFlow applies to the destination folder name.
//
// You can name your destination folders according to the flow frequency and date.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   prefixConfigProperty := &prefixConfigProperty{
//   	pathPrefixHierarchy: []*string{
//   		jsii.String("pathPrefixHierarchy"),
//   	},
//   	prefixFormat: jsii.String("prefixFormat"),
//   	prefixType: jsii.String("prefixType"),
//   }
//
type CfnFlow_PrefixConfigProperty struct {
	// `CfnFlow.PrefixConfigProperty.PathPrefixHierarchy`.
	PathPrefixHierarchy *[]*string `field:"optional" json:"pathPrefixHierarchy" yaml:"pathPrefixHierarchy"`
	// Determines the level of granularity that's included in the prefix.
	PrefixFormat *string `field:"optional" json:"prefixFormat" yaml:"prefixFormat"`
	// Determines the format of the prefix, and whether it applies to the file name, file path, or both.
	PrefixType *string `field:"optional" json:"prefixType" yaml:"prefixType"`
}

