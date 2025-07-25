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
	return fmt.Sprintf("\n%s\n\n%s",
		m.welcomeList.View(),
		"按 Enter 选择，Ctrl+C 退出",
	)
}

// renderConfirmView renders the confirmation screen
func (m Model) renderConfirmView() string {
	var s strings.Builder

	s.WriteString("📋 确认信息\n\n")
	s.WriteString("请确认您的简历信息:\n\n")

	// Personal Information
	s.WriteString("👤 个人信息:\n")
	s.WriteString(fmt.Sprintf("  姓名: %s\n", m.resume.PersonalInfo.Name))
	s.WriteString(fmt.Sprintf("  邮箱: %s\n", m.resume.PersonalInfo.Email))
	if m.resume.PersonalInfo.Phone != "" {
		s.WriteString(fmt.Sprintf("  电话: %s\n", m.resume.PersonalInfo.Phone))
	}
	if m.resume.PersonalInfo.Location != "" {
		s.WriteString(fmt.Sprintf("  地址: %s\n", m.resume.PersonalInfo.Location))
	}
	s.WriteString("\n")

	// Summary
	if m.resume.Summary != "" {
		s.WriteString("📄 个人简介:\n")
		s.WriteString(fmt.Sprintf("  %s\n\n", m.resume.Summary))
	}

	// Education
	if len(m.resume.Education) > 0 {
		s.WriteString("🎓 教育背景:\n")
		for _, edu := range m.resume.Education {
			// 学校名称
			s.WriteString(fmt.Sprintf("  %s\n", edu.Institution))

			// 学位和专业
			if edu.Major != "" {
				s.WriteString(fmt.Sprintf("    %s (%s)\n", edu.Degree, edu.Major))
			} else {
				s.WriteString(fmt.Sprintf("    %s\n", edu.Degree))
			}

			// 地点和时间
			timeStr := fmt.Sprintf("%d-%d", edu.StartDate.Year(), edu.EndDate.Year())
			if edu.Current {
				timeStr = fmt.Sprintf("%d-至今", edu.StartDate.Year())
			}
			s.WriteString(fmt.Sprintf("    %s | %s\n", edu.Location, timeStr))
		}
		s.WriteString("\n")
	}

	// Experience
	if len(m.resume.Experience) > 0 {
		s.WriteString("💼 工作经验:\n")
		for _, exp := range m.resume.Experience {
			// 公司名称
			s.WriteString(fmt.Sprintf("  %s\n", exp.Company))

			// 职位
			s.WriteString(fmt.Sprintf("    %s\n", exp.Position))

			// 地点和时间
			timeStr := fmt.Sprintf("%s-%s", exp.FormatStartDate(), exp.FormatEndDate())
			s.WriteString(fmt.Sprintf("    %s | %s\n", exp.Location, timeStr))
		}
		s.WriteString("\n")
	}

	// Projects
	if len(m.resume.Projects) > 0 {
		s.WriteString("🚀 项目经验:\n")
		for _, proj := range m.resume.Projects {
			// 项目名称
			s.WriteString(fmt.Sprintf("  %s\n", proj.Name))

			// 项目描述
			s.WriteString(fmt.Sprintf("    %s\n", proj.Description))

			// 地点和时间
			timeStr := fmt.Sprintf("%s-%s", proj.FormatStartDate(), proj.FormatEndDate())
			if proj.Location != "" {
				s.WriteString(fmt.Sprintf("    %s | %s\n", proj.Location, timeStr))
			} else {
				s.WriteString(fmt.Sprintf("    %s\n", timeStr))
			}
		}
		s.WriteString("\n")
	}

	// Skills
	if len(m.resume.Skills.Languages) > 0 || len(m.resume.Skills.Frameworks) > 0 {
		s.WriteString("🛠️ 技能:\n")
		if len(m.resume.Skills.Languages) > 0 {
			s.WriteString(fmt.Sprintf("  编程语言: %s\n", strings.Join(m.resume.Skills.Languages, ", ")))
		}
		if len(m.resume.Skills.Frameworks) > 0 {
			s.WriteString(fmt.Sprintf("  框架/库: %s\n", strings.Join(m.resume.Skills.Frameworks, ", ")))
		}
		s.WriteString("\n")
	}

	// Custom sections
	if len(m.resume.Additional) > 0 {
		s.WriteString("✨ 自定义章节:\n")
		for _, section := range m.resume.Additional {
			s.WriteString(fmt.Sprintf("  %s\n", section.Title))
			for _, item := range section.Items {
				s.WriteString(fmt.Sprintf("    • %s\n", item))
			}
		}
		s.WriteString("\n")
	}

	s.WriteString("Enter 保存简历，Esc 返回修改\n")

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
				} else if field.Multiline {
					// Render textarea for multiline fields
					s.WriteString(fmt.Sprintf("  %s\n", m.textArea.View()))
				} else {
					// Render textinput for single-line fields
					s.WriteString(fmt.Sprintf("  %s\n", m.textInputs[i].View()))
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

		s.WriteString("Enter 下一步，↑/↓ 或 Tab(向下)/Shift+Tab(向上) 切换字段，j/k 仅用于输入，Esc 返回上一步\n")
	}

	return s.String()
}
