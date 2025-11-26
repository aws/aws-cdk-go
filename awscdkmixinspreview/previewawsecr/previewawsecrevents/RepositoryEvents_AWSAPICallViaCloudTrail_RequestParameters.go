package previewawsecrevents


// Type definition for RequestParameters.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   requestParameters := &RequestParameters{
//   	AcceptedMediaTypes: []*string{
//   		jsii.String("acceptedMediaTypes"),
//   	},
//   	ImageIds: []RequestParametersItem{
//   		&RequestParametersItem{
//   			ImageTag: []*string{
//   				jsii.String("imageTag"),
//   			},
//   		},
//   	},
//   	RegistryId: []*string{
//   		jsii.String("registryId"),
//   	},
//   	RepositoryName: []*string{
//   		jsii.String("repositoryName"),
//   	},
//   }
//
// Experimental.
type RepositoryEvents_AWSAPICallViaCloudTrail_RequestParameters struct {
	// acceptedMediaTypes property.
	//
	// Specify an array of string values to match this event if the actual value of acceptedMediaTypes is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AcceptedMediaTypes *[]*string `field:"optional" json:"acceptedMediaTypes" yaml:"acceptedMediaTypes"`
	// imageIds property.
	//
	// Specify an array of string values to match this event if the actual value of imageIds is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ImageIds *[]*RepositoryEvents_AWSAPICallViaCloudTrail_RequestParametersItem `field:"optional" json:"imageIds" yaml:"imageIds"`
	// registryId property.
	//
	// Specify an array of string values to match this event if the actual value of registryId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RegistryId *[]*string `field:"optional" json:"registryId" yaml:"registryId"`
	// repositoryName property.
	//
	// Specify an array of string values to match this event if the actual value of repositoryName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Filter with the Repository reference.
	//
	// Experimental.
	RepositoryName *[]*string `field:"optional" json:"repositoryName" yaml:"repositoryName"`
}

