package awsresourceexplorer2


// Properties for defining a `CfnView`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnViewProps := &cfnViewProps{
//   	viewName: jsii.String("viewName"),
//
//   	// the properties below are optional
//   	filters: &filtersProperty{
//   		filterString: jsii.String("filterString"),
//   	},
//   	includedProperties: []interface{}{
//   		&includedPropertyProperty{
//   			name: jsii.String("name"),
//   		},
//   	},
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
type CfnViewProps struct {
	// `AWS::ResourceExplorer2::View.ViewName`.
	ViewName *string `field:"required" json:"viewName" yaml:"viewName"`
	// `AWS::ResourceExplorer2::View.Filters`.
	Filters interface{} `field:"optional" json:"filters" yaml:"filters"`
	// `AWS::ResourceExplorer2::View.IncludedProperties`.
	IncludedProperties interface{} `field:"optional" json:"includedProperties" yaml:"includedProperties"`
	// `AWS::ResourceExplorer2::View.Tags`.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

