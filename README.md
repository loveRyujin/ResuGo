# ResuGo

ResuGo is an interactive command-line tool for creating and managing professional resumes. It uses elegant terminal user interface powered by Bubble Tea and provides multiple output formats.

## Features

- 📝 Interactive resume creation with a beautiful TUI (Terminal User Interface)
- 🎨 Multiple output formats (YAML, Markdown)
- 🚀 Fast and efficient Go-based CLI tool
- 📋 Structured resume data using YAML format
- 🔧 Extensible architecture for adding new features

## Installation

### From Source

```bash
git clone https://github.com/loveRyujin/ResuGo.git
cd ResuGo
go build -o resumgo .
```

### Prerequisites

- Go 1.21 or higher

## Usage

### Commands

#### Create a new resume interactively
```bash
./resumgo create
```

This opens an interactive terminal interface where you can input your resume information step by step.

#### Generate resume from YAML file
```bash
./resumgo generate input.yaml -f markdown -o resume.md
```

Convert a YAML resume file to different formats:
- `-f, --format`: Output format (yaml, markdown)
- `-o, --output`: Output file path

#### Show version
```bash
./resumgo version
```

#### Show help
```bash
./resumgo --help
./resumgo [command] --help
```

### Examples

#### Generate Markdown resume from example
```bash
./resumgo generate templates/example.yaml -f markdown -o my_resume.md
```

#### Generate YAML resume (useful for reformatting)
```bash
./resumgo generate templates/example.yaml -f yaml -o formatted_resume.yaml
```

## Resume Structure

ResuGo uses YAML format for resume data with the following structure:

```yaml
personal_info:
  name: "Your Name"
  title: "Your Professional Title"
  email: "your.email@example.com"
  phone: "+1-555-0123"
  location: "Your City, State"
  website: "https://yourwebsite.com"
  github: "https://github.com/yourusername"
  linkedin: "https://linkedin.com/in/yourusername"
  summary: "Your professional summary"

education:
  - institution: "University Name"
    degree: "Degree Type"
    major: "Your Major"
    start_date: "2010-08-01T00:00:00Z"
    end_date: "2014-05-01T00:00:00Z"
    gpa: "3.7"
    description: "Additional details"

experience:
  - company: "Company Name"
    position: "Your Position"
    location: "City, State"
    start_date: "2020-01-01T00:00:00Z"
    end_date: "2025-01-01T00:00:00Z"
    current: true
    description:
      - "Achievement or responsibility 1"
      - "Achievement or responsibility 2"

skills:
  - category: "Programming Languages"
    items: ["Go", "JavaScript", "Python"]
    level: "expert"  # beginner, intermediate, advanced, expert

projects:
  - name: "Project Name"
    description: "Project description"
    start_date: "2023-01"
    end_date: "2023-08"
    technologies: ["Go", "React", "PostgreSQL"]
    url: "https://project-demo.com"
    repository: "https://github.com/username/project"
    highlights:
      - "Key achievement 1"
      - "Key achievement 2"

languages:
  - name: "English"
    level: "native"  # native, fluent, good, basic
```

## Project Structure

```
.
├── cmd/                    # Cobra commands
│   ├── create.go          # Interactive resume creation
│   ├── generate.go        # Resume generation from YAML
│   ├── root.go            # Root command setup
│   └── version.go         # Version command
├── internal/
│   ├── generator/         # Resume generation logic
│   │   └── generator.go   # Output format generators
│   ├── models/            # Data structures
│   │   └── resume.go      # Resume model definitions
│   └── ui/                # Terminal UI components
│       └── create.go      # Interactive creation UI
├── templates/
│   └── example.yaml       # Example resume template
├── go.mod
├── go.sum
├── main.go               # Application entry point
└── README.md
```

## Dependencies

- [Cobra](https://github.com/spf13/cobra) - CLI framework
- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - Terminal UI framework
- [YAML v3](https://gopkg.in/yaml.v3) - YAML parsing and generation

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## Future Features

- [ ] PDF export support
- [ ] HTML template generation
- [ ] Resume validation and suggestions
- [ ] Multiple resume templates
- [ ] Resume analytics and optimization tips
- [ ] Cloud storage integration
- [ ] Resume comparison tool

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Author

Created by [@loveRyujin](https://github.com/loveRyujin)
一个命令行工具，用来生成简历
