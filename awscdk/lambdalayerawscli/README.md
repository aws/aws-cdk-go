# AWS Lambda Layer with AWS CLI

This module exports a single class called `AwsCliLayer` which is a `lambda.Layer` that bundles the AWS CLI.

Any Lambda Function that uses this layer must use a Python 3.x runtime.

Usage:

```go
// AwsCliLayer bundles the AWS CLI in a lambda layer
import "github.com/aws/aws-cdk-go/awscdk"

var fn function

fn.addLayers(awscdk.NewAwsCliLayer(this, jsii.String("AwsCliLayer")))
```

The CLI will be installed under `/opt/awscli/aws`.
