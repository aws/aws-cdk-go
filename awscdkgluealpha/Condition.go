package awscdkgluealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkgluealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsglue"
)

// A condition that determines when a conditional trigger fires.
//
// A condition watches exactly one target in exactly one state: use
// {@link Condition.job} to watch a job or {@link Condition.crawler} to watch a
// crawler. Because the state is a required argument of each factory, a condition
// can never reference a target without its state, or both a job and a crawler.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkgluealpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var crawlerRef ICrawlerRef
//
//   condition := glue_alpha.Condition_Crawler(crawlerRef, glue_alpha.CrawlerState_RUNNING, &ConditionOptions{
//   	LogicalOperator: glue_alpha.ConditionLogicalOperator_EQUALS,
//   })
//
// Experimental.
type Condition interface {
}

// The jsii proxy struct for Condition
type jsiiProxy_Condition struct {
	_ byte // padding
}

// Experimental.
func NewCondition_Override(c Condition) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-glue-alpha.Condition",
		nil, // no parameters
		c,
	)
}

// Create a condition on the state of a crawler.
// Experimental.
func Condition_Crawler(crawler interfacesawsglue.ICrawlerRef, crawlState CrawlerState, options *ConditionOptions) Condition {
	_init_.Initialize()

	if err := validateCondition_CrawlerParameters(crawler, crawlState, options); err != nil {
		panic(err)
	}
	var returns Condition

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.Condition",
		"crawler",
		[]interface{}{crawler, crawlState, options},
		&returns,
	)

	return returns
}

// Create a condition on the state of a job.
// Experimental.
func Condition_Job(job interfacesawsglue.IJobRef, state JobState, options *ConditionOptions) Condition {
	_init_.Initialize()

	if err := validateCondition_JobParameters(job, state, options); err != nil {
		panic(err)
	}
	var returns Condition

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.Condition",
		"job",
		[]interface{}{job, state, options},
		&returns,
	)

	return returns
}

