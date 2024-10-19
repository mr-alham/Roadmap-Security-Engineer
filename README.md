![Preview of the app](https://i.ibb.co/KKVjcfZ/Roadmap-security-engineer-default.png)

# Roadmap to be a Security Engineer

This project is designed to help me master the **skills** required to become a **Security Engineer** by **studying** and **practicing** one **subtopic** at a time. Developed in **_Golang_**, this platform is structured to follow a **comprehensive roadmap**, integrating **practical projects** that align with **real-world security practices**.

## The Problem

As a **_Computer Science_** student specializing in **Cybersecurity**, To bridge the gap between theoretical focus and the practical skills needed in the industry, I’ve designed a project that provides **hands-on experience** in various areas of security engineering. The goal is to have a portfolio of **real-world projects** that will showcase my **skills** and preparedness for a Security Engineer role by the time I graduate. This **project-based** approach keeps my knowledge **current** and aligned with **industry standards**, even while I continue my studies.

## The Solution

I built a **structured cybersecurity learning platform** that follows a clear roadmap, broken down into **categories** and **subtopics** such as,

- Foundational Security Knowledge

  - Security Principles and CIA Triad
  - Risk Management
  - Authentication, Authorization, and Accounting (AAA)
  - Cryptography

- Linux Security

  - User and Group Management
  - SELinux & AppArmor
  - Network Security in Linux

- Cloud Security (AWS/Azure)

  - Identity and Access Management (IAM)
  - Network Security in Cloud

- Incident Response & Forensics

  - Incident Response Frameworks
  - Digital Forensics

- Offensive Operations
- Advanced Cyber Defense

Each subtopic includes **theoretical** concepts followed by **hands-on projects** that solidify my understanding. For example, learning about **encryption algorithms** includes a project that implements **encryption** and **decryption** using **OpenSSL**.

## Technical Overview

### Language(s)

- **Golang**
- **JSON**

### Platform

- **Linux** - _Fedora_
- **Visual Studio Code** - \*VS code for development

### Focus Areas:

- **Security principles**
  Deep dive into core security concepts like confidentiality, integrity, and availability (CIA Triad).

- **Linux security**
  Mastery over user and group management, SELinux, and network firewall management.

- **Cloud security**
  Expertise in cloud identity management, AWS/Azure security practices, and cloud networking.

- **Incident response**
  Simulated scenarios for detecting and handling security breaches.

- **Offensive operations**
  Experience with penetration testing techniques and red team operations.

This program emphasizes **modular code structure**, leveraging Golang’s powerful **concurrency model** for efficient task execution. It uses **JSON** for dynamic data loading of the security roadmap, ensuring easy **extensibility** for new topics and categories as I **advance**.

## Key Features

- **Structured Roadmap**

  > The program is designed with a logical progression from foundational knowledge to advanced topics, ensuring I build expertise in a step-by-step manner.

- **Hands-On Projects**

  > These projects are not just theoretical exercises but simulate real-world scenarios Security Engineers face daily, ensuring I gain experience in tasks like firewall configuration, encryption practices, and cloud security setups.

- **Progress Tracking**

  > The platform tracks my progress by storing the current topic and learning state, enabling resumption from where I left off.

- **Custom Navigation**

  > Allows jumping to any desired topic or subtopic, enabling focused study on specific areas of interest.

- **Easy Installation & Setup**
  > Built as a command-line tool, it’s easy to install and use, either by downloading the pre-built binary or compiling from source.

## Installation

1. **Run the Program Directly**

- Download the latest `Roadmap-App.zip` file from the `GitHub` repository.

  ```text
  https://github.com/mr-alham/Roadmap-Security-Engineer.git
  ```

- Extract the `Roadmap-App.zip`:

  ```bash
  unzip Roadmap-App.zip
  ```

- Move the `roadmap-sec-engineer.d directory` to `/usr/local/etc/`:

  ```bash
  sudo mv roadmap-sec-engineer.d /usr/local/etc/
  ```

- (Optional) Move the `roadmap-sec-engineer` binary to `/usr/local/bin/` to use it as a regular Linux command:

  ```bash
  sudo mv roadmap-sec-engineer /usr/local/bin/
  ```

- Run the program:

  ```bash
  roadmap-sec-engineer
  ```

2. **Compile the Program Yourself**

   - Ensure you have Golang installed.
   - Clone the repository:

     ```bash
     git clone https://github.com/mr-alham/Roadmap-Security-Engineer.git
     cd Roadmap-Security-Engineer
     ```

   - Install dependencies:

     ```bash
     go get
     ```

   - Build the program:

     ```bash
     go mod tidy
     go build -a -o roadmap-sec-engineer
     ```

   - Run the program:

     ```bash
     ./roadmap-sec-engineer
     ```

## Usage

- **Show Current Topic**

  ```bash
  roadmap-sec-engineer
  ```

  <details>
    <summary>Show the Screenshot</summary>

  ![Preview for current topic](https://i.ibb.co/qpWZ2Lj/Roadmap-security-engineer-done.png)

  </details>

- **Mark Current Topic as Completed**

  ```bash
  roadmap-sec-engineer --done
  ```

  <details>
    <summary>Show the Screenshot</summary>

  ![Preview for current topic](https://i.ibb.co/qpWZ2Lj/Roadmap-security-engineer-done.png)

  </details>

- **Jump to a Specific Topic**

  ```bash
  roadmap-sec-engineer --custom <desired day>
  ```

  <details>
    <summary>Show the Screenshot</summary>

  ![Preview for current topic](https://i.ibb.co/q1Bbrw0/Roadmap-security-engineer-custom.png)

  </details>

- **Display Help Menu**

  ```bash
  roadmap-sec-engineer --help
  ```

  <details>
    <summary>Show the Screenshot</summary>

  ![Preview for current topic](https://i.ibb.co/41DFYqF/Roadmap-security-engineer-help.png)

  </details>
