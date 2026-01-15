# TODO Application (Go + React)

This project is an experimental first step in exploring **Go (Golang)** as a backend for a **React** frontend. It is a full-stack TODO application that utilizes MongoDB for persistent storage, Chakra UI for the interface, and TanStack Query for state management.

## 🚀 Getting Started

### Prerequisites

- **Go** installed on your machine.
- **Node.js** and **npm/yarn** installed.
- **MongoDB** instance (local or Atlas).

- **Air** (Go live-reloading) for a better developer experience.

### Backend Setup

1. Create a `.env` file in the root directory:

```env
MONGODB_URI=your_mongodb_connection_string
PORT=4000

```

2. For ease of use and hot-reloading during development, this project is configured to work with **Air**. Run the following command in your terminal within the root directory:

```bash
air

```

_Note: This automatically rebuilds the Go application whenever you save changes to your `.go` files._

### Frontend Setup

1. Navigate to the client directory:

```bash
cd client

```

2. Install dependencies and start the development server:

```bash
npm install
npm run dev

```

---

## 🛠️ Tech Stack

### Backend

- **Go**: Core language.

- **Fiber**: Express-inspired web framework for Go.

- **MongoDB Driver**: Official Go driver for database interactions.

- **Godotenv**: For managing environment variables.

### Frontend

- **React**: UI library.

- **Chakra UI**: Component library for styling and dark mode support.

- **TanStack Query**: For efficient data fetching and caching.

- **Vite**: Frontend build tool.

---

## 📋 Features

- **CRUD Operations**: Create, Read, Update (mark as completed), and Delete tasks.

- **Dark Mode**: Seamlessly toggle between light and dark themes via the Navbar.

- **Responsive UI**: Optimized for various screen sizes using Chakra UI's layout components.

- **Real-time Updates**: TanStack Query ensures the UI stays in sync with the backend after mutations.

---

## ⚖️ General Ethics & Conduct

This project is developed with the following principles:

- **Data Integrity**: Ensuring user-inputted TODOs are handled securely and stored correctly in MongoDB.

- **Transparency**: This is an open-source experimental project intended for learning and growth.
- **Clean Code**: Adhering to readable formatting and modular component architecture.
