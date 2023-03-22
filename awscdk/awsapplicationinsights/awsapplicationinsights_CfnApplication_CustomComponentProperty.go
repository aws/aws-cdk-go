package awsapplicationinsights


// The `AWS::ApplicationInsights::Application CustomComponent` property type describes a custom component by grouping similar standalone instances to monitor.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customComponentProperty := &customComponentProperty{
//   	componentName: jsii.String("componentName"),
//   	resourceList: []*string{
//   		jsii.String("resourceList"),
//   	},
//   }
//
type CfnApplication_CustomComponentProperty struct {
	// The name of the component.
	ComponentName *string `field:"required" json:"componentName" yaml:"componentName"`
	// The list of resource ARNs that belong to the component.
	ResourceList *[]*string `field:"required" json:"resourceList" yaml:"resourceList"`
}

