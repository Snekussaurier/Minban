# MinBan

> A minimalist Kanban board application built with modern web technologies

MinBan is a sleek, lightweight Kanban board application that helps you organize your tasks and projects efficiently. Built with a Go backend and a Rust/Dioxus frontend, it offers a fast and responsive user experience.

## 🌟 Features

- **Intuitive Kanban Interface**: Drag and drop cards between columns
- **Board Management**: Create and manage multiple boards
- **Card Organization**: Add, edit, and delete cards with ease
- **Tag System**: Organize cards with customizable tags
- **User Authentication**: Secure login system with JWT tokens
- **Responsive Design**: Works seamlessly on desktop and mobile devices
- **Dark Theme**: Built-in dark mode with custom color scheme

## 🛠️ Tech Stack

### Backend
- **Go** with Gin framework
- **SQLite** database with GORM ORM
- **JWT** authentication
- **Docker** containerization
- **RESTful API** design

### Frontend
- **Rust** with Dioxus framework
- **WebAssembly** for optimal performance
- **Tailwind CSS** for styling
- **Modern responsive design**

## 🚀 Quick Start

### Prerequisites
- Go 1.23+
- Rust with cargo
- Docker (optional)

### Development Setup

1. **Clone the repository**
```bash
git clone https://github.com/snekussaurier/minban.git
cd minban
```

2. **Backend Setup**
```bash
cd backend/src
go mod download
go run main.go
```
The backend will start on port 9916.

3. **Frontend Setup**
```bash
cd frontend
dx serve
```
The frontend will be available at `http://localhost:8080`.

### Docker Setup

For a quick deployment using Docker:

```bash
cd backend
docker-compose up
```

This will start the backend with default credentials:
- Username: `snekussaurier`
- Password: `123`

## 📁 Project Structure

```
minban/
├── backend/          # Go backend application
│   ├── src/
│   │   ├── controller/   # API controllers
│   │   ├── database/     # Database models and config
│   │   ├── middleware/   # Authentication & CORS
│   │   ├── routes/       # API routes
│   │   └── utils/        # Utility functions
│   ├── data/             # SQLite database storage
│   ├── static/           # Compiled frontend assets
│   └── docs/             # API documentation
├── frontend/         # Rust/Dioxus frontend
│   ├── src/
│   │   ├── api/          # API client functions
│   │   ├── components/   # UI components
│   │   ├── utils/        # Frontend utilities
│   │   └── main.rs       # Application entry point
│   └── assets/           # Static assets
└── docs/             # Project documentation
```

## 🔧 API Endpoints

The backend provides a RESTful API with the following main endpoints:

- `POST /login` - User authentication
- `GET /boards` - Get all boards
- `POST /boards` - Create a new board
- `PATCH /boards/:id` - Update board
- `GET /cards` - Get cards for a board
- `POST /cards` - Create a new card
- `PATCH /cards/:id` - Update card
- `DELETE /cards/:id` - Delete card
- `GET /tags` - Get all tags
- `GET /states` - Get board states/columns

## 🎨 Customization

### Theme Colors
The application uses a custom color scheme defined in the Tailwind config:
- Primary: `#5a5b70` (minban_dark)
- Accent: `#a294f9` (minban_highlight)
- Font: Poppins

### Configuration
Backend configuration is handled through environment variables:
- `DATABASE_PATH`: SQLite database file path
- `USER_NAME`: Default username
- `USER_PASSWORD`: Default password
- `JWT_SECRET_KEY`: JWT signing secret

## 🧪 Development

### Backend Development
```bash
cd backend/src
go run main.go
```

### Frontend Development
```bash
cd frontend
dx serve --hot-reload
```

### Building for Production
```bash
# Backend
cd backend/src
go build -o minban-backend

# Frontend
cd frontend
dx build --release
```

## 📝 License

This project is licensed under the AGPL 3.0 License - see the [LICENSE](../LICENSE) file for details.

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
