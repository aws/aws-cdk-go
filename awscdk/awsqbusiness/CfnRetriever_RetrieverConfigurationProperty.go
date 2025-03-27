package awsqbusiness


// Provides information on how the retriever used for your Amazon Q Business application is configured.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   retrieverConfigurationProperty := &RetrieverConfigurationProperty{
//   	KendraIndexConfiguration: &KendraIndexConfigurationProperty{
//   		IndexId: jsii.String("indexId"),
//   	},
//   	NativeIndexConfiguration: &NativeIndexConfigurationProperty{
//   		IndexId: jsii.String("indexId"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-retriever-retrieverconfiguration.html
//
type CfnRetriever_RetrieverConfigurationProperty struct {
	// Provides information on how the Amazon Kendra index used as a retriever for your Amazon Q Business application is configured.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-retriever-retrieverconfiguration.html#cfn-qbusiness-retriever-retrieverconfiguration-kendraindexconfiguration
	//
	KendraIndexConfiguration interface{} `field:"optional" json:"kendraIndexConfiguration" yaml:"kendraIndexConfiguration"`
	// Provides information on how a Amazon Q Business index used as a retriever for your Amazon Q Business application is configured.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-retriever-retrieverconfiguration.html#cfn-qbusiness-retriever-retrieverconfiguration-nativeindexconfiguration
	//
	NativeIndexConfiguration interface{} `field:"optional" json:"nativeIndexConfiguration" yaml:"nativeIndexConfiguration"`
}

