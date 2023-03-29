package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   wordCloudOptionsProperty := &WordCloudOptionsProperty{
//   	CloudLayout: jsii.String("cloudLayout"),
//   	MaximumStringLength: jsii.Number(123),
//   	WordCasing: jsii.String("wordCasing"),
//   	WordOrientation: jsii.String("wordOrientation"),
//   	WordPadding: jsii.String("wordPadding"),
//   	WordScaling: jsii.String("wordScaling"),
//   }
//
type CfnDashboard_WordCloudOptionsProperty struct {
	// `CfnDashboard.WordCloudOptionsProperty.CloudLayout`.
	CloudLayout *string `field:"optional" json:"cloudLayout" yaml:"cloudLayout"`
	// `CfnDashboard.WordCloudOptionsProperty.MaximumStringLength`.
	MaximumStringLength *float64 `field:"optional" json:"maximumStringLength" yaml:"maximumStringLength"`
	// `CfnDashboard.WordCloudOptionsProperty.WordCasing`.
	WordCasing *string `field:"optional" json:"wordCasing" yaml:"wordCasing"`
	// `CfnDashboard.WordCloudOptionsProperty.WordOrientation`.
	WordOrientation *string `field:"optional" json:"wordOrientation" yaml:"wordOrientation"`
	// `CfnDashboard.WordCloudOptionsProperty.WordPadding`.
	WordPadding *string `field:"optional" json:"wordPadding" yaml:"wordPadding"`
	// `CfnDashboard.WordCloudOptionsProperty.WordScaling`.
	WordScaling *string `field:"optional" json:"wordScaling" yaml:"wordScaling"`
}

