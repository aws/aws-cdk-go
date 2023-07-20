package awsapplicationinsights


// The `AWS::ApplicationInsights::Application CustomComponent` property type describes a custom component by grouping similar standalone instances to monitor.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customComponentProperty := &CustomComponentProperty{
//   	ComponentName: jsii.String("componentName"),
//   	ResourceList: []*string{
//   		jsii.String("resourceList"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-customcomponent.html
//
type CfnApplication_CustomComponentProperty struct {
	// The name of the component.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-customcomponent.html#cfn-applicationinsights-application-customcomponent-componentname
	//
	ComponentName *string `field:"required" json:"componentName" yaml:"componentName"`
	// The list of resource ARNs that belong to the component.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-customcomponent.html#cfn-applicationinsights-application-customcomponent-resourcelist
	//
	ResourceList *[]*string `field:"required" json:"resourceList" yaml:"resourceList"`
}

