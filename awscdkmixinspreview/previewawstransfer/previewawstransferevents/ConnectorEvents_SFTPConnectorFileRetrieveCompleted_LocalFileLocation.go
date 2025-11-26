package previewawstransferevents


// Type definition for Local-file-location.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   localFileLocation := &LocalFileLocation{
//   	Bucket: []*string{
//   		jsii.String("bucket"),
//   	},
//   	Domain: []*string{
//   		jsii.String("domain"),
//   	},
//   	Key: []*string{
//   		jsii.String("key"),
//   	},
//   }
//
// Experimental.
type ConnectorEvents_SFTPConnectorFileRetrieveCompleted_LocalFileLocation struct {
	// bucket property.
	//
	// Specify an array of string values to match this event if the actual value of bucket is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Bucket *[]*string `field:"optional" json:"bucket" yaml:"bucket"`
	// domain property.
	//
	// Specify an array of string values to match this event if the actual value of domain is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Domain *[]*string `field:"optional" json:"domain" yaml:"domain"`
	// key property.
	//
	// Specify an array of string values to match this event if the actual value of key is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Key *[]*string `field:"optional" json:"key" yaml:"key"`
}

