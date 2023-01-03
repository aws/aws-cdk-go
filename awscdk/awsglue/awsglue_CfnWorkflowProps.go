package awsglue


// Properties for defining a `CfnWorkflow`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var defaultRunProperties interface{}
//   var tags interface{}
//
//   cfnWorkflowProps := &cfnWorkflowProps{
//   	defaultRunProperties: defaultRunProperties,
//   	description: jsii.String("description"),
//   	maxConcurrentRuns: jsii.Number(123),
//   	name: jsii.String("name"),
//   	tags: tags,
//   }
//
type CfnWorkflowProps struct {
	// A collection of properties to be used as part of each execution of the workflow.
	DefaultRunProperties interface{} `field:"optional" json:"defaultRunProperties" yaml:"defaultRunProperties"`
	// A description of the workflow.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// You can use this parameter to prevent unwanted multiple updates to data, to control costs, or in some cases, to prevent exceeding the maximum number of concurrent runs of any of the component jobs.
	//
	// If you leave this parameter blank, there is no limit to the number of concurrent workflow runs.
	MaxConcurrentRuns *float64 `field:"optional" json:"maxConcurrentRuns" yaml:"maxConcurrentRuns"`
	// The name of the workflow representing the flow.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The tags to use with this workflow.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

