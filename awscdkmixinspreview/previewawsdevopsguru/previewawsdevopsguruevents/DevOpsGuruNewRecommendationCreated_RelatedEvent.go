package previewawsdevopsguruevents


// Type definition for RelatedEvent.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   relatedEvent := &RelatedEvent{
//   	Name: []*string{
//   		jsii.String("name"),
//   	},
//   	Resources: []EventResource{
//   		&EventResource{
//   			Name: []*string{
//   				jsii.String("name"),
//   			},
//   			Type: []*string{
//   				jsii.String("type"),
//   			},
//   		},
//   	},
//   }
//
// Experimental.
type DevOpsGuruNewRecommendationCreated_RelatedEvent struct {
	// name property.
	//
	// Specify an array of string values to match this event if the actual value of name is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Name *[]*string `field:"optional" json:"name" yaml:"name"`
	// resources property.
	//
	// Specify an array of string values to match this event if the actual value of resources is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Resources *[]*DevOpsGuruNewRecommendationCreated_EventResource `field:"optional" json:"resources" yaml:"resources"`
}

