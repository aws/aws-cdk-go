package awslambda


// Specifies the tenant isolation mode configuration for a Lambda function.
//
// This allows you to configure specific tenant isolation strategies for your function invocations. Tenant isolation configuration cannot be modified after function creation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tenancyConfigProperty := &TenancyConfigProperty{
//   	TenantIsolationMode: jsii.String("tenantIsolationMode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-tenancyconfig.html
//
type CfnFunction_TenancyConfigProperty struct {
	// Tenant isolation mode allows for invocation to be sent to a corresponding execution environment dedicated to a specific tenant ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-tenancyconfig.html#cfn-lambda-function-tenancyconfig-tenantisolationmode
	//
	TenantIsolationMode *string `field:"required" json:"tenantIsolationMode" yaml:"tenantIsolationMode"`
}

