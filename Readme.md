<div align="center">
<h1 align="center">
<img src="https://raw.githubusercontent.com/CamilleLange/phoenix/main/gopher.png" width="100" />
<br>Phoenix
</h1>
<h3>An easy to use CLI to transfert Go projects to an other git repository</h3>

<p align="center">
<img src="https://img.shields.io/badge/Go-00ADD8.svg?style&logo=Go&logoColor=white" alt="Go" />
</p>
<img src="https://img.shields.io/github/languages/top/CamilleLange/phoenix?style&color=5D6D7E" alt="GitHub top language" />
<img src="https://img.shields.io/github/languages/code-size/CamilleLange/phoenix?style&color=5D6D7E" alt="GitHub code size in bytes" />
<img src="https://img.shields.io/github/commit-activity/m/CamilleLange/phoenix?style&color=5D6D7E" alt="GitHub commit activity" />
<img src="https://img.shields.io/github/license/CamilleLange/phoenix?style&color=5D6D7E" alt="GitHub license" />
</div>

---

## 📖 Table of Contents
- [📖 Table of Contents](#-table-of-contents)
- [📍 Overview](#-overview)
- [📦 Features](#-features)
- [📂 Repository Structure](#-repository-structure)
- [🚀 Getting Started](#-getting-started)
    - [🔧 Installation](#-installation)
    - [🤖 Running phoenix](#-running-phoenix)
- [🤝 Contributing](#-contributing)
- [📄 License](#-license)

---


## 📍 Overview

An easy to use CLI to transfert Go projects to an other git repository

---

## 📦 Features

Phoenix allows you to revive a Go project that has changed Git repository.

Phoenix will modify your Go files to use the new module name.

---


## 📂 Repository Structure

```sh
└── phoenix/
    ├── LICENSE
    ├── cmd/
    │   └── root.go
    ├── go.mod
    ├── go.sum
    ├── main.go
    └── pkg/
        └── phoenix/
            ├── phoenix.go
            └── utils.go
```

---

## 🚀 Getting Started

### 🔧 Installation

1. Clone the phoenix repository:
```sh
git clone https://github.com/CamilleLange/phoenix
```

2. Change to the project directory:
```sh
cd phoenix
```

3. Install the dependencies:
```sh
go build -o phoenix
```

### 🤖 Running phoenix

```sh
./phoenix -h
```

---

## 🤝 Contributing

Contributions are always welcome! Please follow these steps:
1. Fork the project repository. This creates a copy of the project on your account that you can modify without affecting the original project.
2. Clone the forked repository to your local machine using a Git client like Git or GitHub Desktop.
3. Create a new branch with a descriptive name (e.g., `new-feature-branch` or `bugfix-issue-123`).
```sh
git checkout -b new-feature-branch
```
4. Make changes to the project's codebase.
5. Commit your changes to your local branch with a clear commit message that explains the changes you've made.
```sh
git commit -m 'Implemented new feature.'
```
6. Push your changes to your forked repository on GitHub using the following command
```sh
git push origin new-feature-branch
```
7. Create a new pull request to the original project repository. In the pull request, describe the changes you've made and why they're necessary.
The project maintainers will review your changes and provide feedback or merge them into the main branch.

---

## 📄 License

This project is licensed under the `ℹ️  LICENSE-TYPE` License. See the [LICENSE-Type](LICENSE) file for additional info.

---
