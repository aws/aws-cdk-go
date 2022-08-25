package awsecs


// The interface for a service using the Fargate launch type on an ECS cluster.
// Experimental.
type IFargateService interface {
	IService
}

// The jsii proxy for IFargateService
type jsiiProxy_IFargateService struct {
	jsiiProxy_IService
}

