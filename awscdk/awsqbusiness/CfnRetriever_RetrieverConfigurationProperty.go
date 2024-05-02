package awsqbusiness


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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-retriever-retrieverconfiguration.html#cfn-qbusiness-retriever-retrieverconfiguration-kendraindexconfiguration
	//
	KendraIndexConfiguration interface{} `field:"optional" json:"kendraIndexConfiguration" yaml:"kendraIndexConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-retriever-retrieverconfiguration.html#cfn-qbusiness-retriever-retrieverconfiguration-nativeindexconfiguration
	//
	NativeIndexConfiguration interface{} `field:"optional" json:"nativeIndexConfiguration" yaml:"nativeIndexConfiguration"`
}

