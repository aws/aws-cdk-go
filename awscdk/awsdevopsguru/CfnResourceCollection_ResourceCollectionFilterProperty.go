package awsdevopsguru


// Information about a filter used to specify which AWS resources are analyzed for anomalous behavior by DevOps Guru.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourceCollectionFilterProperty := &ResourceCollectionFilterProperty{
//   	CloudFormation: &CloudFormationCollectionFilterProperty{
//   		StackNames: []*string{
//   			jsii.String("stackNames"),
//   		},
//   	},
//   	Tags: []tagCollectionProperty{
//   		&tagCollectionProperty{
//   			AppBoundaryKey: jsii.String("appBoundaryKey"),
//   			TagValues: []*string{
//   				jsii.String("tagValues"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsguru-resourcecollection-resourcecollectionfilter.html
//
type CfnResourceCollection_ResourceCollectionFilterProperty struct {
	// Information about AWS CloudFormation stacks.
	//
	// You can use up to 1000 stacks to specify which AWS resources in your account to analyze. For more information, see [Stacks](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacks.html) in the *AWS CloudFormation User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsguru-resourcecollection-resourcecollectionfilter.html#cfn-devopsguru-resourcecollection-resourcecollectionfilter-cloudformation
	//
	CloudFormation interface{} `field:"optional" json:"cloudFormation" yaml:"cloudFormation"`
	// The AWS tags used to filter the resources in the resource collection.
	//
	// Tags help you identify and organize your AWS resources. Many AWS services support tagging, so you can assign the same tag to resources from different services to indicate that the resources are related. For example, you can assign the same tag to an Amazon DynamoDB table resource that you assign to an AWS Lambda function. For more information about using tags, see the [Tagging best practices](https://docs.aws.amazon.com/whitepapers/latest/tagging-best-practices/tagging-best-practices.html) whitepaper.
	//
	// Each AWS tag has two parts.
	//
	// - A tag *key* (for example, `CostCenter` , `Environment` , `Project` , or `Secret` ). Tag *keys* are case-sensitive.
	// - A field known as a tag *value* (for example, `111122223333` , `Production` , or a team name). Omitting the tag *value* is the same as using an empty string. Like tag *keys* , tag *values* are case-sensitive. The tag value is a required property when AppBoundaryKey is specified.
	//
	// Together these are known as *key* - *value* pairs.
	//
	// > The string used for a *key* in a tag that you use to define your resource coverage must begin with the prefix `Devops-guru-` . The tag *key* might be `DevOps-Guru-deployment-application` or `devops-guru-rds-application` . When you create a *key* , the case of characters in the *key* can be whatever you choose. After you create a *key* , it is case-sensitive. For example, DevOps Guru works with a *key* named `devops-guru-rds` and a *key* named `DevOps-Guru-RDS` , and these act as two different *keys* . Possible *key* / *value* pairs in your application might be `Devops-Guru-production-application/RDS` or `Devops-Guru-production-application/containers` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsguru-resourcecollection-resourcecollectionfilter.html#cfn-devopsguru-resourcecollection-resourcecollectionfilter-tags
	//
	Tags *[]*CfnResourceCollection_TagCollectionProperty `field:"optional" json:"tags" yaml:"tags"`
}

