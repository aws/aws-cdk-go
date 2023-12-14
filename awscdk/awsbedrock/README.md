# Amazon Bedrock Construct Library

Amazon Bedrock is a fully managed service that offers a choice of foundation models (FMs)
along with a broad set of capabilities for building generative AI applications.

CloudFormation does not currently support any Bedrock resource types.
This construct library is a collection of constructs that can look up existing Bedrock models
for use with other services' CDK constructs, such as AWS Step Functions.

To look up a Bedrock base foundation model:

```go
import "github.com/aws/aws-cdk-go/awscdk"


bedrock.FoundationModel_FromFoundationModelId(this, jsii.String("Model"), bedrock.FoundationModelIdentifier_ANTHROPIC_CLAUDE_V2())
```

To look up a Bedrock provisioned throughput model:

```go
import bedrock "github.com/aws/aws-cdk-go/awscdk"


bedrock.ProvisionedModel_FromProvisionedModelArn(this, jsii.String("Model"), jsii.String("arn:aws:bedrock:us-east-2:123456789012:provisioned-model/abc-123"))
```
