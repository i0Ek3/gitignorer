package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	projectTypes := detectProjectTypes(".")

	gitignore := generateGitignore(projectTypes)

	target := ".gitignore"

	// backup existing .gitignore
	if _, err := os.Stat(target); err == nil {
		_ = os.Rename(target, target+".bak")
		fmt.Println("⚠️  已存在 .gitignore，已自动备份为 .gitignore.bak")
	}

	err := os.WriteFile(target, []byte(gitignore), 0644)
	if err != nil {
		fmt.Println("写入 .gitignore 失败:", err)
		return
	}

	fmt.Println("✅ 已生成 .gitignore")
}

// -------------------------------
// 自动识别项目类型
// -------------------------------
func detectProjectTypes(root string) map[string]bool {
	types := map[string]bool{}

	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		name := strings.ToLower(info.Name())

		switch name {
		case "go.mod":
			types["go"] = true
		case "package.json":
			types["node"] = true
		case "requirements.txt", "pyproject.toml":
			types["python"] = true
		case "cargo.toml":
			types["rust"] = true
		case "composer.json":
			types["php"] = true
		case "pubspec.yaml":
			types["flutter"] = true
		case "build.gradle", "pom.xml":
			types["java"] = true
		case "project.swift", "package.swift":
			types["swift"] = true
		}

		if strings.HasSuffix(name, ".xcodeproj") ||
			strings.HasSuffix(name, ".xcworkspace") {
			types["xcode"] = true
		}

		return nil
	})

	return types
}

// -------------------------------
// 根据类型生成丰富的 .gitignore
// -------------------------------
func generateGitignore(types map[string]bool) string {
	var out []string

	// 通用项
	out = append(out, commonGitignore())

	// 语言和框架
	if types["go"] {
		out = append(out, gitignoreGo())
	}
	if types["python"] {
		out = append(out, gitignorePython())
	}
	if types["node"] {
		out = append(out, gitignoreNode())
	}
	if types["rust"] {
		out = append(out, gitignoreRust())
	}
	if types["java"] {
		out = append(out, gitignoreJava())
	}
	if types["php"] {
		out = append(out, gitignorePHP())
	}
	if types["flutter"] {
		out = append(out, gitignoreFlutter())
	}
	if types["swift"] || types["xcode"] {
		out = append(out, gitignoreSwiftXcode())
	}

	// 编辑器等通用补充
	out = append(out, gitignoreVSCode(), gitignoreJetBrains(), gitignoreMacOS())

	return strings.Join(out, "\n\n")
}

// -------------------------------
// 下面是丰富的模板内容
// -------------------------------

func commonGitignore() string {
	return `# =========================================
# COMMON
# =========================================
.DS_Store
Thumbs.db
*.log
*.tmp
*.swp
*.bak
*.temp

# ENV files
.env
.env.*
*.secret

# Git
.gitmodules
.svn/

# System
.idea/
.vscode/
*.iml
`
}

func gitignoreGo() string {
	return `# =========================================
# GO
# =========================================
bin/
vendor/
*.exe
*.exe~
*.dll
*.so
*.dylib
*.test
go.sum
`
}

func gitignorePython() string {
	return `# =========================================
# PYTHON
# =========================================
__pycache__/
*.py[cod]
*.pyo
*.pyd
*.pdb
*.egg
*.egg-info/
dist/
build/
pip-wheel-metadata/
*.sqlite3
*.db
.env/
.venv/
`
}

func gitignoreNode() string {
	return `# =========================================
# NODE / JAVASCRIPT / TYPESCRIPT
# =========================================
node_modules/
npm-debug.log*
yarn-debug.log*
yarn-error.log*
dist/
build/
.cache/
.next/
out/
*.tsbuildinfo
`
}

func gitignoreRust() string {
	return `# =========================================
# RUST
# =========================================
target/
Cargo.lock
*.rs.bk
`
}

func gitignoreJava() string {
	return `# =========================================
# JAVA / KOTLIN / ANDROID
# =========================================
target/
bin/
build/
.gradle/
.gradle-cache/
*.class
*.jar
*.war
*.ear
*.iml
local.properties
`
}

func gitignorePHP() string {
	return `# =========================================
# PHP / LARAVEL
# =========================================
vendor/
storage/
bootstrap/cache/
*.cache
`
}

func gitignoreFlutter() string {
	return `# =========================================
# FLUTTER / DART
# =========================================
.build/
.dart_tool/
.packages
.pub-cache/
pubspec.lock
ios/Pods/
android/.gradle/
android/app/build/
`
}

func gitignoreSwiftXcode() string {
	return `# =========================================
# SWIFT / XCODE / IOS
# =========================================
build/
DerivedData/
*.xcworkspace/xcuserdata/
*.xcodeproj/xcuserdata/
*.xcuserstate
*.swiftpm
.swiftpm/
Carthage/
Pods/
`
}

func gitignoreVSCode() string {
	return `# =========================================
# VS CODE
# =========================================
.vscode/*
!.vscode/settings.json
!.vscode/tasks.json
!.vscode/launch.json
!.vscode/extensions.json
`
}

func gitignoreJetBrains() string {
	return `# =========================================
# JETBRAINS
# =========================================
.idea/
*.iml
out/
`
}

func gitignoreMacOS() string {
	return `# =========================================
# macOS
# =========================================
.DS_Store
.AppleDouble
.LSOverride
`
}
