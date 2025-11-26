package previewawss3events


// Type definition for Bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   bucket := &Bucket{
//   	Name: []*string{
//   		jsii.String("name"),
//   	},
//   }
//
// Experimental.
type BucketEvents_ObjectAccessTierChanged_Bucket struct {
	// name property.
	//
	// Specify an array of string values to match this event if the actual value of name is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Filter with the Bucket reference.
	//
	// Experimental.
	Name *[]*string `field:"optional" json:"name" yaml:"name"`
}

