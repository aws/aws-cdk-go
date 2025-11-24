package mixinsawsqbusiness


// Configuration information for an Amazon Q Business index.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   nativeIndexConfigurationProperty := &NativeIndexConfigurationProperty{
//   	IndexId: jsii.String("indexId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-retriever-nativeindexconfiguration.html
//
type CfnRetrieverPropsMixin_NativeIndexConfigurationProperty struct {
	// The identifier for the Amazon Q Business index.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-retriever-nativeindexconfiguration.html#cfn-qbusiness-retriever-nativeindexconfiguration-indexid
	//
	IndexId *string `field:"optional" json:"indexId" yaml:"indexId"`
}

