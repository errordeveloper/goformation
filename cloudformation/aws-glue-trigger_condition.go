package cloudformation

// AWSGlueTrigger_Condition AWS CloudFormation Resource (AWS::Glue::Trigger.Condition)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-trigger-condition.html
type AWSGlueTrigger_Condition struct {

	// JobName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-trigger-condition.html#cfn-glue-trigger-condition-jobname
	JobName *stringIntrinsic `json:"JobName,omitempty"`

	// LogicalOperator AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-trigger-condition.html#cfn-glue-trigger-condition-logicaloperator
	LogicalOperator *stringIntrinsic `json:"LogicalOperator,omitempty"`

	// State AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-trigger-condition.html#cfn-glue-trigger-condition-state
	State *stringIntrinsic `json:"State,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSGlueTrigger_Condition) AWSCloudFormationType() string {
	return "AWS::Glue::Trigger.Condition"
}
