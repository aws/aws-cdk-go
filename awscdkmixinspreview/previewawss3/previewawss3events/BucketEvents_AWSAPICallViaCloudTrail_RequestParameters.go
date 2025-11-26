package previewawss3events


// Type definition for RequestParameters.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   requestParameters := &RequestParameters{
//   	BucketName: []*string{
//   		jsii.String("bucketName"),
//   	},
//   	Host: []*string{
//   		jsii.String("host"),
//   	},
//   	Key: []*string{
//   		jsii.String("key"),
//   	},
//   	LegalHold: []*string{
//   		jsii.String("legalHold"),
//   	},
//   	Retention: []*string{
//   		jsii.String("retention"),
//   	},
//   }
//
// Experimental.
type BucketEvents_AWSAPICallViaCloudTrail_RequestParameters struct {
	// bucketName property.
	//
	// Specify an array of string values to match this event if the actual value of bucketName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Filter with the Bucket reference.
	//
	// Experimental.
	BucketName *[]*string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// Host property.
	//
	// Specify an array of string values to match this event if the actual value of Host is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Host *[]*string `field:"optional" json:"host" yaml:"host"`
	// key property.
	//
	// Specify an array of string values to match this event if the actual value of key is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Key *[]*string `field:"optional" json:"key" yaml:"key"`
	// legal-hold property.
	//
	// Specify an array of string values to match this event if the actual value of legal-hold is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	LegalHold *[]*string `field:"optional" json:"legalHold" yaml:"legalHold"`
	// retention property.
	//
	// Specify an array of string values to match this event if the actual value of retention is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Retention *[]*string `field:"optional" json:"retention" yaml:"retention"`
}

