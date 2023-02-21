# AWS Lambda Layer with the NPM dependency proxy-agent

This module exports a single class called `NodeProxyAgentLayer` which is a `lambda.Layer` that bundles the NPM dependency [`proxy-agent`](https://www.npmjs.com/package/proxy-agent).

> * proxy-agent Version: 5.0.0

Usage:

```go
import "github.com/aws/aws-cdk-go/awscdk"
import lambda "github.com/aws/aws-cdk-go/awscdk"

var fn function

fn.addLayers(awscdk.NewNodeProxyAgentLayer(this, jsii.String("NodeProxyAgentLayer")))
```

[`proxy-agent`](https://www.npmjs.com/package/proxy-agent) will be installed under `/nodejs/node_modules`.
