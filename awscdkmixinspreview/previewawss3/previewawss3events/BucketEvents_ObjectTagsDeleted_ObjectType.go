package previewawss3events


// Type definition for Object.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   objectType := &ObjectType{
//   	Etag: []*string{
//   		jsii.String("etag"),
//   	},
//   	Key: []*string{
//   		jsii.String("key"),
//   	},
//   	VersionId: []*string{
//   		jsii.String("versionId"),
//   	},
//   }
//
// Experimental.
type BucketEvents_ObjectTagsDeleted_ObjectType struct {
	// etag property.
	//
	// Specify an array of string values to match this event if the actual value of etag is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Etag *[]*string `field:"optional" json:"etag" yaml:"etag"`
	// key property.
	//
	// Specify an array of string values to match this event if the actual value of key is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Key *[]*string `field:"optional" json:"key" yaml:"key"`
	// version-id property.
	//
	// Specify an array of string values to match this event if the actual value of version-id is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	VersionId *[]*string `field:"optional" json:"versionId" yaml:"versionId"`
}

