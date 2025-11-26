package previewawscodecommitevents


// Type definition for AdditionalEventData.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   additionalEventData := &AdditionalEventData{
//   	Capabilities: []*string{
//   		jsii.String("capabilities"),
//   	},
//   	Clone: []*string{
//   		jsii.String("clone"),
//   	},
//   	DataTransferred: []*string{
//   		jsii.String("dataTransferred"),
//   	},
//   	Protocol: []*string{
//   		jsii.String("protocol"),
//   	},
//   	RepositoryId: []*string{
//   		jsii.String("repositoryId"),
//   	},
//   	RepositoryName: []*string{
//   		jsii.String("repositoryName"),
//   	},
//   	Shallow: []*string{
//   		jsii.String("shallow"),
//   	},
//   }
//
// Experimental.
type RepositoryEvents_AWSAPICallViaCloudTrail_AdditionalEventData struct {
	// capabilities property.
	//
	// Specify an array of string values to match this event if the actual value of capabilities is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Capabilities *[]*string `field:"optional" json:"capabilities" yaml:"capabilities"`
	// clone property.
	//
	// Specify an array of string values to match this event if the actual value of clone is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Clone *[]*string `field:"optional" json:"clone" yaml:"clone"`
	// dataTransferred property.
	//
	// Specify an array of string values to match this event if the actual value of dataTransferred is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	DataTransferred *[]*string `field:"optional" json:"dataTransferred" yaml:"dataTransferred"`
	// protocol property.
	//
	// Specify an array of string values to match this event if the actual value of protocol is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Protocol *[]*string `field:"optional" json:"protocol" yaml:"protocol"`
	// repositoryId property.
	//
	// Specify an array of string values to match this event if the actual value of repositoryId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RepositoryId *[]*string `field:"optional" json:"repositoryId" yaml:"repositoryId"`
	// repositoryName property.
	//
	// Specify an array of string values to match this event if the actual value of repositoryName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RepositoryName *[]*string `field:"optional" json:"repositoryName" yaml:"repositoryName"`
	// shallow property.
	//
	// Specify an array of string values to match this event if the actual value of shallow is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Shallow *[]*string `field:"optional" json:"shallow" yaml:"shallow"`
}

