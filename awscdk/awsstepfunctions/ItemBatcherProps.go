package awsstepfunctions


// Interface for ItemBatcher configuration properties.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var batchInput interface{}
//
//   itemBatcherProps := &ItemBatcherProps{
//   	BatchInput: batchInput,
//   	MaxInputBytesPerBatch: jsii.Number(123),
//   	MaxInputBytesPerBatchPath: jsii.String("maxInputBytesPerBatchPath"),
//   	MaxItemsPerBatch: jsii.Number(123),
//   	MaxItemsPerBatchPath: jsii.String("maxItemsPerBatchPath"),
//   }
//
type ItemBatcherProps struct {
	// BatchInput.
	//
	// Fixed JSON input to include in each batch passed to each child workflow execution.
	// Default: - No batchInput.
	//
	BatchInput *map[string]interface{} `field:"optional" json:"batchInput" yaml:"batchInput"`
	// MaxInputBytesPerBatch.
	//
	// Specifies the maximum number of bytes that each child workflow execution processes, as static number.
	// Default: - uses value of `maxInputBytesPerBatchPath` as the max size per batch,
	// no limits on the batch size under the 256KB limit if that property was also not provided.
	//
	MaxInputBytesPerBatch *float64 `field:"optional" json:"maxInputBytesPerBatch" yaml:"maxInputBytesPerBatch"`
	// MaxInputBytesPerBatchPath.
	//
	// Specifies the maximum number of bytes that each child workflow execution processes, as JsonPath.
	// Default: - uses value of `maxInputBytesPerBatch` as the max size per batch,
	// no limits on the batch size under the 256KB limit if that property was also not provided.
	//
	MaxInputBytesPerBatchPath *string `field:"optional" json:"maxInputBytesPerBatchPath" yaml:"maxInputBytesPerBatchPath"`
	// MaxItemsPerBatch.
	//
	// Specifies the maximum number of items that each child workflow execution processes, as static number.
	// Default: - uses value of `maxItemsPerBatchPath` as the max items per batch,
	// no limits on the number of items in a batch under the 256KB limit if that property was also not provided.
	//
	MaxItemsPerBatch *float64 `field:"optional" json:"maxItemsPerBatch" yaml:"maxItemsPerBatch"`
	// MaxItemsPerBatchPath.
	//
	// Specifies the maximum number of items that each child workflow execution processes, as JsonPath.
	// Default: - uses value of `maxItemsPerBatch` as the max items per batch,
	// no limits on the number of items in a batch under the 256KB limit if that property was also not provided.
	//
	MaxItemsPerBatchPath *string `field:"optional" json:"maxItemsPerBatchPath" yaml:"maxItemsPerBatchPath"`
}

