package awscloudfront


// The properties to create a Key Value Store.
//
// Example:
//   store := cloudfront.NewKeyValueStore(this, jsii.String("KeyValueStore"), &KeyValueStoreProps{
//   	KeyValueStoreName: jsii.String("KeyValueStore"),
//   	Source: cloudfront.ImportSource_FromAsset(jsii.String("path-to-data.json")),
//   })
//
type KeyValueStoreProps struct {
	// A comment for the Key Value Store.
	// Default: No comment will be specified.
	//
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// The unique name of the Key Value Store.
	// Default: A generated name.
	//
	KeyValueStoreName *string `field:"optional" json:"keyValueStoreName" yaml:"keyValueStoreName"`
	// The import source for the Key Value Store.
	//
	// This will populate the initial items in the Key Value Store. The
	// source data must be in a valid JSON format.
	// Default: No data will be imported to the store.
	//
	Source ImportSource `field:"optional" json:"source" yaml:"source"`
}

