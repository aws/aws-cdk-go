package awsresourcegroups


// Specifies details within a `ResourceQuery` structure that determines the membership of the resource group.
//
// The contents required in the `Query` structure are determined by the `Type` property of the containing `ResourceQuery` structure.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   queryProperty := &QueryProperty{
//   	ResourceTypeFilters: []*string{
//   		jsii.String("resourceTypeFilters"),
//   	},
//   	StackIdentifier: jsii.String("stackIdentifier"),
//   	TagFilters: []interface{}{
//   		&TagFilterProperty{
//   			Key: jsii.String("key"),
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resourcegroups-group-query.html
//
type CfnGroup_QueryProperty struct {
	// Specifies limits to the types of resources that can be included in the resource group.
	//
	// For example, if `ResourceTypeFilters` is `["AWS::EC2::Instance", "AWS::DynamoDB::Table"]` , only EC2 instances or DynamoDB tables can be members of this resource group. The default value is `["AWS::AllSupported"]` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resourcegroups-group-query.html#cfn-resourcegroups-group-query-resourcetypefilters
	//
	ResourceTypeFilters *[]*string `field:"optional" json:"resourceTypeFilters" yaml:"resourceTypeFilters"`
	// Specifies the ARN of a CloudFormation stack.
	//
	// All supported resources of the CloudFormation stack are members of the resource group. If you don't specify an ARN, this parameter defaults to the current stack that you are defining, which means that all the resources of the current stack are grouped.
	//
	// You can specify a value for `StackIdentifier` only when the `ResourceQuery.Type` property is `CLOUDFORMATION_STACK_1_0.`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resourcegroups-group-query.html#cfn-resourcegroups-group-query-stackidentifier
	//
	StackIdentifier *string `field:"optional" json:"stackIdentifier" yaml:"stackIdentifier"`
	// A list of key-value pair objects that limit which resources can be members of the resource group.
	//
	// This property is required when the `ResourceQuery.Type` property is `TAG_FILTERS_1_0` .
	//
	// A resource must have a tag that matches every filter that is provided in the `TagFilters` list.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resourcegroups-group-query.html#cfn-resourcegroups-group-query-tagfilters
	//
	TagFilters interface{} `field:"optional" json:"tagFilters" yaml:"tagFilters"`
}

