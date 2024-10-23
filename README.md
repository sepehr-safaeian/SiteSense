# SiteSense

SiteSense is a free and open-source command-line interface (CLI) tool developed in Go that analyzes URLs for various essential web metrics. With a focus on SEO, performance, security, usability, and accessibility, SiteSense empowers developers and webmasters to enhance their websites effectively.

## Table of Contents
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Sample Output](#sample-output)
- [Contributing](#contributing)
- [License](#license)
- [Contact](#contact)

## Features

- **SEO Metrics Analysis**: Evaluate essential SEO components such as meta tags, internal/external links, image alt texts, and more.
- **Performance Insights**: Analyze page loading speeds and resource optimization.
- **Security Checks**: Basic security assessments to identify potential vulnerabilities.
- **Usability Evaluation**: Review user interface elements and design usability aspects.
- **Accessibility Checks**: Ensure your website is accessible to all users, including those with disabilities.

## Installation

To install SiteSense, ensure you have Go installed on your system. Then, run the following command in your terminal:

```bash
go run build ./main.go
```


# Usage
To analyze a website, execute the following command:

```bash
./main scan https://example.com
```

The analysis results will be saved in the following path:
```bash
/data/example.com.txt
```


# Sample Output
Here’s a sample output for an SEO analysis:

```bash
SEO Metrics:

URL: https://example.com

Meta Title Tag Present: ✅
Meta Title Tag Text: example
Meta Title Tag Length: 17

Meta Description Present: ✅
Meta Description: Lorem Ipsum
Meta Description Length: 245

Canonical Tag Present: ✅

Sitemap XML Present: ✅

Internal Links Count: 145

External Links Count: 4

Images Count: 14

Images with Alt Text Present: ❌

Count of Images without Alt Text: 3

Author Present: ❌

Language Present: ✅
Language: en-US

ViewPort: ✅
Viewport Content: width=device-width, initial-scale=1

Robots: ✅

Alt Texts:
 - k1
 - k2
 - k3
 - k4

Images without Alt Text:
 - https://example.com/wp-content/uploads/2024/05/5699565.webp
 - https://example.com/wp-content/uploads/2024/04/81.png

```
