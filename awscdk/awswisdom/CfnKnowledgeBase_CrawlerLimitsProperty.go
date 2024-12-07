package awswisdom


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   crawlerLimitsProperty := &CrawlerLimitsProperty{
//   	RateLimit: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-knowledgebase-crawlerlimits.html
//
type CfnKnowledgeBase_CrawlerLimitsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-knowledgebase-crawlerlimits.html#cfn-wisdom-knowledgebase-crawlerlimits-ratelimit
	//
	RateLimit *float64 `field:"optional" json:"rateLimit" yaml:"rateLimit"`
}

