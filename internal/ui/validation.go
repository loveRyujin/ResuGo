package ui

import (
	"fmt"
	"strconv"
	"strings"
)

// validateCurrentStep validates the current step's form data
func (m *Model) validateCurrentStep() bool {
	m.error = ""

	for i, field := range m.fields {
		if field.Required && strings.TrimSpace(field.Value) == "" {
			m.error = fmt.Sprintf("请填写必填项: %s", field.Label)
			m.currentField = i
			return false
		}
	}

	// Special validation for date fields
	if m.currentStep == StepEducation || m.currentStep == StepExperience || m.currentStep == StepProjects {
		// Validate year formats
		for i, field := range m.fields {
			if strings.Contains(field.Label, "年份") || strings.Contains(field.Label, "年月") {
				if field.Value != "current" && field.Value != "" {
					if m.currentStep == StepEducation {
						if _, err := strconv.Atoi(field.Value); err != nil {
							m.error = fmt.Sprintf("年份格式错误: %s (请输入如: 2020)", field.Label)
							m.currentField = i
							return false
						}
					} else {
						// For experience and projects, expect YYYY-MM format
						if field.Value != "current" && !strings.Contains(field.Value, "-") {
							m.error = fmt.Sprintf("日期格式错误: %s (请输入如: 2022-06)", field.Label)
							m.currentField = i
							return false
						}
					}
				}
			}
		}
	}

	return true
}
