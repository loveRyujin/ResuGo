package ui

import (
	"strings"
)

// setupStep configures the form fields for the current step
func (m *Model) setupStep() {
	m.currentField = 0
	m.error = ""
	m.editingList = false

	// Reset management states
	m.managingExperiences = false
	m.managingProjects = false
	m.editingExperience = -1
	m.editingProject = -1
	m.selectedExperience = 0
	m.selectedProject = 0

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

	// Create input components for this step (if fields are set)
	if len(m.fields) > 0 {
		m.createTextInputs()

		// Auto-focus the first non-list field for direct editing
		if !m.fields[0].IsList {
			m.currentField = 0
			m.focusCurrentField()
		}
	}
}

// setupPersonalInfoStep sets up the personal information form fields
func (m *Model) setupPersonalInfoStep() {
	m.fields = []FormField{
		{Label: "姓名", Required: true, Placeholder: "如: 张三"},
		{Label: "邮箱", Required: true, Placeholder: "如: zhangsan@example.com"},
		{Label: "电话", Required: true, Placeholder: "如: 138-0013-8000"},
		{Label: "地址", Required: true, Placeholder: "如: 北京市海淀区"},
		{Label: "网站", Required: false, Placeholder: "如: www.github.com/username (可选)"},
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
		{Label: "个人简介", Required: true, Placeholder: "如: 具有3年软件开发经验的全栈工程师，熟悉React、Node.js等技术栈...", Multiline: true},
	}
	m.fields[0].Value = m.resume.Summary
}

// setupEducationStep sets up the education form fields
func (m *Model) setupEducationStep() {
	m.fields = []FormField{
		{Label: "学校名称", Required: true, Placeholder: "如: 清华大学"},
		{Label: "学位", Required: true, Placeholder: "如: 计算机科学学士、软件工程硕士"},
		{Label: "专业", Required: false, Placeholder: "如: 计算机科学与技术 (可选)"},
		{Label: "地点", Required: true, Placeholder: "如: 北京"},
		{Label: "开始年份", Required: true, Placeholder: "如: 2020"},
		{Label: "结束年份", Required: true, Placeholder: "如: 2024 或 current"},
	}

	// Load existing education data if available
	if len(m.resume.Education) > 0 {
		edu := m.resume.Education[0] // Edit the first education entry
		m.fields[0].Value = edu.Institution
		m.fields[1].Value = edu.Degree
		m.fields[2].Value = edu.Major
		m.fields[3].Value = edu.Location
		m.fields[4].Value = edu.FormatStartDate()
		if edu.Current {
			m.fields[5].Value = "current"
		} else {
			m.fields[5].Value = edu.FormatEndDate()
		}
	}
}

// setupExperienceStep sets up the work experience management or form fields
func (m *Model) setupExperienceStep() {
	m.managingExperiences = true
	m.editingExperience = -1
	m.selectedExperience = 0
	m.fields = nil // Will be set when entering edit mode
}

// enterExperienceEditMode enters edit mode for a specific experience (index -1 for new)
func (m *Model) enterExperienceEditMode(index int) {
	m.managingExperiences = false
	m.editingExperience = index

	m.fields = []FormField{
		{Label: "公司名称", Required: true, Placeholder: "如: 阿里巴巴集团"},
		{Label: "职位", Required: true, Placeholder: "如: 高级软件工程师"},
		{Label: "地点", Required: true, Placeholder: "如: 杭州"},
		{Label: "开始年月", Required: true, Placeholder: "如: 2022-06"},
		{Label: "结束年月", Required: true, Placeholder: "如: 2024-08 或 current"},
		{Label: "工作描述", Required: true, Placeholder: "如: 负责电商平台后端开发\n优化系统性能，提升30%处理速度\n参与微服务架构设计", Multiline: true},
	}

	// Load existing experience data if editing
	if index >= 0 && index < len(m.resume.Experience) {
		exp := m.resume.Experience[index]
		m.fields[0].Value = exp.Company
		m.fields[1].Value = exp.Position
		m.fields[2].Value = exp.Location
		m.fields[3].Value = exp.FormatStartDate()
		if exp.Current {
			m.fields[4].Value = "current"
		} else {
			m.fields[4].Value = exp.FormatEndDate()
		}
		if len(exp.Responsibilities) > 0 {
			m.fields[5].Value = strings.Join(exp.Responsibilities, "\n")
		}
	}

	// Create input components and focus first field
	m.createTextInputs()
	m.currentField = 0
	m.focusCurrentField()
}

// setupProjectsStep sets up the projects management or form fields
func (m *Model) setupProjectsStep() {
	m.managingProjects = true
	m.editingProject = -1
	m.selectedProject = 0
	m.fields = nil // Will be set when entering edit mode
}

// enterProjectEditMode enters edit mode for a specific project (index -1 for new)
func (m *Model) enterProjectEditMode(index int) {
	m.managingProjects = false
	m.editingProject = index

	m.fields = []FormField{
		{Label: "项目名称", Required: true, Placeholder: "如: 在线教育平台"},
		{Label: "项目描述", Required: true, Placeholder: "如: 基于React和Node.js的在线学习系统"},
		{Label: "地点", Required: false, Placeholder: "如: 北京 (可选)"},
		{Label: "开始年月", Required: true, Placeholder: "如: 2023-01"},
		{Label: "结束年月", Required: true, Placeholder: "如: 2023-06 或 current"},
		{Label: "项目详情", Required: true, Placeholder: "如: 负责前端页面开发和API设计\n实现用户认证和课程管理功能\n使用Redis缓存提升系统性能", Multiline: true},
	}

	// Load existing project data if editing
	if index >= 0 && index < len(m.resume.Projects) {
		proj := m.resume.Projects[index]
		m.fields[0].Value = proj.Name
		m.fields[1].Value = proj.Description
		m.fields[2].Value = proj.Location
		m.fields[3].Value = proj.FormatStartDate()
		if proj.Current {
			m.fields[4].Value = "current"
		} else {
			m.fields[4].Value = proj.FormatEndDate()
		}
		if len(proj.Details) > 0 {
			m.fields[5].Value = strings.Join(proj.Details, "\n")
		}
	}

	// Create input components and focus first field
	m.createTextInputs()
	m.currentField = 0
	m.focusCurrentField()
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
		{Label: "自定义章节标题", Required: false, Placeholder: "如: 获得证书、获奖经历、志愿活动 (可留空跳过)"},
		{Label: "章节内容", Required: false, Placeholder: "按Enter编辑列表", IsList: true},
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
