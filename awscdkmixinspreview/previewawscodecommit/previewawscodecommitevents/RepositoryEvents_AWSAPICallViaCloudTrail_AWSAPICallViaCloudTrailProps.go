package previewawscodecommitevents


// Props type for Repository aws.codecommit@AWSAPICallViaCloudTrail event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   aWSAPICallViaCloudTrailProps := &AWSAPICallViaCloudTrailProps{
//   	AdditionalEventData: &AdditionalEventData{
//   		Capabilities: []*string{
//   			jsii.String("capabilities"),
//   		},
//   		Clone: []*string{
//   			jsii.String("clone"),
//   		},
//   		DataTransferred: []*string{
//   			jsii.String("dataTransferred"),
//   		},
//   		Protocol: []*string{
//   			jsii.String("protocol"),
//   		},
//   		RepositoryId: []*string{
//   			jsii.String("repositoryId"),
//   		},
//   		RepositoryName: []*string{
//   			jsii.String("repositoryName"),
//   		},
//   		Shallow: []*string{
//   			jsii.String("shallow"),
//   		},
//   	},
//   }
//
// Experimental.
type RepositoryEvents_AWSAPICallViaCloudTrail_AWSAPICallViaCloudTrailProps struct {
	// additionalEventData property.
	//
	// Specify an array of string values to match this event if the actual value of additionalEventData is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AdditionalEventData *RepositoryEvents_AWSAPICallViaCloudTrail_AdditionalEventData `field:"optional" json:"additionalEventData" yaml:"additionalEventData"`
}

