package previewawscodecommitevents


// Type definition for RepositoryMetadata.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   repositoryMetadata := &RepositoryMetadata{
//   	AccountId: []*string{
//   		jsii.String("accountId"),
//   	},
//   	Arn: []*string{
//   		jsii.String("arn"),
//   	},
//   	CloneUrlHttp: []*string{
//   		jsii.String("cloneUrlHttp"),
//   	},
//   	CloneUrlSsh: []*string{
//   		jsii.String("cloneUrlSsh"),
//   	},
//   	CreationDate: []*string{
//   		jsii.String("creationDate"),
//   	},
//   	LastModifiedDate: []*string{
//   		jsii.String("lastModifiedDate"),
//   	},
//   	RepositoryDescription: []*string{
//   		jsii.String("repositoryDescription"),
//   	},
//   	RepositoryId: []*string{
//   		jsii.String("repositoryId"),
//   	},
//   	RepositoryName: []*string{
//   		jsii.String("repositoryName"),
//   	},
//   }
//
// Experimental.
type AWSAPICallViaCloudTrail_RepositoryMetadata struct {
	// accountId property.
	//
	// Specify an array of string values to match this event if the actual value of accountId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AccountId *[]*string `field:"optional" json:"accountId" yaml:"accountId"`
	// arn property.
	//
	// Specify an array of string values to match this event if the actual value of arn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Arn *[]*string `field:"optional" json:"arn" yaml:"arn"`
	// cloneUrlHttp property.
	//
	// Specify an array of string values to match this event if the actual value of cloneUrlHttp is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CloneUrlHttp *[]*string `field:"optional" json:"cloneUrlHttp" yaml:"cloneUrlHttp"`
	// cloneUrlSsh property.
	//
	// Specify an array of string values to match this event if the actual value of cloneUrlSsh is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CloneUrlSsh *[]*string `field:"optional" json:"cloneUrlSsh" yaml:"cloneUrlSsh"`
	// creationDate property.
	//
	// Specify an array of string values to match this event if the actual value of creationDate is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CreationDate *[]*string `field:"optional" json:"creationDate" yaml:"creationDate"`
	// lastModifiedDate property.
	//
	// Specify an array of string values to match this event if the actual value of lastModifiedDate is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	LastModifiedDate *[]*string `field:"optional" json:"lastModifiedDate" yaml:"lastModifiedDate"`
	// repositoryDescription property.
	//
	// Specify an array of string values to match this event if the actual value of repositoryDescription is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RepositoryDescription *[]*string `field:"optional" json:"repositoryDescription" yaml:"repositoryDescription"`
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
}

