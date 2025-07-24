package ui

import (
	"fmt"
	"strings"
)

// View renders the current view based on the model state
func (m Model) View() string {
	if m.quitting {
		return "再见! 👋\n"
	}

	if m.finished {
		return "🎉 简历创建成功!\n\n" +
			"已保存文件:\n" +
			"• my_resume.yaml (YAML格式)\n" +
			"• my_resume.md (Markdown格式)\n\n" +
			"您可以使用以下命令生成更多格式:\n" +
			"• resumgo generate my_resume.yaml -f markdown\n\n" +
			"按任意键退出..."
	}

	switch m.currentStep {
	case StepWelcome:
		return m.renderWelcomeView()
	case StepConfirm:
		return m.renderConfirmView()
	default:
		return m.renderFormView()
	}
}

// renderWelcomeView renders the welcome screen
func (m Model) renderWelcomeView() string {
	var s strings.Builder

	s.WriteString("✨ 欢迎使用 ResuGo 简历生成工具 ✨\n\n")
	s.WriteString("请选择操作:\n\n")

	for i, choice := range m.choices {
		cursor := "  "
		if m.cursor == i {
			cursor = "▶ "
		}
		s.WriteString(fmt.Sprintf("%s %s\n", cursor, choice))
	}
	s.WriteString("\n使用 ↑/↓ 选择，Enter 确认，Ctrl+C 退出\n")

	return s.String()
}

// renderConfirmView renders the confirmation screen
func (m Model) renderConfirmView() string {
	var s strings.Builder

	s.WriteString("📋 确认信息\n\n")
	s.WriteString("请确认您的简历信息:\n\n")
	s.WriteString(fmt.Sprintf("姓名: %s\n", m.resume.PersonalInfo.Name))
	s.WriteString(fmt.Sprintf("邮箱: %s\n", m.resume.PersonalInfo.Email))
	if m.resume.Summary != "" {
		s.WriteString(fmt.Sprintf("个人简介: %s\n", m.resume.Summary))
	}
	s.WriteString(fmt.Sprintf("教育背景: %d 项\n", len(m.resume.Education)))
	s.WriteString(fmt.Sprintf("工作经验: %d 项\n", len(m.resume.Experience)))
	s.WriteString(fmt.Sprintf("项目经验: %d 项\n", len(m.resume.Projects)))
	s.WriteString(fmt.Sprintf("自定义章节: %d 项\n", len(m.resume.Additional)))
	s.WriteString("\nEnter 保存简历，Esc 返回修改\n")

	return s.String()
}

// renderFormView renders the form input view
func (m Model) renderFormView() string {
	var s strings.Builder

	stepNames := map[int]string{
		StepPersonalInfo:   "📝 个人信息",
		StepSummary:        "📄 个人简介",
		StepEducation:      "🎓 教育背景",
		StepExperience:     "💼 工作经验",
		StepProjects:       "🚀 项目经验",
		StepSkills:         "🛠️ 技能",
		StepCustomSections: "✨ 自定义章节",
	}

	stepName := stepNames[m.currentStep]
	s.WriteString(fmt.Sprintf("%s\n\n", stepName))

	if m.editingList {
		s.WriteString("📝 列表编辑模式\n\n")
		s.WriteString("当前列表项:\n")

		for i, item := range m.listItems {
			cursor := "  "
			if i == m.listIndex {
				cursor = "▶ "
			}
			displayItem := item
			if displayItem == "" {
				displayItem = "(空白项)"
			}
			if i == m.listIndex {
				s.WriteString(fmt.Sprintf("%s[%s_]\n", cursor, displayItem))
			} else {
				s.WriteString(fmt.Sprintf("%s%s\n", cursor, displayItem))
			}
		}
		s.WriteString("\nCtrl+N 新增项，Del 删除项，Enter 完成编辑，Esc 取消\n")
	} else {
		for i, field := range m.fields {
			cursor := "  "
			if i == m.currentField {
				cursor = "▶ "
			}

			required := ""
			if field.Required {
				required = " *"
			}

			s.WriteString(fmt.Sprintf("%s%s%s:\n", cursor, field.Label, required))

			if i == m.currentField {
				if field.IsList {
					s.WriteString(fmt.Sprintf("  [%s] (按Enter编辑)\n", field.Value))
				} else {
					s.WriteString(fmt.Sprintf("  [%s_]\n", field.Value))
				}
			} else {
				value := field.Value
				if value == "" {
					value = fmt.Sprintf("(%s)", field.Placeholder)
				}
				s.WriteString(fmt.Sprintf("  %s\n", value))
			}
			s.WriteString("\n")
		}

		if m.error != "" {
			s.WriteString(fmt.Sprintf("❌ %s\n\n", m.error))
		}

		s.WriteString("Enter 下一步，↑/↓ 切换字段，Tab 快速切换，Esc 返回上一步\n")
	}

	return s.String()
}
