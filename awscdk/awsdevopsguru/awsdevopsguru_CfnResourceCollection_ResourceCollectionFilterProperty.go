package awsdevopsguru


// Information about a filter used to specify which AWS resources are analyzed for anomalous behavior by DevOps Guru.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourceCollectionFilterProperty := &resourceCollectionFilterProperty{
//   	cloudFormation: &cloudFormationCollectionFilterProperty{
//   		stackNames: []*string{
//   			jsii.String("stackNames"),
//   		},
//   	},
//   	tags: []tagCollectionProperty{
//   		&tagCollectionProperty{
//   			appBoundaryKey: jsii.String("appBoundaryKey"),
//   			tagValues: []*string{
//   				jsii.String("tagValues"),
//   			},
//   		},
//   	},
//   }
//
type CfnResourceCollection_ResourceCollectionFilterProperty struct {
	// Information about AWS CloudFormation stacks.
	//
	// You can use up to 500 stacks to specify which AWS resources in your account to analyze. For more information, see [Stacks](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacks.html) in the *AWS CloudFormation User Guide* .
	CloudFormation interface{} `field:"optional" json:"cloudFormation" yaml:"cloudFormation"`
	// The AWS tags used to filter the resources in the resource collection.
	//
	// Tags help you identify and organize your AWS resources. Many AWS services support tagging, so you can assign the same tag to resources from different services to indicate that the resources are related. For example, you can assign the same tag to an Amazon DynamoDB table resource that you assign to an AWS Lambda function. For more information about using tags, see the [Tagging best practices](https://docs.aws.amazon.com/https://d1.awsstatic.com/whitepapers/aws-tagging-best-practices.pdf) whitepaper.
	//
	// Each AWS tag has two parts.
	//
	// - A tag *key* (for example, `CostCenter` , `Environment` , `Project` , or `Secret` ). Tag *keys* are case-sensitive.
	// - An optional field known as a tag *value* (for example, `111122223333` , `Production` , or a team name). Omitting the tag *value* is the same as using an empty string. Like tag *keys* , tag *values* are case-sensitive.
	//
	// Together these are known as *key* - *value* pairs.
	//
	// > The string used for a *key* in a tag that you use to define your resource coverage must begin with the prefix `Devops-guru-` . The tag *key* might be `Devops-guru-deployment-application` or `Devops-guru-rds-application` . While *keys* are case-sensitive, the case of *key* characters don't matter to DevOps Guru. For example, DevOps Guru works with a *key* named `devops-guru-rds` and a *key* named `DevOps-Guru-RDS` . Possible *key* / *value* pairs in your application might be `Devops-Guru-production-application/RDS` or `Devops-Guru-production-application/containers` .
	Tags *[]*CfnResourceCollection_TagCollectionProperty `field:"optional" json:"tags" yaml:"tags"`
}

