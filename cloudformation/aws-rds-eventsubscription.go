package cloudformation

import (
	"encoding/json"
	"errors"
	"fmt"
)

// AWSRDSEventSubscription AWS CloudFormation Resource (AWS::RDS::EventSubscription)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-eventsubscription.html
type AWSRDSEventSubscription struct {

	// Enabled AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-eventsubscription.html#cfn-rds-eventsubscription-enabled
	Enabled *Value `json:"Enabled,omitempty"`

	// EventCategories AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-eventsubscription.html#cfn-rds-eventsubscription-eventcategories
	EventCategories []*Value `json:"EventCategories,omitempty"`

	// SnsTopicArn AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-eventsubscription.html#cfn-rds-eventsubscription-snstopicarn
	SnsTopicArn *Value `json:"SnsTopicArn,omitempty"`

	// SourceIds AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-eventsubscription.html#cfn-rds-eventsubscription-sourceids
	SourceIds []*Value `json:"SourceIds,omitempty"`

	// SourceType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-eventsubscription.html#cfn-rds-eventsubscription-sourcetype
	SourceType *Value `json:"SourceType,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSRDSEventSubscription) AWSCloudFormationType() string {
	return "AWS::RDS::EventSubscription"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r *AWSRDSEventSubscription) MarshalJSON() ([]byte, error) {
	type Properties AWSRDSEventSubscription
	return json.Marshal(&struct {
		Type       string
		Properties Properties
	}{
		Type:       r.AWSCloudFormationType(),
		Properties: (Properties)(*r),
	})
}

// UnmarshalJSON is a custom JSON unmarshalling hook that strips the outer
// AWS CloudFormation resource object, and just keeps the 'Properties' field.
func (r *AWSRDSEventSubscription) UnmarshalJSON(b []byte) error {
	type Properties AWSRDSEventSubscription
	res := &struct {
		Type       string
		Properties *Properties
	}{}
	if err := json.Unmarshal(b, &res); err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return err
	}

	// If the resource has no Properties set, it could be nil
	if res.Properties != nil {
		*r = AWSRDSEventSubscription(*res.Properties)
	}

	return nil
}

// GetAllAWSRDSEventSubscriptionResources retrieves all AWSRDSEventSubscription items from an AWS CloudFormation template
func (t *Template) GetAllAWSRDSEventSubscriptionResources() map[string]AWSRDSEventSubscription {
	results := map[string]AWSRDSEventSubscription{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case AWSRDSEventSubscription:
			// We found a strongly typed resource of the correct type; use it
			results[name] = resource
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::RDS::EventSubscription" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						result := &AWSRDSEventSubscription{}
						if err := result.UnmarshalJSON(b); err == nil {
							results[name] = *result
						}
					}
				}
			}
		}
	}
	return results
}

// GetAWSRDSEventSubscriptionWithName retrieves all AWSRDSEventSubscription items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSRDSEventSubscriptionWithName(name string) (AWSRDSEventSubscription, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case AWSRDSEventSubscription:
			// We found a strongly typed resource of the correct type; use it
			return resource, nil
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::RDS::EventSubscription" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						result := &AWSRDSEventSubscription{}
						if err := result.UnmarshalJSON(b); err == nil {
							return *result, nil
						}
					}
				}
			}
		}
	}
	return AWSRDSEventSubscription{}, errors.New("resource not found")
}
