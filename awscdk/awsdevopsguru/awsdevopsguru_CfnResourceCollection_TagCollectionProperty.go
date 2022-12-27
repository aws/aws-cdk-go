package awsdevopsguru


// A collection of AWS stags.
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
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tagCollectionProperty := &tagCollectionProperty{
//   	appBoundaryKey: jsii.String("appBoundaryKey"),
//   	tagValues: []*string{
//   		jsii.String("tagValues"),
//   	},
//   }
//
type CfnResourceCollection_TagCollectionProperty struct {
	// An AWS tag *key* that is used to identify the AWS resources that DevOps Guru analyzes.
	//
	// All AWS resources in your account and Region tagged with this *key* make up your DevOps Guru application and analysis boundary.
	//
	// > The string used for a *key* in a tag that you use to define your resource coverage must begin with the prefix `Devops-guru-` . The tag *key* might be `Devops-guru-deployment-application` or `Devops-guru-rds-application` . While *keys* are case-sensitive, the case of *key* characters don't matter to DevOps Guru. For example, DevOps Guru works with a *key* named `devops-guru-rds` and a *key* named `DevOps-Guru-RDS` . Possible *key* / *value* pairs in your application might be `Devops-Guru-production-application/RDS` or `Devops-Guru-production-application/containers` .
	AppBoundaryKey *string `field:"optional" json:"appBoundaryKey" yaml:"appBoundaryKey"`
	// The values in an AWS tag collection.
	//
	// The tag's *value* is an optional field used to associate a string with the tag *key* (for example, `111122223333` , `Production` , or a team name). The *key* and *value* are the tag's *key* pair. Omitting the tag *value* is the same as using an empty string. Like tag *keys* , tag *values* are case-sensitive. You can specify a maximum of 256 characters for a tag value.
	TagValues *[]*string `field:"optional" json:"tagValues" yaml:"tagValues"`
}

