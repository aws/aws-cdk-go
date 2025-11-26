package previewawsredshiftserverlessmixins


// A array of parameters to set for more control over a serverless database.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   configParameterProperty := &ConfigParameterProperty{
//   	ParameterKey: jsii.String("parameterKey"),
//   	ParameterValue: jsii.String("parameterValue"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-workgroup-configparameter.html
//
type CfnWorkgroupPropsMixin_ConfigParameterProperty struct {
	// The key of the parameter.
	//
	// The options are `auto_mv` , `datestyle` , `enable_case_sensitive_identifier` , `enable_user_activity_logging` , `query_group` , `search_path` , `require_ssl` , `use_fips_ssl` , and query monitoring metrics that let you define performance boundaries. For more information about query monitoring rules and available metrics, see [Query monitoring metrics for Amazon Redshift Serverless](https://docs.aws.amazon.com/redshift/latest/dg/cm-c-wlm-query-monitoring-rules.html#cm-c-wlm-query-monitoring-metrics-serverless) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-workgroup-configparameter.html#cfn-redshiftserverless-workgroup-configparameter-parameterkey
	//
	ParameterKey *string `field:"optional" json:"parameterKey" yaml:"parameterKey"`
	// The value of the parameter to set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-workgroup-configparameter.html#cfn-redshiftserverless-workgroup-configparameter-parametervalue
	//
	ParameterValue *string `field:"optional" json:"parameterValue" yaml:"parameterValue"`
}

