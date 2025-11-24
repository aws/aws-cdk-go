package mixinsawsqbusiness


// Stores an Amazon Kendra index as a retriever.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   kendraIndexConfigurationProperty := &KendraIndexConfigurationProperty{
//   	IndexId: jsii.String("indexId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-retriever-kendraindexconfiguration.html
//
type CfnRetrieverPropsMixin_KendraIndexConfigurationProperty struct {
	// The identifier of the Amazon Kendra index.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-retriever-kendraindexconfiguration.html#cfn-qbusiness-retriever-kendraindexconfiguration-indexid
	//
	IndexId *string `field:"optional" json:"indexId" yaml:"indexId"`
}

