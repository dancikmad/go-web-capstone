# Restaurant Review Web Application

A simple web application built with Go's `net/http` package that allows users to browse a restaurant menu, submit reviews, and view other users' feedback.

## ✨ Features

- Browse restaurant menu items
- Submit reviews for menu items
- View user reviews in real-time
- RESTful API for menu and review management
- Clean frontend with HTML and CSS
- Docker containerization for consistent deployment

## 🚀 Getting Started

### Option 1: Using Docker Compose (Recommended)

This method ensures consistent operation across all environments.

1. **Clone the repository:**
   ```bash
   git clone <repository_url>
   cd <repository_directory>
   ```

2. **Build and run with Docker Compose:**
   ```bash
   make up
   ```
   This builds the Docker images and starts the containers in detached mode.

3. **Open in browser:**
   ```bash
   make open
   ```
   Opens the application in your default browser at http://localhost:4001.

### Option 2: Running Directly

For local development with Go installed:

1. **Clone the repository:**
   ```bash
   git clone <repository_url>
   cd <repository_directory>
   ```

2. **Start the API server:**
   ```bash
   cd server
   go run requests.go
   ```
   This starts the API server on port 4002.

3. **Start the frontend server:**
   ```bash
   # In a new terminal, from the project root
   go run .
   ```
   This starts the frontend server on port 4001.

4. **Access the application** at http://localhost:4001 in your browser.

## 🛠️ Technology Stack

- Go (`net/http`)
- HTML/CSS
- Docker & Docker Compose

Happy reviewing! 🍽️
