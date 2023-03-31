package awsefs


// A tag is a key-value pair attached to a file system.
//
// Allowed characters in the `Key` and `Value` properties are letters, white space, and numbers that can be represented in UTF-8, and the following characters: `+ - = . _ : /`
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   elasticFileSystemTagProperty := &elasticFileSystemTagProperty{
//   	key: jsii.String("key"),
//   	value: jsii.String("value"),
//   }
//
type CfnFileSystem_ElasticFileSystemTagProperty struct {
	// The tag key (String).
	//
	// The key can't start with `aws:` .
	Key *string `field:"required" json:"key" yaml:"key"`
	// The value of the tag key.
	Value *string `field:"required" json:"value" yaml:"value"`
}

