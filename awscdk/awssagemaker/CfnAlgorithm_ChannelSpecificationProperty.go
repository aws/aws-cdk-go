package awssagemaker


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   channelSpecificationProperty := &ChannelSpecificationProperty{
//   	Name: jsii.String("name"),
//   	SupportedContentTypes: []*string{
//   		jsii.String("supportedContentTypes"),
//   	},
//   	SupportedInputModes: []*string{
//   		jsii.String("supportedInputModes"),
//   	},
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	IsRequired: jsii.Boolean(false),
//   	SupportedCompressionTypes: []*string{
//   		jsii.String("supportedCompressionTypes"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-algorithm-channelspecification.html
//
type CfnAlgorithm_ChannelSpecificationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-algorithm-channelspecification.html#cfn-sagemaker-algorithm-channelspecification-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-algorithm-channelspecification.html#cfn-sagemaker-algorithm-channelspecification-supportedcontenttypes
	//
	SupportedContentTypes *[]*string `field:"required" json:"supportedContentTypes" yaml:"supportedContentTypes"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-algorithm-channelspecification.html#cfn-sagemaker-algorithm-channelspecification-supportedinputmodes
	//
	SupportedInputModes *[]*string `field:"required" json:"supportedInputModes" yaml:"supportedInputModes"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-algorithm-channelspecification.html#cfn-sagemaker-algorithm-channelspecification-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-algorithm-channelspecification.html#cfn-sagemaker-algorithm-channelspecification-isrequired
	//
	IsRequired interface{} `field:"optional" json:"isRequired" yaml:"isRequired"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-algorithm-channelspecification.html#cfn-sagemaker-algorithm-channelspecification-supportedcompressiontypes
	//
	SupportedCompressionTypes *[]*string `field:"optional" json:"supportedCompressionTypes" yaml:"supportedCompressionTypes"`
}

