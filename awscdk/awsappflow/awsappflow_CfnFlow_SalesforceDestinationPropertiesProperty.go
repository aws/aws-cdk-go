package awsappflow


// The properties that are applied when Salesforce is being used as a destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   salesforceDestinationPropertiesProperty := &salesforceDestinationPropertiesProperty{
//   	object: jsii.String("object"),
//
//   	// the properties below are optional
//   	dataTransferApi: jsii.String("dataTransferApi"),
//   	errorHandlingConfig: &errorHandlingConfigProperty{
//   		bucketName: jsii.String("bucketName"),
//   		bucketPrefix: jsii.String("bucketPrefix"),
//   		failOnFirstError: jsii.Boolean(false),
//   	},
//   	idFieldNames: []*string{
//   		jsii.String("idFieldNames"),
//   	},
//   	writeOperationType: jsii.String("writeOperationType"),
//   }
//
type CfnFlow_SalesforceDestinationPropertiesProperty struct {
	// The object specified in the Salesforce flow destination.
	Object *string `field:"required" json:"object" yaml:"object"`
	// Specifies which Salesforce API is used by Amazon AppFlow when your flow transfers data to Salesforce.
	//
	// - **AUTOMATIC** - The default. Amazon AppFlow selects which API to use based on the number of records that your flow transfers to Salesforce. If your flow transfers fewer than 1,000 records, Amazon AppFlow uses Salesforce REST API. If your flow transfers 1,000 records or more, Amazon AppFlow uses Salesforce Bulk API 2.0.
	//
	// Each of these Salesforce APIs structures data differently. If Amazon AppFlow selects the API automatically, be aware that, for recurring flows, the data output might vary from one flow run to the next. For example, if a flow runs daily, it might use REST API on one day to transfer 900 records, and it might use Bulk API 2.0 on the next day to transfer 1,100 records. For each of these flow runs, the respective Salesforce API formats the data differently. Some of the differences include how dates are formatted and null values are represented. Also, Bulk API 2.0 doesn't transfer Salesforce compound fields.
	//
	// By choosing this option, you optimize flow performance for both small and large data transfers, but the tradeoff is inconsistent formatting in the output.
	// - **BULKV2** - Amazon AppFlow uses only Salesforce Bulk API 2.0. This API runs asynchronous data transfers, and it's optimal for large sets of data. By choosing this option, you ensure that your flow writes consistent output, but you optimize performance only for large data transfers.
	//
	// Note that Bulk API 2.0 does not transfer Salesforce compound fields.
	// - **REST_SYNC** - Amazon AppFlow uses only Salesforce REST API. By choosing this option, you ensure that your flow writes consistent output, but you decrease performance for large data transfers that are better suited for Bulk API 2.0. In some cases, if your flow attempts to transfer a vary large set of data, it might fail with a timed out error.
	DataTransferApi *string `field:"optional" json:"dataTransferApi" yaml:"dataTransferApi"`
	// The settings that determine how Amazon AppFlow handles an error when placing data in the Salesforce destination.
	//
	// For example, this setting would determine if the flow should fail after one insertion error, or continue and attempt to insert every record regardless of the initial failure. `ErrorHandlingConfig` is a part of the destination connector details.
	ErrorHandlingConfig interface{} `field:"optional" json:"errorHandlingConfig" yaml:"errorHandlingConfig"`
	// The name of the field that Amazon AppFlow uses as an ID when performing a write operation such as update or delete.
	IdFieldNames *[]*string `field:"optional" json:"idFieldNames" yaml:"idFieldNames"`
	// This specifies the type of write operation to be performed in Salesforce.
	//
	// When the value is `UPSERT` , then `idFieldNames` is required.
	WriteOperationType *string `field:"optional" json:"writeOperationType" yaml:"writeOperationType"`
}

