package awscloudwatch


// A group of actions within an access grant, together with the resource scopes and context conditions that apply to exactly those actions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   scopedActionsProperty := &ScopedActionsProperty{
//   	Actions: []*string{
//   		jsii.String("actions"),
//   	},
//   	ContextConditions: map[string][]*string{
//   		"contextConditionsKey": []*string{
//   			jsii.String("contextConditions"),
//   		},
//   	},
//   	Resources: []interface{}{
//   		&ResourceScopeProperty{
//   			ResourceArns: []*string{
//   				jsii.String("resourceArns"),
//   			},
//   			ResourceType: jsii.String("resourceType"),
//   			RowScopeGroups: []interface{}{
//   				[]interface{}{
//   					&RowScopeProperty{
//   						Field: jsii.String("field"),
//   						Operator: jsii.String("operator"),
//   						Values: []*string{
//   							jsii.String("values"),
//   						},
//   					},
//   				},
//   			},
//   			SignalTypes: []*string{
//   				jsii.String("signalTypes"),
//   			},
//   			Tags: map[string]*string{
//   				"tagsKey": jsii.String("tags"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-accessgrant-scopedactions.html
//
type CfnAccessGrantPropsMixin_ScopedActionsProperty struct {
	// The actions this group applies to, each in vendor-code:Action form.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-accessgrant-scopedactions.html#cfn-cloudwatch-accessgrant-scopedactions-actions
	//
	Actions *[]*string `field:"optional" json:"actions" yaml:"actions"`
	// Context conditions for fine-grained access control on these actions, as a map of condition keys to lists of allowed values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-accessgrant-scopedactions.html#cfn-cloudwatch-accessgrant-scopedactions-contextconditions
	//
	ContextConditions interface{} `field:"optional" json:"contextConditions" yaml:"contextConditions"`
	// Resource scopes constraining these actions to specific resources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-accessgrant-scopedactions.html#cfn-cloudwatch-accessgrant-scopedactions-resources
	//
	Resources interface{} `field:"optional" json:"resources" yaml:"resources"`
}

