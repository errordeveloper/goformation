package cloudformation

// AWSServerlessFunction_ApiEvent AWS CloudFormation Resource (AWS::Serverless::Function.ApiEvent)
// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#api
type AWSServerlessFunction_ApiEvent struct {

	// Method AWS CloudFormation Property
	// Required: true
	// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#api
	Method *StringIntrinsic `json:"Method,omitempty"`

	// Path AWS CloudFormation Property
	// Required: true
	// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#api
	Path *StringIntrinsic `json:"Path,omitempty"`

	// RestApiId AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#api
	RestApiId *StringIntrinsic `json:"RestApiId,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSServerlessFunction_ApiEvent) AWSCloudFormationType() string {
	return "AWS::Serverless::Function.ApiEvent"
}
