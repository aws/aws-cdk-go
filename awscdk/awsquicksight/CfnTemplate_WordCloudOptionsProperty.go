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
type CfnTemplate_WordCloudOptionsProperty struct {
	// `CfnTemplate.WordCloudOptionsProperty.CloudLayout`.
	CloudLayout *string `field:"optional" json:"cloudLayout" yaml:"cloudLayout"`
	// `CfnTemplate.WordCloudOptionsProperty.MaximumStringLength`.
	MaximumStringLength *float64 `field:"optional" json:"maximumStringLength" yaml:"maximumStringLength"`
	// `CfnTemplate.WordCloudOptionsProperty.WordCasing`.
	WordCasing *string `field:"optional" json:"wordCasing" yaml:"wordCasing"`
	// `CfnTemplate.WordCloudOptionsProperty.WordOrientation`.
	WordOrientation *string `field:"optional" json:"wordOrientation" yaml:"wordOrientation"`
	// `CfnTemplate.WordCloudOptionsProperty.WordPadding`.
	WordPadding *string `field:"optional" json:"wordPadding" yaml:"wordPadding"`
	// `CfnTemplate.WordCloudOptionsProperty.WordScaling`.
	WordScaling *string `field:"optional" json:"wordScaling" yaml:"wordScaling"`
}

