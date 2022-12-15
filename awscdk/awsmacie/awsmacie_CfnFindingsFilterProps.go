package awsmacie


// Properties for defining a `CfnFindingsFilter`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFindingsFilterProps := &cfnFindingsFilterProps{
//   	findingCriteria: &findingCriteriaProperty{
//   		criterion: map[string]interface{}{
//   			"criterionKey": &CriterionAdditionalPropertiesProperty{
//   				"eq": []*string{
//   					jsii.String("eq"),
//   				},
//   				"gt": jsii.Number(123),
//   				"gte": jsii.Number(123),
//   				"lt": jsii.Number(123),
//   				"lte": jsii.Number(123),
//   				"neq": []*string{
//   					jsii.String("neq"),
//   				},
//   			},
//   		},
//   	},
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	action: jsii.String("action"),
//   	description: jsii.String("description"),
//   	position: jsii.Number(123),
//   }
//
type CfnFindingsFilterProps struct {
	// The criteria to use to filter findings.
	FindingCriteria interface{} `field:"required" json:"findingCriteria" yaml:"findingCriteria"`
	// A custom name for the findings filter. The name can contain 3-64 characters.
	//
	// Avoid including sensitive data in the name. Users of the account might be able to see the name, depending on the actions that they're allowed to perform in Amazon Macie .
	Name *string `field:"required" json:"name" yaml:"name"`
	// The action to perform on findings that match the filter criteria ( `FindingCriteria` ). Valid values are:.
	//
	// - `ARCHIVE` - Suppress (automatically archive) the findings.
	// - `NOOP` - Don't perform any action on the findings.
	Action *string `field:"optional" json:"action" yaml:"action"`
	// A custom description of the findings filter. The description can contain 1-512 characters.
	//
	// Avoid including sensitive data in the description. Users of the account might be able to see the description, depending on the actions that they're allowed to perform in Amazon Macie .
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The position of the findings filter in the list of saved filters on the Amazon Macie console.
	//
	// This value also determines the order in which the filter is applied to findings, relative to other filters that are also applied to findings.
	Position *float64 `field:"optional" json:"position" yaml:"position"`
}

