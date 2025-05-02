package awscdkgluealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsglue"
)

// Represents a trigger action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import glue_alpha "github.com/aws/aws-cdk-go/awscdkgluealpha"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var cfnCrawler cfnCrawler
//   var job job
//   var securityConfiguration securityConfiguration
//
//   action := &Action{
//   	Arguments: map[string]*string{
//   		"argumentsKey": jsii.String("arguments"),
//   	},
//   	Crawler: cfnCrawler,
//   	Job: job,
//   	SecurityConfiguration: securityConfiguration,
//   	Timeout: cdk.Duration_Minutes(jsii.Number(30)),
//   }
//
// Experimental.
type Action struct {
	// The job arguments used when this trigger fires.
	// Default: - no arguments are passed to the job.
	//
	// Experimental.
	Arguments *map[string]*string `field:"optional" json:"arguments" yaml:"arguments"`
	// The name of the crawler to be used with this action.
	// Default: - no crawler is used.
	//
	// Experimental.
	Crawler awsglue.CfnCrawler `field:"optional" json:"crawler" yaml:"crawler"`
	// The job to be executed.
	// Default: - no job is executed.
	//
	// Experimental.
	Job IJob `field:"optional" json:"job" yaml:"job"`
	// The `SecurityConfiguration` to be used with this action.
	// Default: - no security configuration is used.
	//
	// Experimental.
	SecurityConfiguration ISecurityConfiguration `field:"optional" json:"securityConfiguration" yaml:"securityConfiguration"`
	// The job run timeout.
	//
	// This is the maximum time that a job run can consume resources before it is terminated and enters TIMEOUT status.
	// Default: - the default timeout value set in the job definition.
	//
	// Experimental.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
}

