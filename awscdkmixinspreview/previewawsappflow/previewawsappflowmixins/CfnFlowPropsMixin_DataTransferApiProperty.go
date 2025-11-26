package previewawsappflowmixins


// The API of the connector application that Amazon AppFlow uses to transfer your data.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dataTransferApiProperty := &DataTransferApiProperty{
//   	Name: jsii.String("name"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-datatransferapi.html
//
type CfnFlowPropsMixin_DataTransferApiProperty struct {
	// The name of the connector application API.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-datatransferapi.html#cfn-appflow-flow-datatransferapi-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// You can specify one of the following types:.
	//
	// - **AUTOMATIC** - The default. Optimizes a flow for datasets that fluctuate in size from small to large. For each flow run, Amazon AppFlow chooses to use the SYNC or ASYNC API type based on the amount of data that the run transfers.
	// - **SYNC** - A synchronous API. This type of API optimizes a flow for small to medium-sized datasets.
	// - **ASYNC** - An asynchronous API. This type of API optimizes a flow for large datasets.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-datatransferapi.html#cfn-appflow-flow-datatransferapi-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

