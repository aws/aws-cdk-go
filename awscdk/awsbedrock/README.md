# Amazon Bedrock Construct Library

Amazon Bedrock is a fully managed service that offers a choice of foundation models (FMs)
along with a broad set of capabilities for building generative AI applications.

This construct library provides a collection of constructs that can look up existing Bedrock models
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

There are no official hand-written ([L2](https://docs.aws.amazon.com/cdk/latest/guide/constructs.html#constructs_lib)) constructs for provisioning Bedrock resources yet. Here are some suggestions on how to proceed:

* Search [Construct Hub for Bedrock construct libraries](https://constructs.dev/search?q=bedrock)
* Use the automatically generated [L1](https://docs.aws.amazon.com/cdk/latest/guide/constructs.html#constructs_l1_using) constructs, in the same way you would use [the CloudFormation AWS::Bedrock resources](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/AWS_Bedrock.html) directly.
