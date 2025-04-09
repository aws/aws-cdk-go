package awsdms


// Provides information that defines an OpenSearch endpoint.
//
// This information includes the output format of records applied to the endpoint and details of transaction and control table data information. For more information about the available settings, see [Extra connection attributes when using OpenSearch as a target for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Elasticsearch.html#CHAP_Target.Elasticsearch.Configuration) in the *AWS Database Migration Service User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   elasticsearchSettingsProperty := &ElasticsearchSettingsProperty{
//   	EndpointUri: jsii.String("endpointUri"),
//   	ErrorRetryDuration: jsii.Number(123),
//   	FullLoadErrorPercentage: jsii.Number(123),
//   	ServiceAccessRoleArn: jsii.String("serviceAccessRoleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-elasticsearchsettings.html
//
type CfnEndpoint_ElasticsearchSettingsProperty struct {
	// The endpoint for the OpenSearch cluster.
	//
	// AWS DMS uses HTTPS if a transport protocol (either HTTP or HTTPS) isn't specified.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-elasticsearchsettings.html#cfn-dms-endpoint-elasticsearchsettings-endpointuri
	//
	EndpointUri *string `field:"optional" json:"endpointUri" yaml:"endpointUri"`
	// The maximum number of seconds for which DMS retries failed API requests to the OpenSearch cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-elasticsearchsettings.html#cfn-dms-endpoint-elasticsearchsettings-errorretryduration
	//
	ErrorRetryDuration *float64 `field:"optional" json:"errorRetryDuration" yaml:"errorRetryDuration"`
	// The maximum percentage of records that can fail to be written before a full load operation stops.
	//
	// To avoid early failure, this counter is only effective after 1,000 records are transferred. OpenSearch also has the concept of error monitoring during the last 10 minutes of an Observation Window. If transfer of all records fail in the last 10 minutes, the full load operation stops.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-elasticsearchsettings.html#cfn-dms-endpoint-elasticsearchsettings-fullloaderrorpercentage
	//
	FullLoadErrorPercentage *float64 `field:"optional" json:"fullLoadErrorPercentage" yaml:"fullLoadErrorPercentage"`
	// The Amazon Resource Name (ARN) used by the service to access the IAM role.
	//
	// The role must allow the `iam:PassRole` action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-elasticsearchsettings.html#cfn-dms-endpoint-elasticsearchsettings-serviceaccessrolearn
	//
	ServiceAccessRoleArn *string `field:"optional" json:"serviceAccessRoleArn" yaml:"serviceAccessRoleArn"`
}

