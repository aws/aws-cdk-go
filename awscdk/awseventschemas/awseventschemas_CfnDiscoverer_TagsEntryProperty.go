package awseventschemas


// Tags to associate with the discoverer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tagsEntryProperty := &tagsEntryProperty{
//   	key: jsii.String("key"),
//   	value: jsii.String("value"),
//   }
//
type CfnDiscoverer_TagsEntryProperty struct {
	// They key of a key-value pair.
	Key *string `field:"required" json:"key" yaml:"key"`
	// They value of a key-value pair.
	Value *string `field:"required" json:"value" yaml:"value"`
}

