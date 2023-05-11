package awsmacie


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   findingsFilterListItemProperty := &FindingsFilterListItemProperty{
//   	Id: jsii.String("id"),
//   	Name: jsii.String("name"),
//   }
//
type CfnFindingsFilter_FindingsFilterListItemProperty struct {
	// `CfnFindingsFilter.FindingsFilterListItemProperty.Id`.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// `CfnFindingsFilter.FindingsFilterListItemProperty.Name`.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

