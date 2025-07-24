package ui

import (
	"strings"
)

// setupStep configures the form fields for the current step
func (m *Model) setupStep() {
	m.currentField = 0
	m.error = ""
	m.editingList = false

	switch m.currentStep {
	case StepPersonalInfo:
		m.setupPersonalInfoStep()
	case StepSummary:
		m.setupSummaryStep()
	case StepEducation:
		m.setupEducationStep()
	case StepExperience:
		m.setupExperienceStep()
	case StepProjects:
		m.setupProjectsStep()
	case StepSkills:
		m.setupSkillsStep()
	case StepCustomSections:
		m.setupCustomSectionsStep()
	}
}

// setupPersonalInfoStep sets up the personal information form fields
func (m *Model) setupPersonalInfoStep() {
	m.fields = []FormField{
		{Label: "姓名", Required: true, Placeholder: "请输入您的姓名"},
		{Label: "邮箱", Required: true, Placeholder: "your.email@example.com"},
		{Label: "电话", Required: true, Placeholder: "(123) 456-7890"},
		{Label: "地址", Required: true, Placeholder: "城市, 省份, 国家"},
		{Label: "网站", Required: false, Placeholder: "www.yourwebsite.com (可选)"},
	}
	// Load existing data
	m.fields[0].Value = m.resume.PersonalInfo.Name
	m.fields[1].Value = m.resume.PersonalInfo.Email
	m.fields[2].Value = m.resume.PersonalInfo.Phone
	m.fields[3].Value = m.resume.PersonalInfo.Location
	m.fields[4].Value = m.resume.PersonalInfo.Website
}

// setupSummaryStep sets up the summary form fields
func (m *Model) setupSummaryStep() {
	m.fields = []FormField{
		{Label: "个人简介", Required: true, Placeholder: "请用1-3句话描述您的资历和经验...", Multiline: true},
	}
	m.fields[0].Value = m.resume.Summary
}

// setupEducationStep sets up the education form fields
func (m *Model) setupEducationStep() {
	m.fields = []FormField{
		{Label: "学校名称", Required: true, Placeholder: "请输入学校名称"},
		{Label: "学位", Required: true, Placeholder: "学士/硕士/博士"},
		{Label: "专业", Required: false, Placeholder: "请输入专业 (可选)"},
		{Label: "地点", Required: true, Placeholder: "城市, 省份"},
		{Label: "开始年份", Required: true, Placeholder: "2020"},
		{Label: "结束年份", Required: true, Placeholder: "2024 或 current"},
	}
}

// setupExperienceStep sets up the work experience form fields
func (m *Model) setupExperienceStep() {
	m.fields = []FormField{
		{Label: "公司名称", Required: true, Placeholder: "请输入公司名称"},
		{Label: "职位", Required: true, Placeholder: "请输入职位"},
		{Label: "地点", Required: true, Placeholder: "城市, 省份"},
		{Label: "开始年月", Required: true, Placeholder: "2022-06"},
		{Label: "结束年月", Required: true, Placeholder: "2024-08 或 current"},
		{Label: "工作描述", Required: true, Placeholder: "请描述您的主要职责和成就 (一行一个要点)", Multiline: true},
	}
}

// setupProjectsStep sets up the projects form fields
func (m *Model) setupProjectsStep() {
	m.fields = []FormField{
		{Label: "项目名称", Required: true, Placeholder: "请输入项目名称"},
		{Label: "项目描述", Required: true, Placeholder: "一句话描述项目"},
		{Label: "地点", Required: false, Placeholder: "城市, 省份 (可选)"},
		{Label: "开始年月", Required: true, Placeholder: "2023-01"},
		{Label: "结束年月", Required: true, Placeholder: "2023-06 或 current"},
		{Label: "项目详情", Required: true, Placeholder: "详细描述项目内容和您的贡献 (一行一个要点)", Multiline: true},
	}
}

// setupSkillsStep sets up the skills form fields with default values
func (m *Model) setupSkillsStep() {
	// Default skill categories if none exist
	defaultLanguages := []string{"JavaScript", "Python", "Go", "Java", "TypeScript"}
	defaultFrameworks := []string{"React", "Node.js", "Express", "Vue.js", "Django"}
	defaultDatabases := []string{"PostgreSQL", "MongoDB", "Redis", "MySQL"}
	defaultOther := []string{"Git", "Docker", "Linux", "AWS", "Jenkins"}

	// Use existing skills or defaults
	languages := m.resume.Skills.Languages
	if len(languages) == 0 {
		languages = defaultLanguages
	}

	frameworks := m.resume.Skills.Frameworks
	if len(frameworks) == 0 {
		frameworks = defaultFrameworks
	}

	databases := m.resume.Skills.Databases
	if len(databases) == 0 {
		databases = defaultDatabases
	}

	other := m.resume.Skills.Other
	if len(other) == 0 {
		other = defaultOther
	}

	m.fields = []FormField{
		{Label: "编程语言", Required: false, Placeholder: "Enter编辑列表", IsList: true, Value: strings.Join(languages, ", ")},
		{Label: "框架/库", Required: false, Placeholder: "Enter编辑列表", IsList: true, Value: strings.Join(frameworks, ", ")},
		{Label: "数据库", Required: false, Placeholder: "Enter编辑列表", IsList: true, Value: strings.Join(databases, ", ")},
		{Label: "其他工具", Required: false, Placeholder: "Enter编辑列表", IsList: true, Value: strings.Join(other, ", ")},
	}
}

// setupCustomSectionsStep sets up the custom sections form fields
func (m *Model) setupCustomSectionsStep() {
	m.fields = []FormField{
		{Label: "自定义章节标题", Required: false, Placeholder: "例如: 证书、奖项、志愿经历等 (可留空跳过)"},
		{Label: "章节内容", Required: false, Placeholder: "Enter编辑列表", IsList: true},
	}
}

// nextStep advances to the next step in the creation flow
func (m *Model) nextStep() {
	m.currentStep++
	if m.currentStep <= StepCustomSections {
		m.setupStep()
	} else if m.currentStep == StepConfirm {
		// Setup confirmation step
		m.fields = nil
		m.currentField = 0
	}
}
