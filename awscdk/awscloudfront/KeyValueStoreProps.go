package awscloudfront


// The properties to create a Key Value Store.
//
// Example:
//   storeAsset := cloudfront.NewKeyValueStore(this, jsii.String("KeyValueStoreAsset"), &KeyValueStoreProps{
//   	KeyValueStoreName: jsii.String("KeyValueStoreAsset"),
//   	Source: cloudfront.ImportSource_FromAsset(jsii.String("path-to-data.json")),
//   })
//
//   storeInline := cloudfront.NewKeyValueStore(this, jsii.String("KeyValueStoreInline"), &KeyValueStoreProps{
//   	KeyValueStoreName: jsii.String("KeyValueStoreInline"),
//   	Source: cloudfront.ImportSource_FromInline(jSON.stringify(map[string][]map[string]*string{
//   		"data": []map[string]*string{
//   			map[string]*string{
//   				"key": jsii.String("key1"),
//   				"value": jsii.String("value1"),
//   			},
//   			map[string]*string{
//   				"key": jsii.String("key2"),
//   				"value": jsii.String("value2"),
//   			},
//   		},
//   	})),
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

