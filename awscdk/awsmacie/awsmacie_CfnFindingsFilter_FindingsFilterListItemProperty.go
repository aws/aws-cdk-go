package awsmacie


// Specifies the unique identifier and custom name of a findings filter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   findingsFilterListItemProperty := &findingsFilterListItemProperty{
//   	id: jsii.String("id"),
//   	name: jsii.String("name"),
//   }
//
type CfnFindingsFilter_FindingsFilterListItemProperty struct {
	// The unique identifier for the filter.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The custom name of the filter.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

